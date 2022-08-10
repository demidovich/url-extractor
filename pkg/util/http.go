package util

import (
	"io"
	"log"
	"net/http"
)

func HttpGetRequest(url string) (string, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0G")

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
