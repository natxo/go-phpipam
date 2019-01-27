package main

//https://blog.golang.org/json-and-go
//https://blog.alexellis.io/golang-json-api-client/

// https://markhneedham.com/blog/2017/01/21/go-vs-python-parsing-a-json-response-from-a-http-api/

import (
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"net/http"
	"os"
)

// General variables:
var token string
var url, api, user, pwd = parse_config()

// if the json fields have the same names, no need to tag them with
// `json:"tagname"`, so the struct property is enough
type Sections struct {
	Code    int
	Success bool
	Data    []struct {
		ID          string
		Name        string
		Description string
	}
	Message string
}

func main() {
	token := get_token()
	get_sections(token)
}

func parse_config() (string, string, string, string) {
	cfg, err := ini.Load("api.conf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	url := cfg.Section("").Key("url").String()
	api := cfg.Section("").Key("api").String()
	user := cfg.Section("").Key("user").String()
	pwd := cfg.Section("").Key("password").String()

	return url, api, user, pwd
}

func get_token() string {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url+api+"/user/", nil)
	req.SetBasicAuth(user, pwd)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var f interface{}
	err = json.Unmarshal(body, &f)
	m := f.(map[string]interface{})

	if m["message"] != nil {
		fmt.Println(m["message"])
		os.Exit(1)
	} else {
		switch vv := m["data"].(type) {
		case map[string]interface{}:
			token = vv["token"].(string) // need to type assert as string
			return token
		}
	}
	return token
}

func get_sections(token string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+api+"/sections/", nil)
	req.Header.Set("token", token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)

	f := Sections{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if f.Message != "" {
		fmt.Println(f.Message)
		os.Exit(1)
	}
	fmt.Println("we have", len(f.Data), "sections:")
	for _, v := range f.Data {
		fmt.Printf("Name: %-20s ID: %q\n", v.Name, v.ID)
	}
}

func init() {
	//log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
}
