package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	colorGreen := "\033[32m"
	getURL := "https://cryptoshi.club/home.php?LTC=1&DOGE=1&DGB=1&TRX=1&USDT=1&BCH=1&DASH=1&FEY=1&ZEC=1&SOL=1&redirect=1"
	data := url.Values{}
	getHeader := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36"
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, getURL, strings.NewReader(data.Encode()))
	request.Header.Add("User-Agent", getHeader)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	request.Header.Add("Cookie", "PHPSESSID=6hrr260p5ul468n1q6u1nnnkd3")
	for {
		resp, err:= client.Do(request)
		if err != nil {
		fmt.Println(string(colorGreen), "success Claim")
		}
		time.Sleep(60 * time.Second)
	}

}
