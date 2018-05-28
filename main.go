package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ExploratoryEngineering/go-telenor-auth"
)

var defaultConfig gotelenorauth.ClientConfig
var configPath string

func init() {
	flag.StringVar(&defaultConfig.ClientID, "apigee-client-id", "", "Apigee client ID")
	flag.StringVar(&defaultConfig.ClientSecret, "apigee-client-secret", "", "Apigee client secret")
	flag.StringVar(&configPath, "c", "config.json", "Path to config file as json")
	flag.Parse()
}

func main() {
	config := defaultConfig

	if defaultConfig.ClientID == "" && defaultConfig.ClientSecret == "" {
		config = readConfig(configPath)
	}

	telenorAuth := gotelenorauth.NewTelenorAuth(gotelenorauth.NewDefaultConfig(config))
	// Add login/logout handling
	http.Handle("/auth/", telenorAuth.AuthHandler())

	// Add /api/*
	http.Handle(telenorAuth.Config.ProxyPath, telenorAuth.APIProxyHandler())

	// Add paths you want secured
	http.Handle("/secure/", telenorAuth.NewAuthHandlerFunc(everythingIsOKButSecret))

	// Catch all/fallback
	http.HandleFunc("/", everythingIsOK)
	// Start server
	http.ListenAndServe(":8080", nil)
}

func everythingIsOK(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Everything is OK")
	w.Write([]byte("Everything is OK"))
}

func everythingIsOKButSecret(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Everything is OK, but SECRET")
	w.Write([]byte("Everything is OK, but SECRET"))

}

func readConfig(configPath string) gotelenorauth.ClientConfig {
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Something went wrong when trying to open the JSON config file (%s).\n"+err.Error(), configPath)
		os.Exit(1)
	}

	var fileConfig gotelenorauth.ClientConfig
	err = json.Unmarshal(raw, &fileConfig)

	if err != nil {
		fmt.Println("Something went wrong when trying to load the JSON config file.\n" + err.Error())
		os.Exit(1)
	}

	return fileConfig
}
