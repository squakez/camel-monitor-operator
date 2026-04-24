/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"encoding/json"
	"strings"
)

// RawMetric is used to parse metrics logged by Camel on shutdown.
type RawMetric struct {
	ID struct {
		Name string              `json:"name"`
		Tags []map[string]string `json:"tags"`
	} `json:"id"`

	Value float64 `json:"value,omitempty"`
	Count float64 `json:"count,omitempty"`
}

// ParseLine is used to parse a single metric as a line (json format expected).
func ParseLine(line string) (*RawMetric, bool) {
	start := strings.Index(line, "#METRIC-START#")
	end := strings.Index(line, "#METRIC-END#")

	if start == -1 || end == -1 {
		return nil, false
	}

	jsonPart := line[start+len("#METRIC-START#") : end]
	// handle any possible "nan" string value
	jsonPart = strings.ReplaceAll(jsonPart, "\"value\":\"NaN\"", "\"value\": -1")

	var m RawMetric
	if err := json.Unmarshal([]byte(jsonPart), &m); err != nil {
		return nil, false
	}

	return &m, true
}
