const firestore = require('@google-cloud/firestore');
const client = new firestore.v1.FirestoreAdminClient();
// Replace BUCKET_NAME
const bucket = 'gs://nova-hj-job'

exports.scheduledFirestoreExport = (event, context) => {
  console.log("FirestoreAdminClient 作成");
  const databaseName = client.databasePath(
    'nova-hj',   //GCP_PROJECT
    '(default)'
  );
  console.log("exportDocuments 実行。。。");
  return client
    .exportDocuments({
      name: databaseName,
      outputUriPrefix: bucket,
      // Leave collectionIds empty to export all collections
      // or define a list of collection IDs:
      // collectionIds: ['users', 'posts']
      collectionIds: []
    })
    .then(responses => {
      const response = responses[0];
      console.log("exportDocuments 完了");
      console.log(`操作Name operations ID : ${response['name']}`);
      return response;
    })
    .catch(err => {
      console.error(err);
    });
};