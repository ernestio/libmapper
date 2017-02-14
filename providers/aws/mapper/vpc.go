/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import "github.com/ernestio/libmapper/providers/aws/components"
import "github.com/ernestio/libmapper/providers/aws/definition"

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

		cv.SetDefaultVariables()

		vpcs = append(vpcs, cv)
	}

	return vpcs
}
