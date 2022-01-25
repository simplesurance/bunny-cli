package pullzone

import (
	"fmt"

	"github.com/simplesurance/bunny-cli/internal/global"
	bunny "github.com/simplesurance/bunny-go"
	"github.com/spf13/cobra"
)

type removeCertCmd struct {
	cobra.Command
	pullzoneID int64
	hostname   string
}

func newRemoveCertCmd() *removeCertCmd {
	cmd := removeCertCmd{
		Command: cobra.Command{
			Use:   "remove-cert",
			Short: "remove a hostname ssl certificate",
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
		"The hostname to remove.",
	)

	if err := cmd.MarkFlagRequired("pullzone-id"); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired("hostname"); err != nil {
		panic(err)
	}

	return &cmd
}

func (c *removeCertCmd) runE(_ *cobra.Command, args []string) error {
	ctx := c.Context()
	clt := global.Get().Client

	err := clt.PullZone.RemoveCertificate(ctx, c.pullzoneID, &bunny.RemoveCertificateOptions{
		Hostname: &c.hostname,
	})
	if err != nil {
		return err
	}

	fmt.Println("certificate removed")

	return nil
}
