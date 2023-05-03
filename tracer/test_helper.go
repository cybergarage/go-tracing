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
	"testing"
)

func TestTracer(t *testing.T, tracer Tracer) {
	t.Helper()

	if err := tracer.Start(); err != nil {
		t.Error(err)
		return
	}

	defer func() {
		if err := tracer.Stop(); err != nil {
			t.Error(err)
			return
		}
	}()

	s := tracer.StartSpan("root")
	cs := s.Span().StartSpan("child")
	cs.Span().Finish()
	s.Span().Finish()
}
