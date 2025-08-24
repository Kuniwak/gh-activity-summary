package printer

import (
	"github.com/Kuniwak/gh-activity-summary/summary"
)

type Printer func(summary []summary.Summary) error
