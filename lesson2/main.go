package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	fileParams = "request.json"
)

func main() {
	type reqParams struct {
		Search string
		Sites  []string
	}

	var params reqParams
	reqParamsJSON, _ := openAndReadFile(fileParams)
	err := json.Unmarshal(reqParamsJSON, &params)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(params.Search, params.Sites)

	search(params.Search, params.Sites)

	// router := http.NewServeMux()

	// port := "8080"
	// log.Printf("Start listen on port %v", port)
	// log.Fatal(http.ListenAndServe(":"+port, router))

}

// func reqHandler(wr http.ResponseWriter, req *http.Request) {

// }

// func marshallJSON() {

// }

func openAndReadFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(file)

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
