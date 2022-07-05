/*
Copyright 2022 The Crossplane Authors.

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

package devspace

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure devspace iothub group
func Configure(p *config.Provider) {
	// configured just for preventing package collision in internal/controller/zz_setup.go
	// we will need a proper external name configuration and testing for this resource.
	p.AddResourceConfigurator("azurerm_devspace_controller", func(r *config.Resource) {
		r.Kind = "DevSpaceController"
	})
}
