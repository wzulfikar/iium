package util

import (
	"github.com/wzulfikar/iium/util/parser"
	"strconv"
	"time"
)

type current struct {
	year  int
	month int
}

// format of session: "year1/year2:sem"
// eg. 2015/2016:1 means session 2015/2016 sem 1
func SessionsOfMatric(matric_no string) []string {
	t := time.Now()
	cur := current{
		t.Year(),
		int(t.Month()),
	}

	matric := parser.ParseMatric(matric_no)

	sessions := []string{}

	for y := matric.Year; y <= cur.year; y++ {
		var ses string

		if y < cur.year {
			for sem := 1; sem <= 3; sem++ {
				if y == matric.Year && sem < matric.Sem {
					continue
				}
				ses = makeStrSession(y, sem)
				sessions = append(sessions, ses)
			}
		} else {
			switch cur.month {
			// sem 1
			case 1:
				sessions = sessions[0 : len(sessions)-2]
			// sem 2
			case 2, 3, 4, 5:
				sessions = sessions[0 : len(sessions)-1]
			// sem 1
			case 9, 10, 11, 12:
				ses = makeStrSession(y, 1)
				sessions = append(sessions, ses)
			}
		}
	}
	return sessions
}

func makeStrSession(currentYear int, sem int) string {
	return strconv.Itoa(currentYear) + "/" + strconv.Itoa(currentYear+1) + ":" + strconv.Itoa(sem)
}
