package main

import (
    "fmt"
    "log"
    "net/http"
    "net"
)

func main() {
var port string
    port = "1194"
    fmt.Println("Port: ")
    fmt.Scanf("%s", &port)
    var path string = "./www"
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(path))))
    addrs, err := net.InterfaceAddrs()
    if err != nil {   panic(err)}
    for i, addr := range addrs { fmt.Printf("%d %v\n", i, addr) }
    fmt.Println("Serber running on  " + port);
    
    if err := http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
