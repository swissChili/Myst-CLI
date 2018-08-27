package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"./colors"
	"encoding/json"
)

type Ip struct {
	Loc string `json:"loc"`
}

func main() {
	ip := "http://ipinfo.io"
	api := "http://localhost:8080"

	fmt.Println(colors.Set("cyan", colors.Set("bold", "Myst CLI")))

	b := getJson(ip)

	var user_ip Ip
	json.Unmarshal(b, &user_ip)

	weather := getJson(api)

	fmt.Printf("%s\n", weather)

	fmt.Printf("%s\n", user_ip.Loc)
}	

func getJson ( url string ) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get url info: %v", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read url info: %v", err)
		os.Exit(1)
	}
	return b
}