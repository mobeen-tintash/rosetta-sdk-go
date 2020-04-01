// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package gen

// Version: struct for Version
type Version struct {
	// The `rosetta_version` is the version of the Rosetta interface the implementation adheres to. This can be useful for clients looking to reliably parse responses.
	RosettaVersion string `json:"rosetta_version"`
	// The `node_version` is the cannonical version of the node runtime. This can help clients manage deployments.
	NodeVersion string `json:"node_version"`
	// When a middleware server is used to adhere to the Rosetta interface, it should return its version here. This can help clients manage deployments.
	MiddlewareVersion *string `json:"middleware_version,omitempty"`
	// Any other information that may be useful about versioning of dependent services should be returned here.
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
}
