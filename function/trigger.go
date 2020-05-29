package function

import (
	"context"
	"encoding/json"

	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

var pulsarTrigger *Trigger

func init() {
	_ = trigger.Register(&Trigger{}, &Factory{})
}

type Trigger struct {
	handler trigger.Handler
}

type Factory struct {
}

var logger log.Logger

func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {

	pulsarTrigger = &Trigger{}
	return pulsarTrigger, nil
}

func (f *Factory) Metadata() *trigger.Metadata {
	return nil
}

// Metadata implements trigger.Trigger.Metadata
func (t *Trigger) Metadata() *trigger.Metadata {
	return nil
}

func Invoke(ctx context.Context, in []byte) ([]byte, error) {

	// Get the first Handler
	out := &Output{}
	out.Message = string(in)

	reply, err := pulsarTrigger.handler.Handle(ctx, out)
	if err != nil {
		return nil, err
	}
	return json.Marshal(reply)
}

func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	logger = ctx.Logger()

	// Get First handler
	t.handler = ctx.GetHandlers()[0]

	return nil
}

// Start implements util.Managed.Start
func (t *Trigger) Start() error {
	return nil
}

// Stop implements util.Managed.Stop
func (t *Trigger) Stop() error {

	return nil
}
