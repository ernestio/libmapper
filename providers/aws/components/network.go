/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package components

import (
	"errors"
	"net"
	"reflect"

	graph "gopkg.in/r3labs/graph.v2"
)

// Network : Mapping of a network component
type Network struct {
	ProviderType     string            `json:"_provider"`
	ComponentType    string            `json:"_component"`
	ComponentID      string            `json:"_component_id"`
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
	Vpc              string            `json:"vpc"`
	VpcID            string            `json:"vpc_id"`
	Service          string            `json:"service"`
}

// GetID : returns the component's ID
func (n *Network) GetID() string {
	return n.ComponentID
}

// GetName returns a components name
func (n *Network) GetName() string {
	return n.Name
}

// GetProvider : returns the provider type
func (n *Network) GetProvider() string {
	return n.ProviderType
}

// GetProviderID returns a components provider id
func (n *Network) GetProviderID() string {
	return n.NetworkAWSID
}

// GetType : returns the type of the component
func (n *Network) GetType() string {
	return n.ComponentType
}

// GetState : returns the state of the component
func (n *Network) GetState() string {
	return n.State
}

// SetState : sets the state of the component
func (n *Network) SetState(s string) {
	n.State = s
}

// GetAction : returns the action of the component
func (n *Network) GetAction() string {
	return n.Action
}

// SetAction : Sets the action of the component
func (n *Network) SetAction(s string) {
	n.Action = s
}

// GetGroup : returns the components group
func (n *Network) GetGroup() string {
	return ""
}

// GetTags returns a components tags
func (n *Network) GetTags() map[string]string {
	return n.Tags
}

// GetTag returns a components tag
func (n *Network) GetTag(tag string) string {
	return n.Tags[tag]
}

// Diff : diff's the component against another component of the same type
func (n *Network) Diff(c graph.Component) bool {
	cn, ok := c.(*Network)
	if ok {
		return !reflect.DeepEqual(n.Tags, cn.Tags)
	}

	return false
}

// Update : updates the provider returned values of a component
func (n *Network) Update(c graph.Component) {
	cn, ok := c.(*Network)
	if ok {
		n.NetworkAWSID = cn.NetworkAWSID
		n.AvailabilityZone = cn.AvailabilityZone
	}

	n.SetDefaultVariables()
}

// Rebuild : rebuilds the component's internal state, such as templated values
func (n *Network) Rebuild(g *graph.Graph) {
	if n.Vpc == "" && n.VpcID != "" {
		v := g.GetComponents().ByProviderID(n.VpcID)
		if v != nil {
			n.Vpc = v.GetName()
		}
	}

	if n.Vpc != "" && n.VpcID == "" {
		n.VpcID = templVpcID(n.Vpc)
	}

	n.SetDefaultVariables()
}

// Dependencies : returns a list of component id's upon which the component depends
func (n *Network) Dependencies() []string {
	return []string{"vpc::" + n.Vpc}
}

// Validate : validates the components values
func (n *Network) Validate() error {
	_, _, err := net.ParseCIDR(n.Subnet)
	if err != nil {
		return errors.New("Network CIDR is not valid")
	}

	if n.Name == "" {
		return errors.New("Network name should not be null")
	}

	if n.IsPublic && n.Tags["ernest.nat_gateway"] != "" {
		return errors.New("Public Network should not specify a nat gateway")
	}

	return nil
}

// IsStateful : returns true if the component needs to be actioned to be removed.
func (n *Network) IsStateful() bool {
	return true
}

// SetDefaultVariables : sets up the default template variables for a component
func (n *Network) SetDefaultVariables() {
	n.ComponentType = TYPENETWORK
	n.ComponentID = TYPENETWORK + TYPEDELIMITER + n.Name
	n.ProviderType = PROVIDERTYPE
	n.DatacenterName = DATACENTERNAME
	n.DatacenterType = DATACENTERTYPE
	n.DatacenterRegion = DATACENTERREGION
	n.AccessKeyID = ACCESSKEYID
	n.SecretAccessKey = SECRETACCESSKEY
}
