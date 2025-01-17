package cmd

import (
	"github.com/robocorp/rcc/common"
	"github.com/robocorp/rcc/xviper"

	"github.com/spf13/cobra"
)

var identityCmd = &cobra.Command{
	Use:     "identity",
	Aliases: []string{"i", "id"},
	Short:   "Manage rcc instance identity related things.",
	Long:    "Manage rcc instance identity related things.",
	Run: func(cmd *cobra.Command, args []string) {
		common.Stdout("rcc instance identity is: %v\n", xviper.TrackingIdentity())
		if xviper.CanTrack() {
			common.Stdout("and anonymous health tracking is: enabled\n")
		} else {
			common.Stdout("and anonymous health tracking is: disabled\n")
		}
	},
}

func init() {
	configureCmd.AddCommand(identityCmd)
}
