/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import "github.com/ernestio/libmapper/providers/aws/components"
import "github.com/ernestio/libmapper/providers/aws/definition"

// MapNetworks ...
func MapNetworks(d *definition.Definition) []*components.Network {
	var ns []*components.Network

	for _, network := range d.Networks {
		cn := &components.Network{
			Name:             network.Name,
			Subnet:           network.Subnet,
			IsPublic:         network.Public,
			AvailabilityZone: network.AvailabilityZone,
			Vpc:              network.VPC,
			Tags:             mapNetworkTags(network.Name, d.Name, network.NatGateway),
		}

		cn.SetDefaultVariables()

		ns = append(ns, cn)
	}

	return ns
}

func mapNetworkTags(name, service, gateway string) map[string]string {
	tags := make(map[string]string)

	tags["Name"] = name
	tags["ernest.service"] = service

	if gateway != "" {
		tags["ernest.nat_gateway"] = gateway
	}

	return tags
}
