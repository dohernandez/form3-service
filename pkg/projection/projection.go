package projection

import (
	"context"
	"encoding/json"

	"github.com/hellofresh/goengine"
)

// Projection is a projection that contains state data
type Projection struct {
	name            string
	messageHandlers map[string]goengine.MessageHandler
	streamName      goengine.StreamName
	state           func() interface{}
}

var _ goengine.ProjectionSaga = new(Projection)

// NewProjection creates a new Projection instance
func NewProjection(name string, streamName goengine.StreamName, state func() interface{}) *Projection {
	return &Projection{
		name:            name,
		streamName:      streamName,
		messageHandlers: map[string]goengine.MessageHandler{},
		state:           state,
	}
}

// Init initializes the state of the Query
func (p *Projection) Init(ctx context.Context) (interface{}, error) {
	return nil, nil
}

// Handlers returns the handlers of the messages sent by the notificator
func (p *Projection) Handlers() map[string]goengine.MessageHandler {
	return p.messageHandlers
}

// Name returns the name of the projection
func (p *Projection) Name() string {
	return p.name
}

// FromStream returns the stream of the projection
func (p *Projection) FromStream() goengine.StreamName {
	return p.streamName
}

// RegisterMessageHandlers register message handler
func (p *Projection) RegisterMessageHandlers(messageHandlers map[string]goengine.MessageHandler) {
	p.messageHandlers = messageHandlers
}

// DecodeState reconstitute the projection state based on the provided state data
func (p *Projection) DecodeState(data []byte) (interface{}, error) {
	if data == nil {
		return nil, nil
	}

	v := p.state()

	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}

	return v, nil
}

// EncodeState encode the given object for storage
func (p *Projection) EncodeState(obj interface{}) ([]byte, error) {
	if obj == nil {
		return nil, nil
	}

	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return data, nil
}
