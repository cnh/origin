package config

import (
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubectl/cmd/config"
	"github.com/spf13/cobra"
)

func NewCmdConfig(parentName, name string) *cobra.Command {
	cmd := config.NewCmdConfig(os.Stdout)
	cmd.Short = "Change configuration files for the client"
	cmd.Long = fmt.Sprintf(`Manages the OpenShift config files using subcommands like:

%[1]s %[2]s use-context my-context
%[1]s %[2]s set preferences.some true

Reference: https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/kubeconfig-file.md`, parentName, name)

	return cmd
}
