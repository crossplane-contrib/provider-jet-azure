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
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	tjconfig "github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-azure/apis/rconfig"
)

const (
	// ErrFmtNoAttribute is an error string for not-found attributes
	ErrFmtNoAttribute = `"attribute not found: %s`
	// ErrFmtUnexpectedType is an error string for attribute map values of unexpected type
	ErrFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`

	// nameReferenceKind represents name reference kind
	nameReferenceKind = "name"
	// idReferenceKind represents ID reference kind
	idReferenceKind = "id"
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

// In the regular expressions used when determining the references, we can observe many rules that have common strings.
// Some of these may be more special than others. For example, "ex_subnet$" is a more special case than "subnet$".
// By putting the more specific rules before the general rules in array, processing and capturing the specific
// rules first will be possible. Since there is no fixed index for key-value pairs in maps, it is not possible to place
// rules from specific to general. Therefore, array is used here.
var referenceRules = [][]string{
	{"resource_group$", "/azure/v1alpha1.ResourceGroup"},
	{"subnet$", "/network/v1alpha1.Subnet"},
}

// AddCommonReferences adds some common reference fields.
// This is a part of resource generation pipeline.
func AddCommonReferences(r *tjconfig.Resource) error {
	return addCommonReferences(r.References, r.TerraformResource, r.ShortGroup, []string{})
}

func addCommonReferences(references tjconfig.References, resource *schema.Resource, shortGroup string, nestedFieldNames []string) error {
	for fieldName, s := range resource.Schema {
		if s.Elem != nil {
			e, ok := s.Elem.(*schema.Resource)
			if ok {
				if err := addCommonReferences(references, e, shortGroup, append(nestedFieldNames, fieldName)); err != nil {
					return err
				}
				continue
			}
		}

		referenceName := strings.Join(append(nestedFieldNames, fieldName), ".")
		referenceNameWithoutKind, referenceKind := splitReferenceKindFromReferenceName(referenceName)
		if referenceKind != "" {
			referenceType, err := searchReference(referenceNameWithoutKind)
			if err != nil {
				return err
			}
			if referenceType != "" {
				if references == nil {
					references = make(map[string]tjconfig.Reference)
				}
				if _, ok := references[referenceName]; !ok {
					referenceType = prepareReferenceType(shortGroup, referenceType)
					addReference(references, referenceKind, referenceName, referenceType)
				}
			}
		}
	}
	return nil
}

func splitReferenceKindFromReferenceName(fieldName string) (string, string) {
	p := strings.Split(fieldName, "_")
	if len(p) < 2 {
		return "", ""
	}
	return strings.Join(p[:len(p)-1], "_"), p[len(p)-1]
}

func searchReference(fieldName string) (string, error) {
	for _, rule := range referenceRules {
		r, err := regexp.Compile(rule[0])
		if err != nil {
			return "", err
		}
		if r.MatchString(fieldName) {
			return rule[1], nil
		}
	}
	return "", nil
}

func prepareReferenceType(shortGroup, referenceType string) string {
	p := strings.Split(referenceType, "/")
	if shortGroup == p[1] {
		referenceType = strings.Split(p[2], ".")[1]
	} else {
		referenceType = rconfig.APISPackagePath + referenceType
	}
	return referenceType
}

func addReference(references tjconfig.References, referenceKind, referenceName, referenceType string) {
	switch referenceKind {
	case nameReferenceKind:
		references[referenceName] = tjconfig.Reference{
			Type: referenceType,
		}
	case idReferenceKind:
		references[referenceName] = tjconfig.Reference{
			Type:      referenceType,
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	}
}
