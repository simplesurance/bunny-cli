package pullzone

import (
	"fmt"

	"github.com/simplesurance/bunny-cli/internal/global"
	bunny "github.com/simplesurance/bunny-go"
	"github.com/spf13/cobra"
)

type addHostnameCmd struct {
	cobra.Command
	pullzoneID int64
	hostname   string
}

func newAddHostnameCmd() *addHostnameCmd {
	cmd := addHostnameCmd{
		Command: cobra.Command{
			Use:   "add-hostname",
			Short: "add a hostname to a pullzone",
		},
	}
	cmd.RunE = cmd.runE

	cmd.Flags().Int64Var(
		&cmd.pullzoneID,
		"pullzone-id",
		-1,
		"The ID of the Pull Zone.",
	)
	cmd.Flags().StringVar(
		&cmd.hostname,
		"hostname",
		"",
		"The hostname to add.",
	)

	if err := cmd.MarkFlagRequired("pullzone-id"); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired("hostname"); err != nil {
		panic(err)
	}

	return &cmd
}

func (c *addHostnameCmd) runE(_ *cobra.Command, args []string) error {
	ctx := c.Context()
	clt := global.Get().Client

	err := clt.PullZone.AddCustomHostname(ctx, c.pullzoneID, &bunny.AddCustomHostnameOptions{
		Hostname: &c.hostname,
	})
	if err != nil {
		return err
	}

	fmt.Println("hostname added")

	return nil
}
