package simpleEmbedded

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/objectpartners/continuous-deployment-templates/testing/helper"
)

func TestAccSimpleEmbedded(t *testing.T) {
	helper.Test(t, helper.TestCase{
		Precheck: func() error {
			if os.Getenv("OPI_AWS_PROFILE") == "" {
				return errors.New("Must specify OPI_AWS_PROFILE")
			}
			return nil
		},
		Template: template,
		Inputs: map[string]string{
			"profile": os.Getenv("OPI_AWS_PROFILE"),
		},
		Check: helper.ComposeTestCheckFunc(
			helper.AssertOutputFileContentsEqual("variables.tfvars", renderedVariables),
		),
	})
}

var renderedVariables = fmt.Sprintf(`region              = "us-west-2"
profile             = "%s"
ssh_public_key_path = "~/.ssh/id_rsa.pub"
environment         = "tools"
stack               = "server"
network_cidr        = "10.0.0.0/16"
`, os.Getenv("OPI_AWS_PROFILE"))
