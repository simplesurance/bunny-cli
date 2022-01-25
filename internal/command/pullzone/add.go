package pullzone

import (
	"fmt"

	bunny "github.com/simplesurance/bunny-go"
	"github.com/spf13/cobra"

	"github.com/simplesurance/bunny-cli/internal/cli"
	"github.com/simplesurance/bunny-cli/internal/global"
)

type addCmd struct {
	cobra.Command

	name          string
	originURL     string
	storageZoneID int64
	pullZoneType  int
}

func newAddcmd() *addCmd {
	cmd := addCmd{
		Command: cobra.Command{
			Use:   "add",
			Short: "create a pull zone",
		},
	}
	cmd.RunE = cmd.runE

	cmd.Flags().StringVar(&cmd.name, "name", "", "name of the pull zone")
	cmd.Flags().StringVar(
		&cmd.originURL,
		"origin-url",
		"",
		"The origin URL of the pull zone where the files are fetched from.",
	)
	cmd.Flags().Int64Var(
		&cmd.storageZoneID,
		"id",
		0,
		"The ID of the storage zone that the pull zone is linked to",
	)
	cmd.Flags().IntVar(
		&cmd.pullZoneType,
		"type",
		1,
		"The type of the pull zone. Premium = 0, Volume = 1",
	)

	if err := cmd.MarkFlagRequired("name"); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired("origin-url"); err != nil {
		panic(err)
	}

	return &cmd
}

func (c *addCmd) runE(_ *cobra.Command, args []string) error {
	ctx := c.Context()
	clt := global.Get().Client

	opt := bunny.PullZoneAddOptions{
		Name:      c.name,
		OriginURL: c.originURL,
	}

	if c.Flags().Changed("id") {
		opt.StorageZoneID = &c.storageZoneID
	}
	if c.Flags().Changed("type") {
		opt.Type = c.pullZoneType
	}

	pz, err := clt.PullZone.Add(ctx, &opt)
	if err != nil {
		return err
	}

	fmt.Println("pull zone created")
	fmt.Println(cli.PrettyString(pz))

	return nil
}
