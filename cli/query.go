package main

import (
	"github.com/spf13/cobra"
	"github.com/chaosblade-io/chaosblade/transport"
	"fmt"
)

const DeviceArg = "device"

// QueryCommand defines query command
type QueryCommand struct {
	baseCommand
}

// Init attach command operators includes create instance and bind flags
func (qc *QueryCommand) Init() {
	qc.command = &cobra.Command{
		Use:     "query TARGET TYPE",
		Aliases: []string{"q"},
		Short:   "Query information for chaos experiments",
		Long:    "Query information for chaos experiments",
		RunE: func(cmd *cobra.Command, args []string) error {
			return transport.ReturnFail(transport.Code[transport.IllegalCommand],
				fmt.Sprintf("less TARGE to query"))
		},
		Example: qc.queryExample(),
	}
}

func (qc *QueryCommand) queryExample() string {
	return `query network device`
}
