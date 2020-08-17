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
	foundings, errs, err := search(reqString, pages)
	if err != nil {
		log.Println(err)
	}
	log.Println(foundings)
	log.Println("number of errors:", errs)
	return
}

func search(searchString string, searchPages []string) (foundings []string, errs int, err error) {
	var getBodies = make(map[string]string)
	errs = 0

	for i := range searchPages {
		resp, err := http.Get(searchPages[i])
		if err != nil {
			errs++
			log.Print(err)
			continue
		}
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errs, err
		}

		getBodies[searchPages[i]] = string(bodyBytes)
	}

	for key, value := range getBodies {
		if strings.Contains(value, searchString) {
			foundings = append(foundings, key)
		}
	}
	return
}
