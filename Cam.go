/*
This code Do the following stuff :
1) login to a website called : website.com by submitting password and username on the page with url :- http://website.com/login
2) Now after login using the cookies stored by this webiste access user profile page
3) Now using same client which stored the required cookies make another post request to user profile page present at page :-
http://website.com/upser_profile_page .
4) Now get html of this whole page and print it in log as a string .

*/

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type StudentCred struct {
	matric string
	pwd    string
}

func fetch(END_POINT string, PARAMS url.Values) string {
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

func FetchResult(sc *StudentCred) {
	const END_POINT string = "http://prereg.iium.edu.my/resslip/viewResult.php"
	PARAMS := url.Values{
		"mat_no":   {sc.matric},
		"pin_no":   {sc.pwd},
		"sessi":    {"2015/2016"},
		"semester": {"3"},
		"login":    {"Login"},
	}
	log.Println(fetch(END_POINT, PARAMS))
}

func FetchCam(sc *StudentCred) {
	const END_POINT string = "http://albiruni.iium.edu.my/myapps/StudentOnline/camResult1.php"
	PARAMS := url.Values{
		"mat_no":   {sc.matric},
		"pin":      {sc.pwd},
		"sessi":    {"2015/2016"},
		"semester": {"3"},
		"login":    {"login"},
		"action":   {"view"},
	}
	log.Println(fetch(END_POINT, PARAMS))
}

func main() {
	sc := StudentCred{
		"1222665",
		"Wzulfikar031",
	}
	FetchCam(&sc)
	FetchResult(&sc)
}
