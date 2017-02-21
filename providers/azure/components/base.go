/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package components

// Base : Shared internal component fields
type Base struct {
	ProviderType  string `json:"_provider"`
	ComponentID   string `json:"_component_id"`
	ComponentType string `json:"_component"`
	State         string `json:"_state"`
	Action        string `json:"_action"`
}
