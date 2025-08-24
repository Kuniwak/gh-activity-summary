package printer

import (
	"github.com/Kuniwak/gh-activity-summary/github"
)

type Printer func(events []github.Event) error
