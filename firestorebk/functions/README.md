https://cloud.google.com/functions/docs/writing/write-event-driven-functions?hl=zh-cn#cloudevent-example-go
----
https://qiita.com/iwata@github/items/2c1499d49d7d22d633c1

https://github.com/iwata/go-cloud-functions-examples/tree/master/backup-firestore

https://github.com/sinmetal/firestore_export_debug

https://github.com/ChristianAlexander/FirestoreRestore

----
https://medium.com/swlh/rollbacks-and-infinite-loops-with-firestore-and-cloud-functions-in-golang-263aa76398da


https://xebia.com/blog/how-to-read-firestore-events-with-cloud-functions-and-golang/

# Firestore backup and restore tool

This is a tool to backup and restore a Firestore database. It is written in Go and will run as a function on GCP. The 
tool can be run stand alone or with a scheduler function that backs up the database on a regular basis. The tool can 
also restore a database from a backup file. The tool is designed to be run on a GCP project that has a Firestore database
and a Cloud Storage bucket. The tool will create a folder in the bucket to store the backup files. 

## Prerequisites

* A GCP project with a Firestore database and a Cloud Storage bucket, ensure that the Firestore database is in Native mode.
* A GCP service account with the following roles:
  * Cloud Functions Developer
  * Cloud Functions Invoker
  * Cloud Scheduler Admin
  * Cloud Storage Admin
  * Firestore Admin
  * Service Account User
* Permissions to deploy Cloud Functions and Cloud Scheduler jobs

## Installation

### Deploying the backup function

```shell
export GOOGLE_PROJECT=nova-hj
export REGION=asia-northeast1
export IMAGE_URL=asia-northeast1-docker.pkg.dev/nova-hj/infra/golang-job:latest
export JOB_NAME=golang-job
export project_id="nova-hj"
export sa="golang-job@nova-hj.iam.gserviceaccount.com"
export bucket=nova-hj-job
export furl=""
gcloud config configurations activate $project_id
gcloud config set project $project_id
gcloud auth application-default set-quota-project $project_id

gcloud functions deploy firestore-backup \
  --gen2 \
  --entry-point=HelloPubSub \
  --memory=256MB \
  --region=asia-northeast1 \
  --runtime=go121 \
  --trigger-topic=functions-bk-test  \
  --min-instances 1 \
  --max-instances 3 \
  --allow-unauthenticated \
  --service-account $sa

gcloud functions delete firestore-backup

gcloud functions add-iam-policy-binding firestore-backup \
  --member=serviceAccount:$sa \
  --role=roles/functions.invoker

gcloud scheduler jobs create http firestore-backup \
  --schedule="0 0 * * *" \
  --uri=[function-url] \
  --oidc-service-account-email=$sa \
  --oidc-token-audience=[function-url]
  --body='{"project_id":"$project_id","bucket_name":"$bucket", "action":"backup"}'
```

### Calling the function to restore the database

```shell
curl -X POST -H "Content-Type: application/json" -d '{"project_id":"$project_id","bucket_name":"$bucket", "action":"restore"}' [function-url]
```
