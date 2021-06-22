package client

import (
	"log"
	"os"
	"terraform-provider-hubspot/token"
	"testing"
	"github.com/stretchr/testify/assert"
)

func init() {
	clientId := os.Getenv("HUBSPOT_CLIENT_ID")
	clientSecret := os.Getenv("HUBSPOT_CLIENT_SECRET")
	refreshToken := os.Getenv("HUBSPOT_REFRESH_TOKEN")
	accessToken := token.GenerateToken(clientId, clientSecret, refreshToken)
	os.Setenv("HUBSPOT_TOKEN", accessToken)
}

func TestClient_GetUser(t *testing.T) {
	testCases := []struct {
		testName     string
		userName     string
		seedData     map[string]User
		expectErr    bool
		expectedResp *User
	}{
		{
			testName: "user exists",
			userName: "thesaurabhsaini@gmail.com",
			seedData: map[string]User{
				"user1": {
					Id:     "24791265",
					Email:  "thesaurabhsaini@gmail.com",
					RoleId: "76894",
				},
			},
			expectErr: false,
			expectedResp: &User{
				Id:     "24791265",
				Email:  "thesaurabhsaini@gmail.com",
				RoleId: "76894",
			},
		},
		{
			testName:     "user does not exist",
			userName:     "saurabh.saini@clevertap.com",
			seedData:     nil,
			expectErr:    true,
			expectedResp: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			token := os.Getenv("HUBSPOT_TOKEN")
			client := NewClient(token)
			user, err := client.GetUser(tc.userName)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, user)
		})
	}
}

func TestClient_CreateUser(t *testing.T) {
	testCases := []struct {
		testName  string
		newUser   *User
		seedData  map[string]User
		expectErr bool
	}{
		{
			testName: "success",
			newUser: &User{
				Id:     "24813958",
				Email:  "ravikishandaiya@gmail.com",
				RoleId: "76894",
			},
			seedData:  nil,
			expectErr: false,
		},
		{
			testName: "user already exists",
			newUser: &User{
				Id:     "24813958",
				Email:  "ravikishandaiya@gmail.com",
				RoleId: "76894",
			},
			seedData: map[string]User{
				"user1": {
					Id:     "24813958",
					Email:  "ravikishandaiya@gmail.com",
					RoleId: "76894",
				},
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			token := os.Getenv("HUBSPOT_TOKEN")
			client := NewClient(token)
			err := client.CreateUser(tc.newUser)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestClient_UpdateUser(t *testing.T) {
	testCases := []struct {
		testName    string
		updatedUser *User
		seedData    map[string]User
		expectErr   bool
	}{
		{
			testName: "user exists",
			updatedUser: &User{
				Id:     "24813958",
				Email:  "ravikishandaiya@gmail.com",
				RoleId: "76894",
			},
			seedData: map[string]User{
				"user1": {
					Id:     "24813958",
					Email:  "ravikishandaiya@gmail.com",
					RoleId: "76891",
				},
			},
			expectErr: false,
		},
		{
			testName: "user does not exist",
			updatedUser: &User{
				Id:     "24813958",
				Email:  "saurabh.saini@clevertap.com",
				RoleId: "76891",
			},
			seedData:  nil,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			token := os.Getenv("HUBSPOT_TOKEN")
			client := NewClient(token)
			err := client.UpdateUser(tc.updatedUser)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			user, err := client.GetUser(tc.updatedUser.Email)
			assert.NoError(t, err)
			assert.Equal(t, tc.updatedUser, user)
		})
	}
}

func TestClient_DeleteUser(t *testing.T) {
	testCases := []struct {
		testName  string
		userName  string
		seedData  map[string]User
		expectErr bool
	}{
		{
			testName: "user exists",
			userName: "ravikishandaiya@gmail.com",
			seedData: map[string]User{
				"user1": {
					Id:     "24813958",
					Email:  "ravikishandaiya@gmail.com",
					RoleId: "76891",
				},
			},
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			token := os.Getenv("HUBSPOT_TOKEN")
			client := NewClient(token)
			_, err := client.GetUser(tc.userName)
			if err != nil {
				assert.NoError(t, err)
				return
			}
			err = client.DeleteUser(tc.userName)
			if tc.expectErr {
				log.Println("[DELETE ERROR]: ", err)
				assert.Error(t, err)
				return
			}
		})
	}
}
