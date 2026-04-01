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

package telemetry

import "testing"

func TestGenAIRequestModel(t *testing.T) {
	tests := []struct {
		name  string
		model string
	}{
		{
			name:  "standard model name",
			model: "meta-llama/Llama-3.1-8B-Instruct",
		},
		{
			name:  "empty model name",
			model: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := GenAIRequestModel(tt.model)
			if string(kv.Key) != "gen_ai.request.model" {
				t.Errorf("GenAIRequestModel() key = %q, want %q", kv.Key, "gen_ai.request.model")
			}
			if kv.Value.AsString() != tt.model {
				t.Errorf("GenAIRequestModel() value = %q, want %q", kv.Value.AsString(), tt.model)
			}
		})
	}
}

func TestGenAIRequestID(t *testing.T) {
	tests := []struct {
		name      string
		requestID string
	}{
		{
			name:      "standard request ID",
			requestID: "req-12345-abcde",
		},
		{
			name:      "empty request ID",
			requestID: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := GenAIRequestID(tt.requestID)
			if string(kv.Key) != "gen_ai.request.id" {
				t.Errorf("GenAIRequestID() key = %q, want %q", kv.Key, "gen_ai.request.id")
			}
			if kv.Value.AsString() != tt.requestID {
				t.Errorf("GenAIRequestID() value = %q, want %q", kv.Value.AsString(), tt.requestID)
			}
		})
	}
}

func TestGenAIRequestIDKey(t *testing.T) {
	if genAIRequestIDKey != "gen_ai.request.id" {
		t.Errorf("genAIRequestIDKey = %q, want %q", genAIRequestIDKey, "gen_ai.request.id")
	}
}
