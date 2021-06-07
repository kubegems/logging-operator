// Copyright © 2019 Banzai Cloud
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

package output_test

import (
	"testing"

	"github.com/banzaicloud/logging-operator/pkg/sdk/model/output"
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/render"
	"github.com/ghodss/yaml"
)

func TestElasticSearch(t *testing.T) {
	CONFIG := []byte(`
host: elasticsearch-elasticsearch-cluster.default.svc.cluster.local
port: 9200
scheme: https
ssl_version: TLSv1_2
ssl_verify: false
buffer:
  timekey: 1m
  timekey_wait: 30s
  timekey_use_utc: true
`)
	expected := `
  <match **>
	@type elasticsearch
	@id test
	exception_backup true
	fail_on_putting_template_retry_exceed true
	host elasticsearch-elasticsearch-cluster.default.svc.cluster.local
	port 9200
	reload_connections true
	scheme https
	ssl_verify false
	ssl_version TLSv1_2
	utc_index true
	verify_es_version_at_startup true
    <buffer tag,time>
      @type file
	  chunk_limit_size 8MB
      path /buffers/test.*.buffer
      retry_forever true
      timekey 1m
      timekey_use_utc true
      timekey_wait 30s
    </buffer>
  </match>
`
	es := &output.ElasticsearchOutput{}
	yaml.Unmarshal(CONFIG, es)
	test := render.NewOutputPluginTest(t, es)
	test.DiffResult(expected)
}
