# 生成TS pb
PROTO_DIR=../../../server/proto
OUT_DIR=./proto_gen

path=$1
fileName=$2

if [ $path ] && [ $fileName ]; then
  mkdir -p $OUT_DIR/$fileName
  npx protoc --ts_out $OUT_DIR/$fileName --proto_path $path $path/$fileName.proto
else
  npx protoc --ts_out proto_gen --proto_path $PROTO_DIR $PROTO_DIR/*.proto
fi
