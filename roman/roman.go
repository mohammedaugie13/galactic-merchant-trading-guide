package roman

import (
	"fmt"
	"strings"
)

type Roman struct {
	Romans map[string]int
}

func (r *Roman) ToArabic(roman string) (int, error) {
	num := strings.Split(roman, "")
	arabic := 0
	repeat := 1
	for index := 0; index < len(num); index++ {
		if index > 0 {
			if num[index] == num[index-1] {
				repeat += 1
			} else {
				repeat = 1
			}
		}
		isMaxRepeated := r.CanBeRepeat(num[index], repeat)
		if !isMaxRepeated {
			return 0, fmt.Errorf(
				"This roman number %s can't be repeated more than 3 times",
				num[index],
			)
		}
		if (index < len(num)-1) && (r.Romans[num[index]] < r.Romans[num[index+1]]) {
			canSubstract := r.CanSubstract(num[index], num[index+1])
			if !canSubstract {
				return 0, fmt.Errorf(
					"%s roman number can't be substracted with %s",
					num[index],
					num[index+1],
				)
			}
			arabic -= r.Romans[num[index]]
		} else {
			arabic += r.Romans[num[index]]
		}
	}
	return arabic, nil
}

func (r *Roman) CanSubstract(symb, nextSymb string) bool {
	val := r.Romans[symb]
	nextVal := r.Romans[nextSymb]

	return val*5 == nextVal || val*10 == nextVal
}

func (r *Roman) CanBeRepeat(letters string, repeated int) bool {
	repeatDetails := map[string]int{
		"I": 3,
		"V": 1,
		"X": 3,
		"L": 1,
		"C": 3,
		"D": 1,
		"M": 3,
	}

	if repeatDetails[letters] < repeated {
		return false
	}
	return true
}

func NewRoman() *Roman {
	symbols := make(map[string]int)
	symbols["I"] = 1
	symbols["V"] = 5
	symbols["X"] = 10
	symbols["L"] = 50
	symbols["C"] = 100
	symbols["D"] = 500
	symbols["M"] = 1000

	return &Roman{symbols}
}
