const firestore = require('@google-cloud/firestore');

const BUCKET_NAME = 'nova-hj-job';
const PROJECT_ID = 'nova-hj';
const DATABASE_INSTANCE = '(default)';

async function main() {
  try {
    console.log("FirestoreAdminClient 作成");

    const client = new firestore.v1.FirestoreAdminClient();
    
    const bucket = `gs://${BUCKET_NAME}`;
    const databaseName = client.databasePath(PROJECT_ID, DATABASE_INSTANCE);

    console.log("exportDocuments 実行。。。");

    const responses = await client.exportDocuments({
      name: databaseName,
      outputUriPrefix: bucket,
      collectionIds: []
    });

    const response = responses[0];

    console.log("exportDocuments 完了");
    console.log(`操作Name: ${response['name']}`);
  } catch (err) {
    console.error(err);
    process.exit(1); // Exit with an error code
  }
}

main();
