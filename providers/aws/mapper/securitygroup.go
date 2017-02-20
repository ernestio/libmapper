/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"strconv"

	"github.com/ernestio/libmapper/providers/aws/components"
	"github.com/ernestio/libmapper/providers/aws/definition"
	"github.com/r3labs/graph"
)

// MapSecurityGroups ...
func MapSecurityGroups(d *definition.Definition) []*components.SecurityGroup {
	var sgs []*components.SecurityGroup

	for _, sg := range d.SecurityGroups {

		s := components.SecurityGroup{
			Name: sg.Name,
			Vpc:  sg.Vpc,
			Tags: mapTags(sg.Name, d.Name),
		}

		for _, rule := range sg.Ingress {
			s.Rules.Ingress = append(s.Rules.Ingress, BuildRule(rule))
		}

		for _, rule := range sg.Egress {
			s.Rules.Egress = append(s.Rules.Egress, BuildRule(rule))
		}

		s.SetDefaultVariables()

		sgs = append(sgs, &s)
	}

	return sgs
}

// MapDefinitionSecurityGroups : Maps components security groups into a definition defined security groups
func MapDefinitionSecurityGroups(g *graph.Graph) []definition.SecurityGroup {
	var sgs []definition.SecurityGroup

	for _, c := range g.GetComponents().ByType("security_group") {
		sg := c.(*components.SecurityGroup)

		s := definition.SecurityGroup{
			Name: sg.Name,
		}

		for _, rule := range sg.Rules.Ingress {
			s.Ingress = append(s.Ingress, BuildDefinitionRule(rule))
		}

		for _, rule := range sg.Rules.Egress {
			s.Egress = append(s.Egress, BuildDefinitionRule(rule))
		}

		sgs = append(sgs, s)
	}

	return sgs
}

// BuildRule converts a definition rule into an components rule
func BuildRule(rule definition.SecurityGroupRule) components.SecurityGroupRule {
	from, _ := strconv.Atoi(rule.FromPort)
	to, _ := strconv.Atoi(rule.ToPort)

	return components.SecurityGroupRule{
		IP:       rule.IP,
		From:     from,
		To:       to,
		Protocol: MapProtocol(rule.Protocol),
	}
}

// BuildDefinitionRule converts an components rule into a definition rule
func BuildDefinitionRule(rule components.SecurityGroupRule) definition.SecurityGroupRule {
	from := strconv.Itoa(rule.From)
	to := strconv.Itoa(rule.To)

	return definition.SecurityGroupRule{
		IP:       rule.IP,
		FromPort: from,
		ToPort:   to,
		Protocol: MapDefinitionProtocol(rule.Protocol),
	}
}

// MapProtocol : Maps the security groups protocol to the correct value
func MapProtocol(protocol string) string {
	if protocol == "any" {
		return "-1"
	}
	return protocol
}

// MapDefinitionProtocol : Maps the security groups protocol to the correct definition value
func MapDefinitionProtocol(protocol string) string {
	if protocol == "-1" {
		return "any"
	}
	return protocol
}
