package overlay

// AUTOGENERATED. DO NOT EDIT.

import (
	cdp "github.com/knq/chromedp/cdp"
)

// EventNodeHighlightRequested fired when the node should be highlighted.
// This happens after call to setInspectMode.
type EventNodeHighlightRequested struct {
	NodeID cdp.NodeID `json:"nodeId,omitempty"`
}

// EventInspectNodeRequested fired when the node should be inspected. This
// happens after call to setInspectMode or when user manually inspects an
// element.
type EventInspectNodeRequested struct {
	BackendNodeID cdp.BackendNodeID `json:"backendNodeId,omitempty"` // Id of the node to inspect.
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventOverlayNodeHighlightRequested,
	cdp.EventOverlayInspectNodeRequested,
}
