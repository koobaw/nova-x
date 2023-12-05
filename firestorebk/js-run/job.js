const firestore = require('@google-cloud/firestore');
const { Storage } = require("@google-cloud/storage");

async function main() {
  try {
    console.log("CCCCCCCCCCCCCC");

    const client = new firestore.v1.FirestoreAdminClient();
    const bucket = 'gs://nova-hj-job';
    const databaseName = client.databasePath('nova-hj', '(default)');

    console.log("AAAAAAAAAAAAAAA");

    const responses = await client.exportDocuments({
      name: databaseName,
      outputUriPrefix: bucket,
      collectionIds: []
    });

    const response = responses[0];

    console.log("BBBBBBBBBBBBBB");
    console.log(`Operation Name KKKKKKK: ${response['name']}`);
  } catch (err) {
    console.error(err);
    process.exit(1); // Exit with an error code
  }
}

main();
