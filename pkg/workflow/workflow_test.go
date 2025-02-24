package workflow

import (
	"context"
	"github.com/aiagt/aiagt/pkg/schema"
	"github.com/cloudwego/eino-ext/devops"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func TestWorkflow(t *testing.T) {
	workflow := NewWorkflow()

	//workflow.AppendLambdaNodeStart(func(ctx context.Context, input Object) (Object, error) {
	//	return input, nil
	//})
	workflow.AppendNode(NewStartNode())

	workflow.AppendNode(&Node{
		Name: "read_project_structure",
		InputMapper: ObjectMapper{
			{Name: "owner", RefNode: "start", RefPath: "owner"},
			{Name: "repo", RefNode: "start", RefPath: "repo"},
			{Name: "path", Constant: "/"},
			{Name: "recursion", Constant: true},
		},
		Runner: NewFunctionNodeRunner(func(ctx context.Context, input Object) (Object, error) {
			return Object{"children": []Object{
				{"path": "apps/plugin", "type": "dir"},
				{"path": "apps/plugin/Dockerfile", "type": "file"},
				{"path": "apps/plugin/build.sh", "type": "file"},
				{"path": "apps/plugin/conf", "type": "dir"},
				{"path": "apps/plugin/conf/conf-release.yaml", "type": "file"},
				{"path": "apps/plugin/conf/conf.go", "type": "file"},
				{"path": "apps/plugin/conf/conf.yaml", "type": "file"},
				{"path": "apps/plugin/controller", "type": "dir"},
				{"path": "apps/plugin/controller/controller.go", "type": "file"},
				{"path": "apps/plugin/dal", "type": "dir"},
				{"path": "apps/plugin/dal/db", "type": "dir"},
				{"path": "apps/plugin/dal/db/labels.go", "type": "file"},
				{"path": "apps/plugin/dal/db/plugin.go", "type": "file"},
				{"path": "apps/plugin/dal/db/tool.go", "type": "file"},
				{"path": "apps/plugin/handler", "type": "dir"},
				{"path": "apps/plugin/kitex_info.yaml", "type": "file"},
				{"path": "apps/plugin/main.go", "type": "file"},
				{"path": "apps/plugin/mapper", "type": "dir"},
				{"path": "apps/plugin/mapper/mapper.go", "type": "file"},
				{"path": "apps/plugin/model", "type": "dir"},
				{"path": "apps/plugin/model/base.go", "type": "file"},
				{"path": "apps/plugin/model/label.go", "type": "file"},
				{"path": "apps/plugin/model/plugin.go", "type": "file"},
				{"path": "apps/plugin/model/tool.go", "type": "file"},
				{"path": "apps/plugin/script", "type": "dir"},
				{"path": "apps/plugin/script/bootstrap.sh", "type": "file"},
				{"path": "apps/plugin/test", "type": "dir"},
				{"path": "apps/plugin/test/test.go", "type": "file"},
				{"path": "pkg", "type": "dir"},
				{"path": "pkg/caller", "type": "dir"},
				{"path": "pkg/caller/call.go", "type": "file"},
				{"path": "pkg/caller/types.go", "type": "file"},
				{"path": "pkg/closer", "type": "dir"},
				{"path": "pkg/closer/closer.go", "type": "file"},
				{"path": "pkg/hash", "type": "dir"},
			}}, nil
		}),
	})

	const (
		systemPrompt = `The following is the complete directory structure of warehouse {owner}/{repo}:
{tree}
Now you need to find the code according to the user. Please return the path to the directory where the code may exist in an array in JSON format.`
		userPrompt = `{query}`
	)

	workflow.AppendNode(&Node{
		Name: "extract_target_directory",
		InputMapper: ObjectMapper{
			{Name: "owner", RefNode: "start", RefPath: "owner"},
			{Name: "repo", RefNode: "start", RefPath: "repo"},
			{Name: "query", RefNode: "start", RefPath: "query"},
			{Name: "tree", RefNode: "read_project_structure", RefPath: "children"},
		},
		Runner: NewLLMNodeRunner(OpenaiBaseUrl, OpenaiApiKey, ModelName, systemPrompt, userPrompt, map[string]schema.Definition{
			"dirs": {
				Type:        "array",
				Description: "target dirs",
				Items: &schema.Definition{
					Type: "string",
				},
			},
		}),
	})

	workflow.AppendNode(&Node{
		Name: "read_target_trees",
		InputMapper: ObjectMapper{
			{Name: "owner", RefNode: "start", RefPath: "owner"},
			{Name: "repo", RefNode: "start", RefPath: "repo"},
			{Name: "recursion", Constant: true},
		},
		BatchField: &ObjectField{
			Name:    "dir",
			RefNode: "extract_target_directory",
			RefPath: "dirs",
		},
		Runner: NewFunctionNodeRunner(func(ctx context.Context, input Object) (Object, error) {
			return Object{
				"children": []Object{{"path": input.String("dir"), "type": "file"}},
			}, nil
		}),
	})

	workflow.AppendNode(&Node{
		Name: "extract_files_by_trees",
		InputMapper: ObjectMapper{
			{Name: "trees", RefNode: "read_target_trees", BatchOutput: true},
		},
		Runner: NewFunctionNodeRunner(func(ctx context.Context, input Object) (Object, error) {
			trees := input.ObjectArray("trees")

			var files []string
			for _, tree := range trees {
				for _, child := range tree.ObjectArray("children") {
					if child.String("type") == "file" {
						files = append(files, child.String("path"))
					}
				}
			}

			return Object{"files": files}, nil
		}),
	})

	workflow.AppendNode(&Node{
		Name: "read_files_content",
		InputMapper: ObjectMapper{
			{Name: "owner", RefNode: "start", RefPath: "owner"},
			{Name: "repo", RefNode: "start", RefPath: "repo"},
			{Name: "files", RefNode: "extract_files_by_trees", RefPath: "files"},
		},
		Runner: NewFunctionNodeRunner(func(ctx context.Context, input Object) (Object, error) {
			return Object{
				"files": Object{
					"path":    "README.md",
					"content": "this is README.",
				},
			}, nil
		}),
	})

	workflow.AppendNode(NewEndNode(ObjectMapper{
		{Name: "files", RefNode: "read_files_content", RefPath: "files"},
	}))

	ctx := context.Background()

	type User struct {
		Name string
		Age  int
	}

	err := devops.Init(ctx, devops.AppendType(Object{}))
	if err != nil {
		log.Printf("[eino dev] init failed, err=%v", err)
		return
	}

	//_, _ = workflow.Chain.Compile(ctx)
	//
	//// Blocking process exits
	//sigs := make(chan os.Signal, 1)
	//signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//<-sigs
	//
	//// Exit
	//log.Printf("[eino dev] shutting down\n")

	output, err := workflow.Run(context.TODO(), Object{
		"owner": "aiagt",
		"repo":  "aiagt",
		"query": "how to call plugin",
	})
	if err != nil {
		panic(err)
	}

	log.Println("output:", output.Pretty())
}

func TestT(t *testing.T) {
	file, err := os.ReadFile("./test.json")
	require.NoError(t, err)

	object, err := NewJSONObject(file)
	require.NoError(t, err)

	//children := object.ObjectArray("children")
	//t.Log(children)

	t.Log(pretty(foreach(object)))
}

func foreach(obj Object) []Object {
	var result []Object

	if obj.String("path") != "" {
		result = append(result, Object{
			"path": obj.String("path"),
			"type": obj.String("type"),
		})
	}

	children := obj.ObjectArray("children")
	for _, child := range children {
		result = append(result, foreach(child)...)
	}

	return result
}
