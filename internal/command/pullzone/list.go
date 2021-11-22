package pullzone

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/simplesurance/bunny-cli/internal/cli"
	"github.com/simplesurance/bunny-cli/internal/global"
)

type listCmd struct {
	cobra.Command

	page    uint32
	perPage uint32
}

func newListCmd() *listCmd {
	cmd := listCmd{
		Command: cobra.Command{
			Use:   "list",
			Short: "list pull zones",
		},
	}
	cmd.RunE = cmd.runE

	cmd.Flags().Uint32Var(
		&cmd.page,
		"page",
		1,
		"",
	)

	cmd.Flags().Uint32Var(
		&cmd.perPage,
		"per-page",
		10,
		"",
	)

	return &cmd
}

func (c *listCmd) runE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	clt := global.Get().Client

	zones, err := clt.PullZone.List(ctx, nil)
	if err != nil {
		return err
	}

	fmt.Println(cli.PrettyString(zones))

	return nil
}
