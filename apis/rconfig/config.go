/*
Copyright 2021 The Crossplane Authors.

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

package rconfig

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

// ExtractResourceName extracts the value of `spec.forProvider.name`
// from a managed resource
func ExtractResourceName() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		p, err := fieldpath.PaveObject(mr)
		if err != nil {
			return ""
		}

		name, err := p.GetString("spec.forProvider.name")
		if err != nil {
			return ""
		}
		return name
	}
}
