gh-activity-summary
===================

A command-line tool that generates activity summaries for GitHub users. It aggregates commits, issues, pull requests, reviews, and repository creations by month for a specified period and outputs in TSV format.

## Usage

```console
$ gh-activity-summary -user octocat -since 2018-01-01
month commits issues  pulls   reviews repos
2018-01 0       0       0       0       0
2018-02 0       0       0       0       0
2018-03 0       0       0       0       0
2018-04 21      0       6       1       0
2018-05 45      30      3       2       1
2018-06 39      12      5       3       0
...
```

## Setup

1. Set your GitHub token as an environment variable:
   ```console
   $ export GITHUB_TOKEN="your_github_token_here"
   ```

2. Build the tool:
   ```console
   $ go build -o gh-activity-summary .
   ```

## Options

- `-user <username>`: GitHub username (required)
- `-since <date>`: Start date in YYYY-MM-DD format (required)
- `-until <date>`: End date in YYYY-MM-DD format (defaults to current date if omitted)
- `-github-host <host>`: GitHub API host (default: api.github.com)
- `-debug`: Enable debug logging
- `-quiet`: Show only error logs

## Environment Variables

- `GITHUB_TOKEN`: GitHub API token (required)

## Examples

```console
$ # Basic usage
$ gh-activity-summary -user octocat -since 2023-01-01

$ # Specify date range
$ gh-activity-summary -user octocat -since 2023-01-01 -until 2023-12-31

$ # For GitHub Enterprise
$ gh-activity-summary -user octocat -since 2023-01-01 -github-host github.example.com

$ # Debug mode
$ gh-activity-summary -user octocat -since 2023-01-01 -debug
```

License
-------
MIT License
