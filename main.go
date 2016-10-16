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
	"github.com/gizak/termui"
	"github.com/wzulfikar/iium/student"
	"strings"
)

func main() {
	s := student.Cred{
		"1222665",
		"Wzulfikar031",
	}

	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	//termui.UseTheme("helloworld")

	strs := []string{
		"[Q] [Quit](fg-red)",
		"[R] [View Result](fg-white)",
		"[C] [View Carry Mark](fg-white)",
	}

	ls := termui.NewList()
	ls.Items = strs
	ls.ItemFgColor = termui.ColorYellow
	ls.BorderLabel = "Menu"
	ls.Height = 5
	ls.Width = 25
	ls.Y = 0

	termui.Render(ls)

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		fmt.Println("Closing the app...")
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/r", func(termui.Event) {
		fmt.Println("")
		parse(s.GetResult())
	})

	termui.Handle("/sys/kbd/c", func(termui.Event) {
		fmt.Println("")
		parse(s.GetCam())
	})

	termui.Loop()
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
