package mati

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	path = "/webhooks"
)

var hook *Webhook

func TestMain(m *testing.M) {

	// setup
	var err error
	hook, err = New(Options.Secret("123456"))
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
	// teardown
}

func newServer(handler http.HandlerFunc) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(path, handler)
	return httptest.NewServer(mux)
}

func TestBadRequests(t *testing.T) {
	assert := require.New(t)
	tests := []struct {
		name    string
		event   Event
		payload io.Reader
		headers http.Header
	}{
		{
			name:    "BadNoEventHeader",
			event:   VerificationStartedEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{},
		},
		{
			name:    "UnsubscribedEvent",
			event:   VerificationInputsCompletedEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Mati-Event": []string{"noneexistant_event"},
			},
		},
		{
			name:    "BadBody",
			event:   StepCompletedEvent,
			payload: bytes.NewBuffer([]byte("")),
			headers: http.Header{
				"X-Signature": []string{"sha1=156404ad5f721c53151147f3d3d302329f95a3ab"},
			},
		},
		{
			name:    "BadSignatureLength",
			event:   VerificationUpdatedEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Signature": []string{""},
			},
		},
		{
			name:    "BadSignatureMatch",
			event:   VerificationCompletedEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Signature": []string{"sha1=111"},
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		client := &http.Client{}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var parseError error
			server := newServer(func(w http.ResponseWriter, r *http.Request) {
				_, parseError = hook.Parse(r, tc.event)
			})
			defer server.Close()
			req, err := http.NewRequest(http.MethodPost, server.URL+path, tc.payload)
			assert.NoError(err)
			req.Header = tc.headers
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.Error(parseError)
		})
	}
}
