package p

import (
	"context"
	"fmt"
	"log"
	"time"

	admin "cloud.google.com/go/firestore/apiv1/admin"
	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
	adminpb "google.golang.org/genproto/googleapis/firestore/admin/v1"
)

// Config 包含配置信息
type Config struct {
	ProjectID  string
	DatabaseID string
	BucketName string
}

// HelloPubSub 消费 Pub/Sub 消息
func HelloPubSub(ctx context.Context, m *pubsub.Message) error {
	log.Println(string(m.Data))

	// 配置信息
	config := Config{
		ProjectID:  "nova-hj", // 替换为你的项目 ID
		DatabaseID: "(default)",
		BucketName: "nova-hj-job",
	}

	// 创建 Firestore Admin 客户端
	client, err := admin.NewFirestoreAdminClient(ctx, option.WithCredentials(credentials))
	if err != nil {
		log.Fatalf("Failed to create Firestore Admin client: %v", err)
	}
	defer client.Close()

	// 构建备份请求
	backupFolder := time.Now().Format("2006-01-02-15:04:05")
	request := &adminpb.ExportDocumentsRequest{
		Name:            fmt.Sprintf("projects/%s/databases/%s", config.ProjectID, config.DatabaseID),
		CollectionIds:   nil,
		OutputUriPrefix: fmt.Sprintf("gs://%s/firestore-backup/%s", config.BucketName, backupFolder),
	}

	// 执行备份操作
	log.Print("执行 Firestore 数据库备份...")
	op, err := client.ExportDocuments(ctx, request)
	if err != nil {
		log.Fatalf("Failed to export documents: %v", err)
	}

	// 等待备份操作完成
	if _, err := op.Wait(ctx); err != nil {
		log.Fatalf("Failed to wait for export operation to complete: %v", err)
	}

	log.Print("Firestore 数据库备份完成!")

	return nil
}
