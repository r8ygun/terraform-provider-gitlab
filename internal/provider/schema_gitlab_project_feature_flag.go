package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/xanzy/go-gitlab"
)

func gitlabProjectFeatureFlagGetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description:  "The ID or URL-encoded path of the project.",
			Type:         schema.TypeString,
			ForceNew:     true,
			Required:     true,
			ValidateFunc: StringIsGitlabVariableName,
		},
		"feature_flag_name": {
			Description: "The name of the feature flag.",
			Type:        schema.TypeString,
			ForceNew:    true,
			Required:    true,
		},
	}
}

func gitlabProjectFeatureFlagToStateMap(variable *gitlab.ProjectFeatureFlag) map[string]interface{} {
	stateMap := make(map[string]interface{})
	stateMap["name"] = variable.Name
	stateMap["description"] = variable.Description
	stateMap["active"] = variable.Active
	stateMap["version"] = variable.Version
	stateMap["created_at"] = variable.CreatedAt
	stateMap["updated_at"] = variable.UpdatedAt
	stateMap["scopes"] = variable.Scopes
	stateMap["strategies"] = variable.Strategies
	return stateMap
}
