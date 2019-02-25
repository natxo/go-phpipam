package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	config              = kingpin.Flag("config", "configuration file").Short('c').Required().ExistingFile()
	controller          = kingpin.Flag("controller", "phpipam api controller").Short('t').Default("sections").String()
	token               string
	url, api, user, pwd = parse_config()
)

func main() {
	kingpin.Parse()
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

// if the json fields have the same names, no need to tag them with
// `json:"tagname"`, so the struct property is enough
type Login struct {
	Code int
	//Success bool `json:"success"`, the api returns 0 as false, so no
	Data struct {
		Token   string
		Expires string
	}
	Message string
	Time    float64
}

type Sections struct {
	Code int
	Data []struct {
		ID               string
		Name             string
		Description      string
		masterSection    int
		strictMode       bool
		editDate         time.Time
		showSupernetOnly bool
		showVLAN         bool
		showVRF          bool
	}
	Message string
}

func get_token() string {
	var logindata = new(Login)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url+api+"/user/", nil)
	req.SetBasicAuth(user, pwd)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal([]byte(body), &logindata)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if logindata.Code != 200 {
		fmt.Println(logindata.Message)
		os.Exit(1)
	}
	return logindata.Data.Token
}

func get_sections(token string) {
	var f = new(Sections)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url+api+"/sections/", nil)
	req.Header.Set("token", token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal([]byte(body), &f)
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
		fmt.Printf("Name: %-24s ID: %q\n", v.Name, v.ID)
	}
}
