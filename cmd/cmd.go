package cmd

import (
	"fmt"
	"net/http"

	"github.com/Kuniwak/gh-activity-summary/cli"
	"github.com/Kuniwak/gh-activity-summary/github"
	"github.com/Kuniwak/gh-activity-summary/logging"
	"github.com/Kuniwak/gh-activity-summary/printer"
)

func MainCommandByArgs(args []string, inout *cli.ProcInout) int {
	opts, err := ParseOptions(args, inout)
	if err != nil {
		fmt.Fprintf(inout.Stderr, "error: %v\n", err)
		return 1
	}

	if opts.Help {
		return 0
	}

	if err := MainCommandByOptions(opts, inout); err != nil {
		fmt.Fprintf(inout.Stderr, "error: %v\n", err)
		return 1
	}

	return 0
}

func MainCommandByOptions(opts *Options, inout *cli.ProcInout) error {
	logger := logging.NewWriterLogger(inout.Stderr, opts.Severity)

	client := github.NewClient(opts.GitHubHost, opts.GitHubToken, http.DefaultClient, logger)
	events, err := client.Events(opts.User)
	if err != nil {
		return fmt.Errorf("failed to fetch events: %w", err)
	}

	printer := printer.NewTSV(inout.Stdout)
	if err := printer(events); err != nil {
		return fmt.Errorf("failed to print events: %w", err)
	}

	return nil
}
