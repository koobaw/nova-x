

```
gcloud run jobs deploy job-quickstart \
    --source . \
    --tasks 50 \
    --set-env-vars SLEEP_MS=10000 \
    --set-env-vars FAIL_RATE=0.1 \
    --max-retries 5 \
    --region us-central1 \
    --project=nova-hj


gcloud run jobs execute job-quickstart --region us-central1
```