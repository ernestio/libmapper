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

// Mapper : implements the generic mapper structure
type Mapper struct{}

// New : returns a new aws mapper
func New() *Mapper {
	return &Mapper{}
}

// ConvertDefinition : converts the input yaml definition to a graph format
func (m Mapper) ConvertDefinition(gd libmapper.Definition) (*graph.Graph, error) {
	g := graph.New()

	d, ok := gd.(*def.Definition)
	if ok != true {
		return g, errors.New("Could not convert generic definition into aws format.")
	}

	// Map basic component values from definition
	err := mapComponents(d, g)
	if err != nil {
		return g, err
	}

	for _, component := range g.Components {
		// Build internal & template values
		for _, dep := range component.Dependencies() {
			if g.HasComponent(dep) != true {
				return g, errors.New("Could not resolve component dependency: " + dep)
			}
		}

		component.Rebuild(g)

		// Validate Components
		err := component.Validate()
		if err != nil {
			return g, err
		}

		// Build dependencies
		for _, dep := range component.Dependencies() {
			g.Connect(dep, component.GetID())
		}
	}

	return g, nil
}

// ConvertGraph : converts the service graph into an input yaml format
func (m Mapper) ConvertGraph(g *graph.Graph) (libmapper.Definition, error) {
	var d libmapper.Definition

	return d, nil
}

// LoadDefinition : returns an aws type definition
func (m Mapper) LoadDefinition(gd map[string]interface{}) (libmapper.Definition, error) {
	var d def.Definition

	err := d.LoadMap(gd)

	return &d, err
}

// LoadGraph : returns a generic interal graph
func (m Mapper) LoadGraph(gg map[string]interface{}) (*graph.Graph, error) {
	g := graph.New()

	return g, nil
}

func mapComponents(d *def.Definition, g *graph.Graph) error {
	// Map basic component values from definition

	for _, vpc := range MapVpcs(d) {
		err := g.AddComponent(vpc)
		if err != nil {
			return err
		}
	}

	for _, network := range MapNetworks(d) {
		err := g.AddComponent(network)
		if err != nil {
			return err
		}
	}

	for _, instance := range MapInstances(d) {
		err := g.AddComponent(instance)
		if err != nil {
			return err
		}
	}

	for _, securitygroup := range MapSecurityGroups(d) {
		err := g.AddComponent(securitygroup)
		if err != nil {
			return err
		}
	}

	return nil
}

func mapTags(name, service string) map[string]string {
	tags := make(map[string]string)

	tags["Name"] = name
	tags["ernest.service"] = service

	return tags
}
