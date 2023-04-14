package cmd

import (
	"github.com/spf13/cobra"
	"symeo/cli/packages/constants"
	"symeo/cli/packages/contracts"
	"symeo/cli/packages/utils"
	"symeo/cli/packages/values"
)

var validateCmd = &cobra.Command{
	Use:   "validate [any symeo run command flags]",
	Short: "Used to inject validate environment values against contract",
	Run: func(cmd *cobra.Command, args []string) {
		contractFilePath, _ := cmd.Flags().GetString("contracts-file")
		apiUrl, _ := cmd.Flags().GetString("api-url")

		var rawValues map[string]any

		if cmd.Flags().Changed("api-key") {
			apiKey, _ := cmd.Flags().GetString("api-key")
			fetchedValues, err := values.FetchFromApi(apiUrl, apiKey)

			if err != nil {
				utils.HandleError(err)
			}

			rawValues = fetchedValues
		} else {
			valuesFilePath, _ := cmd.Flags().GetString("values-file")
			fetchedValues, err := values.FetchFromFile(valuesFilePath)

			if err != nil {
				utils.HandleError(err)
			}

			rawValues = fetchedValues
		}

		contract, err := contracts.LoadContractFile(contractFilePath)

		if err != nil {
			utils.HandleError(err)
		}

		initializedValues := values.InitializeValues(contract, rawValues)
		valuesErrors := values.CheckContractCompatibility(contract, initializedValues)

		if len(valuesErrors) > 0 {
			utils.HandleValidationErrors(valuesErrors...)
		}

		utils.HandleSuccess("Configuration values valid")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringP("contracts-file", "c", constants.DefaultContractPath, "Fetch values using the Symeo Token")
	validateCmd.Flags().StringP("api-url", "a", constants.DefaultApiUrl, "Fetch values using the Symeo Token")
	validateCmd.Flags().StringP("values-file", "f", constants.DefaultLocalValuesPath, "Fetch values using the Symeo Token")
	validateCmd.Flags().StringP("api-key", "k", "", "Fetch values using the Symeo Token")
}
