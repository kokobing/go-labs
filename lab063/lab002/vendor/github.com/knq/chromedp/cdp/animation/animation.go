// Package animation provides the Chrome Debugging Protocol
// commands, types, and events for the Animation domain.
//
// Generated by the chromedp-gen command.
package animation

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/cdp/runtime"
)

// EnableParams enables animation domain notifications.
type EnableParams struct{}

// Enable enables animation domain notifications.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Animation.enable against the provided context and
// target handler.
func (p *EnableParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandAnimationEnable, nil, nil)
}

// DisableParams disables animation domain notifications.
type DisableParams struct{}

// Disable disables animation domain notifications.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Animation.disable against the provided context and
// target handler.
func (p *DisableParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandAnimationDisable, nil, nil)
}

// GetPlaybackRateParams gets the playback rate of the document timeline.
type GetPlaybackRateParams struct{}

// GetPlaybackRate gets the playback rate of the document timeline.
func GetPlaybackRate() *GetPlaybackRateParams {
	return &GetPlaybackRateParams{}
}

// GetPlaybackRateReturns return values.
type GetPlaybackRateReturns struct {
	PlaybackRate float64 `json:"playbackRate,omitempty"` // Playback rate for animations on page.
}

// Do executes Animation.getPlaybackRate against the provided context and
// target handler.
//
// returns:
//   playbackRate - Playback rate for animations on page.
func (p *GetPlaybackRateParams) Do(ctxt context.Context, h cdp.Handler) (playbackRate float64, err error) {
	// execute
	var res GetPlaybackRateReturns
	err = h.Execute(ctxt, cdp.CommandAnimationGetPlaybackRate, nil, &res)
	if err != nil {
		return 0, err
	}

	return res.PlaybackRate, nil
}

// SetPlaybackRateParams sets the playback rate of the document timeline.
type SetPlaybackRateParams struct {
	PlaybackRate float64 `json:"playbackRate"` // Playback rate for animations on page
}

// SetPlaybackRate sets the playback rate of the document timeline.
//
// parameters:
//   playbackRate - Playback rate for animations on page
func SetPlaybackRate(playbackRate float64) *SetPlaybackRateParams {
	return &SetPlaybackRateParams{
		PlaybackRate: playbackRate,
	}
}

// Do executes Animation.setPlaybackRate against the provided context and
// target handler.
func (p *SetPlaybackRateParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandAnimationSetPlaybackRate, p, nil)
}

// GetCurrentTimeParams returns the current time of the an animation.
type GetCurrentTimeParams struct {
	ID string `json:"id"` // Id of animation.
}

// GetCurrentTime returns the current time of the an animation.
//
// parameters:
//   id - Id of animation.
func GetCurrentTime(id string) *GetCurrentTimeParams {
	return &GetCurrentTimeParams{
		ID: id,
	}
}

// GetCurrentTimeReturns return values.
type GetCurrentTimeReturns struct {
	CurrentTime float64 `json:"currentTime,omitempty"` // Current time of the page.
}

// Do executes Animation.getCurrentTime against the provided context and
// target handler.
//
// returns:
//   currentTime - Current time of the page.
func (p *GetCurrentTimeParams) Do(ctxt context.Context, h cdp.Handler) (currentTime float64, err error) {
	// execute
	var res GetCurrentTimeReturns
	err = h.Execute(ctxt, cdp.CommandAnimationGetCurrentTime, p, &res)
	if err != nil {
		return 0, err
	}

	return res.CurrentTime, nil
}

// SetPausedParams sets the paused state of a set of animations.
type SetPausedParams struct {
	Animations []string `json:"animations"` // Animations to set the pause state of.
	Paused     bool     `json:"paused"`     // Paused state to set to.
}

// SetPaused sets the paused state of a set of animations.
//
// parameters:
//   animations - Animations to set the pause state of.
//   paused - Paused state to set to.
func SetPaused(animations []string, paused bool) *SetPausedParams {
	return &SetPausedParams{
		Animations: animations,
		Paused:     paused,
	}
}

// Do executes Animation.setPaused against the provided context and
// target handler.
func (p *SetPausedParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandAnimationSetPaused, p, nil)
}

// SetTimingParams sets the timing of an animation node.
type SetTimingParams struct {
	AnimationID string  `json:"animationId"` // Animation id.
	Duration    float64 `json:"duration"`    // Duration of the animation.
	Delay       float64 `json:"delay"`       // Delay of the animation.
}

// SetTiming sets the timing of an animation node.
//
// parameters:
//   animationID - Animation id.
//   duration - Duration of the animation.
//   delay - Delay of the animation.
func SetTiming(animationID string, duration float64, delay float64) *SetTimingParams {
	return &SetTimingParams{
		AnimationID: animationID,
		Duration:    duration,
		Delay:       delay,
	}
}

// Do executes Animation.setTiming against the provided context and
// target handler.
func (p *SetTimingParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandAnimationSetTiming, p, nil)
}

// SeekAnimationsParams seek a set of animations to a particular time within
// each animation.
type SeekAnimationsParams struct {
	Animations  []string `json:"animations"`  // List of animation ids to seek.
	CurrentTime float64  `json:"currentTime"` // Set the current time of each animation.
}

// SeekAnimations seek a set of animations to a particular time within each
// animation.
//
// parameters:
//   animations - List of animation ids to seek.
//   currentTime - Set the current time of each animation.
func SeekAnimations(animations []string, currentTime float64) *SeekAnimationsParams {
	return &SeekAnimationsParams{
		Animations:  animations,
		CurrentTime: currentTime,
	}
}

// Do executes Animation.seekAnimations against the provided context and
// target handler.
func (p *SeekAnimationsParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandAnimationSeekAnimations, p, nil)
}

// ReleaseAnimationsParams releases a set of animations to no longer be
// manipulated.
type ReleaseAnimationsParams struct {
	Animations []string `json:"animations"` // List of animation ids to seek.
}

// ReleaseAnimations releases a set of animations to no longer be
// manipulated.
//
// parameters:
//   animations - List of animation ids to seek.
func ReleaseAnimations(animations []string) *ReleaseAnimationsParams {
	return &ReleaseAnimationsParams{
		Animations: animations,
	}
}

// Do executes Animation.releaseAnimations against the provided context and
// target handler.
func (p *ReleaseAnimationsParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandAnimationReleaseAnimations, p, nil)
}

// ResolveAnimationParams gets the remote object of the Animation.
type ResolveAnimationParams struct {
	AnimationID string `json:"animationId"` // Animation id.
}

// ResolveAnimation gets the remote object of the Animation.
//
// parameters:
//   animationID - Animation id.
func ResolveAnimation(animationID string) *ResolveAnimationParams {
	return &ResolveAnimationParams{
		AnimationID: animationID,
	}
}

// ResolveAnimationReturns return values.
type ResolveAnimationReturns struct {
	RemoteObject *runtime.RemoteObject `json:"remoteObject,omitempty"` // Corresponding remote object.
}

// Do executes Animation.resolveAnimation against the provided context and
// target handler.
//
// returns:
//   remoteObject - Corresponding remote object.
func (p *ResolveAnimationParams) Do(ctxt context.Context, h cdp.Handler) (remoteObject *runtime.RemoteObject, err error) {
	// execute
	var res ResolveAnimationReturns
	err = h.Execute(ctxt, cdp.CommandAnimationResolveAnimation, p, &res)
	if err != nil {
		return nil, err
	}

	return res.RemoteObject, nil
}
