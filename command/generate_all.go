package command

func GenerateAll(arguments []string) {
	GenerateUtils(arguments)
	GenerateEnums(arguments)
	GenerateConfigs(arguments)
	GenerateValidators(arguments)
	GenerateErrors(arguments)
	GenerateRoutes(arguments)

}
