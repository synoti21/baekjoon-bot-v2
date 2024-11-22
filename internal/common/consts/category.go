package consts

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

func ValidateProbCategory(pc ProbCategory) error {
	//TODO: Implement
	return nil
}
