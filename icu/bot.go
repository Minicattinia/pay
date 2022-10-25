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
	getURL := "https://icucu.icu/doge/pkjw-6065-kdrcbifgmtod"
	data := url.Values{}
	getHeader := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36"
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, getURL, strings.NewReader(data.Encode()))
	request.Header.Add("User-Agent", getHeader)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	request.Header.Add("Cookie", "PHPSESSID=6hrr260p5ul468n1q6u1nnnkd3; io=QMAc5JTtGrbTiZXpC3VQ")
	for {
		resp, _ := client.Do(request)
		if resp.StatusCode == 200 {
			fmt.Println(string(colorGreen), "Success Claim")
		}
		time.Sleep(4 * 60 * time.Second)
	}

}
