/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"errors"

	"github.com/ernestio/libmapper"
	def "github.com/ernestio/libmapper/providers/aws/definition"
	"github.com/r3labs/graph"
)

type Mapper struct{}

func (m Mapper) ConvertDefinition(gd libmapper.Definition) (*graph.Graph, error) {
	g := graph.New()

	d, ok := gd.(def.Definition)
	if ok != true {
		return g, errors.New("Could not convert generic definition into aws format.")
	}

	// Map basic component values from definition
	mapComponents(d, g)

	for _, component := range g.Components {
		// Build internal & template values
		for _, dep := range component.Dependencies() {
			if g.HasComponent(dep) != true {
				return g, errors.New("Could not resolve component dependency: " + dep)
			}
		}

		// Rebuild internal values
		err := component.Rebuild(g)
		if err != nil {
			return g, err
		}

		// Validate Components
		err = component.Validate()
		if err != nil {
			return g, err
		}
	}

	return g, nil
}

func (m Mapper) ConvertGraph(g *graph.Graph) (libmapper.Definition, error) {
	var d libmapper.Definition

	return d, nil
}

func (m Mapper) SupportedComponents() []string {
	var supported []string
	return supported
}

func mapComponents(d *def.Definition, g *graph.Graph) error {
	for _, network := range MapNetworks(d) {
		err := g.AddComponent(network)
		if err != nil {
			return err
		}
	}

	// Map basic component values from definition
	for _, instance := range MapInstances(d) {
		err := g.AddComponent(instance)
		if err != nil {
			return err
		}
	}
}
