// Package deviceorientation provides the Chrome Debugging Protocol
// commands, types, and events for the DeviceOrientation domain.
//
// Generated by the chromedp-gen command.
package deviceorientation

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
)

// SetDeviceOrientationOverrideParams overrides the Device Orientation.
type SetDeviceOrientationOverrideParams struct {
	Alpha float64 `json:"alpha"` // Mock alpha
	Beta  float64 `json:"beta"`  // Mock beta
	Gamma float64 `json:"gamma"` // Mock gamma
}

// SetDeviceOrientationOverride overrides the Device Orientation.
//
// parameters:
//   alpha - Mock alpha
//   beta - Mock beta
//   gamma - Mock gamma
func SetDeviceOrientationOverride(alpha float64, beta float64, gamma float64) *SetDeviceOrientationOverrideParams {
	return &SetDeviceOrientationOverrideParams{
		Alpha: alpha,
		Beta:  beta,
		Gamma: gamma,
	}
}

// Do executes DeviceOrientation.setDeviceOrientationOverride against the provided context and
// target handler.
func (p *SetDeviceOrientationOverrideParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDeviceOrientationSetDeviceOrientationOverride, p, nil)
}

// ClearDeviceOrientationOverrideParams clears the overridden Device
// Orientation.
type ClearDeviceOrientationOverrideParams struct{}

// ClearDeviceOrientationOverride clears the overridden Device Orientation.
func ClearDeviceOrientationOverride() *ClearDeviceOrientationOverrideParams {
	return &ClearDeviceOrientationOverrideParams{}
}

// Do executes DeviceOrientation.clearDeviceOrientationOverride against the provided context and
// target handler.
func (p *ClearDeviceOrientationOverrideParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDeviceOrientationClearDeviceOrientationOverride, nil, nil)
}
