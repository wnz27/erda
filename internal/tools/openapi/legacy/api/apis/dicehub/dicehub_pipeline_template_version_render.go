// Copyright (c) 2021 Terminus, Inc.
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

package dicehub

import "github.com/erda-project/erda/internal/tools/openapi/legacy/api/apis"

var DICEHUB_PIPELINE_TEMPLATE_VERSION_RENDER = apis.ApiSpec{
	Path:        "/api/pipeline-templates/<name>/actions/render",
	BackendPath: "/api/pipeline-templates/<name>/actions/render",
	Host:        "erda-server.marathon.l4lb.thisdcos.directory:9095",
	Scheme:      "http",
	Method:      "POST",
	IsOpenAPI:   true,
	CheckLogin:  true,
	CheckToken:  true,
	Doc:         `summary: 渲染模板对应版本的pipeline.yml`,
}
