package ipip

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ipipdotnet/ipdb-go"
)

// Global variables
var (
	DB       *ipdb.City
	Language string = "CN"
)

// Open the ipdb file
func Open(file string) (err error) {
	if DB == nil {
		if file == "" {
			file = "ipipfree.ipdb"
		}
		DB, err = ipdb.NewCity(file)
		if err != nil {
			return
		}
	}
	return
}

// FatalOpen open the ipdb file, if an error occurs the program exit
func FatalOpen(file string) {
	if err := Open(file); err != nil {
		log.Fatalf("Open ipdb file error: %v\n", err)
	}
}

// Get ip location from ipdb file
func Get(ip string) string {
	result, err := DB.Find(ip, Language)
	if err != nil {
		return ""
	}
	return trim(result)
}

// APIGet get ip location from free API
func APIGet(ip string) string {
	resp, err := http.Get("http://freeapi.ipip.net/" + ip)
	if err != nil {
		return "<Http Get Error>"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := make([]string, 5)
	json.Unmarshal(body, &result)
	return trim(result)
}

func trim(result []string) (s string) {
	s = result[1]
	if s != result[2] {
		s = strings.TrimSpace(s + " " + result[2])
	}
	if result[0] != "中国" && result[0] != s {
		s = strings.TrimSpace(result[0] + " " + s)
	}
	return
}
