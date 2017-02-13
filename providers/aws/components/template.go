/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package components

const (
	PROVIDERTYPE     = "$(datacenters.items.0.type)"
	DATACENTERNAME   = "$(datacenters.items.0.name)"
	DATACENTERTYPE   = "$(datacenters.items.0.type)"
	ACCESSKEYID      = "$(datacenters.items.0.aws_access_key_id)"
	SECRETACCESSKEY  = "$(datacenters.items.0.aws_secret_access_key)"
	DATACENTERREGION = "$(datacenters.items.0.region)"
	VPCID            = "$(vpcs.items.0.vpc_id)"
)

func templSecurityGroupID(sg string) string {
	return `$(components.#[_component_id="` + "security_group::" + sg + `"].security_group_aws_id)`
}
