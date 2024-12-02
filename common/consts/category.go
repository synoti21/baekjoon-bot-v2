package consts

import "github.com/synoti21/baekjoon-slack-bot/common/errors"

type ProbCategory int

const (
	IMPLEMENTATION ProbCategory = iota
	DATA_STRUCTURE
	DP
	GRAPH
	TRAVERSE
	STRING
	MATH
	OPTIMIZATION
	GEOGRAPHY
	ADVANCED
)

func ValidateProbCategory(probCategory int) error {
	if probCategory >= int(IMPLEMENTATION) && probCategory <= int(ADVANCED) {
		return errors.NewInternalServerError("Invalid Category Type")
	}
	return nil
}
