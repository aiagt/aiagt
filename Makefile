.PHONY: gen-rpc
gen-rpc:
	mkdir -p app/${svc} && kitex -module github.com/aiagt/aiagt idl/${svc}.thrift && cd app/${svc} && kitex -module github.com/aiagt/aiagt -service ${svc}service -use github.com/aiagt/aiagt/kitex_gen -I ../../idl/ ../../idl/${svc}.thrift && cd ../.. && gen_handler --service_path=app/${svc} --remove_handler=true
