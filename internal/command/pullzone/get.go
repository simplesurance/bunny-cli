package pullzone

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/simplesurance/bunny-cli/internal/cli"
	"github.com/simplesurance/bunny-cli/internal/global"
)

type getCmd struct {
	cobra.Command
}

func newGetCmd() *getCmd {
	cmd := getCmd{
		Command: cobra.Command{
			Use:   "get ID",
			Short: "retrieve a pull zone",
			Args:  cobra.ExactArgs(1),
		},
	}
	cmd.RunE = cmd.runE

	return &cmd
}

func (c *getCmd) runE(_ *cobra.Command, args []string) error {
	ctx := c.Context()
	clt := global.Get().Client

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("parsing argument failed: %w", err)
	}

	pz, err := clt.PullZone.Get(ctx, int64(id))
	if err != nil {
		return err
	}

	fmt.Println("pull zone created")
	fmt.Println(cli.PrettyString(pz))

	return nil
}
