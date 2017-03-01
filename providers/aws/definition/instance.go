/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

// InstanceVolume ...
type InstanceVolume struct {
	Volume string `json:"volume"`
	Device string `json:"device"`
}

// Instance ...
type Instance struct {
	Name           string           `json:"name"`
	Type           string           `json:"type"`
	Image          string           `json:"image"`
	Count          int              `json:"count"`
	Network        string           `json:"network"`
	StartIP        string           `json:"start_ip"`
	KeyPair        string           `json:"key_pair"`
	ElasticIP      bool             `json:"elastic_ip"`
	SecurityGroups []string         `json:"security_groups"`
	Volumes        []InstanceVolume `json:"volumes"`
	UserData       string           `json:"user_data"`
}
