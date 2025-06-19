// Copyright 2025 Clyso GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package env

import "github.com/testcontainers/testcontainers-go"

type ContainerLogType string

const (
	// Represents container standard output log.
	LogStdout ContainerLogType = testcontainers.StdoutLog
	// Represents container standard error log.
	LogStderr ContainerLogType = testcontainers.StderrLog
)

// Option interface to support Functional options for test-containers.
// Allows to pass different options to a function as variadic parameters.
// see: https://github.com/uber-go/guide/blob/master/style.md#functional-options
type Option interface {
	apply(*ComponentCreationConfig)
}

type disableContainerLogOption struct {
	types []ContainerLogType
}

func (o disableContainerLogOption) apply(cfg *ComponentCreationConfig) {
	cfg.DisabledLogs = append(cfg.DisabledLogs, o.types...)
}

// WithDisabledLog returns an Option that disables the specified container logs.
func WithDisabledLog(types ...ContainerLogType) Option {
	return disableContainerLogOption{types: types}
}
