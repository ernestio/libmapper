/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package components

import (
	"errors"

	graph "gopkg.in/r3labs/graph.v2"
)

// SecurityGroupRule ...
type SecurityGroupRule struct {
	IP       string `json:"ip"`
	From     int    `json:"from_port"`
	To       int    `json:"to_port"`
	Protocol string `json:"protocol"`
}

// SecurityGroup : Mapping of a security group component
type SecurityGroup struct {
	ProviderType       string `json:"_provider"`
	ComponentType      string `json:"_component"`
	ComponentID        string `json:"_component_id"`
	State              string `json:"_state"`
	Action             string `json:"_action"`
	SecurityGroupAWSID string `json:"security_group_aws_id"`
	Name               string `json:"name"`
	Rules              struct {
		Ingress []SecurityGroupRule `json:"ingress"`
		Egress  []SecurityGroupRule `json:"egress"`
	} `json:"rules"`
	Tags             map[string]string `json:"tags"`
	DatacenterType   string            `json:"datacenter_type,omitempty"`
	DatacenterName   string            `json:"datacenter_name,omitempty"`
	DatacenterRegion string            `json:"datacenter_region"`
	AccessKeyID      string            `json:"aws_access_key_id"`
	SecretAccessKey  string            `json:"aws_secret_access_key"`
	Vpc              string            `json:"vpc"`
	VpcID            string            `json:"vpc_id"`
	Service          string            `json:"service"`
}

// GetID : returns the component's ID
func (sg *SecurityGroup) GetID() string {
	return sg.ComponentID
}

// GetName returns a components name
func (sg *SecurityGroup) GetName() string {
	return sg.Name
}

// GetProvider : returns the provider type
func (sg *SecurityGroup) GetProvider() string {
	return sg.ProviderType
}

// GetProviderID returns a components provider id
func (sg *SecurityGroup) GetProviderID() string {
	return sg.SecurityGroupAWSID
}

// GetType : returns the type of the component
func (sg *SecurityGroup) GetType() string {
	return sg.ComponentType
}

// GetState : returns the state of the component
func (sg *SecurityGroup) GetState() string {
	return sg.State
}

// SetState : sets the state of the component
func (sg *SecurityGroup) SetState(s string) {
	sg.State = s
}

// GetAction : returns the action of the component
func (sg *SecurityGroup) GetAction() string {
	return sg.Action
}

// SetAction : Sets the action of the component
func (sg *SecurityGroup) SetAction(s string) {
	sg.Action = s
}

// GetGroup : returns the components group
func (sg *SecurityGroup) GetGroup() string {
	return ""
}

// GetTags returns a components tags
func (sg *SecurityGroup) GetTags() map[string]string {
	return sg.Tags
}

// GetTag returns a components tag
func (sg *SecurityGroup) GetTag(tag string) string {
	return sg.Tags[tag]
}

// Diff : diff's the component against another component of the same type
func (sg *SecurityGroup) Diff(c graph.Component) bool {
	csg, ok := c.(*SecurityGroup)
	if ok {
		if len(sg.Rules.Ingress) != len(csg.Rules.Ingress) ||
			len(sg.Rules.Egress) != len(csg.Rules.Egress) {
			return true
		}

		for _, rule := range sg.Rules.Ingress {
			if hasRule(csg.Rules.Ingress, rule) != true {
				return true
			}
		}

		for _, rule := range sg.Rules.Egress {
			if hasRule(csg.Rules.Egress, rule) != true {
				return true
			}
		}
	}

	return false
}

// Update : updates the provider returned values of a component
func (sg *SecurityGroup) Update(c graph.Component) {
	csg, ok := c.(*SecurityGroup)
	if ok {
		sg.SecurityGroupAWSID = csg.SecurityGroupAWSID
	}

	sg.SetDefaultVariables()
}

// Rebuild : rebuilds the component's internal state, such as templated values
func (sg *SecurityGroup) Rebuild(g *graph.Graph) {
	if sg.Vpc == "" && sg.VpcID != "" {
		v := g.GetComponents().ByProviderID(sg.VpcID)
		if v != nil {
			sg.Vpc = v.GetName()
		}
	}

	if sg.Vpc != "" && sg.VpcID == "" {
		sg.VpcID = templVpcID(sg.Vpc)
	}

	sg.SetDefaultVariables()
}

// Dependencies : returns a list of component id's upon which the component depends
func (sg *SecurityGroup) Dependencies() []string {
	return []string{"vpc::" + sg.Vpc}
}

// Validate : validates the components values
func (sg *SecurityGroup) Validate() error {
	if sg.Name == "" {
		return errors.New("Security Group name should not be null")
	}

	for _, rule := range sg.Rules.Ingress {
		err := rule.Validate()
		if err != nil {
			return err
		}
	}

	for _, rule := range sg.Rules.Egress {
		err := rule.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// IsStateful : returns true if the component needs to be actioned to be removed.
func (sg *SecurityGroup) IsStateful() bool {
	return true
}

// SetDefaultVariables : sets up the default template variables for a component
func (sg *SecurityGroup) SetDefaultVariables() {
	sg.ComponentType = TYPESECURITYGROUP
	sg.ComponentID = TYPESECURITYGROUP + TYPEDELIMITER + sg.Name
	sg.ProviderType = PROVIDERTYPE
	sg.DatacenterName = DATACENTERNAME
	sg.DatacenterType = DATACENTERTYPE
	sg.DatacenterRegion = DATACENTERREGION
	sg.AccessKeyID = ACCESSKEYID
	sg.SecretAccessKey = SECRETACCESSKEY
}

// Validate security group rule
func (rule *SecurityGroupRule) Validate() error {
	// Validate FromPort Port
	// Must be: [0 - 65535]
	err := validatePort(rule.From, "Security Group From")
	if err != nil {
		return err
	}

	// Validate ToPort Port
	// Must be: [0 - 65535]
	err = validatePort(rule.To, "Security Group To")
	if err != nil {
		return err
	}

	// Validate Protocol
	// Must be one of: tcp | udp | icmp | any | tcp & udp
	err = validateProtocol(rule.Protocol)
	if err != nil {
		return err
	}

	return nil
}

func hasRule(rules []SecurityGroupRule, rule SecurityGroupRule) bool {
	for _, r := range rules {
		if ruleMatches(r.To, rule.To, r.Protocol, rule.Protocol) &&
			r.Protocol == rule.Protocol &&
			r.IP == rule.IP &&
			ruleMatches(r.From, rule.From, r.Protocol, rule.Protocol) {
			return true
		}
	}

	return false
}

func ruleMatches(nv, ov int, np, op string) bool {
	if np == "-1" && op == "-1" {
		return true
	}

	return nv == ov
}
