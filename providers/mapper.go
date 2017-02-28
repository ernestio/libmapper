/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package providers

import (
	"github.com/ernestio/libmapper"
	aws "github.com/ernestio/libmapper/providers/aws/mapper"
)

// NewMapper : Get a new mapper based on a specified type
func NewMapper(t string) (m libmapper.Mapper) {
	switch t {
	case "aws":
		m = aws.New()
	}

	return m
}
