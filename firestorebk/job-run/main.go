package main

import (
	"context"
	"fmt"
	"log"

	admin "cloud.google.com/go/firestore/apiv1/admin"
	adminpb "google.golang.org/genproto/googleapis/firestore/admin/v1"
)

const (
	projectID  = "nova-hj"
	databaseID = "(default)"
	bucketName = "nova-hj-job"
)

func main() {
	ctx := context.Background()

	// 创建 Firestore Admin 客户端
	client, err := admin.NewFirestoreAdminClient(ctx)
	log.Print("AAAAAAAAAAAAAA")
	if err != nil {
		log.Fatalf("Failed to create Firestore Admin client: %v", err)
	}
	log.Print("BBBBBBBBBBBB")
	defer client.Close()
	log.Print("CCCCCCCCCCCC")
	// 构造备份请求
	request := &adminpb.ExportDocumentsRequest{
		Name:            fmt.Sprintf("projects/%s/databases/%s", projectID, databaseID),
		CollectionIds:   nil, // 备份整个数据库
		OutputUriPrefix: fmt.Sprintf("gs://%s/firestore-backup", bucketName),
	}
	log.Print("AAAAAAAAAAAAAA")
	// 执行备份
	op, err := client.ExportDocuments(ctx, request)
	if err != nil {
		log.Fatalf("Failed to export documents: %v", err)
	}
	log.Print("DDDDDDDDDDDDD")
	// 等待备份完成
	if _, err := op.Wait(ctx); err != nil {
		log.Fatalf("Failed to wait for export operation to complete: %v", err)
	}
	log.Print("EEEEEEEEEEEEE")
	log.Print("Firestore database backup completed successfully!")
}
