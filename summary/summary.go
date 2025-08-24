package summary

import (
	"time"

	"github.com/Kuniwak/gh-activity-summary/daterange"
	"github.com/Kuniwak/gh-activity-summary/github"
)

type Summary struct {
	Month               time.Time
	CommitsCreated      int
	IssuesCreated       int
	PullRequestsCreated int
	ReviewsCreated      int
	RepositoriesCreated int
}

type GetSummaryOfMonthFunc func(user string, month time.Time) (Summary, error)

func NewGetSummaryOfMonth(client *github.Client) GetSummaryOfMonthFunc {
	return func(user string, month time.Time) (Summary, error) {
		since := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.Local)
		until := since.AddDate(0, 1, 0)

		summary, err := client.ContributionCollection(user, since, until)
		if err != nil {
			return Summary{}, err
		}
		return Summary{
			Month:               since,
			CommitsCreated:      summary.TotalCommitContributions,
			IssuesCreated:       summary.TotalIssueContributions,
			PullRequestsCreated: summary.TotalPullRequestContributions,
			ReviewsCreated:      summary.TotalPullRequestReviewContributions,
			RepositoriesCreated: summary.TotalRepositoryContributions,
		}, nil
	}
}

type GetSummaryOfMonthsFunc func(user string, since, until time.Time) ([]Summary, error)

func NewGetSummaryOfMonths(f GetSummaryOfMonthFunc) GetSummaryOfMonthsFunc {
	return func(user string, since, until time.Time) ([]Summary, error) {
		months := daterange.NewDateRange(since, until)
		summaries := make([]Summary, len(months))
		for i, month := range months {
			summary, err := f(user, month)
			if err != nil {
				return nil, err
			}
			summaries[i] = summary
		}
		return summaries, nil
	}
}
