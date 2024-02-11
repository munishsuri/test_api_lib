package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetAPIRouter() ApiRouterSuite {
	return apiRouterSuiteImpl()
}

func apiRouterSuiteImpl() ApiRouterSuite {
	return &apiRouter{
		Router: mux.NewRouter(),
	}
}

type apiRouter struct {
	Router        *mux.Router
	httpBasicAuth string
}

func (a *apiRouter) InitRouter() {
	a.Router = mux.NewRouter().StrictSlash(true)
}

func (a *apiRouter) EnableHTTPBasicAuth(auth string) {
	a.httpBasicAuth = auth
}

func (a *apiRouter) DisableHTTPBasicAuth() {
	a.httpBasicAuth = ""
}

func (a *apiRouter) Start(port string) {
	http.ListenAndServe(port, a.Router)
}

func (a *apiRouter) SetAPIReponse(endpoint string, response []byte, responseCode int) {
	a.Router.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if a.httpBasicAuth != "" {
			if r.Header.Get("Authorization") == a.httpBasicAuth {
				w.Write(response)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			w.Write(response)
		}
	})
}
