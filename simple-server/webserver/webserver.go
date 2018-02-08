package webserver
import (
        "net/http"
        "errors"
        "net"
        "os"
)

type Server struct {
    mux         *http.ServeMux
    listener     net.Listener
}

func New()*Server {
    return &Server {
        mux:    http.NewServeMux(),
    }
}

func (s*Server)HandleFunc(location string, fn func(http.ResponseWriter, *http.Request)) {
    s.mux.HandleFunc(location, fn)
}

func (s *Server) Handle(location string, handler http.Handler) {  
    s.mux.Handle(location, handler)
}

func (s*Server)ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    s.mux.ServeHTTP(rw, req)
}

func (s*Server)Listen(addr string)error{
    if s.listener != nil {
        return nil
    }
    if addr == "" {
        return errors.New("listen addr is empty")
    }

    var err error
    s.listener, err = net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    return nil
}

func (s*Server)Serve() {
    if err := s.Listen(""); err != nil {
        panic("listen error")
    }
    srv := &http.Server{
        Handler : s,
    }
    
    err := srv.Serve(s.listener)
    if err != nil {
        os.Exit(1)
    }
}
