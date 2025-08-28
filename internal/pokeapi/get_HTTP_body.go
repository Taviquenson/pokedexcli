package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getBody(prefix string, name string) ([]byte, error) {

	res, err := http.Get(BaseURL + prefix + "/" + name)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode == 404 {
		return []byte(""), fmt.Errorf("error: %s name not found", prefix)
	} else if res.StatusCode > 299 {
		return []byte(""), fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}
