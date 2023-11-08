// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type DestinationDefinitionRead struct {
	DestinationDefinitionID string  `json:"destinationDefinitionId"`
	DockerImageTag          string  `json:"dockerImageTag"`
	DockerRepository        string  `json:"dockerRepository"`
	DocumentationURL        string  `json:"documentationUrl"`
	Icon                    *string `json:"icon,omitempty"`
	Name                    string  `json:"name"`
	// describes a normalization config for destination definition
	NormalizationConfig NormalizationDestinationDefinitionConfig `json:"normalizationConfig"`
	// The Airbyte Protocol version supported by the connector
	ProtocolVersion *string `json:"protocolVersion,omitempty"`
	// The date when this connector was first released, in yyyy-mm-dd format.
	ReleaseDate  *types.Date   `json:"releaseDate,omitempty"`
	ReleaseStage *ReleaseStage `json:"releaseStage,omitempty"`
	// actor definition specific resource requirements. if default is set, these are the requirements that should be set for ALL jobs run for this actor definition. it is overriden by the job type specific configurations. if not set, the platform will use defaults. these values will be overriden by configuration at the connection level.
	ResourceRequirements *ActorDefinitionResourceRequirements `json:"resourceRequirements,omitempty"`
	// an optional flag indicating whether DBT is used in the normalization. If the flag value is NULL - DBT is not used.
	SupportsDbt bool `json:"supportsDbt"`
}

func (d DestinationDefinitionRead) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDefinitionRead) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDefinitionRead) GetDestinationDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.DestinationDefinitionID
}

func (o *DestinationDefinitionRead) GetDockerImageTag() string {
	if o == nil {
		return ""
	}
	return o.DockerImageTag
}

func (o *DestinationDefinitionRead) GetDockerRepository() string {
	if o == nil {
		return ""
	}
	return o.DockerRepository
}

func (o *DestinationDefinitionRead) GetDocumentationURL() string {
	if o == nil {
		return ""
	}
	return o.DocumentationURL
}

func (o *DestinationDefinitionRead) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *DestinationDefinitionRead) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationDefinitionRead) GetNormalizationConfig() NormalizationDestinationDefinitionConfig {
	if o == nil {
		return NormalizationDestinationDefinitionConfig{}
	}
	return o.NormalizationConfig
}

func (o *DestinationDefinitionRead) GetProtocolVersion() *string {
	if o == nil {
		return nil
	}
	return o.ProtocolVersion
}

func (o *DestinationDefinitionRead) GetReleaseDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.ReleaseDate
}

func (o *DestinationDefinitionRead) GetReleaseStage() *ReleaseStage {
	if o == nil {
		return nil
	}
	return o.ReleaseStage
}

func (o *DestinationDefinitionRead) GetResourceRequirements() *ActorDefinitionResourceRequirements {
	if o == nil {
		return nil
	}
	return o.ResourceRequirements
}

func (o *DestinationDefinitionRead) GetSupportsDbt() bool {
	if o == nil {
		return false
	}
	return o.SupportsDbt
}
