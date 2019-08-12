package issue

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/deciphernow/certifiable/authorities"
	"github.com/deciphernow/certifiable/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Batch returns a command that issues multiple certificates based upon a configuration file.
func Batch() *cobra.Command {

	command := &cobra.Command{
		Use:   "batch",
		Short: "Issue multiple certificates based upon a configuration file.",
		RunE: func(command *cobra.Command, args []string) error {

			file, err := command.Flags().GetString("config")
			if err != nil {
				return errors.Wrap(err, "error fetching [--config] flag")
			}

			viper.SetConfigFile(file)

			err = viper.ReadInConfig()
			if err != nil {
				return errors.Wrapf(err, "error reading configuration file %s", file)
			}

			var configuration []config.Single

			err = viper.Unmarshal(&configuration)
			if err != nil {
				return errors.Wrapf(err, "error unmarshalling configuration")
			}

			timestamp := time.Now().Unix()

			root, err := authorities.Root(fmt.Sprintf("Certifiable Root (%d)", timestamp))
			if err != nil {
				return errors.Wrapf(err, "error creating root certificate authority")
			}

			intermediate, err := root.Branch(fmt.Sprintf("Certifiable Intermediate (%d)", timestamp))
			if err != nil {
				return errors.Wrapf(err, "error creating intermediate certificate authority")
			}

			rootCertificate, _ := root.EncodeCertificate()
			intermediateCertificate, _ := intermediate.EncodeCertificate()

			ioutil.WriteFile("root.crt", []byte(rootCertificate), 0644)
			ioutil.WriteFile("intermediate.crt", []byte(intermediateCertificate), 0644)

			for _, single := range configuration {

				leaf, err := intermediate.Leaf(single.CommonName, single.AlternativeNames, single.Expires)
				if err != nil {
					return errors.Wrapf(err, "error creating leaf certificate")
				}

				leafCertificate, _ := leaf.EncodeCertificate()
				leafKey, _ := leaf.EncodeCertificate()

				ioutil.WriteFile(fmt.Sprintf("%s.crt", single.CommonName), []byte(leafCertificate), 0644)
				ioutil.WriteFile(fmt.Sprintf("%s.key", single.CommonName), []byte(leafKey), 0600)

			}

			return nil
		},
	}

	command.Flags().String("config", "", "the configuration file")

	cobra.MarkFlagRequired(command.Flags(), "config")

	return command
}
