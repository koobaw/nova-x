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

	// Firestore Admin クライアントの作成
	client, err := createFirestoreAdminClient(ctx)
	if err != nil {
		log.Fatalf("Firestore Admin クライアントの作成に失敗しました: %v", err)
	}
	defer client.Close()

	// バックアップの実行
	if err := runFirestoreBackup(ctx, client); err != nil {
		log.Fatalf("Firestore データベースのバックアップに失敗しました: %v", err)
	}

	log.Print("Firestore データベースのバックアップが正常に完了しました！")
}

func createFirestoreAdminClient(ctx context.Context) (*admin.FirestoreAdminClient, error) {
	client, err := admin.NewFirestoreAdminClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Firestore Admin クライアントの作成に失敗しました: %v", err)
	}
	log.Print("Firestore Admin クライアントの作成済み.")
	return client, nil
}

func runFirestoreBackup(ctx context.Context, client *admin.FirestoreAdminClient) error {
	// バックアップリクエストの構築
	request := &adminpb.ExportDocumentsRequest{
		Name:            fmt.Sprintf("projects/%s/databases/%s", projectID, databaseID),
		CollectionIds:   nil, // データベース全体のバックアップ
		OutputUriPrefix: fmt.Sprintf("gs://%s/firestore-backup/%s", bucketName, time.Now().Format("2006-01-02-15:04:05")),
	}

	log.Print("Firestore データベースのバックアップを実行します。")
	// バックアップの実行
	op, err := client.ExportDocuments(ctx, request)
	if err != nil {
		return fmt.Errorf("ドキュメントのエクスポートに失敗しました: %v", err)
	}

	// バックアップの完了を待機
	if _, err := op.Wait(ctx); err != nil {
		return fmt.Errorf("エクスポート操作の完了待機に失敗しました: %v", err)
	}

	return nil
}
