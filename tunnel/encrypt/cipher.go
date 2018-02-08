package encrypt


import(
    "log"
    "crypto/rc4"
    "crypto/md5"
    "crypto/cipher"
)

type Cipher struct {
    enc cipher.Stream
    dec cipher.Stream
}

type chiperCreator func(key []byte) (*Cipher,error)

var cipherMap = map[string]chiperCreator {
    "rc4" : newRC4Cipher,
//    "aes256cfb" : newAES256CFBCipher,
}

func secretToKey(secret []byte, size int) []byte {
    // size mod 16 must be 0
    h := md5.New()
    buf := make([]byte, size)
    count := size / md5.Size
    // repeatly fill the key with the secret
    for i := 0; i < count; i++ {
        h.Write(secret)
        copy(buf[md5.Size*i:md5.Size*(i+1)-1], h.Sum(nil))
    }
    return buf
}

func newRC4Cipher(secret []byte)(*Cipher, error) {
    ec, err := rc4.NewCipher(secretToKey(secret, 16))
    if err != nil {
        return nil, err
    }
    dc := *ec
    return &Cipher{ec, &dc}, nil
}



func NewCipher(cryptoMethod string, secret []byte) *Cipher {
    cc := cipherMap[cryptoMethod]
    if cc == nil {
        log.Fatalf("unsuported crypot method:%s", cryptoMethod)
    }

    c, err := cc(secret)
    if err !=nil {
        log.Fatal(err)
    }

    return c
}



func (c *Cipher) Encrypt(dst, src []byte) {
    c.enc.XORKeyStream(dst, src)
}

func (c *Cipher)decrypt(dst,src []byte) {
    c.dec.XORKeyStream(dst, src)
}
