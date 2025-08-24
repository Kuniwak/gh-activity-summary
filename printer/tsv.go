package printer

import (
	"io"
	"time"

	"github.com/Kuniwak/gh-activity-summary/github"
)

var (
	tab     = []byte{'\t'}
	newline = []byte{'\n'}
)

func NewTSV(w io.Writer) Printer {
	return func(events []github.Event) error {
		for _, event := range events {
			io.WriteString(w, event.ID)
			w.Write(tab)
			io.WriteString(w, event.CreatedAt.Format(time.RFC3339))
			w.Write(tab)
			io.WriteString(w, event.Repo.Name)
			w.Write(tab)
			io.WriteString(w, event.Type)
			w.Write(newline)
		}
		return nil
	}
}
