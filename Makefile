
generate:
	protoc proto/*.proto -I proto --go_out ./internal --go-grpc_out ./internal --openapiv2_out ./docs --openapiv2_opt logtostderr=true --grpc-gateway_out ./internal --grpc-gateway_opt logtostderr=true