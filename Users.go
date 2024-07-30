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

// GetUsersWithFilter - Returns list of users with filter
func (c *Client) GetUsersWithFilter(filterBy string, filterContent string) ([]User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/application/users?filter[%s]=%s", c.HostURL, filterBy, filterContent), nil)
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
func (c *Client) GetUser(userID int32) (User, error) {
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

// GetUserEmail - Returns specific user by email
func (c *Client) GetUserEmail(email string) (User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/application/users/", c.HostURL), nil)
	if err != nil {
		return User{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return User{}, err
	}

	var userList UsersResponse
	err = json.Unmarshal(body, &userList)
	if err != nil {
		return User{}, err
	}

	userResponses := userList.Data

	for _, userResponse := range userResponses {
		if userResponse.Attributes.Email == email {
			return userResponse.Attributes, nil
		}
	}

	return User{}, fmt.Errorf("User with email %s not found", email)
}

// GetUserUsername - Returns specific user by username
func (c *Client) GetUserUsername(username string) (User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/application/users/", c.HostURL), nil)
	if err != nil {
		return User{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return User{}, err
	}

	var userList UsersResponse
	err = json.Unmarshal(body, &userList)
	if err != nil {
		return User{}, err
	}

	userResponses := userList.Data

	for _, userResponse := range userResponses {
		if userResponse.Attributes.Username == username {
			return userResponse.Attributes, nil
		}
	}

	return User{}, fmt.Errorf("User with username %s not found", username)
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
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/application/users", c.HostURL), c.prepareBody(newUser))
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
func (c *Client) UpdateUser(userID int32, userInterface UserInterface) (User, error) {
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
func (c *Client) DeleteUser(userID int32) error {
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
