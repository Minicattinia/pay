package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	var getURL, getHeader, getCookie string
	var getDelay int
	colorGreen := "\033[32m"
	data := url.Values{}
	getHeader = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36"
	fmt.Print("Insert Claim Dealey Minute : ")
	fmt.Scan(&getDelay)
	fmt.Print("Insert web site Url : ")
	fmt.Scan(&getURL)
	fmt.Print("Insert web Cookie : ")
	fmt.Scan(&getCookie)
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, getURL, strings.NewReader(data.Encode()))
	request.Header.Add("User-Agent", getHeader)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	request.Header.Add("Cookie", getCookie)
	for {
		resp, _ := client.Do(request)
		if resp.StatusCode == 200 {
			fmt.Println(string(colorGreen), " Success Claim")
		}
		time.Sleep(time.Duration(getDelay) * 60 * time.Second)
	}

}
