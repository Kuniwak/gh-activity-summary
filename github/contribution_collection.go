package github

import (
	"time"

	"github.com/Kuniwak/gh-activity-summary/httptestable"
)

type ContributionCollectionResponse struct {
	Data ContributionCollectionUser `json:"data"`
}

type ContributionCollectionUser struct {
	User ContributionCollection `json:"user"`
}

type ContributionCollection struct {
	ContributionsCollection ContributionCollectionEntry `json:"contributionsCollection"`
}

type ContributionCollectionEntry struct {
	TotalCommitContributions            int `json:"totalCommitContributions"`
	TotalIssueContributions             int `json:"totalIssueContributions"`
	TotalPullRequestContributions       int `json:"totalPullRequestContributions"`
	TotalPullRequestReviewContributions int `json:"totalPullRequestReviewContributions"`
	TotalRepositoryContributions        int `json:"totalRepositoryContributions"`
}

func (c *Client) ContributionCollection(user string, from, to time.Time) (ContributionCollectionEntry, error) {
	query := `
	query($user: String!, $from: DateTime!, $to: DateTime!) {
		user(login: $user) {
			contributionsCollection(from: $from, to: $to) {
				totalCommitContributions
				totalIssueContributions
				totalPullRequestContributions
				totalPullRequestReviewContributions
				totalRepositoryContributions
			}
		}
	}`

	variables := map[string]any{
		"user": user,
		"from": from.Format(time.RFC3339),
		"to":   to.Format(time.RFC3339),
	}

	graphql := NewDoGraphQL(c.host, httptestable.NewDoJSON[ContributionCollectionResponse](c.do, c.logger))
	contributionCollection, err := graphql(GraphQLRequest{
		Query:     query,
		Variables: variables,
	})
	if err != nil {
		return ContributionCollectionEntry{}, err
	}
	return contributionCollection.Data.User.ContributionsCollection, nil
}
