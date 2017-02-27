/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ernestio/libmapper/providers/aws/components"
	"github.com/ernestio/libmapper/providers/aws/definition"
	graph "gopkg.in/r3labs/graph.v2"
)

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

// MapDefinitionNetworks : Maps output networks into a definition defined networks
func MapDefinitionNetworks(g *graph.Graph) []definition.Network {
	var nws []definition.Network

	for _, c := range g.GetComponents().ByType("network") {
		n := c.(*components.Network)

		nws = append(nws, definition.Network{
			Name:             n.Name,
			Subnet:           n.Subnet,
			Public:           n.IsPublic,
			AvailabilityZone: n.AvailabilityZone,
			NatGateway:       n.Tags["ernest.nat_gateway"],
		})
	}

	return nws
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
