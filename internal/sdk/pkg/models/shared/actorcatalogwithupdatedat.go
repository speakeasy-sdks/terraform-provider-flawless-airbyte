// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ActorCatalogWithUpdatedAtCatalog struct {
}

// ActorCatalogWithUpdatedAt - A source actor catalog with the timestamp it was mostly recently updated
type ActorCatalogWithUpdatedAt struct {
	Catalog   *ActorCatalogWithUpdatedAtCatalog `json:"catalog,omitempty"`
	UpdatedAt *int64                            `json:"updatedAt,omitempty"`
}

func (o *ActorCatalogWithUpdatedAt) GetCatalog() *ActorCatalogWithUpdatedAtCatalog {
	if o == nil {
		return nil
	}
	return o.Catalog
}

func (o *ActorCatalogWithUpdatedAt) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
