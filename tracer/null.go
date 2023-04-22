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

var constNullTracer = nullTacer{}
var constNullSpan = nullSpan{}
var constNullSpanContext = nullSpanCotext{}

type nullTacer struct {
}

type nullSpan struct {
}

type nullSpanCotext struct {
}

func NullTracer() Tracer {
	return &constNullTracer
}

// StartSpan starts a new span.
func (nt *nullTacer) StartSpan(_ string) SpanContext {
	return &constNullSpanContext
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
func (s *nullSpan) StartSpan(_ string) SpanContext {
	return &constNullSpanContext
}

// Span returns the context span.
func (ctx *nullSpanCotext) Span() Span {
	return &constNullSpan
}
