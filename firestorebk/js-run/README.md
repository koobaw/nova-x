# Cloud Run Jobs Sample

## Build

* Set an environment variable with your GCP Project ID:

```
export GOOGLE_PROJECT=nova-hj
export REGION=asia-northeast1
export IMAGE_URL_JS=asia-northeast1-docker.pkg.dev/nova-hj/infra/js-job:latest
export JOB_NAME_JS=run-js-bk
export project_id="nova-hj"
```

* Use a [Buildpack](https://github.com/GoogleCloudPlatform/buildpacks) to build the container:

```sh
docker image build .  --platform amd64 -t $IMAGE_URL_JS
docker image push $IMAGE_URL_JS

docker build -t $IMAGE_URL_JS .
docker push $IMAGE_URL_JS
```

## Run Locally

```sh
docker run --rm $IMAGE_URL_JS
```

## Create a Job

```
gcloud run jobs create $JOB_NAME_JS \
    --set-env-vars SLEEP_MS=10000 \
    --set-env-vars FAIL_RATE=0.1 \
    --max-retries 5 \
    --task-timeout 4h \
    --image $IMAGE_URL_JS \
    --project $GOOGLE_PROJECT \
    --region $REGION

gcloud run jobs execute $JOB_NAME_JS --region $REGION --project $GOOGLE_PROJECT

gcloud run jobs delete $JOB_NAME_JS \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

## Run the Job
```
gcloud run jobs run job-quickstart
```