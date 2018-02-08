package main


import (
    "log"
    "os"
    "flag"
    "fmt"
    "strconv"
    "net"
    "io"
    "errors"
    "encoding/binary"
    ss "ss-mock/shadowsocks"
)

var debug ss.DebugLog

var (
    errAddrType         = errors.New("socks addr type not supported")        
    errVer              = errors.New("socks version not supported")        
    errAuthExtraData    = errors.New("socks authentication get extra data")        
    errReqExtraData  = errors.New("socks request get extra data")
    errCmd           = errors.New("socks command not supported")
)

const (
    socksVer5       = 5 //version
    socksCmdConnect = 1 //cmd
)

type ServerCipher struct {
    server string
    cipher  *ss.Cipher
}

var servers struct {
    srvCipher   []*ServerCipher
    failCnt     []int // failed connection count
}

func enoughOptions(config *ss.Config) bool {
    return config.Server != nil && config.ServerPort != 0 &&
        config.LocalPort != 0 && config.Password != "" 
}

func parseServerConfig(config *ss.Config) {
    hasPort := func(s string) bool {
        _, port, err := net.SplitHostPort(s)
        if err != nil {
            return false
        }
        return port != ""
    }
    
    if len(config.ServerPassword) == 0 {
        method := config.Method
        //auth
        cipher, err := ss.NewCipher(method, config.Password)
        if err != nil {
            log.Fatal("failed generating ciphers:", err)
        }
        srvPort := strconv.Itoa(config.ServerPort)
        srvArr := config.GetServerArray()
        n := len(srvArr)
        servers.srvCipher = make([]*ServerCipher, n)
        for i, s := range srvArr {
            if hasPort(s) {
                log.Println("ignore server_port option for server", s)
                servers.srvCipher[i] = &ServerCipher{s, cipher}
            } else {
                servers.srvCipher[i] = &ServerCipher{net.JoinHostPort(s, srvPort), cipher}
            }
        }
    } //TODO else
    
    servers.failCnt = make([]int, len(servers.srvCipher))
    for _, se := range servers.srvCipher {
        log.Println("avaliable remote server", se.server)
    }
}

func handShake(conn net.Conn)(err error) {
    const (
        idVer = 0        //version byte index
        idNmethod = 1   //method byte index
    )

    buf := make([]byte, 258)
    var n int
    ss.SetReadTimeout(conn)

    if n, err = io.ReadAtLeast(conn, buf, idNmethod+1); err != nil {
        return
    }

    //buf_s := string(buf[:])
    //println(buf_s, n)

    if buf[idVer] != socksVer5 {
        return errVer
    }
    nmethod := int(buf[idNmethod]) //how many bytes method will occupy
    msgLen := nmethod + 2 //plus 2 means: version + nmethod
    if n == msgLen { //handshake done, common case
        //do nothing, jump directly to send confirmation
    } else if n < msgLen { //has more methods to read, rare case
        if _, err = io.ReadFull(conn, buf[n:msgLen]); err != nil {
            return
        }
    } else {//error, should not get extr data
        return errAuthExtraData
    }
    
    //send confirmation: version 5, no authentication required
    _, err = conn.Write([]byte{socksVer5, 0})
    return
}

func getRequest(conn net.Conn)(rawaddr []byte, host string, err error) {
    const (
        idVer = 0
        idCmd = 1
        idType = 3
        idIP0 = 4
        idDmLen = 4
        idDm0 = 5

        typeIPv4 = 1
        typeDm = 3
        typeIPv6 =4 

        lenIPv4 = 3 + 1 + net.IPv4len + 2 //3(ver+cmd+rsv) +1(addrtype) + ipv4 + 2(port)
        lenIPv6 = 3 + 1 + net.IPv6len + 2 //3(ver+cmd+rsv) +1(addrtype) + ipv6 +2(port)
        lenDmBase = 3 + 1 + 1 + 2  // 3(ver+cdm+rsv) +1(addrtype)+1(addrlen)+2(port), plus addrlen
    )

    buf := make([]byte, 263)
    var n int 
    ss.SetReadTimeout(conn)

    if n, err = io.ReadAtLeast(conn, buf, idDmLen+1); err != nil {
        return
    }

    //check version and cmd
    if buf[idVer] != socksVer5 {
        err = errVer
        return
    }

    if buf[idCmd] != socksCmdConnect {
        err = errCmd
        return
    }

    reqLen := -1
    switch buf[idType] {
        case typeIPv4:
            reqLen = lenIPv4
        case typeIPv6:
            reqLen = lenIPv6
        case typeDm:
            reqLen = int(buf[idDmLen]) + lenDmBase
        default:
            err = errAddrType
            return
    }
    
    if n == reqLen {
        //common case, do nothing
    } else if n < reqLen { // rare case 
        if _, err = io.ReadFull(conn, buf[n:reqLen]); err != nil {
            return
        }
    } else {
        err = errReqExtraData 
        return
    }

    rawaddr = buf[idType:reqLen]

    if debug {
        switch buf[idType] {
            case typeIPv4:
                host = net.IP(buf[idIP0 : idIP0+net.IPv4len]).String()
            case typeIPv6:
                host = net.IP(buf[idIP0 : idIP0+net.IPv6len]).String()
            case typeDm:
                host = string(buf[idDm0 : idDm0+buf[idDmLen]])
        }
        port := binary.BigEndian.Uint16(buf[reqLen-2:reqLen])
        host = net.JoinHostPort(host, strconv.Itoa(int(port)))
    }

    return

}

