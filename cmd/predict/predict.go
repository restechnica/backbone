package predict

import (
	"github.com/urfave/cli/v2"

	"github.com/restechnica/anyreleaser/cmd/predict/version"
)

// NewCommand a command to predict the current semver version.
// Returns the CLI command.
func NewCommand(app *cli.App) *cli.Command {
	var command = "predict"
	var description = "get future information from the cli"
	var aliases = []string{"p"}

	var subcommands = []*cli.Command{
		version.NewCommand(app),
	}

	return &cli.Command{
		Aliases:         aliases,
		HideHelp:        app.HideHelp,
		HideHelpCommand: app.HideHelpCommand,
		Name:            command,
		Subcommands:     subcommands,
		Usage:           description,
	}
}
