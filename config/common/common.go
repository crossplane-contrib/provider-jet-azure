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

package common

import (
	"fmt"
	"strings"

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
)

// Examples of fully qualifiers:
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMySQL/servers/server1
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkInterfaces/nic1

func GetNameFromFullyQualifiedID(tfstate map[string]interface{}) string {
	id, ok := tfstate["id"].(string)
	if !ok {
		return ""
	}
	words := strings.Split(id, "/")
	return words[len(words)-1]
}

func GetFullyQualifiedIDFn(servicePath string) tjconfig.GetIDFn {
	return func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
		subID, ok := providerConfig["subscriptionId"].(string)
		if !ok {
			return ""
		}
		rg, ok := parameters["resource_group_name"].(string)
		if !ok {
			return ""
		}
		return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers%s/%s", subID, rg, servicePath, name)
	}
}
