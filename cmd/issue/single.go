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

// Single returns a command that issues a certificate based upon command line arguments.
func Single() *cobra.Command {

	command := &cobra.Command{
		Use:   "single",
		Short: "Issue a certificate based upon command line arguments.",
		RunE: func(command *cobra.Command, args []string) error {

			file, err := command.Flags().GetString("config")
			if err != nil {
				return errors.Wrap(err, "error fetching [--config] flag")
			}

			if file != "" {
				viper.SetConfigFile(file)

				err = viper.ReadInConfig()
				if err != nil {
					return errors.Wrapf(err, "error reading configuration file %s", file)
				}
			}

			var configuration config.Single

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

			leaf, err := intermediate.Leaf(configuration.CommonName, configuration.AlternativeNames, configuration.Expires)
			if err != nil {
				return errors.Wrapf(err, "error creating leaf certificate")
			}

			rootCertificate, _ := root.EncodeCertificate()
			intermediateCertificate, _ := intermediate.EncodeCertificate()
			leafCertificate, _ := leaf.EncodeCertificate()
			leafKey, _ := leaf.EncodeCertificate()

			ioutil.WriteFile("root.crt", []byte(rootCertificate), 0644)
			ioutil.WriteFile("intermediate.crt", []byte(intermediateCertificate), 0644)
			ioutil.WriteFile(fmt.Sprintf("%s.crt", configuration.CommonName), []byte(leafCertificate), 0644)
			ioutil.WriteFile(fmt.Sprintf("%s.key", configuration.CommonName), []byte(leafKey), 0600)

			return nil
		},
	}

	command.Flags().String("config", "", "configuration file")

	command.Flags().DurationP("expires", "e", (time.Hour * 24 * 365), "duration for which the certificate is valid.")
	command.Flags().StringP("common", "c", "localhost", "common name of the certificate")
	command.Flags().StringSliceP("alternative", "a", []string{"localhost"}, "subject alternative name of the certificate")

	viper.BindPFlag("expires", command.Flags().Lookup("expires"))
	viper.BindPFlag("commonName", command.Flags().Lookup("common"))
	viper.BindPFlag("alternativeNames", command.Flags().Lookup("alternative"))

	return command
}
