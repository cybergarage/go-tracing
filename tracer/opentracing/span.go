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
	"context"

	ot "github.com/opentracing/opentracing-go"
)

type span struct {
	ot.Span
	ctx context.Context
}

// SetTag sets a tag on the span.
func (s *span) SetTag(key string, value any) {
	s.Span.SetTag(key, value)
}

// Finish marks the end of the span.
func (s *span) Finish() {
	s.Span.Finish()
}

// Context returns the span's context.
func (s *span) Context() context.Context {
	return s.ctx
}