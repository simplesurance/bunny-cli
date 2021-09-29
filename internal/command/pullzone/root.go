package pullzone

import "github.com/spf13/cobra"

func Register(rootCmd *cobra.Command) {
	root := &cobra.Command{
		Use: "pullzone",
	}

	root.AddCommand(&newListCmd().Command)
	root.AddCommand(&newAddcmd().Command)
	root.AddCommand(&newDelCmd().Command)
	root.AddCommand(&newGetCmd().Command)

	root.AddCommand(&newEdgeRuleSetmd().Command)
	root.AddCommand(&newEdgeRuleEnableCmd().Command)

	rootCmd.AddCommand(root)
}
