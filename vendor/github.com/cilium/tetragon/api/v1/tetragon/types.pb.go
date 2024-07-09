// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

// Code generated by protoc-gen-go-tetragon. DO NOT EDIT

package tetragon

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

// IsGetEventsResponse_Event encapulates isGetEventsResponse_Event
type IsGetEventsResponse_Event = isGetEventsResponse_Event

// Event represents a Tetragon event
type Event interface {
	Encapsulate() IsGetEventsResponse_Event
}

// ProcessEvent represents a Tetragon event that has a Process field
type ProcessEvent interface {
	Event
	SetProcess(p *Process)
}

// ParentEvent represents a Tetragon event that has a Parent field
type ParentEvent interface {
	Event
	SetParent(p *Process)
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *ProcessExec) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_ProcessExec{
		ProcessExec: event,
	}
}

// SetProcess implements the ProcessEvent interface.
// Sets the Process field of an event.
func (event *ProcessExec) SetProcess(p *Process) {
	event.Process = p
}

// SetParent implements the ParentEvent interface.
// Sets the Parent field of an event.
func (event *ProcessExec) SetParent(p *Process) {
	event.Parent = p
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *ProcessExit) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_ProcessExit{
		ProcessExit: event,
	}
}

// SetProcess implements the ProcessEvent interface.
// Sets the Process field of an event.
func (event *ProcessExit) SetProcess(p *Process) {
	event.Process = p
}

// SetParent implements the ParentEvent interface.
// Sets the Parent field of an event.
func (event *ProcessExit) SetParent(p *Process) {
	event.Parent = p
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *ProcessKprobe) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_ProcessKprobe{
		ProcessKprobe: event,
	}
}

// SetProcess implements the ProcessEvent interface.
// Sets the Process field of an event.
func (event *ProcessKprobe) SetProcess(p *Process) {
	event.Process = p
}

// SetParent implements the ParentEvent interface.
// Sets the Parent field of an event.
func (event *ProcessKprobe) SetParent(p *Process) {
	event.Parent = p
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *ProcessTracepoint) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_ProcessTracepoint{
		ProcessTracepoint: event,
	}
}

// SetProcess implements the ProcessEvent interface.
// Sets the Process field of an event.
func (event *ProcessTracepoint) SetProcess(p *Process) {
	event.Process = p
}

// SetParent implements the ParentEvent interface.
// Sets the Parent field of an event.
func (event *ProcessTracepoint) SetParent(p *Process) {
	event.Parent = p
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *ProcessUprobe) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_ProcessUprobe{
		ProcessUprobe: event,
	}
}

// SetProcess implements the ProcessEvent interface.
// Sets the Process field of an event.
func (event *ProcessUprobe) SetProcess(p *Process) {
	event.Process = p
}

// SetParent implements the ParentEvent interface.
// Sets the Parent field of an event.
func (event *ProcessUprobe) SetParent(p *Process) {
	event.Parent = p
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *ProcessLsm) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_ProcessLsm{
		ProcessLsm: event,
	}
}

// SetProcess implements the ProcessEvent interface.
// Sets the Process field of an event.
func (event *ProcessLsm) SetProcess(p *Process) {
	event.Process = p
}

// SetParent implements the ParentEvent interface.
// Sets the Parent field of an event.
func (event *ProcessLsm) SetParent(p *Process) {
	event.Parent = p
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *Test) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_Test{
		Test: event,
	}
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *ProcessLoader) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_ProcessLoader{
		ProcessLoader: event,
	}
}

// SetProcess implements the ProcessEvent interface.
// Sets the Process field of an event.
func (event *ProcessLoader) SetProcess(p *Process) {
	event.Process = p
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *RateLimitInfo) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_RateLimitInfo{
		RateLimitInfo: event,
	}
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *ProcessThrottle) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_ProcessThrottle{
		ProcessThrottle: event,
	}
}

// UnwrapGetEventsResponse gets the inner event type from a GetEventsResponse
func UnwrapGetEventsResponse(response *GetEventsResponse) interface{} {
	event := response.GetEvent()
	if event == nil {
		return nil
	}
	switch ev := event.(type) {
	case *GetEventsResponse_ProcessExec:
		return ev.ProcessExec
	case *GetEventsResponse_ProcessExit:
		return ev.ProcessExit
	case *GetEventsResponse_ProcessKprobe:
		return ev.ProcessKprobe
	case *GetEventsResponse_ProcessTracepoint:
		return ev.ProcessTracepoint
	case *GetEventsResponse_ProcessUprobe:
		return ev.ProcessUprobe
	case *GetEventsResponse_ProcessLsm:
		return ev.ProcessLsm
	case *GetEventsResponse_Test:
		return ev.Test
	case *GetEventsResponse_ProcessLoader:
		return ev.ProcessLoader
	case *GetEventsResponse_RateLimitInfo:
		return ev.RateLimitInfo
	case *GetEventsResponse_ProcessThrottle:
		return ev.ProcessThrottle
	}
	return nil
}

// ResponseIsType checks whether the GetEventsResponse is of the type specified by this EventType
func (type_ EventType) ResponseIsType(response *GetEventsResponse) bool {
	if response == nil {
		return false
	}

	eventProtoNum := response.EventType()
	return eventProtoNum == type_
}

// EventIsType checks whether the Event is of the type specified by this EventType
func (type_ EventType) EventIsType(event Event) bool {
	if event == nil {
		return false
	}

	eventWrapper := event.Encapsulate()
	ger := GetEventsResponse{
		Event: eventWrapper,
	}

	return type_.ResponseIsType(&ger)
}

// EventType gets the EventType for a GetEventsResponse
func (response *GetEventsResponse) EventType() EventType {
	eventProtoNum := EventType_UNDEF

	if response == nil {
		return eventProtoNum
	}

	// Find the protobuf number for the set oneof field, if it exists.
	// Later on, we use this number to figure out if the set oneof field matches
	// our expected event type.
	rft := response.ProtoReflect()
	rft.Range(func(eventDesc protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if eventDesc.ContainingOneof() == nil || !rft.Has(eventDesc) {
			return true
		}

		eventProtoNum = EventType(eventDesc.Number())
		return false
	})

	return eventProtoNum
}
