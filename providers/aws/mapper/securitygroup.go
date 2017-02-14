/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"strconv"

	"github.com/ernestio/libmapper/providers/aws/components"
	"github.com/ernestio/libmapper/providers/aws/definition"
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

// BuildRule converts a definition rule into an output rule
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

// MapProtocol : Maps the security groups protocol to the correct value
func MapProtocol(protocol string) string {
	if protocol == "any" {
		return "-1"
	}
	return protocol
}
