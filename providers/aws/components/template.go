/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package components

const (
	PROVIDERTYPE     = "$(datacenters.items.0.type)"
	DATACENTERNAME   = "$(datacenters.items.0.name)"
	DATACENTERTYPE   = "$(datacenters.items.0.type)"
	ACCESSKEYID      = `$(components.#[_component_id="credentials::aws"].aws_access_key_id)`
	SECRETACCESSKEY  = `$(components.#[_component_id="credentials::aws"].aws_secret_access_key)`
	DATACENTERREGION = `$(components.#[_component_id="credentials::aws"].region)`
)

func templVpcID(vpc string) string {
	return `$(components.#[_component_id="` + "vpc::" + vpc + `"].vpc_aws_id)`
}

func templSecurityGroupID(sg string) string {
	return `$(components.#[_component_id="` + "security_group::" + sg + `"].security_group_aws_id)`
}

func templSubnetID(nw string) string {
	return `$(components.#[_component_id="` + "network::" + nw + `"].network_aws_id)`
}
