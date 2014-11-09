package main

import (
    "fmt"
    "log"
    "net/http"
    "net"
    "os"
    "github.com/gorilla/handlers"
)

func main() {
var port string
    port = "1194"
    fmt.Println("Port: ")
    fmt.Scanf("%s", &port)
    var path string = "./www"

    logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }

    http.Handle("/", handlers.CombinedLoggingHandler(logFile, http.FileServer(http.Dir(path))))
    	//http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir("."))))
    	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(path))))
    addrs, err := net.InterfaceAddrs()
    if err != nil {   panic(err)}
    for i, addr := range addrs { fmt.Printf("%d %v\n", i, addr) }
    fmt.Println("Serber running on  " + port);
    
    if err := http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
