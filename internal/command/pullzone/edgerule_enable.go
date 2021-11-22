package pullzone

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/simplesurance/bunny-cli/internal/global"
)

type edgeRuleEnableCmd struct {
	cobra.Command
	pullzoneID   int64
	edgeRuleGUID string
	enable       bool
	disable      bool
}

func newEdgeRuleEnableCmd() *edgeRuleEnableCmd {
	cmd := edgeRuleEnableCmd{
		Command: cobra.Command{
			Use:   "edgerule-enable",
			Short: "enable/disable an edge rule of a pull zone",
		},
	}
	cmd.RunE = cmd.runE

	cmd.Flags().Int64Var(
		&cmd.pullzoneID,
		"pullzone-id",
		-1,
		"The ID of the Pull Zone to that the Edge Rule belongs.",
	)
	cmd.Flags().StringVar(
		&cmd.edgeRuleGUID,
		"edgerule-guid",
		"",
		"The GUID of the edge-rule that is enabled/disabled.",
	)
	cmd.Flags().BoolVar(
		&cmd.enable,
		"enable",
		false,
		"enable the edge rule",
	)

	cmd.Flags().BoolVar(
		&cmd.enable,
		"disable",
		false,
		"disable the edge rule",
	)

	return &cmd
}

func (c *edgeRuleEnableCmd) runE(_ *cobra.Command, args []string) error {
	ctx := c.Context()
	clt := global.Get().Client

	if c.pullzoneID == -1 {
		return errors.New("--pullzone-id must be specified")
	}

	if c.edgeRuleGUID == "" {
		return errors.New("--edge-rule-guid must be specified")
	}

	if !c.enable && !c.disable {
		return errors.New("--enable or --disable must be specified")
	}

	if c.enable && c.disable {
		return errors.New("--enable or --disable must be specified, not both")
	}

	err := clt.PullZone.EdgeRuleService(c.pullzoneID).Enable(ctx, c.edgeRuleGUID, c.enable)
	if err != nil {
		return err
	}

	if c.enable {
		fmt.Println("edge rule enabled")
	} else {
		fmt.Println("edge rule disabled")
	}

	return nil
}
