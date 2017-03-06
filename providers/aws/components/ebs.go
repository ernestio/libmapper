/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, e. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package components

import (
	"errors"

	graph "gopkg.in/r3labs/graph.v2"
)

// EBSVolume ...
type EBSVolume struct {
	ProviderType     string            `json:"_provider"`
	ComponentType    string            `json:"_component"`
	ComponentID      string            `json:"_component_id"`
	State            string            `json:"_state"`
	Action           string            `json:"_action"`
	VolumeAWSID      string            `json:"volume_aws_id"`
	Name             string            `json:"name"`
	AvailabilityZone string            `json:"availability_zone"`
	VolumeType       string            `json:"volume_type"`
	Size             *int64            `json:"size"`
	Iops             *int64            `json:"iops"`
	Encrypted        bool              `json:"encrypted"`
	EncryptionKeyID  *string           `json:"encryption_key_id"`
	Tags             map[string]string `json:"tags"`
	DatacenterType   string            `json:"datacenter_type,omitempty"`
	DatacenterName   string            `json:"datacenter_name,omitempty"`
	DatacenterRegion string            `json:"datacenter_region"`
	AccessKeyID      string            `json:"aws_access_key_id"`
	SecretAccessKey  string            `json:"aws_secret_access_key"`
	Service          string            `json:"service"`
}

// GetID : returns the component's ID
func (e *EBSVolume) GetID() string {
	return e.ComponentID
}

// GetName returns a components name
func (e *EBSVolume) GetName() string {
	return e.Name
}

// GetProvider : returns the provider type
func (e *EBSVolume) GetProvider() string {
	return e.ProviderType
}

// GetProviderID returns a components provider id
func (e *EBSVolume) GetProviderID() string {
	return e.VolumeAWSID
}

// GetType : returns the type of the component
func (e *EBSVolume) GetType() string {
	return e.ComponentType
}

// GetState : returns the state of the component
func (e *EBSVolume) GetState() string {
	return e.State
}

// SetState : sets the state of the component
func (e *EBSVolume) SetState(s string) {
	e.State = s
}

// GetAction : returns the action of the component
func (e *EBSVolume) GetAction() string {
	return e.Action
}

// SetAction : Sets the action of the component
func (e *EBSVolume) SetAction(s string) {
	e.Action = s
}

// GetGroup : returns the components group
func (e *EBSVolume) GetGroup() string {
	return ""
}

// GetTags returns a components tags
func (e *EBSVolume) GetTags() map[string]string {
	return e.Tags
}

// GetTag returns a components tag
func (e *EBSVolume) GetTag(tag string) string {
	return e.Tags[tag]
}

// Diff : diff's the component against another component of the same type
func (e *EBSVolume) Diff(c graph.Component) bool {
	return false
}

// Update : updates the provider returned values of a component
func (e *EBSVolume) Update(c graph.Component) {
	ce, ok := c.(*EBSVolume)
	if ok {
		e.VolumeAWSID = ce.VolumeAWSID
	}

	e.SetDefaultVariables()
}

// Rebuild : rebuilds the component's internal state, such as templated values
func (e *EBSVolume) Rebuild(g *graph.Graph) {
	if e.VolumeType != "io1" {
		e.Iops = nil
	}

	e.SetDefaultVariables()
}

// Dependencies : returns a list of component id's upon which the component depends
func (e *EBSVolume) Dependencies() []string {
	return []string{}
}

// Validate : validates the components values
func (e *EBSVolume) Validate() error {
	if e.Name == "" {
		return errors.New("EBS Volume name should not be null")
	}

	if e.AvailabilityZone == "" {
		return errors.New("EBS Volume availability zone name should not be null")
	}

	if e.VolumeType == "" {
		return errors.New("EBS Volume type should not be null")
	}

	if e.Encrypted && e.EncryptionKeyID == nil {
		return errors.New("EBS Volume encryption key id (KMS key id) should be set if volume is encrypted")
	}

	if e.VolumeType != "io1" && e.Iops != nil {
		return errors.New("EBS Volume type must be 'io1' when specifying iops")
	}

	if e.Size != nil {
		if *e.Size < 1 || *e.Size > 16384 {
			return errors.New("EBS Volume size should be between 1 - 16385 (GB)")
		}
	}

	return nil
}

// IsStateful : returns true if the component needs to be actioned to be removed.
func (e *EBSVolume) IsStateful() bool {
	return true
}

// SetDefaultVariables : sets up the default template variables for a component
func (e *EBSVolume) SetDefaultVariables() {
	e.ComponentType = TYPEEBSVOLUME
	e.ComponentID = TYPEEBSVOLUME + TYPEDELIMITER + e.Name
	e.ProviderType = PROVIDERTYPE
	e.DatacenterName = DATACENTERNAME
	e.DatacenterType = DATACENTERTYPE
	e.DatacenterRegion = DATACENTERREGION
	e.AccessKeyID = ACCESSKEYID
	e.SecretAccessKey = SECRETACCESSKEY
}
