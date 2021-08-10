package generic

import (
	"errors"
	"net/http"
)

var HttpClient *http.Client
var ErrNon200Response = errors.New("Error: GET submission link return NON-ACCEPTED response")

func httpClientInit() error {
	HttpClient = &http.Client{}
	HttpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return nil
}