func connectToServer(serverId int, rawaddr []byte, addr string)(remote *ss.Conn, err error) {
    se := servers.srvCipher[serverId]
    remote, err = ss.DialWithRawAddr(rawaddr, se.server, se.cipher.Copy())
    if err != nil {
        log.Println("error connecting to shadowsockss server: ", err)
        //TODO retry
        return nil, err
    }
    debug.Printf("connected to %s via %s\n", addr, se.server)
    //TODO 
}
    

func createServerConn(rawaddr []byte, addr string) (remote *ss.Conn, err error) {
    const baseFailCnt = 20
    n := len(servers.srvCipher)
    skipped := make([]int, 0)
    for i := 0, i < n; i++ {
        //TODO try

        remote, err = connectToServer(i, rawaddr, addr)
        if err == nil { //success
            return 
        }            
    }
    return nil, err
}

    

func handleConnection(conn net.Conn) {
    if debug {
        debug.Printf("socks connect from %s\n", conn.RemoteAddr().String())
    }
    closed := false
    defer func() {
        if !closed {
            conn.Close()
        }
    }()

    var err error = nil
    if err = handShake(conn); err != nil {
        log.Println("socks handshake:", err)
        return
    }

    rawaddr, addr, err := getRequest(conn)
    if err != nil {
        log.Println("error getting request:", err)
        return
    }
    //Sending connection established message immediately to client.
    _, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x08, 0x43})
    if err != nil {
        debug.Println("send connection confiration:",err)
        return
    }

    //now socks handshake complete successfully

    remote, err := createServerConn(rawaddr, addr)
    if err != nil {
        if len(servers.srvCipher) > 1 {
            log.Println("failed connect to all avaialbe shadowsocks server")
        }
        return
    }
    defer func() {
        if !closed {
            remote.Close()
        }
    }()
        
}

func run(listenAddr string) {
    ln, err := net.Listen("tcp", listenAddr)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("starting local socks5 server at %v .. \n", listenAddr)
    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println("accept: ", err)
                continue;
        }
        go handleConnection(conn)
    }
}

func main() {
    log.SetOutput(os.Stdout)
    var printVer bool

    var configFile ,cmdLocal string
    //var cmdConfig ss.Config

    flag.StringVar(&configFile, "c", "config.json", "sepcify config file")
    flag.BoolVar(&printVer, "v", false, "print version")

    flag.Parse()

    if printVer {
        log.Printf("hello, print")
        ss.PrintVersion()
        os.Exit(0)
    }

    debug = true // has been declared
    ss.SetDebug(debug)

    exists, err := ss.IsFileExists(configFile)
    if err != nil {
        log.Printf("not found config file: ", configFile)
        print(exists)
        os.Exit(0)
    }

    config, err := ss.ParseConfig(configFile)
    if err != nil {
        print(config)
    } else {
        println("parse config success")
        //ss.UpdateConfig(config, &cmdConfig)
    }

    if config.Method == "" {
        config.Method = "aes-256-cfb"
    }

    if len(config.ServerPassword) == 0 {
        if !enoughOptions(config) {
            fmt.Fprintln(os.Stderr, "must specify server address, password and both server/local port")
            os.Exit(1)
        }
    }


    //initialization var `servers`
    parseServerConfig(config)

    run(cmdLocal + ":" + strconv.Itoa(config.LocalPort))

}





