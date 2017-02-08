/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import "encoding/json"

// Definition ...
type Definition struct {
	Name           string          `json:"name"`
	Datacenter     string          `json:"datacenter"`
	Vpcs           []Vpc           `json:"vpcs"`
	Networks       []Network       `json:"networks,omitempty"`
	Instances      []Instance      `json:"instances,omitempty"`
	SecurityGroups []SecurityGroup `json:"security_groups,omitempty"`
	ELBs           []ELB           `json:"loadbalancers,omitempty"`
	EBSVolumes     []EBSVolume     `json:"ebs_volumes,omitempty"`
	//S3Buckets         []S3            `json:"s3_buckets,omitempty"`
	//Route53Zones      []Route53Zone   `json:"route53_zones,omitempty"`
	//RDSClusters       []RDSCluster    `json:"rds_clusters,omitempty"`
	//RDSInstances      []RDSInstance   `json:"rds_instances,omitempty"`
	//NatGateways       []NatGateway    `json:"nat_gateways,omitempty"`
}

// New returns a new Definition
func New() *Definition {
	return &Definition{}
}

// FromJSON creates a definition from json
func FromJSON(data []byte) (*Definition, error) {
	var d Definition

	err := json.Unmarshal(data, d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}
