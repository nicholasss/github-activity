package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/nicholasss/github-activity/internal/format"
	"github.com/nicholasss/github-activity/internal/net"
)

func userCommand(ctx context.Context, cmd *cli.Command) error {
	username := cmd.StringArg("username")

	// get the events from net module
	// fmt.Println("DEBUG: looking at user:", username)
	userEvents, err := net.FetchUserEvents(username)
	if err != nil {
		fmt.Printf("We ran into an error!\nError: %q\n", err)
		return err
	}

	// print events with format module
	err = format.PrintEvents(userEvents)
	if err != nil {
		fmt.Printf("We ran into an error!\nError: %q\n", err)
		return err
	}

	return nil
}

func repoCommand(ctx context.Context, cmd *cli.Command) error {
	repository := cmd.StringArg("repository")
	// perform repo lookup

	fmt.Println("DEBUG: looking at repo:", repository)
	return nil
}

func main() {
	rootCmd := &cli.Command{
		Name:  "github-activity",
		Usage: "Lookup the activity of someone or something on GitHub",
		Commands: []*cli.Command{
			{
				Name:    "user",
				Aliases: []string{"u"},
				Usage:   "list activity of a user",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "username",
					},
				},
				Action: userCommand,
			},
			{
				Name:    "repo",
				Aliases: []string{"r"},
				Usage:   "list activity of a repo",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "repository",
					},
				},
				Action: repoCommand,
			},
		},
	}

	// actually start the root command
	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		// fmt.Println(err)
		os.Exit(1)
	}
}
