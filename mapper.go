package libmapper

import "github.com/r3labs/graph"

// Mapper : interface for each provider mapper to satisfy
type Mapper interface {
	ConvertDefinition(Definition) (*graph.Graph, error)
	ConvertGraph(*graph.Graph) (Definition, error)
	LoadDefinition(map[string]interface{}) (Definition, error)
	LoadGraph(map[string]interface{}) (*graph.Graph, error)
	//SupportedComponents() []string // returns a list of supported components for constructing an import query
}
