package hubspot

import (
	"fmt"
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccUser_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckUserBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("hubspot_user.user1", "email", "saurabh.saini@clevertap.com"),
					resource.TestCheckResourceAttr("hubspot_user.user1", "role_id", "76891"),
				),
			},
		},
	})
}

func testAccCheckUserBasic() string {
	return fmt.Sprintf(`
	resource "hubspot_user" "user1" {
		email  = "saurabh.saini@clevertap.com"
		role_id = "76891"
	}
	`)
}

func TestAccUser_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckUserUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("hubspot_user.user1", "email", "saurabh.saini@clevertap.com"),
					resource.TestCheckResourceAttr("hubspot_user.user1", "role_id", "76891"),
				),
			},
			{
				Config: testAccCheckUserUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("hubspot_user.user1", "email", "saurabh.saini@clevertap.com"),
					resource.TestCheckResourceAttr("hubspot_user.user1", "role_id", "76894"),
				),
			},
		},
	})
}

func testAccCheckUserUpdatePre() string {
	return fmt.Sprintf(`
	resource "hubspot_user" "user1" {
		email  = "saurabh.saini@clevertap.com"
		role_id = "76891"
	}
	`)
}

func testAccCheckUserUpdatePost() string {
	return fmt.Sprintf(`
	resource "hubspot_user" "user1" {
		email  = "saurabh.saini@clevertap.com"
		role_id = "76894"
	}
	`)
}
