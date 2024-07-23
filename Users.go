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

	userResponses := userList.Data

	users := make([]User, len(userResponses))
	for i, userResponse := range userResponses {
		users[i] = userResponse.Attributes
	}

	return users, nil
}

// GetUser - Returns specific user
func (c *Client) GetUser(userID int) (User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/application/users/%d", c.HostURL, userID), nil)
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
func (c *Client) CreateUser(newUser PartialUser) (User, error) {
	marshalled_user, err := json.Marshal(PartialUser{Email: newUser.Email, Username: newUser.Username, FirstName: newUser.FirstName, LastName: newUser.LastName})
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
func (c *Client) UpdateUser(userID int, userInterface UserInterface) (User, error) {
	marshalled_user, err := json.Marshal(PartialUser{Email: userInterface.GetEmail(), Username: userInterface.GetUsername(), FirstName: userInterface.GetFirstName(), LastName: userInterface.GetLastName()})
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/application/users/%d", c.HostURL, userID), bytes.NewReader(marshalled_user))
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
func (c *Client) DeleteUser(userID int) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/application/users/%d", c.HostURL, userID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}
