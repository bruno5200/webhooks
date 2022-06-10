package mati

import (
	"time"
)

type VerificationStartedPayload struct {
	Resource  string    `json:"resource"`
	EventName string    `json:"eventName"`
	FlowID    string    `json:"flowId"`
	Timestamp time.Time `json:"timestamp"`
}

type VerificationInputsCompletedPayload struct {
	Resource  string    `json:"resource"`
	EventName string    `json:"eventName"`
	FlowID    string    `json:"flowId"`
	Timestamp time.Time `json:"timestamp"`
}

type StepCompletedPayload struct {
	Resource string `json:"resource"`
	Step     struct {
		Status int    `json:"status"`
		ID     string `json:"id"`
		Data   struct {
			PhoneNumber string `json:"phoneNumber"`
			CountryCode string `json:"countryCode"`
			DialingCode int    `json:"dialingCode"`
		} `json:"data"`
		Error interface{} `json:"error"`
	} `json:"step"`
	EventName string    `json:"eventName"`
	FlowID    string    `json:"flowId"`
	Timestamp time.Time `json:"timestamp"`
}

type VerificationCompletedPayload struct {
	Resource          string `json:"resource"`
	DeviceFingerprint struct {
		Ua      string `json:"ua"`
		Browser struct {
			Name    string `json:"name"`
			Version string `json:"version"`
			Major   string `json:"major"`
		} `json:"browser"`
		Engine struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"engine"`
		Os struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"os"`
		CPU struct {
			Architecture string `json:"architecture"`
		} `json:"cpu"`
		App struct {
			Platform string `json:"platform"`
			Version  string `json:"version"`
		} `json:"app"`
		IP                  string `json:"ip"`
		VpnDetectionEnabled bool   `json:"vpnDetectionEnabled"`
	} `json:"deviceFingerprint"`
	IdentityStatus string `json:"identityStatus"`
	Details        struct {
		IsDocumentExpired struct {
			Data struct {
			} `json:"data"`
		} `json:"isDocumentExpired"`
	} `json:"details"`
	MatiDashboardURL string    `json:"matiDashboardUrl"`
	Status           string    `json:"status"`
	EventName        string    `json:"eventName"`
	FlowID           string    `json:"flowId"`
	Timestamp        time.Time `json:"timestamp"`
}

type VerificationUpdatedPayload struct {
	Resource          string `json:"resource"`
	DeviceFingerprint struct {
		Ua      string `json:"ua"`
		Browser struct {
			Name    string `json:"name"`
			Version string `json:"version"`
			Major   string `json:"major"`
		} `json:"browser"`
		Engine struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"engine"`
		Os struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"os"`
		CPU struct {
			Architecture string `json:"architecture"`
		} `json:"cpu"`
		App struct {
			Platform string `json:"platform"`
			Version  string `json:"version"`
		} `json:"app"`
		IP                  string `json:"ip"`
		VpnDetectionEnabled bool   `json:"vpnDetectionEnabled"`
	} `json:"deviceFingerprint"`
	IdentityStatus   string    `json:"identityStatus"`
	MatiDashboardURL string    `json:"matiDashboardUrl"`
	Status           string    `json:"status"`
	EventName        string    `json:"eventName"`
	FlowID           string    `json:"flowId"`
	Timestamp        time.Time `json:"timestamp"`
}
