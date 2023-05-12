### コンテナのArtifact Registryへのpush
project が制限されているようで、リージョンが特定の箇所しかダメだった。
https://cloud.google.com/build/docs/locations?hl=ja#restricted_regions_for_some_projects
cloud buildsをyamlで制御できないか考える。
gcloud builds submit --region=us-central1 --tag us-central1-docker.pkg.dev/learngcp-361805/chi-authentication/backend:latest

### デプロイコマンド
gcloud run deploy chi-authentication --region us-central1 --image us-central1-docker.pkg.dev/learngcp-361805/chi-authentication/backend:latest --allow-unauthenticated

*次はウェブトラフィックの処理から読む