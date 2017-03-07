/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ernestio/libmapper/providers/aws/components"
	"github.com/ernestio/libmapper/providers/aws/definition"
	graph "gopkg.in/r3labs/graph.v2"
)

// MapRDSClusters : Maps the rds clusters for the input payload on a ernest internal format
func MapRDSClusters(d *definition.Definition) []*components.RDSCluster {
	var clusters []*components.RDSCluster

	for _, cluster := range d.RDSClusters {
		rc := &components.RDSCluster{
			Name:              cluster.Name,
			Engine:            cluster.Engine,
			EngineVersion:     cluster.EngineVersion,
			Port:              cluster.Port,
			AvailabilityZones: cluster.AvailabilityZones,
			SecurityGroups:    cluster.SecurityGroups,
			Networks:          cluster.Networks,
			DatabaseName:      cluster.DatabaseName,
			DatabaseUsername:  cluster.DatabaseUsername,
			DatabasePassword:  cluster.DatabasePassword,
			BackupRetention:   cluster.Backups.Retention,
			BackupWindow:      cluster.Backups.Window,
			MaintenanceWindow: cluster.MaintenanceWindow,
			ReplicationSource: cluster.ReplicationSource,
			FinalSnapshot:     cluster.FinalSnapshot,
			Tags:              mapTagsServiceOnly(d.Name),
		}

		rc.SetDefaultVariables()

		clusters = append(clusters, rc)
	}

	return clusters
}

// MapDefinitionRDSClusters : Maps the rds clusters for the internal ernest format to the input definition format
func MapDefinitionRDSClusters(g *graph.Graph) []definition.RDSCluster {
	var clusters []definition.RDSCluster

	for _, gc := range g.GetComponents().ByType("rds_cluster") {
		cluster := gc.(*components.RDSCluster)
		c := definition.RDSCluster{
			Name:              cluster.Name,
			Engine:            cluster.Engine,
			EngineVersion:     cluster.EngineVersion,
			Port:              cluster.Port,
			AvailabilityZones: cluster.AvailabilityZones,
			SecurityGroups:    cluster.SecurityGroups,
			Networks:          cluster.Networks,
			DatabaseName:      cluster.DatabaseName,
			DatabaseUsername:  cluster.DatabaseUsername,
			DatabasePassword:  cluster.DatabasePassword,
			MaintenanceWindow: cluster.MaintenanceWindow,
			ReplicationSource: cluster.ReplicationSource,
			FinalSnapshot:     cluster.FinalSnapshot,
		}

		c.Backups.Retention = cluster.BackupRetention
		c.Backups.Window = cluster.BackupWindow

		clusters = append(clusters, c)
	}

	return clusters
}
