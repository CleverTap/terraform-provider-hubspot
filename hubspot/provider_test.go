package hubspot

import (
	"log"
	"os"
	"terraform-provider-hubspot/token"
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	clientId := os.Getenv("HUBSPOT_CLIENT_ID")
	clientSecret := os.Getenv("HUBSPOT_CLIENT_SECRET")
	refreshToken := os.Getenv("HUBSPOT_REFRESH_TOKEN")
	accessToken := token.GenerateToken(clientId, clientSecret, refreshToken)
	os.Setenv("HUBSPOT_TOKEN", string(accessToken))
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"hubspot": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		log.Println("[ERROR]: ", err)
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("HUBSPOT_TOKEN"); v == "" {
		t.Fatal("HUBSPOT_TOKEN must be set for acceptance tests")
	}
}
