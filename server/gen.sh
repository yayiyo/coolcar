PROTO_PATH=./auth/api
GO_OUT_PATH=./auth/api/gen/v1
PROTO_FILE=auth

if [ $1 ]; then
    PROTO_PATH=$1
fi

if [ $2 ]; then
    GO_OUT_PATH=$2
fi

if [ $3 ]; then
    PROTO_FILE=$3
fi

mkdir -p $GO_OUT_PATH

# proto.pb.go
protoc -I $PROTO_PATH --go_out $GO_OUT_PATH \
  --go_opt paths=source_relative \
  --go-grpc_out $GO_OUT_PATH \
  --go-grpc_opt paths=source_relative \
  $PROTO_FILE.proto

## proto.pb.gw.go
#protoc -I .  --grpc-gateway_out ./gen/go \
#  --grpc-gateway_opt logtostderr=true \
#  --grpc-gateway_opt paths=source_relative \
#  --grpc-gateway_opt standalone=true \
#  --grpc-gateway_opt grpc_api_configuration=trip.yaml \
#  trip.proto

# V1 生成 gateway
protoc -I=$PROTO_PATH --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/$PROTO_FILE.yaml:$GO_OUT_PATH $PROTO_FILE.proto