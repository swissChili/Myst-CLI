package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"./colors"
	"encoding/json"
	"strings"
	"strconv"
)

type Ip struct {
	Loc string `json:"loc"`
}

type Day struct {
	Summary string `json:"summary"`
	TemperatureHigh float64 `json:"temperatureHigh"`
	TemperatureLow float64 `json:"temperatureLow"`
}

func main() {
	ip := "http://ipinfo.io"
	api := "http://localhost:8080/::long::/::lat::/daily"

	fmt.Println(colors.Set("blue", colors.Set("bold", `  ███╗   ███╗██╗   ██╗███████╗████████╗
  ████╗ ████║╚██╗ ██╔╝██╔════╝╚══██╔══╝
  ██╔████╔██║ ╚████╔╝ ███████╗   ██║   
  ██║╚██╔╝██║  ╚██╔╝  ╚════██║   ██║   
  ██║ ╚═╝ ██║   ██║   ███████║   ██║   
  ╚═╝     ╚═╝   ╚═╝   ╚══════╝   ╚═╝ `)))

	b := getJson(ip)

	var user_ip Ip
	json.Unmarshal(b, &user_ip)
	user_coords := strings.Split(user_ip.Loc, ",")
	long := user_coords[0]
	lat := user_coords[1]

	weather := getJson(format(api, "::long::", long, "::lat::", lat))

	var days []Day

	json.Unmarshal(weather, &days)
	fmt.Println(colors.Set("cyan", "  Forecast:"))
	for i := 0; i < len(days); i++ {
		fmt.Println("  -", days[i].Summary, colors.Set("red", strconv.FormatFloat(days[i].TemperatureHigh, 'f', 2, 32)), "\u00B0F")
	}
}	

func putSpacing (str string, size int) string {
	if len(str) <= size {
		spacing := size - len(str)
		spaces := strings.Repeat(" ", spacing)
		return str + spaces
	} else {
		return str
	}
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
