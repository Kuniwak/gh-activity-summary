package httptestable

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Kuniwak/gh-activity-summary/logging"
)

type Do func(req *http.Request) (*http.Response, error)

func NewDo(h *http.Client, logger logging.Logger) Do {
	return func(req *http.Request) (*http.Response, error) {
		sb := &strings.Builder{}
		sb.WriteString("NewDo: curl -sSfL")
		for k, vs := range req.Header {
			for _, v := range vs {
				sb.WriteString(" -H \"")
				sb.WriteString(k)
				sb.WriteString(": ")
				sb.WriteString(v)
				sb.WriteString("\"")
			}
		}
		sb.WriteString(" ")
		sb.WriteString(req.URL.String())

		logger.Debug(sb.String())

		resp, err := h.Do(req)
		if err != nil {
			return nil, fmt.Errorf("NewDo: %s --> %w", req.URL.String(), err)
		}

		logger.Debug(fmt.Sprintf("NewDo: %s --> %s", req.URL.String(), resp.Status))

		return resp, nil
	}
}

type DoJSON[T any] func(req *http.Request) (T, error)

func NewDoJSON[T any](f Do, logger logging.Logger) DoJSON[T] {
	return func(req *http.Request) (T, error) {
		var t T

		req.Header.Set("Content-Type", "application/json")

		resp, err := f(req)
		if err != nil {
			return t, fmt.Errorf("NewDOJSON: %s --> %w", req.URL.String(), err)
		}

		if resp.StatusCode != http.StatusOK {
			return t, fmt.Errorf("NewDOJSON: %s --> %s", req.URL.String(), resp.Status)
		}

		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
			return t, fmt.Errorf("NewDOJSON: %s --> %w", req.URL.String(), err)
		}

		logger.Debug(fmt.Sprintf("NewDoJSON: %s --> Valid JSON", req.URL.String()))

		return t, nil
	}
}
