package main

import (
	"github.com/munishsuri/test_api_lib/api"
)

func main() {
	router := api.GetAPIRouter()
	router.InitRouter()
	router.EnableHTTPBasicAuth("auth")
	router.SetAPIReponse("/checkHealth", []byte("check"), 200)
	router.Start(":3000")
	for {

	}
}
