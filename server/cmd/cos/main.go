package main

import (
	"context"
	"fmt"
	"log"

	blobpb "coolcar/blob/api/gen/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := blobpb.NewBlobServiceClient(conn)
	ctx := context.Background()
	//cbr, err := client.CreateBlob(ctx, &blobpb.CreateBlobRequest{
	//	AccountId:           "account-1",
	//	UploadUrlTimeoutSec: 1000,
	//})

	//resp, err := client.GetBlob(ctx, &blobpb.GetBlobRequest{
	//	Id: "63ef3143edd8b65e78b7c23c",
	//})

	resp, err := client.GetBlobURL(ctx, &blobpb.GetBlobURLRequest{
		Id:         "63ef3143edd8b65e78b7c23c",
		TimeoutSec: 100,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
