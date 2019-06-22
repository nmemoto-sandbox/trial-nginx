package main

import (
	"bytes"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Method] " + r.Method)
		fmt.Println("[URL] " + r.URL.String())
		fmt.Println("[Proto] " + r.Proto)
		fmt.Println("[ProtoMajor] " + strconv.Itoa(r.ProtoMajor))
		fmt.Println("[ProtoMiner] " + strconv.Itoa(r.ProtoMinor))
		for k, v := range r.Header {
			fmt.Print("[Header] " + k)
			fmt.Println(": " + strings.Join(v, ","))
		}
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(r.Body)
		body := bufbody.String()
		fmt.Println("[Body] " + body)
		fmt.Println("[ContentLength] " + strconv.FormatInt(r.ContentLength, 10))
		fmt.Println("[TransferEncoding] " + strings.Join(r.TransferEncoding, ","))
		fmt.Println("[Close] " + strconv.FormatBool(r.Close))
		fmt.Println("[Host] " + r.Host)
		for k, v := range r.Trailer {
			fmt.Print("[Trailer] " + k)
			fmt.Println(": " + strings.Join(v, ","))
		}
		fmt.Println("[RemoteAddr] " + r.RemoteAddr)
		fmt.Println("[RequestURI] " + r.RequestURI)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
