/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"strings"

	"github.com/ernestio/libmapper/providers/aws/components"
	"github.com/ernestio/libmapper/providers/aws/definition"
	graph "gopkg.in/r3labs/graph.v2"
)

// MapELBs : Maps the elbs from a given input payload.
func MapELBs(d *definition.Definition) []*components.ELB {
	var elbs []*components.ELB

	for _, elb := range d.ELBs {
		e := components.ELB{
			Name:           elb.Name,
			IsPrivate:      elb.Private,
			Instances:      elb.Instances,
			Networks:       elb.Subnets,
			SecurityGroups: elb.SecurityGroups,
			Tags:           mapTagsServiceOnly(d.Name),
		}

		for _, listener := range elb.Listeners {
			e.Listeners = append(e.Listeners, components.ELBListener{
				FromPort: listener.FromPort,
				ToPort:   listener.ToPort,
				Protocol: strings.ToUpper(listener.Protocol),
				SSLCert:  listener.SSLCert,
			})
		}

		e.SetDefaultVariables()

		elbs = append(elbs, &e)
	}

	return elbs
}

// MapDefinitionELBs : Maps output elbs into a definition defined elbs
func MapDefinitionELBs(g *graph.Graph) []definition.ELB {
	var elbs []definition.ELB

	for _, gelb := range g.GetComponents().ByType("elb") {
		elb := gelb.(*components.ELB)

		e := definition.ELB{
			Name:           elb.Name,
			Private:        elb.IsPrivate,
			Subnets:        elb.Networks,
			Instances:      elb.Instances,
			SecurityGroups: elb.SecurityGroups,
		}

		for _, l := range elb.Listeners {
			e.Listeners = append(e.Listeners, definition.ELBListener{
				FromPort: l.FromPort,
				ToPort:   l.ToPort,
				Protocol: strings.ToLower(l.Protocol),
				SSLCert:  l.SSLCert,
			})
		}

		elbs = append(elbs, e)
	}

	return elbs
}
