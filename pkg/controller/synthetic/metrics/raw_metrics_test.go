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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseValue(t *testing.T) {
	line := `2026-04-24 10:34:20.144  WARN 156781 --- [           main] icrometer.json.AbstractMicrometerService : #METRIC-START#{"id":{"name":"jvm.memory.used","tags":[{"area":"nonheap"},{"id":"Compressed Class Space"}]},"value":4528752.0}#METRIC-END#`
	metric, ok := ParseLine(line)
	assert.True(t, ok)
	require.NotNil(t, metric)
	assert.Equal(t, float64(4528752), metric.Value)
}

func TestParseCount(t *testing.T) {
	line := `2026-04-24 10:34:20.147  WARN 156781 --- [           main] icrometer.json.AbstractMicrometerService : #METRIC-START#{"id":{"name":"camel.exchanges.external.redeliveries","tags":[{"camelContext":"camel-1"},{"eventType":"context"},{"kind":"CamelRoute"},{"routeId":""}]},"count":0.0}#METRIC-END#`
	metric, ok := ParseLine(line)
	assert.True(t, ok)
	require.NotNil(t, metric)
	assert.Equal(t, float64(0), metric.Count)
}

func TestParseLabels(t *testing.T) {
	line := `2026-04-24 10:34:20.146  WARN 156781 --- [           main] icrometer.json.AbstractMicrometerService : #METRIC-START#{"id":{"name":"app.info","tags":[{"camel.context":"camel-1"},{"camel.runtime.provider":"Main"},{"camel.runtime.version":"4.20.0-SNAPSHOT"},{"camel.version":"4.20.0-SNAPSHOT"}]},"value":"NaN"}#METRIC-END#`
	metric, ok := ParseLine(line)
	assert.True(t, ok)
	require.NotNil(t, metric)
	assert.Contains(t, metric.ID.Tags, map[string]string{"camel.runtime.provider": "Main"})
	assert.Contains(t, metric.ID.Tags, map[string]string{"camel.runtime.version": "4.20.0-SNAPSHOT"})
}
