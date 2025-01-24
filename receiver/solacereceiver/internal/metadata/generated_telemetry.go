// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/otel/metric"
	noopmetric "go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                                                      metric.Meter
	SolacereceiverDroppedEgressSpans                           metric.Int64Counter
	SolacereceiverDroppedSpanMessages                          metric.Int64Counter
	SolacereceiverFailedReconnections                          metric.Int64Counter
	SolacereceiverFatalUnmarshallingErrors                     metric.Int64Counter
	SolacereceiverNeedUpgrade                                  metric.Int64Gauge
	SolacereceiverReceivedSpanMessages                         metric.Int64Counter
	SolacereceiverReceiverFlowControlRecentRetries             metric.Int64Gauge
	SolacereceiverReceiverFlowControlStatus                    metric.Int64Gauge
	SolacereceiverReceiverFlowControlTotal                     metric.Int64Counter
	SolacereceiverReceiverFlowControlWithSingleSuccessfulRetry metric.Int64Counter
	SolacereceiverReceiverStatus                               metric.Int64Gauge
	SolacereceiverRecoverableUnmarshallingErrors               metric.Int64Counter
	SolacereceiverReportedSpans                                metric.Int64Counter
}

// TelemetryBuilderOption applies changes to default builder.
type TelemetryBuilderOption interface {
	apply(*TelemetryBuilder)
}

type telemetryBuilderOptionFunc func(mb *TelemetryBuilder)

func (tbof telemetryBuilderOptionFunc) apply(mb *TelemetryBuilder) {
	tbof(mb)
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...TelemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{}
	for _, op := range options {
		op.apply(&builder)
	}
	builder.meter = Meter(settings)
	var err, errs error
	builder.SolacereceiverDroppedEgressSpans, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_solacereceiver_dropped_egress_spans",
		metric.WithDescription("Number of dropped egress spans"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.SolacereceiverDroppedSpanMessages, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_solacereceiver_dropped_span_messages",
		metric.WithDescription("Number of dropped span messages"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.SolacereceiverFailedReconnections, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_solacereceiver_failed_reconnections",
		metric.WithDescription("Number of failed broker reconnections"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.SolacereceiverFatalUnmarshallingErrors, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_solacereceiver_fatal_unmarshalling_errors",
		metric.WithDescription("Number of fatal message unmarshalling errors"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.SolacereceiverNeedUpgrade, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Gauge(
		"otelcol_solacereceiver_need_upgrade",
		metric.WithDescription("Indicates with value 1 that receiver requires an upgrade and is not compatible with messages received from a broker"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.SolacereceiverReceivedSpanMessages, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_solacereceiver_received_span_messages",
		metric.WithDescription("Number of received span messages"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.SolacereceiverReceiverFlowControlRecentRetries, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Gauge(
		"otelcol_solacereceiver_receiver_flow_control_recent_retries",
		metric.WithDescription("Most recent/current retry count when flow controlled"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.SolacereceiverReceiverFlowControlStatus, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Gauge(
		"otelcol_solacereceiver_receiver_flow_control_status",
		metric.WithDescription("Indicates the flow control status of the receiver. 0 = not flow controlled, 1 = currently flow controlled"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.SolacereceiverReceiverFlowControlTotal, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_solacereceiver_receiver_flow_control_total",
		metric.WithDescription("Number of times the receiver instance became flow controlled"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.SolacereceiverReceiverFlowControlWithSingleSuccessfulRetry, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_solacereceiver_receiver_flow_control_with_single_successful_retry",
		metric.WithDescription("Number of times the receiver instance became flow controlled and resolved situations after the first retry"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.SolacereceiverReceiverStatus, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Gauge(
		"otelcol_solacereceiver_receiver_status",
		metric.WithDescription("Indicates the status of the receiver as an enum. 0 = starting, 1 = connecting, 2 = connected, 3 = disabled (often paired with needs_upgrade), 4 = terminating, 5 = terminated"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.SolacereceiverRecoverableUnmarshallingErrors, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_solacereceiver_recoverable_unmarshalling_errors",
		metric.WithDescription("Number of recoverable message unmarshalling errors"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.SolacereceiverReportedSpans, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64Counter(
		"otelcol_solacereceiver_reported_spans",
		metric.WithDescription("Number of reported spans"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}

func getLeveledMeter(meter metric.Meter, cfgLevel, srvLevel configtelemetry.Level) metric.Meter {
	if cfgLevel <= srvLevel {
		return meter
	}
	return noopmetric.Meter{}
}
