package httptestable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Kuniwak/gh-activity-summary/logging"
)

type Do func(req *http.Request) (*http.Response, error)

func NewDo(f Do, logger logging.Logger) Do {
	return func(req *http.Request) (*http.Response, error) {
		resp, err := f(req)
		if err != nil {
			return nil, fmt.Errorf("NewDo: %s --> %w", req.URL.String(), err)
		}

		return resp, nil
	}
}

func NewDebugDo(f Do, logger logging.Logger) Do {
	return func(req *http.Request) (*http.Response, error) {
		logger.Debug(fmt.Sprintf("NewDebugDo: %s", req.URL.String()))

		resp, err := f(req)
		if err != nil {
			return nil, fmt.Errorf("NewDebugDo: %s --> %w", req.URL.String(), err)
		}

		body := &bytes.Buffer{}
		_, err = io.Copy(body, resp.Body)
		if err != nil {
			return nil, fmt.Errorf("NewDebugDo: %s --> %w", req.URL.String(), err)
		}

		logger.Debug(fmt.Sprintf("NewDebugDo: %s --> %s", req.URL.String(), resp.Status))
		logger.Debug(body.String())

		resp.Body = io.NopCloser(body)
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

		logger.Debug(fmt.Sprintf("NewDoJSON: %s --> %v", req.URL.String(), t))

		return t, nil
	}
}
