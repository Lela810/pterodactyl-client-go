package pterodactyl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetUsers - Returns list of users
func (c *Client) GetUsers() ([]User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/application/users", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	var userList UsersResponse
	err = json.Unmarshal(body, &userList)
	if err != nil {
		return nil, err
	}

	users := userList.Data

	return users, nil
}

// GetUser - Returns specific user
func (c *Client) GetUser(userID string) (User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/application/users/%s", c.HostURL, userID), nil)
	if err != nil {
		return User{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return User{}, err
	}

	var response UserResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return User{}, err
	}

	user := response.Attributes

	return user, nil
}

// GetUserExternalID - Returns specific user by external ID
func (c *Client) GetUserExternalID(externalID string) (User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/application/users/external/%s", c.HostURL, externalID), nil)
	if err != nil {
		return User{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return User{}, err
	}

	var response UserResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return User{}, err
	}

	user := response.Attributes

	return user, nil
}

// CreateUser - Create new user
func (c *Client) CreateUser(newUser User) (User, error) {
	marshalled_user, err := json.Marshal(User{Email: newUser.Email, Username: newUser.Username, FirstName: newUser.FirstName, LastName: newUser.LastName})
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/application/users", c.HostURL), bytes.NewReader(marshalled_user))
	if err != nil {
		return User{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return User{}, err
	}

	var response UserResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return User{}, err
	}

	user := response.Attributes

	return user, nil
}

// UpdateUser - Update user
func (c *Client) UpdateUser(userID string, updatedUser User) (User, error) {
	marshalled_user, err := json.Marshal(User{Email: updatedUser.Email, Username: updatedUser.Username, FirstName: updatedUser.FirstName, LastName: updatedUser.LastName})
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/application/users/%s", c.HostURL, userID), bytes.NewReader(marshalled_user))
	if err != nil {
		return User{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return User{}, err
	}

	var response UserResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return User{}, err
	}

	user := response.Attributes

	return user, nil
}

// DeleteUser - Delete user
func (c *Client) DeleteUser(userID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/application/users/%s", c.HostURL, userID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}
