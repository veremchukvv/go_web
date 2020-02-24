package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	fileParams = "request.json"
	port       = "8080"
)

func main() {
	reqParamsJSON, err := openAndReadFile(fileParams)
	if err != nil {
		log.Println(err)
	}
	router := http.NewServeMux()

	router.HandleFunc("/get", func(wr http.ResponseWriter, req *http.Request) {
		type reqParams struct {
			Search string
			Sites  []string
		}

		var params reqParams

		reqParamsPOST, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(reqParamsPOST, &params)
		if err != nil {
			http.Error(wr, "can't read parameters from POST request", http.StatusInternalServerError)
			return
		}

		foundingsJSON, err := search(params.Search, params.Sites)
		if err != nil {
			fmt.Println(err)
		}

		log.Println(string(foundingsJSON))
		_, _ = wr.Write(foundingsJSON)

	})

	router.HandleFunc("/post", func(wr http.ResponseWriter, req *http.Request) {
		resp, err := http.Post("http://127.0.0.1:8080/get", "application/json", bytes.NewBuffer(reqParamsJSON))
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
	})

	router.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		http.SetCookie(wr, &http.Cookie{
			Name:    "Vladimir",
			Expires: time.Now().Add(time.Minute * 10),
		})
	})

	router.HandleFunc("/cookie", func(wr http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("Vladimir")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(cookie.Name, cookie.Value)
		}
	})

	log.Printf("Start listen on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func openAndReadFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(file)
}

func search(searchString string, searchPages []string) (foundingsJSON []byte, err error) {
	var getBodies = make(map[string]string)
	var foundings []string

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
	foundingsJSON, _ = json.Marshal(foundings)

	return
}
