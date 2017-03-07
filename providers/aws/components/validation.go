/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package components

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	// PROTOCOLTCP : TCP Protocol value
	PROTOCOLTCP = "tcp"
	// PROTOCOLUDP : UDP Protocol value
	PROTOCOLUDP = "udp"
	// PROTOCOLANY : Any Protocol value
	PROTOCOLANY = "-1"
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

func validateDateTimeFormat(t string) error {
	// ddd:hh24:mi
	var days = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun", "mon", "tue", "wed", "thu", "fri", "sat", "sun"}

	parts := strings.Split(t, ":")
	if len(parts) != 3 {
		return errors.New("Date format must take the form of 'ddd:hh24:mi'. i.e. 'Mon:21:30'")
	}

	// is valid day
	if isOneOf(days, parts[0]) != true {
		return fmt.Errorf("Date format invalid. Day must be one of %s", strings.Join(days, ", "))
	}

	// is valid hour
	d, err := strconv.Atoi(parts[1])
	if err != nil || d < 0 || d > 23 {
		return errors.New("Date format invalid. Hour must be between 0 and 23 hours")
	}

	// is valid minute
	d, err = strconv.Atoi(parts[2])
	if err != nil || d < 0 || d > 59 {
		return errors.New("Date format invalid. Minute must be between 0 and 59 minutes")
	}

	return nil
}

func validateTimeFormat(t string) error {
	parts := strings.Split(t, ":")
	if len(parts) != 2 {
		return errors.New("Time format must take the form of 'hh24:mi-hh24:mi'. i.e. '21:30-22:00'")
	}
	// is valid hour
	d, err := strconv.Atoi(parts[0])
	if err != nil || d < 0 || d > 23 {
		return errors.New("Time format invalid. Hour must be between 0 and 23 hours")
	}

	// is valid minute
	d, err = strconv.Atoi(parts[1])
	if err != nil || d < 0 || d > 59 {
		return errors.New("Time format invalid. Minute must be between 0 and 59 minutes")
	}

	return nil
}

func validateTimeWindow(w string) error {
	p := strings.Split(w, "-")
	if len(p) != 2 {
		return errors.New("Window format must take the form of 'ddd:hh24:mi-ddd:hh24:mi'. i.e. 'Mon:21:30-Mon:22:00'")
	}

	err := validateDateTimeFormat(p[0])
	if err != nil {
		return err
	}

	return validateDateTimeFormat(p[1])
}

func appendUnique(s []string, v string) []string {
	for _, x := range s {
		if x == v {
			return s
		}
	}
	return append(s, v)
}

func isOneOf(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
