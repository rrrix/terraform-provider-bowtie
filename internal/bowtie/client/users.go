package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type BowtieUser struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	AuthzDevices      bool   `json:"authz_devices"`
	AuthzPolicies     bool   `json:"authz_policies"`
	AuthzControlPanel bool   `json:"authz_control_plane"`
	AuthzUsers        bool   `json:"authz_users"`
	Status            string `json:"status"`
}

func (c *Client) GetUsers() (map[string]BowtieUser, error) {
	req, err := http.NewRequest(http.MethodGet, c.getHostURL("/users"), nil)
	if err != nil {
		return nil, err
	}

	responseBody, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var users map[string]BowtieUser = map[string]BowtieUser{}
	err = json.Unmarshal(responseBody, &users)
	return users, err
}

func (c *Client) GetUserByEmail(ctx context.Context, email string) (BowtieUser, error) {
	users, err := c.GetUsers()
	tflog.Info(ctx, fmt.Sprintf("%+v", users))
	if err != nil {
		return BowtieUser{}, err
	}

	for _, user := range users {
		fmt.Printf("%s - %s\n\n", user.Email, email)
		if user.Email == strings.TrimSpace(email) {
			return user, nil
		}
	}

	return BowtieUser{}, fmt.Errorf("user not found")
}

func (c *Client) DeleteUser(id string) error {
	req, err := http.NewRequest(http.MethodDelete, c.getHostURL(fmt.Sprintf("/user/%s", id)), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

func (c *Client) DisableUser(id string) error {
	payload := BowtieUser{
		Status: "Disabled",
		ID:     id,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, c.getHostURL("/user/upsert"), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}
