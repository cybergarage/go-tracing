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
	"fmt"
	"io"
	"net/url"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	"github.com/cybergarage/go-tracing/tracer"
)

type otracer struct {
	io.Closer
	serviceName string
	endpoint    string
	host        string
	port        int
}

func New() tracer.Tracer {
	return &otracer{
		Closer:      nil,
		serviceName: "",
		host:        "",
		port:        0,
	}
}

// SetServiceName sets a service name.
func (ot *otracer) SetServiceName(name string) {
	ot.serviceName = name
}

// SetAgentHost sets an agent host.
func (ot *otracer) SetAgentHost(host string) {
	ot.host = host
}

// SetAgentPort sets an agent port.
func (ot *otracer) SetAgentPort(port int) {
	ot.port = port
}

// SetEndpoint sets an endpoint.
func (ot *otracer) SetEndpoint(endpoint string) {
	ot.endpoint = endpoint
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

	if ot.endpoint != "" {
		u, err := url.ParseRequestURI(ot.endpoint)
		if err != nil {
			return fmt.Errorf("invalid endpoint: %w", err)
		}
		cfg.Reporter.CollectorEndpoint = u.String()
	} else {
		host := jaeger.DefaultUDPSpanServerHost
		if ot.host != "" {
			host = ot.host
		}
		port := jaeger.DefaultUDPSpanServerPort
		if ot.port != 0 {
			port = ot.port
		}
		cfg.Reporter.CollectorEndpoint = fmt.Sprintf("%s:%d", host, port)
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return err
	}

	ot.Closer = closer
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
