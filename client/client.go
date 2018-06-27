//Adapted from https://github.com/jcbsmpsn/golang-https-example
package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"crypto/rsa"
	//_ "github.com/nvn1729/badimportdemo/dontimportme"
)

func main() {
	log.Println("crypto/rsa.ErrVerification value:", rsa.ErrVerification)
	
	caCert, _ := ioutil.ReadFile("ca.crt")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
			},
		},
	}

	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		log.Println(err)
		return
	}

	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	
	log.Printf("%v\n", resp.Status)
	log.Printf("Server response: " + string(htmlData))
}
