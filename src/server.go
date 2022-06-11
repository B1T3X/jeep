package main

import (
        "net"
        "net/http"
        "fmt"
        //"encoding/json"
        "io/ioutil"
        "log"
        "time"
        "crypto/tls"

	"github.com/gorilla/mux"
        )


// This function is needed in order to bypass Go only listening on IPv6 by default
func listenOnIPv4(portToListenTo string) (router *mux.Router, server *http.Server, listener net.Listener, err error) {
	router = mux.NewRouter()
	address := fmt.Sprintf("0.0.0.0:%v", portToListenTo)
	log.Printf("Going to listen on %v", address)
	server = &http.Server{
		Handler: router,
		Addr:    address,
		TLSConfig: &tls.Config{
			MaxVersion: tls.VersionTLS13,
			MinVersion: tls.VersionTLS12,
		},
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	listener, err = net.Listen("tcp4", address)
	return
}

func (conf *jeepConfig) runServerWithConfig() {
        r, srv, listener, err := listenOnIPv4(conf.HttpsConfig.Port)
                if err != nil {
                        panic(err)
                }
                r.HandleFunc("/issuePrinter", conf.issueCreatedHandler)
                log.Fatal(srv.ServeTLS(listener, conf.HttpsConfig.CertificatePath, conf.HttpsConfig.PrivateKeyPath))
}

// This function needs to be used on a the jeepConfig struct because we want to
// pass the PrinterPath file from the configuration to the function.
func (conf *jeepConfig) issueCreatedHandler(w http.ResponseWriter, r *http.Request) {
        defer r.Body.Close()
        b, err := ioutil.ReadAll(r.Body)
        if err != nil {
                panic(err)
        }
        issue, err := parseIssue(b)
        if err != nil {
                panic(err)
        }
        fmt.Println("Issue details:")
        fmt.Printf("Issue Key: %v\n", issue.Key)
        fmt.Printf("Issue Description: %v\n", issue.Description)
        fmt.Printf("Issue Type: %v\n", issue.Type)
        fmt.Printf("Issue Reporter: %v\n", issue.Reporter)
        err = conf.printIssue(issue)
        if err != nil {
                panic(err)
        }
}


