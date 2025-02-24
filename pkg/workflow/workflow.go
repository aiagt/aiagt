package workflow

import (
	"context"
	"github.com/cloudwego/eino/compose"
)

type Workflow struct {
	Chain *compose.Chain[Object, Object]
}

func NewWorkflow() *Workflow {
	return &Workflow{Chain: compose.NewChain[Object, Object]()}
}

func (f *Workflow) AppendEinoLambda(lambda *compose.Lambda) {
	f.Chain.AppendLambda(lambda)
}

func (f *Workflow) AppendLambdaNode(node string, mapper ObjectMapper, lambda func(ctx context.Context, input Object) (Object, error)) {
	f.Chain.AppendLambda(NodeLambda(node, mapper, lambda))
}

func (f *Workflow) AppendLambdaNodeStart(lambda func(ctx context.Context, input Object) (Object, error)) {
	f.Chain.AppendLambda(NodeLambdaStart(lambda))
}

func (f *Workflow) AppendLambdaNodeEnd(mapper ObjectMapper, lambda func(ctx context.Context, input Object) (Object, error)) {
	f.Chain.AppendLambda(NodeLambdaEnd(mapper, lambda))
}

func (f *Workflow) AppendLambdaNodeBatch(node string, mapper ObjectMapper, splitter Splitter, lambda func(ctx context.Context, input Object) (Object, error)) {
	f.Chain.AppendLambda(NodeLambdaBatch(node, mapper, splitter, lambda))
}

func (f *Workflow) AppendNode(node *Node) {
	f.Chain.AppendLambda(node.Lambda())
}

func (f *Workflow) AppendParallelNodes(nodes ...*Node) {
	parallel := compose.NewParallel()
	for _, node := range nodes {
		parallel.AddLambda(node.Name, node.Lambda())
	}

	f.Chain.AppendParallel(parallel)
}

func (f *Workflow) Run(ctx context.Context, input Object) (Object, error) {
	r, err := f.Chain.Compile(ctx)
	if err != nil {
		return nil, err
	}

	output, err := r.Invoke(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}
