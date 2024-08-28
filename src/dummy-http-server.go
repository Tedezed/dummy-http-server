// Dummy HTTP Server
// By Tedezed

package main

import (
    "fmt"
    "os"
    "net/http"
    "strings"
)

var host string
var port string
var config string

func ReadUserIP(req *http.Request) string {
    IPAddress := req.Header.Get("X-Real-Ip")
    if IPAddress == "" {
        IPAddress = req.Header.Get("X-Forwarded-For")
    }
    if IPAddress == "" {
        IPAddress = req.RemoteAddr
    }
    return IPAddress
}

func LogRequest(req *http.Request) {
    fmt.Println("[INFO] " + req.Method + " "  + req.URL.Path + " - IP: " + ReadUserIP(req))
}

func hello(w http.ResponseWriter, req *http.Request) {

    LogRequest(req)
    fmt.Fprintf(w, "dummy server\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    LogRequest(req)
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func getEnv(envVar string) string {
    return strings.TrimPrefix(os.Getenv(envVar), "= ")
}

func main() {
    
    http.HandleFunc("/", hello)
    http.HandleFunc("/headers", headers)
    host = getEnv("DUMMY_HOST")
    port = getEnv("DUMMY_PORT")
    config = host + ":" + port

    fmt.Println("[INFO] Start Dummy HTTP Server: " + config)
    http.ListenAndServe(config, nil)
}
