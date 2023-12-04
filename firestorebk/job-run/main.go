package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	// Firestore Adminクライアントの作成
	client, err := admin.NewFirestoreAdminClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Firestore Admin client: %v", err)
	}
	log.Print("Firestore Adminクライアントの作成済み.")
	defer client.Close()
	// バックアップリクエストの構築
	request := &adminpb.ExportDocumentsRequest{
		Name:            fmt.Sprintf("projects/%s/databases/%s", projectID, databaseID),
		CollectionIds:   nil, // データベース全体のバックアップ
		OutputUriPrefix: fmt.Sprintf("gs://%s/firestore-backup/%s", bucketName, time.Now().Format("2006-01-02-15:04:05")),
	}
	// バックアップの実行
	log.Print("Firestore database backup バックアップの実行する.")
	op, err := client.ExportDocuments(ctx, request)
	if err != nil {
		log.Fatalf("Failed to export documents: %v", err)
	}
	// バックアップの完了を待機
	if _, err := op.Wait(ctx); err != nil {
		log.Fatalf("Failed to wait for export operation to complete: %v", err)
	}
	log.Print("Firestore database backup completed successfully!")
}
