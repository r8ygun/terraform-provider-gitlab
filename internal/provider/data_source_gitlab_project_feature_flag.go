package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/xanzy/go-gitlab"
)

var _ = registerDataSource("gitlab_project_feature_flag", func() *schema.Resource {
	return &schema.Resource{
		Description: `The ` + "`gitlab_project_feature_flag`" + ` data source allows to retrieve details about an instance-level CI/CD variable.

		**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/feature_flags.html)`,

		ReadContext: dataSourceGitlabProjectFeatureFlagsRead,
		Schema:      datasourceSchemaFromResourceSchema(gitlabProjectFeatureFlagGetSchema(), []string{"id", "feature_flag_name"}, nil),
	}
})

func dataSourceGitlabProjectFeatureFlagsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)
	projectId := d.Get("id").(string)
	feature_flag_name := d.Get("feature_flag_name").(string)

	featureFlag, _, err := client.ProjectFeatureFlags.GetProjectFeatureFlag(projectId, feature_flag_name, gitlab.WithContext(ctx))

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s:%s", projectId, feature_flag_name))
	stateMap := gitlabProjectFeatureFlagToStateMap(featureFlag)

	if err := setStateMapInResourceData(stateMap, d); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
