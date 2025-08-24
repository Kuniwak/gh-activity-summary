package printer

import (
	"io"
	"strconv"

	"github.com/Kuniwak/gh-activity-summary/summary"
)

var (
	tab     = []byte{'\t'}
	newline = []byte{'\n'}
	header  = []byte("month\tcommits\tissues\tpulls\treviews\trepos")
)

func NewTSV(w io.Writer) Printer {
	return func(summaries []summary.Summary) error {
		w.Write(header)
		w.Write(newline)

		for _, summary := range summaries {
			io.WriteString(w, summary.Month.Format("2006-01"))
			w.Write(tab)
			io.WriteString(w, strconv.Itoa(summary.CommitsCreated))
			w.Write(tab)
			io.WriteString(w, strconv.Itoa(summary.IssuesCreated))
			w.Write(tab)
			io.WriteString(w, strconv.Itoa(summary.PullRequestsCreated))
			w.Write(tab)
			io.WriteString(w, strconv.Itoa(summary.ReviewsCreated))
			w.Write(tab)
			io.WriteString(w, strconv.Itoa(summary.RepositoriesCreated))
			w.Write(newline)
		}
		return nil
	}
}
