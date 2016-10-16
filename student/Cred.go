/*
This code Do the following stuff :
1) login to a website called : website.com by submitting password and username on the page with url :- http://website.com/login
2) Now after login using the cookies stored by this webiste access user profile page
3) Now using same client which stored the required cookies make another post request to user profile page present at page :-
http://website.com/upser_profile_page .
4) Now get html of this whole page and print it in log as a string .

*/

package student

import (
	"github.com/wzulfikar/iium/util"
	"net/url"
)

type Cred struct {
	Matric string
	Pwd    string
}

var parser = ""

func (c *Cred) GetResult(session ...string) string {
	const END_POINT string = "http://prereg.iium.edu.my/resslip/viewResult.php"
	PARAMS := url.Values{
		"mat_no":   {c.Matric},
		"pin_no":   {c.Pwd},
		"sessi":    {"2015/2016"},
		"semester": {"3"},
		"login":    {"Login"},
	}
	return get(END_POINT, PARAMS)
}

func (c *Cred) GetLatestCam() string {
	const END_POINT string = "http://albiruni.iium.edu.my/myapps/StudentOnline/camResult1.php"
	PARAMS := url.Values{
		"mat_no": {c.Matric},
		"pin":    {c.Pwd},
	}
	return get(END_POINT, PARAMS)
}

func get(END_POINT string, PARAMS url.Values) string {
	return util.Fetcher(END_POINT, PARAMS)
}
