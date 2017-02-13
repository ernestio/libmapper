/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

// SecurityGroup ...
type SecurityGroup struct {
	Name    string              `json:"name"`
	Vpc     string              `json:"vpc"`
	Ingress []SecurityGroupRule `json:"ingress"`
	Egress  []SecurityGroupRule `json:"egress"`
}

// SecurityGroupRule ...
type SecurityGroupRule struct {
	IP       string `json:"ip"`
	FromPort string `json:"from_port"`
	ToPort   string `json:"to_port"`
	Protocol string `json:"protocol"`
}
