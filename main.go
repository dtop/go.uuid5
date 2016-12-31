package main

import (
    "os"
    "fmt"
    "path/filepath"
    "github.com/dtop/go.uuid5/uuid5"
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: %s [namespace] [name]", filepath.Base(os.Args[0]))
    fmt.Println(".")
    fmt.Println("Defaults:")
    fmt.Println("ns:DNS: 6ba7b810-9dad-11d1-80b4-00c04fd430c8")
    fmt.Println("ns:URL: 6ba7b810-9dad-11d1-80b4-00c04fd430c8")
    fmt.Println("ns:OID: 6ba7b810-9dad-11d1-80b4-00c04fd430c8")
    os.Exit(1)
}

func main() {


    if len(os.Args) == 3 {

        namespace := os.Args[1]
        name      := os.Args[2]

        uuidgen := &uuid5.UUIDGenerator{Namespace: namespace}
        generated, err := uuidgen.GenerateWithName(name)
        if err != nil {
            println(err.Error())
            usage()
            os.Exit(1)
        }

        println(generated)
        os.Exit(0)
    }

    usage()
}
