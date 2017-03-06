/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"strconv"

	"github.com/ernestio/libmapper/providers/aws/components"
	"github.com/ernestio/libmapper/providers/aws/definition"
	graph "gopkg.in/r3labs/graph.v2"
)

// MapEBSVolumes : Maps the ebs volumes from a given input payload.
func MapEBSVolumes(d *definition.Definition) []*components.EBSVolume {
	var volumes []*components.EBSVolume

	for _, vol := range d.EBSVolumes {
		for i := 0; i < vol.Count; i++ {
			name := vol.Name + "-" + strconv.Itoa(i+1)

			v := &components.EBSVolume{
				Name:             name,
				AvailabilityZone: vol.AvailabilityZone,
				VolumeType:       vol.Type,
				Size:             vol.Size,
				Iops:             vol.Iops,
				Encrypted:        vol.Encrypted,
				EncryptionKeyID:  vol.EncryptionKeyID,
				Tags:             mapEBSTags(name, d.Name, vol.Name),
			}

			v.SetDefaultVariables()

			volumes = append(volumes, v)
		}
	}

	return volumes
}

// MapDefinitionEBSVolumes : Maps components ebs volumes into a definition defined ebs volumes
func MapDefinitionEBSVolumes(g *graph.Graph) []definition.EBSVolume {
	var vols []definition.EBSVolume

	ci := g.GetComponents().ByType("ebs_volume")

	for _, vg := range ci.TagValues("ernest.volume_group") {
		vs := ci.ByGroup("ernest.volume_group", vg)

		if len(vs) < 1 {
			continue
		}

		firstVolume := vs[0].(*components.EBSVolume)

		vols = append(vols, definition.EBSVolume{
			Name:             vg,
			Type:             firstVolume.VolumeType,
			Size:             firstVolume.Size,
			Iops:             firstVolume.Iops,
			AvailabilityZone: firstVolume.AvailabilityZone,
			Encrypted:        firstVolume.Encrypted,
			EncryptionKeyID:  firstVolume.EncryptionKeyID,
			Count:            len(vs),
		})

	}

	return vols
}
func mapEBSTags(name, service, volumeGroup string) map[string]string {
	tags := make(map[string]string)

	tags["Name"] = name
	tags["ernest.service"] = service
	tags["ernest.volume_group"] = volumeGroup

	return tags
}
