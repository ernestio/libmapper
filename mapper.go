package libmapper

import "github.com/r3labs/graph"

// Mapper : interface for each provider mapper to satisfy
type Mapper interface {
	ConvertDefinition(Definition) (graph.Graph, error)
	ConvertGraph(graph.Graph) (Definition, error)
	SupportedComponents() []string // for ??? cant remember :-@
}
