package pullzone

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/simplesurance/bunny-cli/internal/global"
)

type delCmd struct {
	cobra.Command
}

func newDelCmd() *delCmd {
	cmd := delCmd{
		Command: cobra.Command{
			Use:   "del PULL-ZONE-ID",
			Short: "delete a pull zone",
			Args:  cobra.ExactArgs(1),
		},
	}
	cmd.RunE = cmd.runE

	return &cmd
}

func (c *delCmd) runE(_ *cobra.Command, args []string) error {
	ctx := c.Context()
	clt := global.Get().Client

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid commandline argument: %w", err)
	}

	err = clt.PullZone.Delete(ctx, int64(id))
	if err != nil {
		return err
	}

	fmt.Println("success")

	return nil
}
