//go:build acceptance
// +build acceptance

package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/xanzy/go-gitlab"
)

func TestAccDataSourceGitlabProjectFeatureFlag_basic(t *testing.T) {
	project := testAccCreateProject(t)
	flagName := fmt.Sprintf("feature_flag_name_%d", acctest.RandInt())
	scopes := make([]*gitlab.ProjectFeatureFlagScope, 0)
	scopes = append(scopes, &gitlab.ProjectFeatureFlagScope{
		ID:               acctest.RandInt(),
		EnvironmentScope: "production",
	})

	testFlag := testAccCreateProjectFeatureFlag(t, project.ID,
		withFlagName(flagName),
		withFlagDescription(acctest.RandString(30)),
		withFlagActive(acctest.RandInt()%2 == 0),
		withFlagVersion("new_version_flag"),
		withFlagStrategy(&gitlab.FeatureFlagStrategyOptions{
			Name: gitlab.String("userWithId"),
			Parameters: &gitlab.ProjectFeatureFlagStrategyParameter{
				UserIDs: "10",
			},
			Scopes: &scopes,
		}),
	)

	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitlabProjectFeatureFlagConfig_basic(project.ID, flagName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.gitlab_project_feature_flag.foo",
						"name",
						testFlag.Name,
					),
					resource.TestCheckResourceAttr(
						"data.gitlab_project_feature_flag.foo",
						"description",
						testFlag.Description,
					),
					resource.TestCheckResourceAttr(
						"data.gitlab_project_feature_flag.foo",
						"active",
						strconv.FormatBool(testFlag.Active),
					),
					resource.TestCheckResourceAttr(
						"data.gitlab_project_feature_flag.foo",
						"version",
						testFlag.Version,
					),
					// resource.TestCheckResourceAttr(
					// 	"data.gitlab_project_feature_flag.foo",
					// 	"strategies.#",
					// 	fmt.Sprintf("%d", len(testFlag.Strategies)),
					// ),
					// resource.TestCheckResourceAttr(
					// 	"data.gitlab_project_feature_flag.foo",
					// 	"strategies.0.name",
					// 	testFlag.Strategies[0].Name,
					// ),
					// resource.TestCheckResourceAttr(
					// 	"data.gitlab_project_feature_flag.foo",
					// 	"strategies.0.scopes.#",
					// 	fmt.Sprintf("%d", len(testFlag.Strategies[0].Scopes)),
					// ),
					// resource.TestCheckResourceAttr(
					// 	"data.gitlab_project_feature_flag.foo",
					// 	"strategies.0.scopes.0.id",
					// 	fmt.Sprintf("%d", testFlag.Strategies[0].Scopes[0].ID),
					// ),
					// resource.TestCheckResourceAttr(
					// 	"data.gitlab_project_feature_flag.foo",
					// 	"strategies.0.scopes.0.environment_scope",
					// 	testFlag.Strategies[0].Scopes[0].EnvironmentScope,
					// ),
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
