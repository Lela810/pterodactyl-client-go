package pterodactyl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetLocations - Returns list of locations
func (c *Client) GetLocations() ([]Location, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/application/locations", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	var locationList LocationsResponse
	err = json.Unmarshal(body, &locationList)
	if err != nil {
		return nil, err
	}

	locationsData := locationList.Data

	locations := make([]Location, len(locationsData))
	for i, locationData := range locationsData {
		locations[i] = locationData.Attributes
	}

	return locations, nil
}

// GetLocation - Returns information about a specific location
func (c *Client) GetLocation(locationID int32) (Location, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/application/locations/%d", c.HostURL, locationID), nil)
	if err != nil {
		return Location{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return Location{}, err
	}

	var response LocationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Location{}, err
	}

	location := response.Attributes

	return location, nil
}

// CreateLocation - Creates a new location
func (c *Client) CreateLocation(location LocationInterface) (Location, error) {
	partialLocation := PartialLocation{
		Short: location.GetShort(),
		Long:  location.GetLong(),
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/application/locations", c.HostURL), c.prepareBody(partialLocation))
	if err != nil {
		return Location{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return Location{}, err
	}

	var response LocationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Location{}, err
	}

	newLocation := response.Attributes

	return newLocation, nil
}

// UpdateLocation - Updates a location
func (c *Client) UpdateLocation(locationID int32, location LocationInterface) (Location, error) {
	partialLocation := PartialLocation{
		Short: location.GetShort(),
		Long:  location.GetLong(),
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/application/locations/%d", c.HostURL, locationID), c.prepareBody(partialLocation))
	if err != nil {
		return Location{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return Location{}, err
	}

	var response LocationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Location{}, err
	}

	updatedLocation := response.Attributes

	return updatedLocation, nil
}

// DeleteLocation - Deletes a location
func (c *Client) DeleteLocation(locationID int32) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/application/locations/%d", c.HostURL, locationID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}
