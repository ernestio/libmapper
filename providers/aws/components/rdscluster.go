/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package components

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	graph "gopkg.in/r3labs/graph.v2"
)

// RDSCluster ...
type RDSCluster struct {
	ProviderType        string            `json:"_provider"`
	ComponentType       string            `json:"_component"`
	ComponentID         string            `json:"_component_id"`
	State               string            `json:"_state"`
	Action              string            `json:"_action"`
	ARN                 string            `json:"arn"`
	Name                string            `json:"name"`
	Engine              string            `json:"engine"`
	EngineVersion       string            `json:"engine_version,omitempty"`
	Port                *int64            `json:"port,omitempty"`
	Endpoint            string            `json:"endpoint,omitempty"`
	AvailabilityZones   []string          `json:"availability_zones"`
	SecurityGroups      []string          `json:"security_groups"`
	SecurityGroupAWSIDs []string          `json:"security_group_aws_ids"`
	Networks            []string          `json:"networks"`
	NetworkAWSIDs       []string          `json:"network_aws_ids"`
	DatabaseName        string            `json:"database_name,omitempty"`
	DatabaseUsername    string            `json:"database_username,omitempty"`
	DatabasePassword    string            `json:"database_password,omitempty"`
	BackupRetention     *int64            `json:"backup_retention,omitempty"`
	BackupWindow        string            `json:"backup_window,omitempty"`
	MaintenanceWindow   string            `json:"maintenance_window,omitempty"`
	ReplicationSource   string            `json:"replication_source,omitempty"`
	FinalSnapshot       bool              `json:"final_snapshot"`
	Tags                map[string]string `json:"tags"`
	DatacenterType      string            `json:"datacenter_type"`
	DatacenterName      string            `json:"datacenter_name"`
	DatacenterRegion    string            `json:"datacenter_region"`
	AccessKeyID         string            `json:"aws_access_key_id"`
	SecretAccessKey     string            `json:"aws_secret_access_key"`
	Service             string            `json:"service"`
}

// GetID : returns the component's ID
func (r *RDSCluster) GetID() string {
	return r.ComponentID
}

// GetName returns a components name
func (r *RDSCluster) GetName() string {
	return r.Name
}

// GetProvider : returns the provider type
func (r *RDSCluster) GetProvider() string {
	return r.ProviderType
}

// GetProviderID returns a components provider id
func (r *RDSCluster) GetProviderID() string {
	return r.ARN
}

// GetType : returns the type of the component
func (r *RDSCluster) GetType() string {
	return r.ComponentType
}

// GetState : returns the state of the component
func (r *RDSCluster) GetState() string {
	return r.State
}

// SetState : sets the state of the component
func (r *RDSCluster) SetState(s string) {
	r.State = s
}

// GetAction : returns the action of the component
func (r *RDSCluster) GetAction() string {
	return r.Action
}

// SetAction : Sets the action of the component
func (r *RDSCluster) SetAction(s string) {
	r.Action = s
}

// GetGroup : returns the components group
func (r *RDSCluster) GetGroup() string {
	return ""
}

// GetTags returns a components tags
func (r *RDSCluster) GetTags() map[string]string {
	return r.Tags
}

// GetTag returns a components tag
func (r *RDSCluster) GetTag(tag string) string {
	return r.Tags[tag]
}

// Diff : diff's the component against another component of the same type
func (r *RDSCluster) Diff(c graph.Component) bool {
	cr, ok := c.(*RDSCluster)
	if ok {
		if r.Port != nil && cr.Port != nil {
			if *r.Port != *cr.Port {
				return true
			}
		}

		if r.DatabasePassword != cr.DatabasePassword {
			return true
		}

		if r.BackupRetention != nil && cr.BackupRetention != nil {
			if *r.BackupRetention != *cr.BackupRetention {
				return true
			}
		}

		if r.BackupWindow != cr.BackupWindow {
			return true
		}

		if r.MaintenanceWindow != cr.MaintenanceWindow {
			return true
		}

		if reflect.DeepEqual(r.Networks, cr.Networks) != true {
			return true
		}

		return !reflect.DeepEqual(r.SecurityGroups, cr.SecurityGroups)
	}

	return false
}

// Update : updates the provider returned values of a component
func (r *RDSCluster) Update(c graph.Component) {
	cr, ok := c.(*RDSCluster)
	if ok {
		r.ARN = cr.ARN
		r.Endpoint = cr.Endpoint
	}

	r.SetDefaultVariables()
}

