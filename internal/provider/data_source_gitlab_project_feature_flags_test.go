//go:build acceptance
// +build acceptance

package provider

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitlabProjectFeatureFlag_basic(t *testing.T) {
	project := testAccCreateProject(t)
	testFlag := testAccCreateProjectFeatureFlag(t, project.ID)

	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `data "gitlab_project_feature_flag" "this"  {}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.gitlab_project_feature_flag.this", "Name", testFlag.Name),
					resource.TestCheckResourceAttr("data.gitlab_project_feature_flag.this", "Description", testFlag.Description),
					resource.TestCheckResourceAttr("data.gitlab_project_feature_flag.this", "Active", strconv.FormatBool(testFlag.Active)),
				),
			},
		},
	})

}
