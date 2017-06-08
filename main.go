package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/hbavtq/go-ntlm-auth/ntlm"
)

var (
	proxy  = "http://localhost:8080"
	target = "http://google.com"
)

func main() {
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	proxyUrl, err := url.Parse(proxy)
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	res, err := ntlm.DoNTLMRequest(myClient, req)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}
