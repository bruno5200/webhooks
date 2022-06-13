package mati

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	ErrEventNotSpecifiedToParse   = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod          = errors.New("invalid HTTP Method")
	ErrMissingEventKeyHeader      = errors.New("missing X-Event-Key Header")
	ErrMissingMatiSignatureHeader = errors.New("missing X-Mati-Signature Header")
	ErrEventNotFound              = errors.New("event not defined to be parsed")
	ErrParsingPayload             = errors.New("error parsing payload")
	ErrHMACVerificationFailed     = errors.New("HMAC verification failed")
)

type Event string

const (
	VerificationStartedEvent          Event = "verification_started"
	VerificationInputsCompletedEvent Event = "verification_inputs_completed"
	StepCompletedEvent               Event = "step_completed"
	VerificationUpdatedEvent         Event = "verification_updated"
	VerificationCompletedEvent       Event = "verification_completed"
)

// Option is a configuration option for the webhook
type Option func(*Webhook) error

// Options is a namespace var for configuration options
var Options = WebhookOptions{}

// WebhookOptions is a namespace for configuration option methods
type WebhookOptions struct{}

// Secret registers the GitHub secret
func (WebhookOptions) Secret(secret string) Option {
	return func(hook *Webhook) error {
		hook.secret = secret
		return nil
	}
}

// Webhook instance contains all methods needed to process events
type Webhook struct {
	secret string
}

// New creates and returns a WebHook instance denoted by the Provider type
func New(options ...Option) (*Webhook, error) {
	hook := new(Webhook)
	for _, opt := range options {
		if err := opt(hook); err != nil {
			return nil, errors.New("error applying option")
		}
	}
	return hook, nil
}

// Parse verifies and parses the events specified and returns the payload object or an error
func (hook Webhook) Parse(r *http.Request, events ...Event) (interface{}, error) {
	defer func() {
		_, _ = io.Copy(ioutil.Discard, r.Body)
		_ = r.Body.Close()
	}()

	if len(events) == 0 {
		return nil, ErrEventNotSpecifiedToParse
	}
	if r.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	event := r.Header.Get("x-signature")
	if event == "" {
		return nil, ErrMissingMatiSignatureHeader
	}

	matiEvent := Event(event)

	var found bool
	for _, evt := range events {
		if evt == matiEvent {
			found = true
			break
		}
	}
	// event not defined to be parsed
	if !found {
		return nil, ErrEventNotFound
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {
		signature := r.Header.Get("x-signature")
		if len(signature) == 0 {
			return nil, ErrMissingMatiSignatureHeader
		}
		mac := hmac.New(sha1.New, []byte(hook.secret))
		_, _ = mac.Write(payload)
		expectedMAC := hex.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {
			return nil, ErrHMACVerificationFailed
		}
	}

	switch matiEvent {
	case VerificationStartedEvent:
		var pl VerificationStartedPayload
		return pl, json.Unmarshal([]byte(payload), &pl)
	case VerificationInputsCompletedEvent:
		var pl VerificationInputsCompletedPayload
		return pl, json.Unmarshal([]byte(payload), &pl)
	case StepCompletedEvent:
		var pl StepCompletedPayload
		return pl, json.Unmarshal([]byte(payload), &pl)
	case VerificationCompletedEvent:
		var pl VerificationCompletedPayload
		return pl, json.Unmarshal([]byte(payload), &pl)
	case VerificationUpdatedEvent:
		var pl VerificationUpdatedPayload
		return pl, json.Unmarshal([]byte(payload), &pl)
	default:
		return nil, fmt.Errorf("unknown event %s", matiEvent)
	}
}
