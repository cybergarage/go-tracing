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
	"github.com/cybergarage/go-tracing/tracer"
)

type spanContext struct {
	spans []tracer.Span
}

// SetSpan sets a new root span context.
func NewSpanContextWith(span tracer.Span) tracer.SpanContext {
	ctx := &spanContext{
		spans: make([]tracer.Span, 0),
	}
	ctx.pushSpan(span)
	return ctx
}

// Span returns the current top span on the span stack.
func (ctx *spanContext) Span() tracer.Span {
	if len(ctx.spans) == 0 {
		return nil
	}
	return ctx.spans[len(ctx.spans)-1]
}

// StartSpan starts a new child span and pushes it onto the span stack.
func (ctx *spanContext) StartSpan(name string) bool {
	if len(ctx.spans) == 0 {
		return false
	}
	span := ctx.spans[len(ctx.spans)-1]
	childSpan := span.StartSpan(name)
	ctx.pushSpan(childSpan.Span())
	return true
}

// FinishSpan ends the current top span and pops it from the span stack.
func (ctx *spanContext) FinishSpan() bool {
	span := ctx.popSpan()
	if span == nil {
		return false
	}
	span.Finish()
	return true
}

func (ctx *spanContext) popSpan() tracer.Span {
	if len(ctx.spans) == 0 {
		return nil
	}
	lastSpanIndex := len(ctx.spans) - 1
	span := ctx.spans[lastSpanIndex]
	ctx.spans = ctx.spans[:lastSpanIndex]
	return span
}

func (ctx *spanContext) pushSpan(span tracer.Span) {
	ctx.spans = append(ctx.spans, span)
}
