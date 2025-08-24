package github

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestContributionCollection(t *testing.T) {
	body := `{"data":{"user":{"contributionsCollection":{"totalCommitContributions":1,"totalIssueContributions":2,"totalPullRequestContributions":3,"totalPullRequestReviewContributions":4,"totalRepositoryContributions":5}}}}`
	var actual ContributionCollectionResponse
	if err := json.Unmarshal([]byte(body), &actual); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	expected := ContributionCollectionEntry{
		TotalCommitContributions:            1,
		TotalIssueContributions:             2,
		TotalPullRequestContributions:       3,
		TotalPullRequestReviewContributions: 4,
		TotalRepositoryContributions:        5,
	}

	if !reflect.DeepEqual(actual.Data.User.ContributionsCollection, expected) {
		t.Fatal(cmp.Diff(expected, actual.Data.User.ContributionsCollection))
	}
}
