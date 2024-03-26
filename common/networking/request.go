package networking

import (
	"log"
	"net/http"
)

func RequestGetToService(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	log.Print("req : ")
	log.Println(req)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	return resp, nil
}
