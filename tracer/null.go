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

package tracer

import (
	"context"
)

// NullTracer is a null tracing tracer.
var NullTracer = &nullTacer{}

// NullSpan is a null tracing span.
var NullSpan = nullSpan{}

// NullSpanContext is a null tracing span context.
var NullSpanContext = nullSpanCotext{}

type nullTacer struct {
}

// NewNullTracer returns a new null tracing tracer.
func NewNullTracer() Tracer {
	return NullTracer
}

// SetPackageName sets a package name.
func (nt *nullTacer) SetPackageName(_ string) {
}

// PackageName returns the package name.
func (nt *nullTacer) PackageName() string {
	return ""
}

// SetServiceName sets a service name.
func (nt *nullTacer) SetServiceName(_ string) {
}

// SetEndpoint sets an endpoint.
func (nt *nullTacer) SetEndpoint(_ string) {
}

// ServiceName returns the service name.
func (nt *nullTacer) ServiceName() string {
	return ""
}

// Endpoint returns the endpoint.
func (nt *nullTacer) Endpoint() string {
	return ""
}

// StartSpan starts a new span.
func (nt *nullTacer) StartSpan(_ string) Context {
	return &NullSpanContext
}

// Start starts a tracer.
func (nt *nullTacer) Start() error {
	return nil
}

// Stop stops a tracer.
func (nt *nullTacer) Stop() error {
	return nil
}

type nullSpan struct {
}

// SetTag sets a tag on the span.
func (s *nullSpan) SetTag(_ string, _ any) {
}

// Finish marks the end of the span.
func (s *nullSpan) Finish() {
}

// Context returns the span's context.
func (s *nullSpan) Context() context.Context {
	return nil
}

// StartSpan starts a new child span.
func (s *nullSpan) StartSpan(_ string) Context {
	return &NullSpanContext
}

type nullSpanCotext struct {
}

// Span returns the current top span on the span stack.
func (ctx *nullSpanCotext) Span() Span {
	return &NullSpan
}

// StartSpan starts a new child span and pushes it onto the span stack.
func (ctx *nullSpanCotext) StartSpan(_ string) bool {
	return true
}

// FinishSpan ends the current top span and pops it from the span stack.
func (ctx *nullSpanCotext) FinishSpan() bool {
	return true
}
