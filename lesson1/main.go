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

	var getBodies = make(map[string]string)

	for i := range searchPages {
		resp, err := http.Get(searchPages[i])
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		BodyString := string(bodyBytes)
		getBodies[searchPages[i]] = BodyString
	}

	for key, value := range getBodies {
		if strings.Contains(value, searchString) {
			foundings = append(foundings, key)
		}
	}
	log.Println(foundings)
	return

}
