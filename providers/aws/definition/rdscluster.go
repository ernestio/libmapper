/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

// RDSBackup ...
type RDSBackup struct {
	Window    string `json:"window"`
	Retention *int64 `json:"retention"`
}

// RDSCluster ...
type RDSCluster struct {
	Name              string    `json:"name"`
	Engine            string    `json:"engine"`
	EngineVersion     string    `json:"engine_version"`
	Port              *int64    `json:"port"`
	AvailabilityZones []string  `json:"availability_zones"`
	SecurityGroups    []string  `json:"security_groups"`
	Networks          []string  `json:"networks"`
	DatabaseName      string    `json:"database_name"`
	DatabaseUsername  string    `json:"database_username"`
	DatabasePassword  string    `json:"database_password"`
	Backups           RDSBackup `json:"backups"`
	MaintenanceWindow string    `json:"maintenance_window"`
	ReplicationSource string    `json:"replication_source"`
	FinalSnapshot     bool      `json:"final_snapshot"`
}
