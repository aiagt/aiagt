package workflow

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino/compose"
	"log"
	"reflect"
)

const (
	NodeNameStart       = "start"
	NodeNameEnd         = "end"
	BatchNodeOutputName = "output_array"
)

func NodeLambda(node string, mapper ObjectMapper, lambda func(ctx context.Context, input Object) (Object, error)) *compose.Lambda {
	return compose.InvokableLambda(func(ctx context.Context, chainContext ChainContext) (ChainContext, error) {
		input := chainContext.MapperObject(mapper)
		log.Printf("Node: %s, input: %s", node, input.Pretty())

		output, err := lambda(ctx, input)
		if err != nil {
			return nil, err
		}
		log.Printf("Node: %s, output: %s", node, output.Pretty())

		return chainContext.AddNote(node, output), nil
	})
}

func NodeLambdaStart(lambda func(ctx context.Context, input Object) (Object, error)) *compose.Lambda {
	return compose.InvokableLambda(func(ctx context.Context, input Object) (ChainContext, error) {
		log.Printf("Node: %s, input: %s", NodeNameStart, input.Pretty())

		output, err := lambda(ctx, input)
		if err != nil {
			return nil, err
		}
		log.Printf("Node: %s, output: %s", NodeNameStart, output.Pretty())

		return NewChainContext().AddNote(NodeNameStart, output), nil
	})
}

func NodeLambdaEnd(mapper ObjectMapper, lambda func(ctx context.Context, input Object) (Object, error)) *compose.Lambda {
	return compose.InvokableLambda(func(ctx context.Context, chainContext ChainContext) (Object, error) {
		input := chainContext.MapperObject(mapper)
		log.Printf("Node: %s, input: %s", NodeNameEnd, input.Pretty())

		output, err := lambda(ctx, input)
		if err != nil {
			return nil, err
		}
		log.Printf("Node: %s, output: %s", NodeNameEnd, output.Pretty())

		return output, nil
	})
}

func NodeLambdaBatch(node string, mapper ObjectMapper, splitter Splitter, lambda func(ctx context.Context, input Object) (Object, error)) *compose.Lambda {
	return compose.InvokableLambda(func(ctx context.Context, chainContext ChainContext) (ChainContext, error) {
		input := chainContext.MapperObject(mapper)
		log.Printf("Node: %s, input: %s", node, input.Pretty())

		inputs, err := splitter(ctx, chainContext, input)
		if err != nil {
			return nil, err
		}

		log.Printf("Node: %s, batch_inputs: %s", node, pretty(inputs))

		outputs := make([]Object, 0, len(inputs))
		for _, in := range inputs {
			if out, err := lambda(ctx, in); err != nil {
				return nil, err
			} else {
				outputs = append(outputs, out)
			}
		}

		output := Object{BatchNodeOutputName: outputs}
		log.Printf("Node: %s, output: %s", node, output.Pretty())

		return chainContext.AddNote(node, output), nil
	})
}

type Splitter func(ctx context.Context, chainCtx ChainContext, input Object) ([]Object, error)

func ArraySplitter(field *ObjectField) Splitter {
	return func(ctx context.Context, chainCtx ChainContext, input Object) ([]Object, error) {
		var result []Object

		value := chainCtx.Node(field.RefNode).Get(field.RefPath)
		if value == nil {
			return nil, nil
		}

		switch arr := value.(type) {
		case []int:
			result = splitArrayToObjects(input, arr, field.Name)
		case []float64:
			result = splitArrayToObjects(input, arr, field.Name)
		case []string:
			result = splitArrayToObjects(input, arr, field.Name)
		case []bool:
			result = splitArrayToObjects(input, arr, field.Name)
		case []Object:
			result = splitArrayToObjects(input, arr, field.Name)
		case []any:
			result = splitArrayToObjects(input, arr, field.Name)
		default:
			if reflect.TypeOf(value).Kind() == reflect.Slice {
				result = splitArrayToObjectsWithReflect(input, value, field.Name)
			} else {
				return nil, fmt.Errorf("the type %T is not an array type", value)
			}
		}

		return result, nil
	}
}

func splitArrayToObjects[T any](input Object, array []T, field string) []Object {
	result := make([]Object, 0, len(array))

	for _, item := range array {
		input.With(field, item)

		result = append(result, input.Copy().With(field, item))
	}

	return result
}

func splitArrayToObjectsWithReflect(input Object, value any, field string) []Object {
	v := reflect.ValueOf(value)

	if v.Kind() != reflect.Slice {
		return nil
	}

	result := make([]Object, 0, v.Len())

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i).Interface()
		result = append(result, input.Copy().With(field, item))
	}

	return result
}
