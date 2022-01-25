package pullzone

import (
	"fmt"
	"os"

	"github.com/simplesurance/bunny-cli/internal/global"
	bunny "github.com/simplesurance/bunny-go"
	"github.com/spf13/cobra"
)

type addCertCmd struct {
	cobra.Command
	pullzoneID         int64
	hostname           string
	certificatePath    string
	certificateKeyPath string
}

func newAddCustomCertcmd() *addCertCmd {
	cmd := addCertCmd{
		Command: cobra.Command{
			Use:   "add-cert",
			Short: "add a custom certificate for hostname",
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
		"The hostname for which the certificate is added.",
	)
	cmd.Flags().StringVar(
		&cmd.certificatePath,
		"path",
		"",
		"The path to the certicate file.",
	)
	cmd.Flags().StringVar(
		&cmd.certificateKeyPath,
		"key",
		"",
		"The path to the certicate key file.",
	)

	if err := cmd.MarkFlagRequired("pullzone-id"); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired("hostname"); err != nil {
		panic(err)
	}

	if err := cmd.MarkFlagRequired("path"); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired("key"); err != nil {
		panic(err)
	}

	return &cmd
}

func (c *addCertCmd) runE(_ *cobra.Command, args []string) error {
	ctx := c.Context()
	clt := global.Get().Client

	cert, err := os.ReadFile(c.certificatePath)
	if err != nil {
		return fmt.Errorf("reading certificate file failed: %w", err)
	}

	key, err := os.ReadFile(c.certificateKeyPath)
	if err != nil {
		return fmt.Errorf("reading ertificate key file failed: %w", err)
	}

	err = clt.PullZone.AddCustomCertificate(ctx, c.pullzoneID, &bunny.PullZoneAddCustomCertificateOptions{
		Hostname:       c.hostname,
		Certificate:    cert,
		CertificateKey: key,
	})

	if err != nil {
		return err
	}

	fmt.Println("certificate added")

	return nil
}
