/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package components

import (
	"github.com/ernestio/ernestprovider/event"
	"github.com/ernestio/ernestprovider/providers/azure/virtualmachine"
	"github.com/r3labs/graph"
)

// VirtualMachine : A resource group a container that holds
// related resources for an Azure solution.
type VirtualMachine struct {
	*virtualmachine.Event
	*Base
}

// GetID : returns the component's ID
func (i *VirtualMachine) GetID() string {
	return i.ComponentID
}

// GetName returns a components name
func (i *VirtualMachine) GetName() string {
	return i.Name
}

// GetProvider : returns the provider type
func (i *VirtualMachine) GetProvider() string {
	return i.ProviderType
}

// GetProviderID returns a components provider id
func (i *VirtualMachine) GetProviderID() string {
	return i.ID
}

// GetType : returns the type of the component
func (i *VirtualMachine) GetType() string {
	return i.ComponentType
}

// GetState : returns the state of the component
func (i *VirtualMachine) GetState() string {
	return i.State
}

// SetState : sets the state of the component
func (i *VirtualMachine) SetState(s string) {
	i.State = s
}

// GetAction : returns the action of the component
func (i *VirtualMachine) GetAction() string {
	return i.Action
}

// SetAction : Sets the action of the component
func (i *VirtualMachine) SetAction(s string) {
	i.Action = s
}

// GetGroup : returns the components group
func (i *VirtualMachine) GetGroup() string {
	return ""
}

// GetTags returns a components tags
func (i *VirtualMachine) GetTags() map[string]string {
	return i.Tags
}

// GetTag returns a components tag
func (i *VirtualMachine) GetTag(tag string) string {
	return ""
}

// Diff : diff's the component against another component of the same type
func (i *VirtualMachine) Diff(c graph.Component) bool {

	return false
}

// Update : updates the provider returned values of a component
func (i *VirtualMachine) Update(c graph.Component) {
}

// Rebuild : rebuilds the component's internal state, such as templated values
func (i *VirtualMachine) Rebuild(g *graph.Graph) {
}

// Dependencies : returns a list of component id's upon which the component depends
func (i *VirtualMachine) Dependencies() (deps []string) {
	return
}

// Validate : validates the components values
func (i *VirtualMachine) Validate() error {
	val := event.NewValidator()
	return val.Validate(i)
}

// IsStateful : returns true if the component needs to be actioned to be removed.
func (i *VirtualMachine) IsStateful() bool {
	return true
}

// SetDefaultVariables : sets up the default template variables for a component
func (i *VirtualMachine) SetDefaultVariables() {
}
