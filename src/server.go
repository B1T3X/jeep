package main

import (
        "net/http"
        "fmt"
        //"encoding/json"
        "io/ioutil"
        "log"
        )

type issue struct {
	key string
	description string
	type string
	reporter string

}

func issueCreatedHandler(w http.ResponseWriter, r *http.Request) {
        //var t test_struct
        defer r.Body.Close()
        //err := json.NewDecoder(r.Body).Decode(&t)
        b, err := ioutil.ReadAll(r.Body)
        if err != nil {
                //fmt.Println(r.Body)
                panic(err)
        }
        //fmt.Println(t.Test)
        fmt.Println(string(b))
}

func main() {
        http.HandleFunc("/test", issueCreatedHandler)
        log.Fatal(http.ListenAndServeTLS(":8443", "/home/orel/fich.is/certificate.crt", "/home/orel/fich.is/private.key", nil))
}
