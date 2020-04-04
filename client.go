/**
* Copyright (c) 2019 Tencent.
* @File   : client.go
* @Author : vicfan
* @email  : vicfan@tencent.com
* @Date   : 2020/4/4 22:19
 */

package main

import (
    "fmt"
    "net/http"
)

func main()  {
    resp, err := http.Get("http://127.0.0.1:8888")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    for {
        cc := make([]byte, 10)
        l, err := resp.Body.Read(cc)
        if err != nil {
            fmt.Println(err.Error())
            break
        }
        fmt.Println("readn = ", l)
        fmt.Println("content = ", string(cc))
    }
}