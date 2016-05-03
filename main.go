package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "os/user"
)

func main() {
    usr, err := user.Current()
    resp, err := http.Get("https://google.com")
    check(err)
    body, err := ioutil.ReadAll(resp.Body)
    check(err)
    fmt.Printf("Hey %s, you fetched %d bytes from google.\n", usr.Username, len(body))
}


func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
