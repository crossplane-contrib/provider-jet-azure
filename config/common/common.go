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
	"context"
	"fmt"
	"path/filepath"
	"strings"

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/pkg/errors"
)

const (
	// ErrFmtNoAttribute is an error string for not-found attributes
	ErrFmtNoAttribute = `"attribute not found: %s`
	// ErrFmtUnexpectedType is an error string for attribute map values of unexpected type
	ErrFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
)

// GetNameFromFullyQualifiedID extracts external-name from Azure ID
// Examples of fully qualifiers:
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMySQL/servers/server1
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkInterfaces/nic1
func GetNameFromFullyQualifiedID(tfstate map[string]interface{}) (string, error) {
	id, ok := tfstate["id"]
	if !ok {
		return "", errors.Errorf(ErrFmtNoAttribute, "id")
	}
	idStr, ok := id.(string)
	if !ok {
		return "", errors.Errorf(ErrFmtUnexpectedType, "id")
	}
	words := strings.Split(idStr, "/")
	return words[len(words)-1], nil
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
	return func(ctx context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
		subID, ok := providerConfig["subscription_id"]
		if !ok {
			return "", errors.Errorf(ErrFmtNoAttribute, "subscription_id")
		}
		subIDStr, ok := subID.(string)
		if !ok {
			return "", errors.Errorf(ErrFmtUnexpectedType, "subscription_id")
		}
		rg, ok := parameters["resource_group_name"]
		if !ok {
			return "", errors.Errorf(ErrFmtNoAttribute, "resource_group_name")
		}
		rgStr, ok := rg.(string)
		if !ok {
			return "", errors.Errorf(ErrFmtUnexpectedType, "resource_group_name")
		}

		path := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/%s", subIDStr, rgStr, serviceProvider)
		for i := 0; i < len(keyPairs); i += 2 {
			val, ok := parameters[keyPairs[i+1]]
			if !ok {
				return "", errors.Errorf(ErrFmtNoAttribute, keyPairs[i+1])
			}
			valStr, ok := val.(string)
			if !ok {
				return "", errors.Errorf(ErrFmtUnexpectedType, keyPairs[i+1])
			}
			path = filepath.Join(path, keyPairs[i], valStr)
		}
		return path, nil
	}
}
