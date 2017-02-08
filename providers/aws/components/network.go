/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package components

import "github.com/r3labs/graph"

// Network : Mapping of a network component
type Network struct {
	ProviderType     string            `json:"_type"`
	ComponentType    string            `json:"_component"`
	State            string            `json:"_state"`
	Action           string            `json:"_action"`
	NetworkAWSID     string            `json:"network_aws_id"`
	Name             string            `json:"name"`
	Subnet           string            `json:"range"`
	IsPublic         bool              `json:"is_public"`
	Tags             map[string]string `json:"tags"`
	AvailabilityZone string            `json:"availability_zone"`
	DatacenterType   string            `json:"datacenter_type"`
	DatacenterName   string            `json:"datacenter_name"`
	DatacenterRegion string            `json:"datacenter_region"`
	AccessKeyID      string            `json:"aws_access_key_id"`
	SecretAccessKey  string            `json:"aws_secret_access_key"`
	VpcID            string            `json:"vpc_id"`
	Service          string            `json:"service"`
	Status           string            `json:"status"`
	Exists           bool
}

// GetID : returns the component's ID
func (n *Network) GetID() string {
	return n.ComponentType + "::" + n.Name
}

// GetName returns a components name
func (n Network) GetName() string {
	return n.Name
}

func (n Network) GetProvider() string {
	return n.ProviderType
}

// GetProviderID returns a components provider id
func (n *Network) GetProviderID() string {
	return n.NetworkAWSID
}

func (n Network) GetType() string {
	return n.ComponentType
}

func (n Network) GetState() string {
	return n.State
}

func (n Network) SetState(s string) {
	n.State = s
}

func (n Network) GetAction() string {
	return n.Action
}

func (n Network) SetAction(s string) {
	n.Action = s
}

func (n Network) GetGroup() string {
	return ""
}

// GetTags returns a components tags
func (n Network) GetTags() map[string]string {
	return n.Tags
}

func (n Network) Diff(c graph.Vertex) {

}

func (n *Network) Update(c graph.Vertex) bool {
	cn := c.(Network)

	return false
}

func (n *Network) Rebuild() {

}

func (n Network) Dependencies() []string {
	return []string{}
}

func (n *Network) IsStateful() bool {
	return true
}
