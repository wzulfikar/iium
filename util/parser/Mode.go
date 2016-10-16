package parser

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Fetch(END_POINT string, PARAMS url.Values) string {
	client := http.Client{}

	resp, err := client.PostForm(END_POINT, PARAMS)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	// return whole html of user profile data
	return string(data)
}
