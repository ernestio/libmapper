/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"net"
	"strconv"

	"github.com/ernestio/libmapper/providers/aws/components"
	"github.com/ernestio/libmapper/providers/aws/definition"
)

// MapInstances ...
func MapInstances(d definition.Definition) []*components.Instance {
	var is []*components.Instance

	for _, instance := range d.Instances {
		ip := make(net.IP, net.IPv4len)
		copy(ip, instance.StartIP.To4())

		for i := 0; i < instance.Count; i++ {
			name := instance.Name + "-" + strconv.Itoa(i+1)

			is = append(is, &components.Instance{
				Name:            name,
				Type:            instance.Type,
				Image:           instance.Image,
				Network:         instance.Network,
				IP:              net.ParseIP(ip.String()),
				KeyPair:         instance.KeyPair,
				AssignElasticIP: instance.ElasticIP,
				SecurityGroups:  instance.SecurityGroups,
				UserData:        instance.UserData,
				Tags:            mapInstanceTags(name, d.Name, instance.Name),
			})
		}
	}

	return is
}

// SecurityGroupAWSIDs: mapInstanceSecurityGroupIDs(sgroups)
// NetworkAWSID:        `$(networks.items.#[name="` + d.GeneratedName() + instance.Network + `"].network_aws_id)`,

func mapInstanceTags(name, service, instanceGroup string) map[string]string {
	tags := make(map[string]string)

	tags["Name"] = name
	tags["ernest.service"] = service
	tags["ernest.instance_group"] = instanceGroup

	return tags
}
