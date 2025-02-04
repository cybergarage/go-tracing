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

package ot

import (
	"io"

	"github.com/cybergarage/go-logger/log"

	opentracing "github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	"github.com/cybergarage/go-tracing/tracer"
)

const (
	serviceName     = "go-opentracing"
	defaultEndpoint = "localhost:6831"
)

type otracer struct {
	io.Closer
	pkgName     string
	serviceName string
	endpoint    string
}

func NewTracer() tracer.Tracer {
	return &otracer{
		Closer:      nil,
		serviceName: tracer.PackageName,
		endpoint:    defaultEndpoint,
	}
}

// SetPackageName sets a package name.
func (ot *otracer) SetPackageName(name string) {
	ot.pkgName = name
}

// PackageName returns the package name.
func (ot *otracer) PackageName() string {
	return ot.pkgName
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

// StartSpan starts a new span.
func (ot *otracer) StartSpan(name string) tracer.Context {
	gt := opentracing.GlobalTracer()
	span := &span{
		Tracer: gt,
		Span:   gt.StartSpan(name),
		ctx:    nil,
	}
	return NewSpanContextWith(span)
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
		cfg.Reporter.CollectorEndpoint = ot.endpoint
	}

	jtracer, closer, err := cfg.NewTracer()
	if err != nil {
		return err
	}

	ot.Closer = closer
	opentracing.SetGlobalTracer(jtracer)

	log.Infof("%s(%s)/%s (%s) started", tracer.PackageName, serviceName, tracer.Version, ot.endpoint)

	return nil
}

// Stop stops a tracer.
func (ot *otracer) Stop() error {
	if ot.Closer != nil {
		if err := ot.Closer.Close(); err != nil {
			return err
		}
	}

	log.Infof("%s(%s)/%s (%s) terminated", tracer.PackageName, serviceName, tracer.Version, ot.endpoint)

	return nil
}
