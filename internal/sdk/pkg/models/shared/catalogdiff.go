// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CatalogDiff - Describes the difference between two Airbyte catalogs.
type CatalogDiff struct {
	// list of stream transformations. order does not matter.
	Transforms []StreamTransform `json:"transforms"`
}