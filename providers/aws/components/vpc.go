/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package components

import (
	graph "gopkg.in/r3labs/graph.v2"
)

// Vpc : mapping of an instance component
type Vpc struct {
	ProviderType     string            `json:"_provider"`
	ComponentType    string            `json:"_component"`
	ComponentID      string            `json:"_component_id"`
	State            string            `json:"_state"`
	Action           string            `json:"_action"`
	VpcAWSID         string            `json:"vpc_aws_id"`
	Name             string            `json:"name"`
	Subnet           string            `json:"subnet"`
	AutoRemove       bool              `json:"auto_remove"`
	Tags             map[string]string `json:"tags"`
	DatacenterType   string            `json:"datacenter_type,omitempty"`
	DatacenterName   string            `json:"datacenter_name,omitempty"`
	DatacenterRegion string            `json:"datacenter_region"`
	AccessKeyID      string            `json:"aws_access_key_id"`
	SecretAccessKey  string            `json:"aws_secret_access_key"`
	Service          string            `json:"service"`
}

// GetID : returns the component's ID
func (v *Vpc) GetID() string {
	return v.ComponentID
}

// GetName returns a components name
func (v *Vpc) GetName() string {
	return v.Name
}

// GetProvider : returns the provider type
func (v *Vpc) GetProvider() string {
	return v.ProviderType
}

// GetProviderID returns a components provider id
func (v *Vpc) GetProviderID() string {
	return v.VpcAWSID
}

// GetType : returns the type of the component
func (v *Vpc) GetType() string {
	return v.ComponentType
}

// GetState : returns the state of the component
func (v *Vpc) GetState() string {
	return v.State
}

// SetState : sets the state of the component
func (v *Vpc) SetState(s string) {
	v.State = s
}

// GetAction : returns the action of the component
func (v *Vpc) GetAction() string {
	return v.Action
}

// SetAction : Sets the action of the component
func (v *Vpc) SetAction(s string) {
	v.Action = s
}

// GetGroup : returns the components group
func (v *Vpc) GetGroup() string {
	return ""
}

// GetTags returns a components tags
func (v *Vpc) GetTags() map[string]string {
	return v.Tags
}

// GetTag returns a components tag
func (v *Vpc) GetTag(tag string) string {
	return v.Tags[tag]
}

// Diff : diff's the component against another component of the same type
func (v *Vpc) Diff(c graph.Component) bool {
	return false
}

// Update : updates the provider returned values of a component
func (v *Vpc) Update(c graph.Component) {
	v.SetDefaultVariables()
}

// Rebuild : rebuilds the component's internal state, such as templated values
func (v *Vpc) Rebuild(g *graph.Graph) {
	v.SetDefaultVariables()
}

// Dependencies : returns a list of component id's upon which the component depends
func (v *Vpc) Dependencies() []string {
	return []string{}
}

// Validate : validates the components values
func (v *Vpc) Validate() error {
	return nil
}

// IsStateful : returns true if the component needs to be actioned to be removed.
func (v *Vpc) IsStateful() bool {
	return true
}

// SetDefaultVariables : sets up the default template variables for a component
func (v *Vpc) SetDefaultVariables() {
	v.ComponentType = TYPEVPC
	v.ComponentID = TYPEVPC + TYPEDELIMITER + v.Name
	v.ProviderType = PROVIDERTYPE
	v.DatacenterType = DATACENTERTYPE
	v.DatacenterRegion = DATACENTERREGION
	v.AccessKeyID = ACCESSKEYID
	v.SecretAccessKey = SECRETACCESSKEY
}
