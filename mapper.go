package libmapper

import (
	graph "gopkg.in/r3labs/graph.v2"
)

// Mapper : interface for each provider mapper to satisfy. A mapper
// is basically managing all translations between a Definition (the
// user interface of an ernest service) and the Graph (the internal
// representation of this service).
// Additionally Graph object includes the necessary steps to build a
// specific service
type Mapper interface {

	// ConvertDefinition : Given the input Definition it returns a valid Graph ("service") object
	ConvertDefinition(Definition) (*graph.Graph, error)

	// ConvertGraph : Given a valid Graph("service") object it will build a valid Definition
	ConvertGraph(*graph.Graph) (Definition, error)

	// LoadDefinition : ...
	LoadDefinition(map[string]interface{}) (Definition, error)

	// LoadGraph : ...
	LoadGraph(map[string]interface{}) (*graph.Graph, error)

	// CreateImportGraph : Creates an import graph based on the specified service_id
	CreateImportGraph([]string) *graph.Graph

	// ProviderCredentials : Returns a provider specific mapped component
	ProviderCredentials(map[string]interface{}) graph.Component
}
