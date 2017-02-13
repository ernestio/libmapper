/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package components

import (
	"errors"
	"fmt"
)

const (
	// PROTOCOLTCP : TCP Protocol value
	PROTOCOLTCP = "tcp"
	// PROTOCOLUDP : UDP Protocol value
	PROTOCOLUDP = "udp"
	// PROTOCOLANY : Any Protocol value
	PROTOCOLANY = "any"
	// PROTOCOLICMP : ICMP Protocol value
	PROTOCOLICMP = "icmp"
	// TARGETEXTERNAL : External target
	TARGETEXTERNAL = "external"
	// TARGETINTERNAL : Internal target
	TARGETINTERNAL = "internal"
	// TARGETANY : Any Target
	TARGETANY = "any"
	// AWSMAXNAME : Maximum size of an aws character
	AWSMAXNAME = 50
)

func validateProtocol(p string) error {
	switch p {
	case PROTOCOLTCP, PROTOCOLUDP, PROTOCOLICMP, PROTOCOLANY:
		return nil
	}
	return errors.New("Protocol is invalid")
}

// ValidatePort checks an string to be a valid TCP port
func validatePort(port int, ptype string) error {
	if port < 0 || port > 65535 {
		return fmt.Errorf("%s Port (%d) is out of range [0 - 65535]", ptype, port)
	}

	return nil
}
