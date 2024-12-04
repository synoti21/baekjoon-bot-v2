package consts

import (
	"strconv"
)

type ProbCategory int

const (
	IMPLEMENTATION ProbCategory = iota
	DATA_STRUCTURE
	DP
	GRAPH
	TRAVERSE
	STRING
	MATH
	OPTIMIZATIO9
	GEOGRAPHY
	ADVANCED
)

func ValidateProbCategory(arg string) (ProbCategory, error) {
	probCategory, err := strconv.ParseInt(arg, 0, 64)
	if err != nil {
		return -1, err
	}

	if int(probCategory) >= int(IMPLEMENTATION) && int(probCategory) <= int(ADVANCED) {
		return -1, err
	}
	return ProbCategory(probCategory), nil
}
