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
    //var port string
    port := "1194"
    fmt.Print("Port: ")
    fmt.Scanf("%s", &port)
    var wwwDir string = "./www"

    logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }

    http.Handle("/", handlers.CombinedLoggingHandler(logFile, http.FileServer(http.Dir(wwwDir))))
    addrs, err := net.InterfaceAddrs()
    if err != nil { panic(err) }
    for i, addr := range addrs { fmt.Printf("%d %v\n", i, addr) }

    serveDir, err := os.Getwd()
    fmt.Println("Serber running on port " + port);
    fmt.Println("wwwDir = " + serveDir);
    
    if err := http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
