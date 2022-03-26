package agent

import (
	"fmt"
	"log"
	"net/http"
)

const (
	serverPath  = "http://localhost:8080"
	contentType = "text/plain"
)

func request(statString string) {
	resp, err := http.Post(serverPath+"/update/counter/", contentType, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = resp.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Request error with %s, http status %d", statString, resp.StatusCode)
	}
}
