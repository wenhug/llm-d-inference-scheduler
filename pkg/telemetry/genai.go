/*
Copyright 2025 The llm-d Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package telemetry provides OpenTelemetry semantic convention helpers
// for GenAI inference observability.
package telemetry

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.40.0"
)

// genAIRequestIDKey is a custom attribute key for the inference request ID.
// This is not part of the official OTel GenAI semantic conventions but is
// used by llm-d to correlate traces with specific inference requests.
//
// Unexported to prevent accidental mutation; use GenAIRequestID() instead.
const genAIRequestIDKey = attribute.Key("gen_ai.request.id")

// GenAIRequestModel returns an attribute KeyValue for the requested model name,
// conforming to the OTel GenAI semantic conventions.
// See: https://opentelemetry.io/docs/specs/semconv/gen-ai/gen-ai-spans/
func GenAIRequestModel(val string) attribute.KeyValue {
	return semconv.GenAIRequestModel(val)
}

// GenAIRequestID returns an attribute KeyValue for the inference request ID.
func GenAIRequestID(val string) attribute.KeyValue {
	return genAIRequestIDKey.String(val)
}
