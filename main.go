package main

import (
    "fmt"
    "os"
    "net/http"
)

func main() {
    fileServe := http.FileServer(http.Dir("./site"))
    http.Handle("/", fileServe)

    port := os.Getenv("PORT")

    err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
    if err != nil {
        fmt.Println("CANNOT RUN SERVER")
        return 
    }
   

    fmt.Println("Listening on port:", port)
}
