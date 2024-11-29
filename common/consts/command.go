package consts

type SlashCommand string

const (
	SCRegister                SlashCommand = SlashCommand("/register")
	SCWithdraw                SlashCommand = SlashCommand("/quit")
	SCRecommandSingleProblem  SlashCommand = SlashCommand("/prob")
	SCRecommandByCategory     SlashCommand = SlashCommand("/category")
	SCRecommandSimilarProblem SlashCommand = SlashCommand("/similarid")
	SCScheduleDailyProblem    SlashCommand = SlashCommand("/daily")
	SCUnscheduleDailyProblem  SlashCommand = SlashCommand("/deactivate")
	SCShowCategoryList        SlashCommand = SlashCommand("/categorylist")
	SCShowHelpGuide           SlashCommand = SlashCommand("/help")
)
