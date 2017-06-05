package security

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// AUTOGENERATED. DO NOT EDIT.

// CertificateID an internal certificate ID value.
type CertificateID int64

// Int64 returns the CertificateID as int64 value.
func (t CertificateID) Int64() int64 {
	return int64(t)
}

// State the security level of a page or resource.
type State string

// String returns the State as string value.
func (t State) String() string {
	return string(t)
}

// State values.
const (
	StateUnknown  State = "unknown"
	StateNeutral  State = "neutral"
	StateInsecure State = "insecure"
	StateWarning  State = "warning"
	StateSecure   State = "secure"
	StateInfo     State = "info"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t State) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t State) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *State) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch State(in.String()) {
	case StateUnknown:
		*t = StateUnknown
	case StateNeutral:
		*t = StateNeutral
	case StateInsecure:
		*t = StateInsecure
	case StateWarning:
		*t = StateWarning
	case StateSecure:
		*t = StateSecure
	case StateInfo:
		*t = StateInfo

	default:
		in.AddError(errors.New("unknown State value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *State) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// StateExplanation an explanation of an factor contributing to the security
// state.
type StateExplanation struct {
	SecurityState  State  `json:"securityState,omitempty"`  // Security state representing the severity of the factor being explained.
	Summary        string `json:"summary,omitempty"`        // Short phrase describing the type of factor.
	Description    string `json:"description,omitempty"`    // Full text explanation of the factor.
	HasCertificate bool   `json:"hasCertificate,omitempty"` // True if the page has a certificate.
}

// InsecureContentStatus information about insecure content on the page.
type InsecureContentStatus struct {
	RanMixedContent                bool  `json:"ranMixedContent,omitempty"`                // True if the page was loaded over HTTPS and ran mixed (HTTP) content such as scripts.
	DisplayedMixedContent          bool  `json:"displayedMixedContent,omitempty"`          // True if the page was loaded over HTTPS and displayed mixed (HTTP) content such as images.
	ContainedMixedForm             bool  `json:"containedMixedForm,omitempty"`             // True if the page was loaded over HTTPS and contained a form targeting an insecure url.
	RanContentWithCertErrors       bool  `json:"ranContentWithCertErrors,omitempty"`       // True if the page was loaded over HTTPS without certificate errors, and ran content such as scripts that were loaded with certificate errors.
	DisplayedContentWithCertErrors bool  `json:"displayedContentWithCertErrors,omitempty"` // True if the page was loaded over HTTPS without certificate errors, and displayed content such as images that were loaded with certificate errors.
	RanInsecureContentStyle        State `json:"ranInsecureContentStyle,omitempty"`        // Security state representing a page that ran insecure content.
	DisplayedInsecureContentStyle  State `json:"displayedInsecureContentStyle,omitempty"`  // Security state representing a page that displayed insecure content.
}

// CertificateErrorAction the action to take when a certificate error occurs.
// continue will continue processing the request and cancel will cancel the
// request.
type CertificateErrorAction string

// String returns the CertificateErrorAction as string value.
func (t CertificateErrorAction) String() string {
	return string(t)
}

// CertificateErrorAction values.
const (
	CertificateErrorActionContinue CertificateErrorAction = "continue"
	CertificateErrorActionCancel   CertificateErrorAction = "cancel"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t CertificateErrorAction) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t CertificateErrorAction) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *CertificateErrorAction) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch CertificateErrorAction(in.String()) {
	case CertificateErrorActionContinue:
		*t = CertificateErrorActionContinue
	case CertificateErrorActionCancel:
		*t = CertificateErrorActionCancel

	default:
		in.AddError(errors.New("unknown CertificateErrorAction value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *CertificateErrorAction) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
