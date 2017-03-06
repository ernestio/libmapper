/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package components

import (
	"errors"
	"reflect"

	graph "gopkg.in/r3labs/graph.v2"
)

// NatGateway : mapping of a nat component
type NatGateway struct {
	ProviderType           string            `json:"_provider"`
	ComponentType          string            `json:"_component"`
	ComponentID            string            `json:"_component_id"`
	State                  string            `json:"_state"`
	Action                 string            `json:"_action"`
	NatGatewayAWSID        string            `json:"nat_gateway_aws_id"`
	Name                   string            `json:"name"`
	PublicNetwork          string            `json:"public_network"`
	RoutedNetworks         []string          `json:"routed_networks"`
	RoutedNetworkAWSIDs    []string          `json:"routed_networks_aws_ids"`
	PublicNetworkAWSID     string            `json:"public_network_aws_id"`
	NatGatewayAllocationID string            `json:"nat_gateway_allocation_id"`
	NatGatewayAllocationIP string            `json:"nat_gateway_allocation_ip"`
	DatacenterType         string            `json:"datacenter_type"`
	DatacenterName         string            `json:"datacenter_name"`
	DatacenterRegion       string            `json:"datacenter_region"`
	AccessKeyID            string            `json:"aws_access_key_id"`
	SecretAccessKey        string            `json:"aws_secret_access_key"`
	Tags                   map[string]string `json:"tags"`
	Service                string            `json:"service"`
}

// GetID : returns the component's ID
func (n *NatGateway) GetID() string {
	return n.ComponentID
}

// GetName returns a components name
func (n *NatGateway) GetName() string {
	return n.Name
}

// GetProvider : returns the provider type
func (n *NatGateway) GetProvider() string {
	return n.ProviderType
}

// GetProviderID returns a components provider id
func (n *NatGateway) GetProviderID() string {
	return n.NatGatewayAWSID
}

// GetType : returns the type of the component
func (n *NatGateway) GetType() string {
	return n.ComponentType
}

// GetState : returns the state of the component
func (n *NatGateway) GetState() string {
	return n.State
}

// SetState : sets the state of the component
func (n *NatGateway) SetState(s string) {
	n.State = s
}

// GetAction : returns the action of the component
func (n *NatGateway) GetAction() string {
	return n.Action
}

// SetAction : Sets the action of the component
func (n *NatGateway) SetAction(s string) {
	n.Action = s
}

// GetGroup : returns the components group
func (n *NatGateway) GetGroup() string {
	return ""
}

// GetTags returns a components tags
func (n *NatGateway) GetTags() map[string]string {
	return n.Tags
}

// GetTag returns a components tag
func (n *NatGateway) GetTag(tag string) string {
	return n.Tags[tag]
}

// Diff : diff's the component against another component of the same type
func (n *NatGateway) Diff(c graph.Component) bool {
	cn, ok := c.(*NatGateway)
	if ok {
		return !reflect.DeepEqual(n.RoutedNetworks, cn.RoutedNetworks)
	}

	return false
}

// Update : updates the provider returned values of a component
func (n *NatGateway) Update(c graph.Component) {
	cn, ok := c.(*NatGateway)
	if ok {
		n.NatGatewayAWSID = cn.NatGatewayAWSID
		n.NatGatewayAllocationID = cn.NatGatewayAllocationID
		n.NatGatewayAllocationIP = cn.NatGatewayAllocationIP
	}

	n.SetDefaultVariables()
}

// Rebuild : rebuilds the component's internal state, such as templated values
func (n *NatGateway) Rebuild(g *graph.Graph) {
	if n.PublicNetwork == "" && n.PublicNetworkAWSID != "" {
		pn := g.GetComponents().ByProviderID(n.PublicNetworkAWSID)
		if pn != nil {
			n.PublicNetwork = pn.GetName()
		}
	}

	if n.PublicNetworkAWSID != "" && n.PublicNetwork != "" {
		n.PublicNetworkAWSID = templSubnetID(n.PublicNetwork)
	}

	if len(n.RoutedNetworks) > len(n.RoutedNetworkAWSIDs) {
		for _, nw := range n.RoutedNetworks {
			n.RoutedNetworkAWSIDs = append(n.RoutedNetworkAWSIDs, templSubnetID(nw))
		}
	}

	if len(n.RoutedNetworkAWSIDs) > len(n.RoutedNetworks) {
		for _, nwid := range n.RoutedNetworkAWSIDs {
			nw := g.GetComponents().ByProviderID(nwid)
			if nw != nil {
				n.RoutedNetworks = append(n.RoutedNetworks, nw.GetName())
			}
		}
	}

	n.SetDefaultVariables()
}

// Dependencies : returns a list of component id's upon which the component depends
func (n *NatGateway) Dependencies() []string {
	var deps []string

	for _, nw := range n.RoutedNetworkAWSIDs {
		deps = append(deps, TYPENETWORK+TYPEDELIMITER+nw)
	}

	deps = append(deps, TYPENETWORK+TYPEDELIMITER+n.PublicNetwork)

	return deps
}

// Validate : validates the components values
func (n *NatGateway) Validate() error {
	if n.Name == "" {
		return errors.New("Nat Gateway name should not be null")
	}

	if n.PublicNetwork == "" {
		return errors.New("Nat Gateway should specify a public network")
	}

	return nil
}

// IsStateful : returns true if the component needs to be actioned to be removed.
func (n *NatGateway) IsStateful() bool {
	return true
}

// SetDefaultVariables : sets up the default template variables for a component
func (n *NatGateway) SetDefaultVariables() {
	n.ComponentType = TYPENATGATEWAY
	n.ComponentID = TYPENATGATEWAY + TYPEDELIMITER + n.Name
	n.ProviderType = PROVIDERTYPE
	n.DatacenterName = DATACENTERNAME
	n.DatacenterType = DATACENTERTYPE
	n.DatacenterRegion = DATACENTERREGION
	n.AccessKeyID = ACCESSKEYID
	n.SecretAccessKey = SECRETACCESSKEY
}
