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

type Tracer interface {
	// SetPackageName sets a package name.
	SetPackageName(name string)
	// SetServiceName sets a service name.
	SetServiceName(name string)
	// SetEndpoint sets an endpoint.
	SetEndpoint(endpoint string)
	// PackageName returns the package name.
	PackageName() string
	// ServiceName returns the service name.
	ServiceName() string
	// Endpoint returns the endpoint.
	Endpoint() string
	// StartSpan starts a new span.
	StartSpan(name string) Context
	// Start starts a tracer.
	Start() error
	// Stop stops a tracer.
	Stop() error
}
