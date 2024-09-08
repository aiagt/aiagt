.PHONY: gen-rpc install init-rpc

install:
	@go install mvdan.cc/gofumpt@v0.6.0 && \
	cd tools/gen_handler && go install gen_handler.go && cd ../.. && \
	cd tools/init_service && go install . && cd ../..

gen-rpc:
	@mkdir -p app/${svc} && \
	kitex -module github.com/aiagt/aiagt idl/${svc}.thrift && \
	cd app/${svc} && \
	kitex -module github.com/aiagt/aiagt -service ${svc}service -use github.com/aiagt/aiagt/kitex_gen -I ../../idl/ ../../idl/${svc}.thrift && \
	gen_handler --service_path=. --remove_handler=true

model_list = $(foreach model,$(models),-model $(model))

init-rpc:
	@cd app/${svc} && \
	gen_handler --service_path=. --remove_handler=true && \
	init_service --service_path=. --service_name=${svc} ${model_list}
