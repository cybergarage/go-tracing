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
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	"github.com/cybergarage/go-tracing/tracer"
)

type otracer struct {
	io.Closer
	serviceName
}

func New() tracer.Tracer {
	return &otracer{
		Closer:      nil,
		serviceName: "",
	}
}

// SetServiceName sets a service name.
func (ot *otracer) SetServiceName(name string) {
	ot.serviceName = name
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
	gt := opentracing.GlobalTracer()
	return &spanContext{
		span: &span{
			Tracer: gt,
			Span:   gt.StartSpan(name),
			ctx:    nil,
		},
	}
}

// Start starts a tracer.
func (ot *otracer) Start() error {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		return err
	}

	if ot.serviceName != "" {
		cfg.ServiceName = ot.serviceName
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return err
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)
	return nil
}

// Stop stops a tracer.
func (ot *otracer) Stop() error {
	if ot.Closer != nil {
		return ot.Closer.Close()
	}
	return nil
}
