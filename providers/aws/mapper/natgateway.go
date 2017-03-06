/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ernestio/libmapper/providers/aws/components"
	"github.com/ernestio/libmapper/providers/aws/definition"
	graph "gopkg.in/r3labs/graph.v2"
)

// MapNats : Generates necessary nats rules for input networks
func MapNats(d *definition.Definition) []*components.NatGateway {
	var nats []*components.NatGateway

	for _, ng := range d.NatGateways {
		nats = append(nats, &components.NatGateway{
			Name:           ng.Name,
			PublicNetwork:  ng.PublicNetwork,
			RoutedNetworks: mapNetworkNames(d, ng.Name),
		})
	}

	return nats
}

// MapDefinitionNats : Maps components nat gateways into a definition defined nat gateways
func MapDefinitionNats(g *graph.Graph) []definition.NatGateway {
	var nts []definition.NatGateway

	for _, ng := range g.GetComponents().ByType("nat") {
		nc := ng.(*components.NatGateway)

		nts = append(nts, definition.NatGateway{
			Name:          nc.Name,
			PublicNetwork: nc.PublicNetwork,
		})
	}

	return nts
}

func mapNetworkNames(d *definition.Definition, name string) []string {
	var nws []string
	for _, network := range d.Networks {
		if network.NatGateway == name {
			nws = append(nws, network.Name)
		}
	}

	return nws
}
