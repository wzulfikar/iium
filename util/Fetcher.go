/*
This code Do the following stuff :
1) login to a website called : website.com by submitting password and username on the page with url :- http://website.com/login
2) Now after login using the cookies stored by this webiste access user profile page
3) Now using same client which stored the required cookies make another post request to user profile page present at page :-
http://website.com/upser_profile_page .
4) Now get html of this whole page and print it in log as a string .

*/

package util

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Fetcher(END_POINT string, PARAMS url.Values) string {
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
