package cmd

import (
	"fmt"
	"net/http"

	"github.com/Kuniwak/gh-activity-summary/cli"
	"github.com/Kuniwak/gh-activity-summary/github"
	"github.com/Kuniwak/gh-activity-summary/logging"
	"github.com/Kuniwak/gh-activity-summary/printer"
	"github.com/Kuniwak/gh-activity-summary/summary"
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

	verbose := opts.Severity == logging.SeverityDebug
	client := github.NewClient(opts.GitHubHost, opts.GitHubToken, http.DefaultClient, verbose, logger)
	getSummaryOfMonths := summary.NewGetSummaryOfMonths(summary.NewGetSummaryOfMonth(client))
	summaries, err := getSummaryOfMonths(opts.User, opts.Since, opts.Until)
	if err != nil {
		return fmt.Errorf("failed to fetch summaries: %w", err)
	}

	printer := printer.NewTSV(inout.Stdout)
	if err := printer(summaries); err != nil {
		return fmt.Errorf("failed to print summaries: %w", err)
	}

	return nil
}
