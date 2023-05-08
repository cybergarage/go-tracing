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

package otel

import (
	"context"
	"time"

	"github.com/cybergarage/go-tracing/tracer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

const (
	defaultServiceName = "go-opentelemetry"
	defaultEndpoint    = "http://localhost:14268/api/traces"
)

type otracer struct {
	serviceName string
	endpoint    string
	tp          *tracesdk.TracerProvider
}

func NewTracer() tracer.Tracer {
	return &otracer{
		serviceName: defaultServiceName,
		endpoint:    defaultEndpoint,
		tp:          nil,
	}
}

// SetServiceName sets a service name.
func (ot *otracer) SetServiceName(name string) {
	ot.serviceName = name
}

// SetEndpoint sets an endpoint.
func (ot *otracer) SetEndpoint(endpoint string) {
	ot.endpoint = endpoint
}

// ServiceName returns the service name.
func (ot *otracer) ServiceName() string {
	return ot.serviceName
}

// Endpoint returns the endpoint.
func (ot *otracer) Endpoint() string {
	return ot.endpoint
}

func (ot *otracer) StartSpan(name string) tracer.Context {
	ctx := context.Background()
	tr := ot.tp.Tracer("")
	ctx, ots := tr.Start(ctx, name)
	span := &span{
		name: name,
		Span: ots,
		ctx:  ctx,
	}
	return NewSpanContextWith(span)
}

// Start starts a tracer.
func (ot *otracer) Start() error {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(ot.endpoint)))
	if err != nil {
		return err
	}
	ot.tp = tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(ot.serviceName),
		)),
	)
	otel.SetTracerProvider(ot.tp)
	return nil
}

// Stop stops a tracer.
func (ot *otracer) Stop() error {
	if ot.tp == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := ot.tp.Shutdown(ctx); err != nil {
		return err
	}
	ot.tp = nil
	return nil
}
