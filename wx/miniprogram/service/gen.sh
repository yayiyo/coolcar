# 生成TS pb
PROTO_DIR=../../../server/proto

npx protoc --ts_out proto_gen  --proto_path $PROTO_DIR $PROTO_DIR/trip.proto