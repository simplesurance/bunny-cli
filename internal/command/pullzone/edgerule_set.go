package pullzone

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	bunny "github.com/simplesurance/bunny-go"
	"github.com/spf13/cobra"

	"github.com/simplesurance/bunny-cli/internal/global"
)

type edgeRuleSetmd struct {
	cobra.Command
}

func newEdgeRuleSetmd() *edgeRuleSetmd {
	const longHelp = `
Add or Update an Edge Rule of a Pull Zone.
The Edge Rule definition must be supplied in API JSON format on STDIN.
The JSON format are the Body Params defined in the API documentation
(https://docs.bunny.net/reference/pullzonepublic_addedgerule).`

	cmd := edgeRuleSetmd{
		Command: cobra.Command{
			Use:   "edgerule-set PULL-ZONE-ID",
			Short: "add or Update an Edge Rule of a Pull Zone.",
			Long:  strings.TrimSpace(longHelp),
			Args:  cobra.ExactArgs(1),
		},
	}
	cmd.RunE = cmd.runE

	return &cmd
}

func (c *edgeRuleSetmd) runE(_ *cobra.Command, args []string) error {
	ctx := c.Context()
	clt := global.Get().Client

	pullZoneId, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid commandline argument: %w", err)
	}

	fmt.Println("reading edge rule definition from STDIN")

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("reading from stdin failed: %w", err)
	}

	data = []byte(strings.TrimSpace(string(data)))

	if len(data) == 0 {
		return fmt.Errorf("edge rule definition must be supplied on stdin")
	}

	var updateOpts bunny.EdgeRule

	if err := json.Unmarshal(data, &updateOpts); err != nil {
		return fmt.Errorf("edge rule format is invalid: %w", err)
	}

	err = clt.PullZone.EdgeRuleService(int64(pullZoneId)).AddOrUpdate(ctx, &updateOpts)
	if err != nil {
		return err
	}

	fmt.Println("edge rule set")

	return nil
}
