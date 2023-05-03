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

package opentelemetry

import (
	"context"

	"github.com/cybergarage/go-tracing/tracer"
	"go.opentelemetry.io/otel"
)

type otracer struct {
	serviceName string
	endpoint    string
}

func NewTracer() tracer.Tracer {
	return &otracer{
		serviceName: "",
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

func (ot *otracer) StartSpan(name string) tracer.SpanContext {
	ctx := context.Background()
	ctx, s := otel.Tracer(ot.serviceName).Start(ctx, name)
	return &spanContext{
		span: &span{
			name: name,
			Span: s,
			ctx:  ctx,
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
