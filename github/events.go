package github

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/Kuniwak/gh-activity-summary/httptestable"
)

type Repo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Repo      Repo      `json:"repo"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Client) Events(user string) ([]Event, error) {
	paging := NewPagingDo(httptestable.NewDoJSON[[]Event](c.do, c.logger), 100, c.logger)

	u := &url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join("/users", user, "events"),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Client#Events: %w", err)
	}

	events, err := paging(req)
	if err != nil {
		return nil, fmt.Errorf("Client#Events: %w", err)
	}

	return events, nil
}
