package command

import (
	"context"
	"fmt"
	"log"
	"os"

	bunny "github.com/simplesurance/bunny-go"
	"github.com/spf13/cobra"

	"github.com/simplesurance/bunny-cli/internal/command/pullzone"
	"github.com/simplesurance/bunny-cli/internal/global"
)

const shortDesc = "A commandline tool for the bunny.net CDN API"
const envVarAPIKey = "BUNNY_API_KEY"

// version contains the version number. It is set during compilation.
var version = "undefined"

type rootCmd struct {
	*cobra.Command
	printVersion bool
}

func newRootCmd() *rootCmd {
	var apiKeyParam string
	var debugEnabled bool

	longHelp := fmt.Sprintf(shortDesc+`
Supported Environment Variables:
	%s - default:
`, envVarAPIKey)

	cmd := rootCmd{
		Command: &cobra.Command{
			Use:   "admin-cli",
			Short: shortDesc,
			Long:  longHelp,
			PersistentPreRun: func(_ *cobra.Command, _ []string) {
				globalArgs := global.Get()
				if apiKeyParam != "" {
					globalArgs.APIKey = apiKeyParam
				} else {
					globalArgs.APIKey = os.Getenv(envVarAPIKey)
				}

				var opts []bunny.Option
				if debugEnabled {
					logger := log.New(os.Stderr, "http-request: ", log.LstdFlags|log.Lmsgprefix)
					opts = append(opts, bunny.WithHTTPRequestLogger(logger.Printf))
				}

				clt := bunny.NewClient(globalArgs.APIKey, opts...)

				globalArgs.Client = clt
				global.Set(globalArgs)
			},
			SilenceUsage: true,
		},
	}

	cmd.RunE = cmd.runE

	cmd.PersistentFlags().StringVar(&apiKeyParam, "api-key", "", "API key")
	cmd.PersistentFlags().BoolVar(&debugEnabled, "debug", false, "print debug message")
	cmd.Flags().BoolVar(&cmd.printVersion, "version", false, "print the version and exit")

	return &cmd
}

// Execute runs the subcommand matching the commandline args
func Execute() {
	root := newRootCmd()
	pullzone.Register(root.Command)

	// cobra prints the returned error and exits with code != 0
	_ = root.ExecuteContext(context.Background())

}

func (c *rootCmd) runE(_ *cobra.Command, _ []string) error {
	if c.printVersion {
		fmt.Printf("bunny-cli version %s\n", version)
	}

	return nil
}
