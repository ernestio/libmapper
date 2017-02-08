/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

// Network ...
type Network struct {
	Name             string `json:"name"`
	Subnet           string `json:"subnet"`
	Public           bool   `json:"public"`
	NatGateway       string `json:"nat_gateway"`
	AvailabilityZone string `json:"availability_zone"`
	VPC              string `json:"vpc"`
}
