# proto.pb.go
protoc -I . --go_out ./gen/go/ \
  --go_opt paths=source_relative \
  --go-grpc_out ./gen/go/ \
  --go-grpc_opt paths=source_relative \
  trip.proto

## proto.pb.gw.go
#protoc -I .  --grpc-gateway_out ./gen/go \
#  --grpc-gateway_opt logtostderr=true \
#  --grpc-gateway_opt paths=source_relative \
#  --grpc-gateway_opt standalone=true \
#  --grpc-gateway_opt grpc_api_configuration=trip.yaml \
#  trip.proto

# V1 生成 gateway
protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=trip.yaml:gen/go trip.proto
