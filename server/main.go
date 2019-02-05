package main

import (
    "os"
	"fmt"
	"log"
	"net/http"
    "bytes"
    "strings"
    "net/url"
    "github.com/ishiikurisu/wireworld"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        buf := new(bytes.Buffer)
        buf.ReadFrom(r.Body)
        s := buf.String()
        s = strings.Split(s, "=")[1]
        s, oops := url.PathUnescape(s)
        if oops != nil {
            http.Error(w, oops.Error(), 400)
            return
        }
        ww, _ := wireworld.LoadCsv(s)
        ww = wireworld.Update(ww)
        s = wireworld.ConvertToCsv(ww)
        fmt.Fprintf(w, "%s", s)
    })
	log.Fatal(http.ListenAndServe(":" + os.Args[1], nil))
}
