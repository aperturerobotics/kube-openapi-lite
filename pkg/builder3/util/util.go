/*
Copyright 2022 The Kubernetes Authors.

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

package util

import (
	"k8s.io/kube-openapi/pkg/schemamutation"
	"k8s.io/kube-openapi/pkg/validation/spec"
)

// WrapRefs wraps OpenAPI V3 Schema refs that contain sibling elements.
//
// AllOf is used to wrap the Ref to prevent references from having sibling elements
// Please see https://github.com/kubernetes/kubernetes/issues/106387#issuecomment-967640388
func WrapRefs(schema *spec.Schema) *spec.Schema {
	walker := schemamutation.Walker{
		SchemaCallback: func(schema *spec.Schema) *spec.Schema {
			// Leave inline schemas and bare references unchanged.
			if schema.Ref.String() == "" || !hasRefSiblings(schema) {
				return schema
			}

			// Preserve the caller's schema while moving its reference into AllOf.
			wrapped := *schema
			wrapped.Ref = spec.Ref{}
			wrapped.AllOf = []spec.Schema{{SchemaProps: spec.SchemaProps{Ref: schema.Ref}}}
			return &wrapped
		},
		RefCallback: schemamutation.RefCallbackNoop,
	}
	return walker.WalkSchema(schema)
}

// hasRefSiblings reports fields other than Ref, including explicitly empty collections.
func hasRefSiblings(schema *spec.Schema) bool {
	return schema.ID != "" ||
		schema.Schema != "" ||
		schema.Description != "" ||
		schema.Type != nil ||
		schema.Nullable ||
		schema.Format != "" ||
		schema.Title != "" ||
		schema.Default != nil ||
		schema.Maximum != nil ||
		schema.ExclusiveMaximum ||
		schema.Minimum != nil ||
		schema.ExclusiveMinimum ||
		schema.MaxLength != nil ||
		schema.MinLength != nil ||
		schema.Pattern != "" ||
		schema.MaxItems != nil ||
		schema.MinItems != nil ||
		schema.UniqueItems ||
		schema.MultipleOf != nil ||
		schema.Enum != nil ||
		schema.MaxProperties != nil ||
		schema.MinProperties != nil ||
		schema.Required != nil ||
		schema.Items != nil ||
		schema.AllOf != nil ||
		schema.OneOf != nil ||
		schema.AnyOf != nil ||
		schema.Not != nil ||
		schema.Properties != nil ||
		schema.AdditionalProperties != nil ||
		schema.PatternProperties != nil ||
		schema.Dependencies != nil ||
		schema.AdditionalItems != nil ||
		schema.Definitions != nil ||
		schema.Discriminator != "" ||
		schema.ReadOnly ||
		schema.ExternalDocs != nil ||
		schema.Example != nil ||
		schema.Extensions != nil ||
		schema.ExtraProps != nil
}
