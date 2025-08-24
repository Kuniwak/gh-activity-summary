package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/Kuniwak/gh-activity-summary/httptestable"
)

type GraphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables"`
}

type DoGraphQL[T any] func(req GraphQLRequest) (T, error)

func NewDoGraphQL[T any](host string, do httptestable.DoJSON[T]) DoGraphQL[T] {
	return func(req GraphQLRequest) (T, error) {
		var t T
		body, err := json.Marshal(req)
		if err != nil {
			return t, fmt.Errorf("NewDoGraphQL: %w", err)
		}

		u := &url.URL{
			Scheme: "https",
			Host:   host,
			Path:   path.Join("/graphql"),
		}

		httpReq, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
		if err != nil {
			return t, fmt.Errorf("NewDoGraphQL: %w", err)
		}

		t, err = do(httpReq)
		if err != nil {
			return t, fmt.Errorf("NewDoGraphQL: %w", err)
		}

		return t, nil
	}
}
