const firestore = require('@google-cloud/firestore');

// Firestoreクライアントの作成
function createFirestoreClient() {
  return new firestore.v1.FirestoreAdminClient();
}

// バケット名の定義
const BUCKET_NAME = 'nova-hj-job'; // バケット名

// バケット名の作成
function createBucketName() {
  return `gs://${BUCKET_NAME}`;
}

// データベース名の作成
function createDatabaseName() {
  const client = createFirestoreClient();
  return client.databasePath('nova-hj', '(default)');
}

// ドキュメントのエクスポート
async function exportFirestoreDocuments(client, databaseName, bucket) {
  console.log("FirestoreAdminClient 作成");

  console.log("exportDocuments 実行。。。");
  try {
    const responses = await client.exportDocuments({
      name: databaseName,
      outputUriPrefix: bucket,
      collectionIds: []
    });

    const response = responses[0];

    if (!response || !response.name) {
      throw new Error("エクスポートの応答が無効です。");
    }

    console.log("exportDocuments 完了");
    console.log(`操作Name operations ID : ${response.name}`);
    return response;
  } catch (err) {
    console.error("エクスポート中にエラーが発生しました:", err.message);
    throw err; // エラーを再スロー
  }
}

// Cloud Functionsエクスポートトリガー
exports.scheduledFirestoreExport = async (event, context) => {
  try {
    // Firestoreクライアントの作成
    const client = createFirestoreClient();

    // バケットとデータベース名の作成
    const bucket = createBucketName();
    const databaseName = createDatabaseName();

    // ドキュメントのエクスポート
    const response = await exportFirestoreDocuments(client, databaseName, bucket);
    return response;
  } catch (err) {
    console.error("Cloud Function でエラーが発生しました:", err.message);
    throw err; // エラーを再スロー
  }
};
