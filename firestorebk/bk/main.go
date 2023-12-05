package main

import (
	"context"
	"fmt"
	"log"
	"time"

	admin "cloud.google.com/go/firestore/apiv1/admin"
	"cloud.google.com/go/storage"
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
	client, err := admin.NewFirestoreAdminClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Firestore Admin client: %v", err)
	}
	log.Print("Firestore Admin クライアントの作成済み.")
	defer client.Close()

	// バックアップリクエストの構築
	backupFolder := time.Now().Format("2006-01-02-15:04:05")
	request := &adminpb.ExportDocumentsRequest{
		Name:            fmt.Sprintf("projects/%s/databases/%s", projectID, databaseID),
		CollectionIds:   nil, // データベース全体のバックアップ
		OutputUriPrefix: fmt.Sprintf("gs://%s/firestore-backup/%s", bucketName, backupFolder),
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

	// GCS ストレージクライアントの作成
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Storage client: %v", err)
	}
	defer storageClient.Close()

	// 一週間以上前のバックアップフォルダを削除
	if err := deleteOldBackupFolders(ctx, storageClient, bucketName, backupFolder); err != nil {
		log.Fatalf("Failed to delete old backup folders: %v", err)
	}
}

// deleteOldBackupFolders は一週間以上前のバックアップフォルダを削除します。
func deleteOldBackupFolders(ctx context.Context, client *storage.Client, bucketName, currentBackupFolder string) error {
	iter := client.Bucket(bucketName).Objects(ctx, &storage.Query{
		Prefix: "firestore-backup/",
	})

	// 一週間以上前の時刻を計算
	oneWeekAgo := time.Now().Add(-7 * 24 * time.Hour)

	for {
		objAttrs, err := iter.Next()
		if err == storage.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Failed to iterate over GCS objects: %v", err)
		}

		// 解析できる形式の時刻フォーマットかどうか確認
		timestamp, err := time.Parse("2006-01-02-15:04:05", objAttrs.Prefix)
		if err == nil && timestamp.Before(oneWeekAgo) {
			// 一週間以上前のフォルダを削除
			if err := deleteFolder(ctx, client, bucketName, objAttrs.Prefix); err != nil {
				return fmt.Errorf("Failed to delete folder %s: %v", objAttrs.Prefix, err)
			}
		}
	}

	return nil
}

// deleteFolder は指定されたフォルダを削除します。
func deleteFolder(ctx context.Context, client *storage.Client, bucketName, folderName string) error {
	bucket := client.Bucket(bucketName)

	// フォルダ内のオブジェクトを削除
	iter := bucket.Objects(ctx, &storage.Query{Prefix: folderName})
	for {
		objAttrs, err := iter.Next()
		if err == storage.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Failed to iterate over GCS objects in folder %s: %v", folderName, err)
		}
		if err := bucket.Object(objAttrs.Name).Delete(ctx); err != nil {
			return fmt.Errorf("Failed to delete GCS object %s: %v", objAttrs.Name, err)
		}
	}

	// フォルダを削除
	if err := bucket.Object(folderName).Delete(ctx); err != nil {
		return fmt.Errorf("Failed to delete GCS folder %s: %v", folderName, err)
	}

	return nil
}
