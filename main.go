package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)

// HTTP basic authentication example in Golang using the RTC Server RESTful API
func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")

	// Customer ID
	username := "YYYYYYY"
	// Customer secret
	password := "ZZZZZZZZ"

	// Concatenate customer key and customer secret and use base64 to encode the concatenated string
	plainCredentials := username + ":" + password
	base64Credentials := base64.StdEncoding.EncodeToString([]byte(plainCredentials))

// enter the tenant Name of your SPM Tenant
	url := "https://YYYY.callidusondemand.com/api/v2/creditTypes"
	skip := "?skip=0&top=100"
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url+skip, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	// Add Authorization header
	req.Header.Add("Authorization", "Basic "+base64Credentials)
	req.Header.Add("Content-Type", "application/json")

	// Send HTTP request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))


	w.Write([]byte(string(body)))
	

}


  
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cors",Cors)
	http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}
