/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"net"
	"strconv"

	"github.com/ernestio/libmapper/providers/aws/components"
	"github.com/ernestio/libmapper/providers/aws/definition"
	"github.com/r3labs/graph"
)

// MapInstances ...
func MapInstances(d *definition.Definition) []*components.Instance {
	var is []*components.Instance

	for _, instance := range d.Instances {
		ip := make(net.IP, net.IPv4len)
		copy(ip, instance.StartIP.To4())

		for i := 0; i < instance.Count; i++ {
			name := instance.Name + "-" + strconv.Itoa(i+1)

			ci := &components.Instance{
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
			}

			ci.SetDefaultVariables()

			is = append(is, ci)
		}
	}

	return is
}

// MapDefinitionInstances : Maps output instances into a definition defined instances
func MapDefinitionInstances(g *graph.Graph) []definition.Instance {
	var instances []definition.Instance

	ci := g.GetComponents().ByType("instance")

	for _, ig := range ci.TagValues("ernest.instance_group") {
		is := ci.ByGroup("ernest.instance_group", ig)

		if len(is) < 1 {
			continue
		}

		firstInstance := is[0].(*components.Instance)
		elastic := false

		if firstInstance.ElasticIP != "" {
			elastic = true
		}

		instance := definition.Instance{
			Name:           ig,
			Type:           firstInstance.Type,
			Image:          firstInstance.Image,
			Network:        firstInstance.Network,
			StartIP:        firstInstance.IP,
			KeyPair:        firstInstance.KeyPair,
			SecurityGroups: firstInstance.SecurityGroups,
			ElasticIP:      elastic,
			Count:          len(is),
		}

		for _, vol := range firstInstance.Volumes {
			vc := g.GetComponents().ByProviderID(vol.VolumeAWSID)
			if vc == nil {
				continue
			}

			instance.Volumes = append(instance.Volumes, definition.InstanceVolume{
				Device: vol.Device,
				Volume: vc.GetTag("ernest.volume_group"),
			})
		}

		instances = append(instances, instance)

	}

	return instances
}

func mapInstanceTags(name, service, instanceGroup string) map[string]string {
	tags := make(map[string]string)

	tags["Name"] = name
	tags["ernest.service"] = service
	tags["ernest.instance_group"] = instanceGroup

	return tags
}
