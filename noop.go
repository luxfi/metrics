// Copyright (C) 2020-2025, Lux Industries Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metrics

import "time"

// noopCounter is a counter that does nothing
type noopCounter struct {
	value float64
}

func (n *noopCounter) Inc()          { n.value++ }
func (n *noopCounter) Add(v float64)  { n.value += v }
func (n *noopCounter) Get() float64  { return n.value }

// noopGauge is a gauge that does nothing
type noopGauge struct {
	value float64
}

func (n *noopGauge) Set(v float64)   { n.value = v }
func (n *noopGauge) Inc()            { n.value++ }
func (n *noopGauge) Dec()            { n.value-- }
func (n *noopGauge) Add(v float64)   { n.value += v }
func (n *noopGauge) Sub(v float64)   { n.value -= v }
func (n *noopGauge) Get() float64    { return n.value }

// noopHistogram is a histogram that does nothing
type noopHistogram struct{}

func (n *noopHistogram) Observe(v float64) {}

// noopSummary is a summary that does nothing
type noopSummary struct{}

func (n *noopSummary) Observe(v float64) {}

// noopTimer is a timer that does nothing
type noopTimer struct{}

func (n *noopTimer) Start() func()            { return func() {} }
func (n *noopTimer) ObserveTime(d time.Duration) {}

// noopCounterVec is a counter vector that does nothing
type noopCounterVec struct{}

func (n *noopCounterVec) With(Labels) Counter                       { return &noopCounter{} }
func (n *noopCounterVec) WithLabelValues(...string) Counter         { return &noopCounter{} }

// noopGaugeVec is a gauge vector that does nothing
type noopGaugeVec struct{}

func (n *noopGaugeVec) With(Labels) Gauge                           { return &noopGauge{} }
func (n *noopGaugeVec) WithLabelValues(...string) Gauge             { return &noopGauge{} }

// noopHistogramVec is a histogram vector that does nothing
type noopHistogramVec struct{}

func (n *noopHistogramVec) With(Labels) Histogram                   { return &noopHistogram{} }
func (n *noopHistogramVec) WithLabelValues(...string) Histogram     { return &noopHistogram{} }

// noopSummaryVec is a summary vector that does nothing
type noopSummaryVec struct{}

func (n *noopSummaryVec) With(Labels) Summary                       { return &noopSummary{} }
func (n *noopSummaryVec) WithLabelValues(...string) Summary         { return &noopSummary{} }

// noopRegistry is a registry that does nothing
type noopRegistry struct{}

func (n *noopRegistry) Register(Collector) error                    { return nil }
func (n *noopRegistry) MustRegister(Collector)                      {}
func (n *noopRegistry) Unregister(Collector) bool                   { return true }
func (n *noopRegistry) Gather() ([]*MetricFamily, error)            { return nil, nil }

// noopMetrics is a metrics implementation that does nothing
type noopMetrics struct {
	registry Registry
}

func (n *noopMetrics) NewCounter(name, help string) Counter {
	return &noopCounter{}
}

func (n *noopMetrics) NewCounterVec(name, help string, labelNames []string) CounterVec {
	return &noopCounterVec{}
}

func (n *noopMetrics) NewGauge(name, help string) Gauge {
	return &noopGauge{}
}

func (n *noopMetrics) NewGaugeVec(name, help string, labelNames []string) GaugeVec {
	return &noopGaugeVec{}
}

func (n *noopMetrics) NewHistogram(name, help string, buckets []float64) Histogram {
	return &noopHistogram{}
}

func (n *noopMetrics) NewHistogramVec(name, help string, labelNames []string, buckets []float64) HistogramVec {
	return &noopHistogramVec{}
}

func (n *noopMetrics) NewSummary(name, help string, objectives map[float64]float64) Summary {
	return &noopSummary{}
}

func (n *noopMetrics) NewSummaryVec(name, help string, labelNames []string, objectives map[float64]float64) SummaryVec {
	return &noopSummaryVec{}
}

func (n *noopMetrics) Registry() Registry {
	return n.registry
}

// noopFactory creates noop metrics
type noopFactory struct{}

// NewNoOpFactory creates a factory that produces noop metrics
func NewNoOpFactory() Factory {
	return &noopFactory{}
}

func (f *noopFactory) New(namespace string) Metrics {
	return &noopMetrics{
		registry: &noopRegistry{},
	}
}

func (f *noopFactory) NewWithRegistry(namespace string, registry Registry) Metrics {
	return &noopMetrics{
		registry: registry,
	}
}

// NewNoOpMetrics creates a metrics instance that does nothing
func NewNoOpMetrics(namespace string) Metrics {
	return &noopMetrics{
		registry: &noopRegistry{},
	}
}