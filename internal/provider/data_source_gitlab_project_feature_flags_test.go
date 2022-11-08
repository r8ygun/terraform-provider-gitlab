//go:build acceptance
// +build acceptance

package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitlabProjectFeatureFlag_basic(t *testing.T) {
	project := testAccCreateProject(t)
	flagName := fmt.Sprintf("feature_flag_name_%d", acctest.RandInt())
	testFlag := testAccCreateProjectFeatureFlag(t, project.ID, flagName)

	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitlabProjectFeatureFlagConfig_basic(project.ID, flagName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.gitlab_project_feature_flag.foo", "name", testFlag.Name),
					resource.TestCheckResourceAttr("data.gitlab_project_feature_flag.foo", "description", testFlag.Description),
					resource.TestCheckResourceAttr("data.gitlab_project_feature_flag.foo", "active", strconv.FormatBool(testFlag.Active)),
				),
			},
		},
	})
}

func testAccDataSourceGitlabProjectFeatureFlagConfig_basic(projectId int, flagName string) string {
	return fmt.Sprintf(`
data "gitlab_project_feature_flag" "foo" {
	id = "%d"
	feature_flag_name = "%s"
}`, projectId, flagName)
}
