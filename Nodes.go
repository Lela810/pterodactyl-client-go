package pterodactyl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetNodes - Returns list of nodes
func (c *Client) GetNodes() ([]Node, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/application/nodes", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	var nodeList NodesResponse
	err = json.Unmarshal(body, &nodeList)
	if err != nil {
		return nil, err
	}

	nodesData := nodeList.Data

	nodes := make([]Node, len(nodesData))
	for i, nodeData := range nodesData {
		nodes[i] = nodeData.Attributes
	}

	return nodes, nil
}

// GetNode - Returns specific node
func (c *Client) GetNode(nodeID int32) (Node, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/application/nodes/%d", c.HostURL, nodeID), nil)
	if err != nil {
		return Node{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return Node{}, err
	}

	var response NodeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Node{}, err
	}

	node := response.Attributes

	return node, nil
}

// GetNodeConfiguration - Returns node configuration
func (c *Client) GetNodeConfiguration(nodeID int32) (NodeConfiguration, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/application/nodes/%d/configuration", c.HostURL, nodeID), nil)
	if err != nil {
		return NodeConfiguration{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return NodeConfiguration{}, err
	}

	var response NodeConfigurationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return NodeConfiguration{}, err
	}

	nodeConfiguration := response.Attributes

	return nodeConfiguration, nil
}

// CreateNode - Creates a new node
func (c *Client) CreateNode(node NodesInterface) (Node, error) {
	createNode := PartialNode{
		Name:               node.GetName(),
		Description:        node.GetDescription(),
		Public:             node.GetPublic(),
		BehindProxy:        node.GetBehindProxy(),
		MaintenanceMode:    node.GetMaintenanceMode(),
		LocationID:         node.GetLocationID(),
		FQDN:               node.GetFQDN(),
		Scheme:             node.GetScheme(),
		Memory:             node.GetMemory(),
		MemoryOverallocate: node.GetMemoryOverallocate(),
		Disk:               node.GetDisk(),
		DiskOverallocate:   node.GetDiskOverallocate(),
		UploadSize:         node.GetUploadSize(),
		DaemonListen:       node.GetDaemonListen(),
		DaemonSFTP:         node.GetDaemonSFTP(),
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/application/nodes", c.HostURL), c.prepareBody(createNode))
	if err != nil {
		return Node{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return Node{}, err
	}

	var response NodeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Node{}, err
	}

	newNode := response.Attributes

	return newNode, nil
}

// UpdateNode - Updates a node
func (c *Client) UpdateNode(nodeID int32, node NodesInterface) (Node, error) {
	updatedNode := PartialNode{
		Name:               node.GetName(),
		Description:        node.GetDescription(),
		Public:             node.GetPublic(),
		BehindProxy:        node.GetBehindProxy(),
		MaintenanceMode:    node.GetMaintenanceMode(),
		LocationID:         node.GetLocationID(),
		FQDN:               node.GetFQDN(),
		Scheme:             node.GetScheme(),
		Memory:             node.GetMemory(),
		MemoryOverallocate: node.GetMemoryOverallocate(),
		Disk:               node.GetDisk(),
		DiskOverallocate:   node.GetDiskOverallocate(),
		UploadSize:         node.GetUploadSize(),
		DaemonListen:       node.GetDaemonListen(),
		DaemonSFTP:         node.GetDaemonSFTP(),
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/application/nodes/%d", c.HostURL, nodeID), c.prepareBody(updatedNode))
	if err != nil {
		return Node{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return Node{}, err
	}

	var response NodeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Node{}, err
	}

	newNode := response.Attributes

	return newNode, nil
}

// DeleteNode - Deletes a node
func (c *Client) DeleteNode(nodeID int32) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/application/nodes/%d", c.HostURL, nodeID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetNodeAllocations - Returns list of allocations added to a node
func (c *Client) GetNodeAllocations(nodeID int32) ([]Allocation, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/application/nodes/%d/allocations", c.HostURL, nodeID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	var allocationList AllocationsResponse
	err = json.Unmarshal(body, &allocationList)
	if err != nil {
		return nil, err
	}

	allocationsData := allocationList.Data

	allocations := make([]Allocation, len(allocationsData))
	for i, allocationData := range allocationsData {
		allocations[i] = allocationData.Attributes
	}

	return allocations, nil
}

// CreateAllocation - Adds an allocation to a node
func (c *Client) CreateAllocation(nodeID int32, allocation PartialAllocation) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/application/nodes/%d/allocations", c.HostURL, nodeID), c.prepareBody(allocation))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}

// DeleteAllocation - Deletes an allocation from a node
func (c *Client) DeleteAllocation(nodeID, allocationID int32) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/application/nodes/%d/allocations/%d", c.HostURL, nodeID, allocationID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}
