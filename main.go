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
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/wzulfikar/iium/student"
	"strings"
)

func main() {
	s := student.Cred{
		"1222665",
		"Wzulfikar031",
	}
	parse(s.GetResult())
	// parse(s.GetCam())
}

// type Result struct {
// 	course string
// 	mark   string
// }

func parse(html string) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		txt := s.Text()
		fmt.Println(txt)
	})
}
