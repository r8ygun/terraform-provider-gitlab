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
			Required:    true,
		},
		"name": {
			Description: "The name of the feature flag.",
			Type:        schema.TypeString,
		},
		"description": {
			Description: "The description of the feature flag.",
			Type:        schema.TypeString,
		},
		"active": {
			Description: "The active state of the flag. Defaults to true. Supported in GitLab 13.3 and later.",
			Type:        schema.TypeBool,
		},
		"version": {
			Description: "The version of the feature flag.",
			Type:        schema.TypeString,
		},
		"created_at": {
			Description: "The date/time that the feature flag was created.",
			Type:        schema.TypeString,
		},
		"updated_at": {
			Description: "The date/time that the feature flag was last updated.",
			Type:        schema.TypeString,
		},
		"scopes": {
			Description: "The scopes for the strategy.",
			Type:        schema.TypeList,
			Elem: &schema.Resource{
				Schema: datasourceSchemaFromResourceSchema(gitlabProjectFeatuerFlagScopeSchema(), nil, nil),
			},
		},
		"strategies": {
			Description: "The feature flag strategies.",
			Type:        schema.TypeList,
			Elem: &schema.Resource{
				Schema: datasourceSchemaFromResourceSchema(gitlabProjectFeatureFlagStrategySchema(), nil, nil),
			},
		},
	}
}

func gitlabProjectFeatuerFlagScopeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: "The ID of the scope.",
			Type:        schema.TypeInt,
		},
		"environment_scope": {
			Description: "The environment scope.",
			Type:        schema.TypeString,
		},
	}
}

func gitlabProjectFeatureFlagStrategySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: "The ID of the strategy.",
			Type:        schema.TypeString,
		},
		"name": {
			Description: "The strategy name.",
			Type:        schema.TypeString,
		},
		"parameters": {
			Description: "Parameters given to the strategy",
			Type:        schema.TypeSet,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"user_ids": {
						Description: "The UserIds strategy type",
						Type:        schema.TypeString,
					},
					"group_id": {
						Description: "The GroupId strategy type",
						Type:        schema.TypeString,
					},
					"percentage": {
						Description: "The Percentage strategy type",
						Type:        schema.TypeString,
					},
				},
			},
		},
		"scopes": {
			Description: "The scopes for the strategy.",
			Type:        schema.TypeList,
			Elem: &schema.Resource{
				Schema: datasourceSchemaFromResourceSchema(gitlabProjectFeatuerFlagScopeSchema(), nil, nil),
			},
		},
	}
}

func gitlabProjectFeatureFlagToStateMap(variable *gitlab.ProjectFeatureFlag) map[string]interface{} {
	stateMap := make(map[string]interface{})
	stateMap["name"] = variable.Name
	stateMap["feature_flag_name"] = variable.Name
	stateMap["description"] = variable.Description
	stateMap["active"] = variable.Active
	stateMap["version"] = variable.Version

	if variable.CreatedAt != nil {
		stateMap["created_at"] = variable.CreatedAt.String()
	}
	if variable.UpdatedAt != nil {
		stateMap["updated_at"] = variable.UpdatedAt.String()
	}
	if variable.Scopes != nil {
		stateMap["scopes"] = variable.Scopes
	}
	if variable.Strategies != nil {
		stateMap["strategies"] = flattenStrategies(variable.Strategies)
	}

	return stateMap
}

func flattenStrategies(strategies []*gitlab.ProjectFeatureFlagStrategy) []map[string]interface{} {
	if len(strategies) == 0 {
		return []map[string]interface{}{}
	}

	// Map the strategy to the lowercase equivalent - parameters, etc
}
