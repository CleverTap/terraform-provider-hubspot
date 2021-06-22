package hubspot

import (
	"fmt"
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccUserDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.hubspot_user.user1", "id", "thesaurabhsaini@gmail.com"),
				),
			},
		},
	})
}

func testAccUserDataSourceConfig() string {
	return fmt.Sprintf(`	  
	resource "hubspot_user" "user4" {
		email        = "abhishekdon70@gmail.com"
		role_id       = "76891"
	  }
	data "hubspot_user" "user1" {
		id = "thesaurabhsaini@gmail.com"
	}
	`)
}
