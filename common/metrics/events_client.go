// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package metrics

import (
	"sync"
	"time"
)

type (
	eventsClient struct {
		provider   MetricProvider
		serviceIdx ServiceIdx

		scopes *sync.Map
	}
)

var _ Client = (*eventsClient)(nil)

func NewEventsClient(provider MetricProvider, idx ServiceIdx) *eventsClient {
	serviceTypeTagValue, _ := MetricsServiceIdxToServiceName(idx)

	return &eventsClient{
		provider:   provider.WithTags(ServiceTypeTag(serviceTypeTagValue), NamespaceTag(namespaceAllValue)),
		serviceIdx: idx,
		scopes:     new(sync.Map),
	}
}

// IncCounter increments a counter metric
func (e eventsClient) IncCounter(scope int, counter int) {
	e.Scope(scope).IncCounter(counter)
}

// AddCounter adds delta to the counter metric
func (e eventsClient) AddCounter(scope int, counter int, delta int64) {
	e.Scope(scope).AddCounter(counter, delta)
}

// StartTimer starts a timer for the given
// metric name. Time will be recorded when stopwatch is stopped.
func (e eventsClient) StartTimer(scope int, timer int) Stopwatch {
	return e.Scope(scope).StartTimer(timer)
}

// RecordTimer starts a timer for the given
// metric name
func (e eventsClient) RecordTimer(scope int, timer int, d time.Duration) {
	e.Scope(scope).RecordTimer(timer, d)
}

// RecordDistribution records and emits a distribution (wrapper on top of timer) for the given
// metric name
func (e eventsClient) RecordDistribution(scope int, timer int, d int) {
	e.Scope(scope).RecordDistribution(timer, d)
}

// UpdateGauge reports Gauge type absolute value metric
func (e eventsClient) UpdateGauge(scope int, gauge int, value float64) {
	e.Scope(scope).UpdateGauge(gauge, value)
}

// Scope returns an internal scope that can be used to add additional
// information to metrics
func (e eventsClient) Scope(scope int, tags ...Tag) Scope {
	s, ok := e.scopes.Load(scope)

	if !ok {
		if s, ok = e.scopes.Load(scope); !ok {
			s = newEventsScope(e.provider.WithTags(tags...), e.serviceIdx, scope)
			e.scopes.Store(scope, s.(Scope))
		}
	}

	return s.(Scope)
}
