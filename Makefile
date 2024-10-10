# Copyright (C) 2023 The go-tracing Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL := bash

MODULE_ROOT=github.com/cybergarage/go-tracing

PKG_NAME=tracer
PKG_ID=${MODULE_ROOT}/${PKG_NAME}
PKG_ROOT=${PKG_NAME}
PKGS=${PKG_ID}

.PHONY: format vet lint clean version

all: test

version:
	@pushd ${PKG_ROOT} && ./version.gen > version.go && popd
	-git commit ${PKG_ROOT}/version.go -m "Update version"

format: version
	gofmt -w ${PKG_ROOT}

vet: format
	go vet ${PKG_ID}

lint: vet
	golangci-lint run ${PKG_ROOT}

test: lint
	go test -v -cover -timeout 60s ${PKGS}/...

clean:
	go clean -i ${PKGS}
