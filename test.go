/**
* Copyright (c) 2019 Tencent.
* @File   : test.go
* @Author : vicfan
* @email  : vicfan@tencent.com
* @Date   : 2020/3/29 22:04
 */

package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Transfer-Encoding", "chunked")
        w.Header().Set("Content-Type", "text/html")
        f, _ := w.(http.Flusher)
        for i := 0; i < 1000; i++ {
            _, _ = w.Write([]byte("abc<br>"))
            f.Flush()
            time.Sleep(time.Second)
        }
    })
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r.RemoteAddr)
        _, _ = w.Write([]byte("22222"))
    })
    _ = http.ListenAndServe(":8888", nil)
}