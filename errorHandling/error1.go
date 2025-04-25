package errorhandling

import (
	"fmt"
	"strconv"
)

/*
1. to handle error. Writing a program to convert string in to an integer
eg: "123" => valid
"abc" => invalid
*/

func ParseStringInToInt(str string) (int, error) {
	num, err := strconv.Atoi(str)

	if err != nil {
		return 0, fmt.Errorf("cannot convert string into int: %v ", err)
	}
	return num, nil
}
