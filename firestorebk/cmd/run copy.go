package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore/apiv1/admin"
)

const (
	projectID    = "nova-hj"
	bucketName   = "nova-hj-job"
	databaseName = "projects/nova-hj/databases/(default)"
)

func backupFirestoreToGCS() error {
	ctx := context.Background()

	// 创建 Firestore Admin 客户端
	client, err := admin.NewFirestoreAdminClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	// 创建备份请求
	request := &admin.ExportDocumentsRequest{
		Name:            databaseName,
		OutputUriPrefix: "gs://" + bucketName + "/firestore-backup",
	}

	// 发起备份操作
	operation, err := client.ExportDocuments(ctx, request)
	if err != nil {
		return err
	}

	// 等待备份操作完成
	_, err = operation.Wait(ctx)
	if err != nil {
		return err
	}

	log.Println("Firestore backup completed successfully.")
	return nil
}

func main() {
	log.Print("Hello, Cloud Run Job!")

	// 执行 Firestore 备份到 GCS 操作
	if err := backupFirestoreToGCS(); err != nil {
		log.Printf("Error backing up Firestore: %v", err)
	}

	log.Print("Your task has been successfully completed!!!")
}
