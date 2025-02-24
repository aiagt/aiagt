package workflow

type ChainContext map[string]Object

func NewChainContext() ChainContext {
	return make(map[string]Object)
}

func (c ChainContext) AddNote(name string, note Object) ChainContext {
	c[name] = note
	return c
}

func (c ChainContext) Node(name string) Object {
	return c[name]
}

func (c ChainContext) MapperObject(mapper ObjectMapper) Object {
	result := NewObject()

	for _, field := range mapper {
		if field.Constant != nil {
			result.With(field.Name, field.Constant)
		} else if field.BatchOutput {
			result.With(field.Name, c.Node(field.RefNode).GetByPaths([]string{BatchNodeOutputName}))
		} else {
			result.With(field.Name, c.Node(field.RefNode).GetByPaths(field.Paths()))
		}
	}

	return result
}

func (c ChainContext) Pretty() string {
	return pretty(c)
}

func (c ChainContext) PrettyIndent() string {
	return prettyIndent(c)
}
