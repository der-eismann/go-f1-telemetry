package cmd

import (
	"github.com/rebuy-de/rebuy-go-sdk/v2/pkg/cmdutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	app := new(App)

	return cmdutil.New(
		"telemetry", "A tool for receiving & viewing F1 2020 telemetry",
		cmdutil.WithLogVerboseFlag(),
		cmdutil.WithVersionCommand(),
		cmdutil.WithVersionLog(logrus.DebugLevel),

		cmdutil.WithSubCommand(cmdutil.New(
			"dummy-packet", "Send dummy packets to debug",
			cmdutil.WithRun(app.DummyPacket),
		)),
		cmdutil.WithRun(app.Listen),
	)
}
