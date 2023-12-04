package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore/apiv1/admin"
	"google.golang.org/api/option"
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
	client, err := admin.NewFirestoreAdminClient(ctx, option.WithCredentialsFile("path/to/your/service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Firestore Admin client: %v", err)
	}
	defer client.Close()

	// 构造备份请求
	request := &adminpb.ExportDocumentsRequest{
		Name:            fmt.Sprintf("projects/%s/databases/%s", projectID, databaseID),
		CollectionIds:   nil, // 备份整个数据库
		OutputUriPrefix: fmt.Sprintf("gs://%s/firestore-backup", bucketName),
	}

	// 执行备份
	op, err := client.ExportDocuments(ctx, request)
	if err != nil {
		log.Fatalf("Failed to export documents: %v", err)
	}

	// 等待备份完成
	if _, err := op.Wait(ctx); err != nil {
		log.Fatalf("Failed to wait for export operation to complete: %v", err)
	}

	log.Print("Firestore database backup completed successfully!")
}
