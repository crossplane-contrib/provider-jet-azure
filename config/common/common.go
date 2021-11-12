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
	"path/filepath"
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

// GetFullyQualifiedIDFn returns a GetIDFn that can parse any Azure fully qualifier.
// An example identifier is as follows:
// /subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/sqlDatabases/%s/containers/%s
// An input to this function to parse it would be:
// serviceProvider: "Microsoft.DocumentDB"
// keyPairs: []string{"databaseAccounts", "account_name", "sqlDatabases", "database_name", "containers", "name"}
func GetFullyQualifiedIDFn(serviceProvider string, keyPairs ...string) tjconfig.GetIDFn {
	if len(keyPairs)%2 != 0 {
		panic("each service name has to have a key")
	}
	return func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
		subID, ok := providerConfig["subscriptionId"].(string)
		if !ok {
			return ""
		}
		rg, ok := parameters["resource_group_name"].(string)
		if !ok {
			return ""
		}
		path := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/%s", subID, rg, serviceProvider)
		for i := 0; i < len(keyPairs); i += 2 {
			val, ok := parameters[keyPairs[i+1]].(string)
			if !ok {
				return ""
			}
			path = filepath.Join(path, keyPairs[i], val)
		}
		return path
	}
}
