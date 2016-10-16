package parser

import (
	"strconv"
)

type Matric struct {
	Year int
	Sem  int
}

func ParseMatric(matric_no string) Matric {
	if len(matric_no) == 8 {
		matric_no = matric_no[1:8]
	}
	year, _ := strconv.Atoi(matric_no[0:2])
	sem, _ := strconv.Atoi(matric_no[2:3])

	return Matric{
		2000 + year,
		sem,
	}
}
