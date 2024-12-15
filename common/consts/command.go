package consts

type SlashCommand string

const (
	SCRegister                SlashCommand = SlashCommand("/register")
	SCWithdraw                SlashCommand = SlashCommand("/quit")
	SCRecommendSingleProblem  SlashCommand = SlashCommand("/prob")
	SCRecommendByCategory     SlashCommand = SlashCommand("/category")
	SCRecommendSimilarProblem SlashCommand = SlashCommand("/similarid")
	SCScheduleDailyProblem    SlashCommand = SlashCommand("/daily")
	SCUnscheduleDailyProblem  SlashCommand = SlashCommand("/deactivate")
	SCShowCategoryList        SlashCommand = SlashCommand("/categorylist")
	SCShowHelpGuide           SlashCommand = SlashCommand("/help")
)

var commandMap map[string]SlashCommand = map[string]SlashCommand{
	"/register":     SCRegister,
	"/quit":         SCWithdraw,
	"/prob":         SCRecommendSingleProblem,
	"/category":     SCRecommendByCategory,
	"/simliarid":    SCRecommendSimilarProblem,
	"/daily":        SCScheduleDailyProblem,
	"/deactivate":   SCUnscheduleDailyProblem,
	"/categorylist": SCShowCategoryList,
	"/help":         SCShowHelpGuide,
}

func ValidateSlashCommand(command string) (SlashCommand, bool) {
	if sc, ok := commandMap[command]; !ok {
		return "", false
	} else {
		return sc, true
	}
}
