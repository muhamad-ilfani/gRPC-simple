gen_server:
	protoc -I server/proto server/proto/proto.proto --go_out=server/proto --go-grpc_out=require_unimplemented_servers=false:server/proto
rm_server:
	rm server/proto/*.go
gen_client:
	protoc -I client/proto client/proto/proto.proto --go_out=client/proto --go-grpc_out=require_unimplemented_servers=false:client/proto
rm_client:
	rm client/proto/*.go