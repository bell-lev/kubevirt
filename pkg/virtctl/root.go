package virtctl

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"flag"

	"kubevirt.io/kubevirt/pkg/kubecli"
	"kubevirt.io/kubevirt/pkg/virtctl/console"
	"kubevirt.io/kubevirt/pkg/virtctl/offlinevm"
	"kubevirt.io/kubevirt/pkg/virtctl/vnc"
)

func NewVirtctlCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "virtctl",
		Short:         "virtctl controls virtual machine related operations on your kubernetes cluster.",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	//TODO: Add a ClientConfigFactory which allows substituting the KubeVirt client with a mock for unit testing
	clientConfig := kubecli.DefaultClientConfig(rootCmd.PersistentFlags())
	flag.CommandLine.Set("logtostderr", "true")
	AddGlogFlags(rootCmd.PersistentFlags())
	rootCmd.AddCommand(console.NewCommand(clientConfig), vnc.NewCommand(clientConfig), offlinevm.NewStartCommand(clientConfig), offlinevm.NewStopCommand(clientConfig))
	return rootCmd
}

func Execute() {
	if err := NewVirtctlCommand().Execute(); err != nil {
		fmt.Println(strings.TrimSpace(err.Error()))
		os.Exit(1)
	}
}
