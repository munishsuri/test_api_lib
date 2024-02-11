package api

type ApiRouterSuite interface {
	InitRouter()
	EnableHTTPBasicAuth(string)
	DisableHTTPBasicAuth()
	SetAPIReponse(string, []byte, int)
	Start(string)
}
