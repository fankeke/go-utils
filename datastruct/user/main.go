package main

import("fmt")


type Person struct {
    ID      string
    Name    string
    Address string
}

func main() {
    userDB := make(map[string](*User))
    
    u := &User{}
    u.SetAge(12)
    u.SetName("bar")
    u.SetSex("male")
    u.SetPhone("222")

    userDB["u1"] = u

    v, ok := userDB["u1"]
    if !ok {
        fmt.Println(" failed to found info")
        return
    }

    fmt.Println(v.GenInfo())
}
    





