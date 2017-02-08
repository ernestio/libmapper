/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

// ELBListener ...
type ELBListener struct {
	FromPort int    `json:"from_port"`
	ToPort   int    `json:"to_port"`
	Protocol string `json:"protocol"`
	SSLCert  string `json:"ssl_cert"`
}

// ELB ...
type ELB struct {
	Name           string        `json:"name"`
	Private        bool          `json:"private"`
	Subnets        []string      `json:"networks"`
	Instances      []string      `json:"instances"`
	SecurityGroups []string      `json:"security_groups"`
	Listeners      []ELBListener `json:"listeners"`
}
