package github

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kuniwak/gh-activity-summary/httptestable"
	"github.com/Kuniwak/gh-activity-summary/logging"
)

func NewDo(f httptestable.Do, token string) httptestable.Do {
	return func(req *http.Request) (*http.Response, error) {
		req.Header.Set("Accept", "application/vnd.github.v3+json")
		req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return f(req)
	}
}

func NewPagingDo[T any](f httptestable.DoJSON[[]T], perPage int, logger logging.Logger) httptestable.DoJSON[[]T] {
	return func(req *http.Request) ([]T, error) {
		page := 1
		ts := make([]T, 0)
		for {
			q := req.URL.Query()
			q.Set("per_page", strconv.Itoa(perPage))
			q.Set("page", strconv.Itoa(page))
			req.URL.RawQuery = q.Encode()

			logger.Debug(fmt.Sprintf("NewPagingDo: %s", req.URL.String()))

			t, err := f(req)
			if err != nil {
				return nil, fmt.Errorf("NewPagingDo: %w", err)
			}

			ts = append(ts, t...)
			if len(t) < perPage {
				return ts, nil
			}
			page++
		}
	}
}

type Client struct {
	host   string
	do     httptestable.Do
	logger logging.Logger
}

func NewClient(host, token string, http *http.Client, logger logging.Logger) *Client {
	return &Client{
		host:   host,
		do:     NewDo(httptestable.NewDo(http, logger), token),
		logger: logger,
	}
}
