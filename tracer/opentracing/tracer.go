// Copyright (C) 2023 The go-tracing Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package opentracing

import (
	"github.com/cybergarage/go-tracing/tracer"
	ot "github.com/opentracing/opentracing-go"
)

type otracer struct {
	ot.Tracer
}

func New() tracer.Tracer {
	return &otracer{
		Tracer: ot.GlobalTracer(),
	}
}

// SetServiceName sets a service name.
func (ot *otracer) SetServiceName(_ string) {
}

// SetAgentHost sets an agent host.
func (ot *otracer) SetAgentHost(_ string) {
}

// SetAgentPort sets an agent port.
func (ot *otracer) SetAgentPort(_ int) {
}

// SetEndpoint sets an endpoint.
func (ot *otracer) SetEndpoint(_ string) {
}

// StartSpan starts a new span.
func (ot *otracer) StartSpan(name string) tracer.SpanContext {
	return &spanContext{
		span: &span{
			Tracer: ot.Tracer,
			Span:   ot.Tracer.StartSpan(name),
			ctx:    nil,
		},
	}
}

// Start starts a tracer.
func (ot *otracer) Start() error {
	return nil
}

// Stop stops a tracer.
func (ot *otracer) Stop() error {
	return nil
}
