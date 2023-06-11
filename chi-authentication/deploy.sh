#!/bin/bash

PROJECT="learngcp-361805"
REGION="us-central1"

RED='\e[31m'
BLUE='\e[34m'

function cancel() {
    printf "${RED}デプロイを中断します。"
    exit 1
}

# mainブランチかどうか確認する
if [[ $(git rev-parse --abbrev-ref HEAD) != "main" ]]; then
    printf "${RED}デプロイできるのはmainブランチのみです。"
    cancel
fi

# GCPの認証情報があるかどうか確認する
echo "GCPの認証情報を表示します。"
echo "******************************"
gcloud auth list
echo "******************************"

echo "デプロイしたいGCPプロジェクトに所属するアカウントでしたか？"
read -p "Y/n> " answer2
if [[ "${answer2}" == "n" ]]; then
    printf "${RED}gcloud auth loginでアカウントを切り替えてください。"
    cancel
fi

echo "デプロイを開始します。"

# ProjectとRegionを設定する
gcloud config set project "${PROJECT}"
gcloud config set run/region "${REGION}"

# コンテナをpushするRegistoryが存在するか確認する
gcloud artifacts repositories list | grep "chi-authentication"
if [[ $? == 1 ]]; then
    printf "${RED}コンテナをpushするArtifact Registryが存在しません。"
    cancel
fi

# Cloud Buildでビルドする
gcloud builds submit --region="${REGION}" --config cloudbuild.yaml
if [[ $? == 1 ]]; then
    printf "${RED}コンテナのビルドに失敗しました。"
    cancel
fi

# 認証なしにユーザーがコールできるようにする
# --authenticatedのコマンドラインバージョン
gcloud run services set-iam-policy chi-authentication policy.yaml

printf "${BLUE}デプロイが完了しました。"