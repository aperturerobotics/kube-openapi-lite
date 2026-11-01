package util

import (
	"testing"

	"k8s.io/kube-openapi/pkg/validation/spec"
)

// TestWrapRefs preserves reference siblings and leaves bare references unchanged.
func TestWrapRefs(t *testing.T) {
	// A bare reference needs no wrapper or copy.
	ref := spec.MustCreateRef("#/components/schemas/Example")
	bare := &spec.Schema{SchemaProps: spec.SchemaProps{Ref: ref}}
	if WrapRefs(bare) != bare {
		t.Fatal("bare reference was copied")
	}

	// Even explicitly empty containers count as siblings under the schema contract.
	withSiblings := &spec.Schema{SchemaProps: spec.SchemaProps{Ref: ref, Title: "example", Required: []string{}}}
	wrapped := WrapRefs(withSiblings)
	if wrapped == withSiblings || wrapped.Ref.String() != "" || len(wrapped.AllOf) != 1 || wrapped.AllOf[0].Ref.String() != ref.String() {
		t.Fatal("reference siblings were not wrapped")
	}

	// Wrapping preserves the sibling fields and the caller's original schema.
	if wrapped.Title != "example" || wrapped.Required == nil || withSiblings.Ref.String() != ref.String() {
		t.Fatal("wrapping changed schema siblings or the original reference")
	}
}
