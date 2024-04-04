package tailtracer

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var (
	typeStr  = component.MustNewType("tailtracer")
	defaultInterval = 1 * time.Minute
)

func createDefaultConfig() component.Config {
	return &Config{
		Interval: string(defaultInterval),
	}
}

func createTracesReceiver(_ context.Context, params receiver.CreateSettings, baseCfg component.Config, consumer consumer.Traces) (receiver.Traces, error) {
	
	// ErrNilNextConsumer is now deprecated (https://github.com/open-telemetry/opentelemetry-collector/pull/9779)
	/* if consumer == nil {
		return nil, component.ErrNilNextConsumer
	} */

	logger := params.Logger
	tailtracerCfg := baseCfg.(*Config)

	traceRcvr := &tailtracerReceiver{
		logger: logger,
		nextConsumer: consumer,
		config: tailtracerCfg,
	}

	return traceRcvr, nil

}

// NewFactory creates a factory for tailtracer receiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithTraces(createTracesReceiver, component.StabilityLevelAlpha))
}
