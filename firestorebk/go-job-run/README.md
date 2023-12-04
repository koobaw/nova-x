https://tech.buysell-technologies.com/entry/adventcalendar2022-12-04

https://qiita.com/kyohei_tsuno/items/e4a5975899f09a1fa2b6  区别

https://zenn.dev/nekoshita/articles/cf39a31f3052bf

# cloud run job example

# やること
Cloud Run Jobの動作確認をする
https://cloud.google.com/run/docs/create-jobs


# Cloud Run Jobsだけで何ができるの？
## ドキュメント
- https://cloud.google.com/run/docs/create-jobs
- https://cloud.google.com/run/docs/execute/jobs
- https://cloud.google.com/run/docs/configuring/containers#configure-entrypoint
- https://cloud.google.com/run/docs/configuring/environment-variables#setting

## 環境変数を用意
```
export GOOGLE_PROJECT=nova-hj
export REGION=asia-northeast1
export IMAGE_URL=asia-northeast1-docker.pkg.dev/nova-hj/infra/golang-job:latest
export JOB_NAME=golang-job
export project_id="nova-hj"
gcloud config configurations activate $project_id
gcloud config set project $project_id
gcloud auth application-default set-quota-project $project_id
```

## イメージを用意
```
docker image build .  --platform amd64 -t $IMAGE_URL
docker container run --rm $IMAGE_URL
docker image push $IMAGE_URL
```


## コマンドフラグなし -> 動作する
```
# https://cloud.google.com/run/docs/create-jobs?hl=zh-cn#command-line

gcloud run jobs create $JOB_NAME \
    --set-env-vars SLEEP_MS=10000 \
    --set-env-vars FAIL_RATE=0.1 \
    --max-retries 5 \
    --task-timeout 4h \
    --image $IMAGE_URL \
    --project $GOOGLE_PROJECT \
    --region $REGION

gcloud run jobs execute $JOB_NAME --region $REGION

gcloud run jobs update $JOB_NAME
    --set-env-vars SLEEP_MS=10000 \
    --set-env-vars FAIL_RATE=0.1 \
    --max-retries 5 \
    --task-timeout 6h \
    --image $IMAGE_URL \
    --project $GOOGLE_PROJECT \
    --region $REGION

gcloud run jobs delete $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

## コマンドと引数は指定できる
```
gcloud run jobs create $JOB_NAME \
    --image $IMAGE_URL \
    --command "/app" \
    --args a,b,c,100 \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

## フラグっぽい引数は指定する -> エラー
```
gcloud run jobs create $JOB_NAME \
    --image $IMAGE_URL \
    --command "/app" \
    --args --user,hoge \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

##  yamlで無理やり更新を試みる -> servicesと違ってjobsはreplaceコマンド使えない
なので、今のところフラグっぽい引数を指定することは不可能っぽい
```
gcloud run jobs describe $JOB_NAME --format export \
    --project $GOOGLE_PROJECT \
    --region $REGION
gcloud run jobs replace sample_job.yaml \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

## 環境変数を指定できる
```
gcloud run jobs create $JOB_NAME \
    --image $IMAGE_URL \
    --set-env-vars FOO=foo,BAR=bar \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

## tasksを指定する -> 並列に実行される
```
gcloud run jobs create $JOB_NAME \
    --tasks 3 \
    --image $IMAGE_URL \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

## tasksとparallelismを指定する -> 並列数を制御できる
```
gcloud run jobs create $JOB_NAME \
    --tasks 3 \
    --parallelism 1 \
    --image $IMAGE_URL \
    --project $GOOGLE_PROJECT \
    --region $REGION
```


## その他使ったコマンド
```
gcloud run jobs list \
    --project $GOOGLE_PROJECT \
    --region $REGION

gcloud run jobs executions list\
    --project $GOOGLE_PROJECT \
    --region $REGION

gcloud run jobs executions describe sample-job-m9xkv \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

# 他のサービスと組み合わせで何ができるの？
## Cloud Scheduler
Cloud Schedulerとの組み合わせで定期実行ができるらしい
https://cloud.google.com/run/docs/execute/jobs-on-schedule#console

## Cloud Workflow
https://cloud.google.com/run/docs/create-jobs

>After you create or update a job, you can execute the job as a one-off, on a schedule or as part of a workflow. You can manage individual job executions and view the execution logs.

Cloud Workflowで実行できるとの記載があるが、ドキュメントはまだないようで、yamlの書き方が不明だったので試せなかった
