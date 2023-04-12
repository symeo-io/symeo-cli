package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"os/signal"
	"symeo/cli/packages/constants"
	"symeo/cli/packages/contracts"
	"symeo/cli/packages/utils"
	"symeo/cli/packages/values"
	"syscall"
)

var startCmd = &cobra.Command{
	Use:   "start [any symeo run command flags] -- [your application start command]",
	Short: "Used to inject environments variables into your application process",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("at least one argument is required after the run command, received %d", len(args))
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		contractFilePath, _ := cmd.Flags().GetString("contracts-file")
		apiUrl, _ := cmd.Flags().GetString("api-url")

		var rawValues values.Values

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

		systemEnv := utils.GetSystemEnv()
		env := values.ValuesToEnv(initializedValues)

		err = executeSingleCommandWithEnvs(args, len(env), append(systemEnv, env...))
		if err != nil {
			utils.HandleError(err, "Unable to execute your single command")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("contracts-file", "c", constants.DefaultContractPath, "Fetch values using the Symeo Token")
	startCmd.Flags().StringP("api-url", "a", constants.DefaultApiUrl, "Fetch values using the Symeo Token")
	startCmd.Flags().StringP("values-file", "f", constants.DefaultLocalValuesPath, "Fetch values using the Symeo Token")
	startCmd.Flags().StringP("api-key", "k", "", "Fetch values using the Symeo Token")
}

func executeSingleCommandWithEnvs(args []string, secretsCount int, env []string) error {
	command := args[0]
	argsForCommand := args[1:]
	color.Green("Injecting %v values into your application process", secretsCount)

	cmd := exec.Command(command, argsForCommand...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env

	return execCmd(cmd)
}

func execCmd(cmd *exec.Cmd) error {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel)

	if err := cmd.Start(); err != nil {
		return err
	}

	go func() {
		for {
			sig := <-sigChannel
			_ = cmd.Process.Signal(sig) // process all sigs
		}
	}()

	if err := cmd.Wait(); err != nil {
		_ = cmd.Process.Signal(os.Kill)
		return fmt.Errorf("failed to wait for command termination: %v", err)
	}

	waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
	os.Exit(waitStatus.ExitStatus())
	return nil
}
