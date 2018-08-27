package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"./colors"
	"encoding/json"
	"strings"
)

type Ip struct {
	Loc string `json:"loc"`
}

func main() {
	ip := "http://ipinfo.io"
	api := "http://localhost:8080/::long::/::lat::/forecast"

	fmt.Println(colors.Set("cyan", colors.Set("bold", "Myst CLI")))

	b := getJson(ip)

	var user_ip Ip
	json.Unmarshal(b, &user_ip)
	user_coords := strings.Split(user_ip.Loc, ",")
	long := user_coords[0]
	lat := user_coords[1]

	weather := getJson(format(api, "::long::", long, "::lat::", lat))

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

func format(format string, args ...string) string {
    r := strings.NewReplacer(args...)
    return r.Replace(format)
}
