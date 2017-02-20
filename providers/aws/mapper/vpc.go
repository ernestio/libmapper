/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ernestio/libmapper/providers/aws/components"
	"github.com/ernestio/libmapper/providers/aws/definition"
	"github.com/r3labs/graph"
)

// MapVpcs ...
func MapVpcs(d *definition.Definition) []*components.Vpc {
	var vpcs []*components.Vpc

	for _, vpc := range d.Vpcs {
		cv := &components.Vpc{
			Name:       vpc.Name,
			VpcAWSID:   vpc.ID,
			Subnet:     vpc.Subnet,
			AutoRemove: vpc.AutoRemove,
		}

		if vpc.ID != "" {
			cv.SetAction("none")
		}

		cv.SetDefaultVariables()

		vpcs = append(vpcs, cv)
	}

	return vpcs
}

// MapDefinitionVpcs : Maps output networks into a definition defined networks
func MapDefinitionVpcs(g *graph.Graph) []definition.Vpc {
	var vpcs []definition.Vpc

	for _, c := range g.GetComponents().ByType("vpc") {
		v := c.(*components.Vpc)

		vpcs = append(vpcs, definition.Vpc{
			ID:         v.VpcAWSID,
			Name:       v.Name,
			Subnet:     v.Subnet,
			AutoRemove: false,
		})
	}

	return vpcs
}
