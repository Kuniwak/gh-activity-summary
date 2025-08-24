package cmd

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/Kuniwak/gh-activity-summary/cli"
	"github.com/Kuniwak/gh-activity-summary/logging"
)

type Options struct {
	User        string
	GitHubToken string
	GitHubHost  string
	Since       time.Time
	Until       time.Time
	Severity    logging.Severity
	Help        bool
}

func ParseOptions(args []string, inout *cli.ProcInout) (*Options, error) {
	flags := flag.NewFlagSet("gh-activity-summary", flag.ContinueOnError)
	flags.SetOutput(inout.Stderr)
	flags.Usage = func() {
		fmt.Fprintln(inout.Stderr, "usage: gh-activity-summary [-github-host <host>] -user <username>")
		fmt.Fprintln(inout.Stderr, "ENVIRONMENT:")
		fmt.Fprintln(inout.Stderr, "  GITHUB_TOKEN: GitHub API token (required)")
		fmt.Fprintln(inout.Stderr)
		fmt.Fprintln(inout.Stderr, "OPTIONS:")
		flags.PrintDefaults()
	}

	options := &Options{}
	flags.StringVar(&options.User, "user", "", "GitHub username (required)")
	flags.StringVar(&options.GitHubHost, "github-host", "api.github.com", "GitHub API host")
	var since string
	flags.StringVar(&since, "since", "", "Since date (YYYY-MM-DD)")
	var until string
	flags.StringVar(&until, "until", "", "Until date (YYYY-MM-DD)")
	var debug bool
	flags.BoolVar(&debug, "debug", false, "Debug logging")
	var quiet bool
	flags.BoolVar(&quiet, "quiet", false, "Quiet logging")

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			options.Help = true
			return options, nil
		}
		return nil, err
	}

	if options.User == "" {
		return nil, errors.New("-user is required")
	}

	token := inout.Env("GITHUB_TOKEN")
	if token == "" {
		return nil, errors.New("GITHUB_TOKEN is required")
	}
	options.GitHubToken = token

	if debug {
		options.Severity = logging.SeverityDebug
	} else if quiet {
		options.Severity = logging.SeverityError
	} else {
		options.Severity = logging.SeverityInfo
	}

	if since == "" {
		return nil, errors.New("-since is required")
	}

	sinceTime, err := time.Parse("2006-01-02", since)
	if err != nil {
		return nil, fmt.Errorf("failed to parse since: %w", err)
	}
	options.Since = sinceTime

	var untilTime time.Time
	if until == "" {
		untilTime = time.Now()
	} else {
		untilTime, err = time.Parse("2006-01-02", until)
		if err != nil {
			return nil, fmt.Errorf("failed to parse until: %w", err)
		}
	}
	options.Until = untilTime

	return options, nil
}
