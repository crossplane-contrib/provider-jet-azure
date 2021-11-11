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
	"strings"

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/pkg/errors"
)

const (
	fmtID                = "/subscriptions/%s/resourceGroups/%s/providers%s/%s"
	errFmtNoAttribute    = `"attribute not found: %s`
	errFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
)

// Examples of fully qualifiers:
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMySQL/servers/server1
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkInterfaces/nic1

func GetNameFromFullyQualifiedID(tfstate map[string]interface{}) (string, error) {
	id, ok := tfstate["id"]
	if !ok {
		return "", errors.Errorf(errFmtNoAttribute, "id")
	}
	idStr, ok := id.(string)
	if !ok {
		return "", errors.Errorf(errFmtUnexpectedType, "id")
	}
	words := strings.Split(idStr, "/")
	return words[len(words)-1], nil
}

func GetFullyQualifiedIDFn(servicePath string) tjconfig.GetIDFn {
	return func(ctx context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
		subID, ok := providerConfig["subscriptionId"]
		if !ok {
			return "", errors.Errorf(errFmtNoAttribute, "subscriptionId")
		}
		subIDStr, ok := subID.(string)
		if !ok {
			return "", errors.Errorf(errFmtUnexpectedType, "subscriptionId")
		}
		rg, ok := parameters["resource_group_name"]
		if !ok {
			return "", errors.Errorf(errFmtNoAttribute, "resource_group_name")
		}
		rgStr, ok := rg.(string)
		if !ok {
			return "", errors.Errorf(errFmtUnexpectedType, "resource_group_name")
		}
		return fmt.Sprintf(fmtID, subIDStr, rgStr, servicePath, externalName), nil
	}
}
