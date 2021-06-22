package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const HostURL string = "https://api.hubapi.com"

type User struct {
	Id     string `json:"id"`
	Email  string `json:"email"`
	RoleId string `json:"roleId"`
}

type CreateUserRequestWithRole struct {
	Email            string `json:"email"`
	RoleId           string `json:"roleId"`
	SendWelcomeEmail bool   `json:"sendWelcomeEmail"`
}

type CreateUserRequestWithNoRole struct {
	Email            string `json:"email"`
	SendWelcomeEmail bool   `json:"sendWelcomeEmail"`
}

type UpdateUserRequest struct {
	RoleId string `json:"roleId"`
}

var (
	Errors = make(map[int]string)
)

func init() {
	Errors[400] = "Bad Request, StatusCode = 400"
	Errors[404] = "User Does Not Exist , StatusCode = 404"
	Errors[409] = "User Already Exist, StatusCode = 409"
	Errors[401] = "Unauthorized Access, StatusCode = 401"
	Errors[429] = "User Has Sent Too Many Request, StatusCode = 429"
}

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

func NewClient(token string) *Client {
	c := Client{
		HTTPClient: &http.Client{},
		HostURL:    HostURL,
		Token:      token,
	}
	return &c
}

func (c *Client) GetUser(userId string) (*User, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/settings/v3/users/%s?idProperty=EMAIL", c.HostURL, userId), nil)
	if err != nil {
		log.Println("[READ ERROR]: ", err)
		return nil, err
	}
	request.Header.Add("Authorization", "Bearer "+c.Token)
	request.Header.Add("Accept", "application/json")
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		log.Println("[READ ERROR]: ", err)
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("READ ERROR : %v", Errors[response.StatusCode])
	}
	user := &User{}
	err = json.NewDecoder(response.Body).Decode(user)
	if err != nil {
		log.Println("[READ ERROR]: ", err)
		return nil, err
	}
	return user, nil
}

func (c *Client) CreateUser(user *User) error {
	if user.RoleId == "" {
		createUserRequest := CreateUserRequestWithNoRole{
			Email:            user.Email,
			SendWelcomeEmail: true,
		}
		reqjson, err := json.Marshal(createUserRequest)
		if err != nil {
			log.Println("[CREATE ERROR]: ", err)
			return err
		}
		request, err := http.NewRequest("POST", fmt.Sprintf("%s/settings/v3/users/", c.HostURL), strings.NewReader(string(reqjson)))
		if err != nil {
			log.Println("[CREATE ERROR]: ", err)
			return err
		}
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Authorization", "Bearer "+c.Token)
		request.Header.Add("Accept", "application/json")
		response, err := c.HTTPClient.Do(request)
		if err != nil {
			log.Println("[CREATE ERROR]: ", err)
			return err
		}
		if response.StatusCode >= 200 && response.StatusCode <= 299 {
			return nil
		} else {
			return fmt.Errorf("CREATE ERROR : %v", Errors[response.StatusCode])
		}
	} else {
		createUserRequest := CreateUserRequestWithRole{
			Email:            user.Email,
			RoleId:           user.RoleId,
			SendWelcomeEmail: true,
		}
		reqjson, err := json.Marshal(createUserRequest)
		if err != nil {
			log.Println("[CREATE ERROR]: ", err)
			return err
		}
		request, err := http.NewRequest("POST", fmt.Sprintf("%s/settings/v3/users/", c.HostURL), strings.NewReader(string(reqjson)))
		if err != nil {
			log.Println("[CREATE ERROR]: ", err)
			return err
		}
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Authorization", "Bearer "+c.Token)
		request.Header.Add("Accept", "application/json")
		response, err := c.HTTPClient.Do(request)
		if err != nil {
			log.Println("[CREATE ERROR]: ", err)
			return err
		}
		if response.StatusCode >= 200 && response.StatusCode <= 299 {
			return nil
		} else {
			return fmt.Errorf("CREATE ERROR : %v", Errors[response.StatusCode])
		}
	}
}

func (c *Client) UpdateUser(user *User) error {
	updateUserRequest := UpdateUserRequest{
		RoleId: user.RoleId,
	}
	updatejson, err := json.Marshal(updateUserRequest)
	if err != nil {
		log.Println("[UPDATE ERROR]: ", err)
		return err
	}
	request, err := http.NewRequest("PUT", fmt.Sprintf("%s/settings/v3/users/%s?idProperty=EMAIL", c.HostURL, user.Email), strings.NewReader(string(updatejson)))
	if err != nil {
		log.Println("[UPDATE ERROR]: ", err)
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+c.Token)
	request.Header.Add("Accept", "application/json")
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		log.Println("[UPDATE ERROR]: ", err)
		return err
	}
	if response.StatusCode >= 200 && response.StatusCode < 400 {
		return nil
	} else {
		return fmt.Errorf("UPDATE Error : %v", Errors[response.StatusCode])
	}
}

func (c *Client) DeleteUser(userId string) error {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/settings/v3/users/%s?idProperty=EMAIL", c.HostURL, userId), nil)
	if err != nil {
		log.Println("[DELETE ERROR]: ", err)
		return err
	}
	request.Header.Add("Authorization", "Bearer "+c.Token)
	request.Header.Add("Accept", "application/json")
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		log.Println("[DELETE ERROR]: ", err)
		return err
	}
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		return nil
	} else {
		log.Println("Broken Request")
		return fmt.Errorf("DELETE ERROR : %v", Errors[response.StatusCode])
	}
}

func (c *Client) IsRetry(err error) bool {
	if err != nil {
		if strings.Contains(err.Error(), "429") == true {
			return true
		}
	}
	return false
}
