const firestore = require('@google-cloud/firestore');

// 定数を定義
const BUCKET_NAME = 'nova-hj-job'; // バケット名
const PROJECT_ID = 'nova-hj';      // プロジェクトID
const DATABASE_INSTANCE = '(default)'; // データベースインスタンス

// Firestoreクライアントの作成
function createFirestoreClient() {
  return new firestore.v1.FirestoreAdminClient();
}

// バケット名の作成
function createBucketName() {
  return `gs://${BUCKET_NAME}`;
}

// データベース名の作成
function createDatabaseName(client) {
  return client.databasePath(PROJECT_ID, DATABASE_INSTANCE);
}

// ドキュメントのエクスポート
async function exportFirestoreDocuments(client, databaseName, bucket) {
  console.log("exportDocuments 実行。。。");
  const responses = await client.exportDocuments({
    name: databaseName,
    outputUriPrefix: bucket,
    collectionIds: []
  });
  const response = responses[0];
  console.log("exportDocuments 完了");
  console.log(`操作Name operations ID : ${response['name']}`);
}

async function main() {
  try {
    console.log("FirestoreAdminClient 作成");

    // Firestoreクライアントの作成
    const client = createFirestoreClient();
    
    // バケットとデータベース名の作成
    const bucket = createBucketName();
    const databaseName = createDatabaseName(client);

    // ドキュメントのエクスポート
    await exportFirestoreDocuments(client, databaseName, bucket);
  } catch (err) {
    console.error(err);
    process.exit(1); // エラーコードでプロセスを終了
  }
}

main();
