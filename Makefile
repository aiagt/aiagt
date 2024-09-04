.PHONY: gen-rpc install init-rpc

install:
	@go install mvdan.cc/gofumpt@latest && \
	go install ./tool/gen_handler/gen_handler.go

gen-rpc:
	@mkdir -p app/${svc} && \
	kitex -module github.com/aiagt/aiagt idl/${svc}.thrift && \
	cd app/${svc} && \
	kitex -module github.com/aiagt/aiagt -service ${svc}service -use github.com/aiagt/aiagt/kitex_gen -I ../../idl/ ../../idl/${svc}.thrift && \
	gen_handler --service_path=. --remove_handler=true

init-rpc:
	@cd app/${svc} && \
	gen_handler --service_path=. --remove_handler=true
