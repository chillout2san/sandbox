### コンテナのArtifact Registryへのpush
project が制限されているようで、リージョンが特定の箇所しかダメだった。
https://cloud.google.com/build/docs/locations?hl=ja#restricted_regions_for_some_projects

Cloud Runデプロイしようとするとサービスアカウントにロールを渡さないといけない。
https://cloud.google.com/build/docs/deploy-containerized-application-cloud-run?hl=ja

### ロードバランサー
gcloud compute addresses create admin-ip \
    --network-tier=PREMIUM \
    --ip-version=IPV4 \
    --global

gcloud compute addresses describe admin-ip \
    --format="get(address)" \
    --global