/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

// EBSVolume ...
type EBSVolume struct {
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	Size             *int64  `json:"size"`
	Iops             *int64  `json:"iops"`
	Count            int     `json:"count"`
	Encrypted        bool    `json:"encrypted"`
	EncryptionKeyID  *string `json:"encryption_key_id"`
	AvailabilityZone string  `json:"availability_zone"`
}