// Rebuild : rebuilds the component's internal state, such as templated values
func (r *RDSCluster) Rebuild(g *graph.Graph) {
	if len(r.Networks) > len(r.NetworkAWSIDs) {
		for _, nw := range r.Networks {
			r.NetworkAWSIDs = append(r.NetworkAWSIDs, templSubnetID(nw))
		}
	}

	if len(r.NetworkAWSIDs) > len(r.Networks) {
		for _, nwid := range r.NetworkAWSIDs {
			nw := g.GetComponents().ByProviderID(nwid)
			if nw != nil {
				r.Networks = append(r.Networks, nw.GetName())
			}
		}
	}

	if len(r.SecurityGroups) > len(r.SecurityGroupAWSIDs) {
		for _, sg := range r.SecurityGroups {
			r.SecurityGroupAWSIDs = append(r.SecurityGroupAWSIDs, templSecurityGroupID(sg))
		}
	}

	if len(r.SecurityGroupAWSIDs) > len(r.SecurityGroups) {
		for _, sgid := range r.SecurityGroupAWSIDs {
			sg := g.GetComponents().ByProviderID(sgid)
			if sg != nil {
				r.SecurityGroups = append(r.SecurityGroups, sg.GetName())
			}
		}
	}

	r.SetDefaultVariables()
}

// Dependencies : returns a list of component id's upon which the component depends
func (r *RDSCluster) Dependencies() []string {
	var deps []string

	for _, sg := range r.SecurityGroups {
		deps = append(deps, TYPESECURITYGROUP+TYPEDELIMITER+sg)
	}

	for _, nw := range r.Networks {
		deps = append(deps, TYPENETWORK+TYPEDELIMITER+nw)
	}

	return deps
}

// Validate : validates the components values
func (r *RDSCluster) Validate() error {
	if r.Name == "" {
		return errors.New("RDS Cluster name should not be null")
	}

	if len(r.Name) > 255 {
		return errors.New("RDS Cluster name should not exceed 255 characters")
	}

	if r.Engine == "" {
		return errors.New("RDS Cluster engine type should not be null")
	}

	if r.ReplicationSource != "" {
		if len(r.ReplicationSource) < 12 || r.ReplicationSource[:12] != "arn:aws:rds:" {
			return errors.New("RDS Cluster replication source should be a valid amazon resource name (ARN), i.e. 'arn:aws:rds:us-east-1:123456789012:cluster:my-aurora-cluster'")
		}
	}

	if r.DatabaseName == "" {
		return errors.New("RDS Cluster database name should not be null")
	}

	if len(r.DatabaseName) > 64 {
		return errors.New("RDS Cluster database name should not exceed 64 characters")
	}

	for _, c := range r.DatabaseName {
		if unicode.IsLetter(c) != true && unicode.IsNumber(c) != true {
			return errors.New("RDS Cluster database name can only contain alphanumeric characters")
		}
	}

	if r.DatabaseUsername == "" {
		return errors.New("RDS Cluster database username should not be null")
	}

	if len(r.DatabaseUsername) > 16 {
		return errors.New("RDS Cluster database username should not exceed 16 characters")
	}

	if r.DatabasePassword == "" {
		return errors.New("RDS Cluster database password should not be null")
	}

	if len(r.DatabasePassword) < 8 || len(r.DatabasePassword) > 41 {
		return errors.New("RDS Cluster database password should be between 8 and 41 characters")
	}

	for _, c := range r.DatabasePassword {
		if unicode.IsSymbol(c) || unicode.IsMark(c) {
			return fmt.Errorf("RDS Cluster database password contains an offending character: '%c'", c)
		}
	}

	if r.Port != nil {
		if *r.Port < 1150 || *r.Port > 65535 {
			return errors.New("RDS Cluster port number should be between 1150 and 65535")
		}
	}

	if r.BackupRetention != nil {
		if *r.BackupRetention < 1 || *r.BackupRetention > 35 {
			return errors.New("RDS Cluster backup retention should be between 1 and 35 days")
		}
	}

	if r.BackupWindow != "" {
		parts := strings.Split(r.BackupWindow, "-")

		err := validateTimeFormat(parts[0])
		if err != nil {
			return errors.New("RDS Cluster backup window: " + err.Error())
		}

		err = validateTimeFormat(parts[1])
		if err != nil {
			return errors.New("RDS Cluster backup window: " + err.Error())
		}
	}

	if mwerr := validateTimeWindow(r.MaintenanceWindow); r.MaintenanceWindow != "" && mwerr != nil {
		return fmt.Errorf("RDS Cluster maintenance window: %s", mwerr.Error())
	}

	return nil

}

// IsStateful : returns true if the component needs to be actioned to be removed.
func (r *RDSCluster) IsStateful() bool {
	return true
}

// SetDefaultVariables : sets up the default template variables for a component
func (r *RDSCluster) SetDefaultVariables() {
	r.ComponentType = TYPERDSCLUSTER
	r.ComponentID = TYPERDSCLUSTER + TYPEDELIMITER + r.Name
	r.ProviderType = PROVIDERTYPE
	r.DatacenterName = DATACENTERNAME
	r.DatacenterType = DATACENTERTYPE
	r.DatacenterRegion = DATACENTERREGION
	r.AccessKeyID = ACCESSKEYID
	r.SecretAccessKey = SECRETACCESSKEY
}
