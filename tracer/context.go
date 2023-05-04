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

// Context represents a span context.
type Context interface {
	// Span returns the current top tracer span on the tracer span stack.
	Span() Span
	// StartSpan starts a new child tracer span and pushes it onto the tracer span stack.
	StartSpan(name string) bool
	// FinishSpan ends the current top tracer span and pops it from the tracer span stack.
	FinishSpan() bool
}
