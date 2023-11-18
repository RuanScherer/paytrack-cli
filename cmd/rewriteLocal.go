package cmd

import (
	"os"

	"github.com/RuanScherer/paytrack-cli/application/ui"
	"github.com/spf13/cobra"
)

// rewriteLocalCmd represents the rewriteLocal command
var rewriteLocalCmd = &cobra.Command{
	Use:   "rewriteLocal",
	Short: "Sobrescreve a dependência da UI Library na node_modules do projeto frontend local com a versão local da UI Library",
	Run: func(cmd *cobra.Command, args []string) {
		frontendProject, err := cmd.Flags().GetString("frontend-project")
		if err != nil {
			os.Exit(1)
		}

		command, err := ui.NewRewriteLocalCommand(frontendProject)
		if err != nil {
			os.Exit(1)
		}

		err = command.Execute()
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	uiCmd.AddCommand(rewriteLocalCmd)

	// Here you will define your flags and configuration settings.

	rewriteLocalCmd.Flags().StringP(
		"frontend-project",
		"f",
		"",
		`Nome do projeto frontend local. Parâmetro obrigatório.		
Exemplo: patrack-cli ui rewriteLocal -f paytrack-frontend`,
	)
}
