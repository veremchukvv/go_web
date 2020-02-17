package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const reqString = "noindex"

var pages = []string{"https://www.google.com", "https://www.yandex.ru", "https://www.rambler.ru"}

func main() {
	search(reqString, pages)
}

func search(searchString string, searchPages []string) (foundings []string, err error) {

	// var getBodies = make([]string, len(searchPages))
	var getBodies = make(map[string]string)
	// getBodies = make([]string, len(searchPages))

	for i := range searchPages {
		resp, err := http.Get(searchPages[i])
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// log.Printf("Status: %v;\n StatusCode: %v;\n Header: %v\n", resp.Status, resp.StatusCode, resp.Header)
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		BodyString := string(bodyBytes)
		// getBodies = append(getBodies, searchPages[i], BodyString)
		getBodies[searchPages[i]] = BodyString
		// for key := range getBodies {
		// 	fmt.Println(key)
		// }

	}

	// log.Println(getBodies)
	for key, value := range getBodies {
		// log.Println(value)
		if strings.Contains(value, searchString) {
			foundings = append(foundings, key)
			// log.Println("+")
		}
	}
	log.Println(foundings)
	return

}
