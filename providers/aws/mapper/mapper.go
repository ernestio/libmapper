/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"errors"

	"github.com/ernestio/libmapper"
	"github.com/ernestio/libmapper/providers/aws/components"
	def "github.com/ernestio/libmapper/providers/aws/definition"
	"github.com/mitchellh/mapstructure"
	"github.com/r3labs/graph"
)

// SUPPORTEDCOMPONENTS represents all component types supported by ernest
var SUPPORTEDCOMPONENTS = []string{"vpc", "network", "instance", "security_group", "nat_gateway", "elb", "ebs", "s3", "route53", "rds_instance", "rds_cluster"}

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

	for _, c := range g.Components {
		// Build internal & template values
		for _, dep := range c.Dependencies() {
			if g.HasComponent(dep) != true {
				return g, errors.New("Could not resolve component dependency: " + dep)
			}
		}

		c.Rebuild(g)

		// Validate Components
		err := c.Validate()
		if err != nil {
			return g, err
		}

		// Build dependencies
		for _, dep := range c.Dependencies() {
			g.Connect(dep, c.GetID())
		}
	}

	return g, nil
}

// ConvertGraph : converts the service graph into an input yaml format
func (m Mapper) ConvertGraph(g *graph.Graph) (libmapper.Definition, error) {
	var d def.Definition

	for _, c := range g.Components {
		c.Rebuild(g)

		for _, dep := range c.Dependencies() {
			if g.HasComponent(dep) != true {
				return g, errors.New("Could not resolve component dependency: " + dep)
			}
		}

		err := c.Validate()
		if err != nil {
			return d, err
		}
	}

	d.Vpcs = MapDefinitionVpcs(g)
	d.Networks = MapDefinitionNetworks(g)
	d.Instances = MapDefinitionInstances(g)
	d.SecurityGroups = MapDefinitionSecurityGroups(g)

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

	g.Load(gg)

	for i := 0; i < len(g.Components); i++ {
		gc := g.Components[i].(*graph.GenericComponent)

		var c graph.Component

		switch gc.GetType() {
		case "vpc":
			c = &components.Vpc{}
		case "network":
			c = &components.Network{}
		case "instance":
			c = &components.Instance{}
		case "security_group":
			c = &components.SecurityGroup{}
		}

		err := mapstructure.Decode(i, c)
		if err != nil {
			return g, err
		}

		g.Components[i] = c
	}

	return g, nil
}

// CreateImportGraph : creates a new graph with component queries used to import components from a provider
func (m Mapper) CreateImportGraph(service string, credentials map[string]interface{}) *graph.Graph {
	g := graph.New()
	params := make(map[string]string)

	params["ernest.service"] = service

	for _, ctype := range SUPPORTEDCOMPONENTS {
		q := MapQuery(ctype, params)
		g.AddComponent(q)
	}

	creds := convertCredentials(credentials)
	g.AddComponent(&creds)

	return g
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

func convertCredentials(credentials map[string]interface{}) graph.GenericComponent {
	return credentials
}
