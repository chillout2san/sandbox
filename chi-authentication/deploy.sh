#!/bin/bash

# 定数を置く
DEPLOY_FILE=deploy.yaml
PROJECT="learngcp-361805"

# mainブランチかどうか確認する
if [[ $(git rev-parse --abbrev-ref HEAD) != "main" ]]; then
    echo "デプロイできるのはmainブランチのみです。"
    exit 1
fi

# deploy.yamlが準備されているか確認する。
if [ ! -e $DEPLOY_FILE ]; then
    cp app.yaml deploy.yaml
    echo "deploy.yamlを生成しました。"
    echo "環境変数を記載した後、再度実行してください。"
    exit 1
fi

# deploy.yamlに環境変数を記載しているか確認する。
echo "deploy.yamlに環境変数を記載しているか確認してください。"
read -p "Y/n> " answer1
if [[ "${answer1}" == "n" ]]; then
    echo "環境変数を記載した後、再度実行してください。"
    exit 1
fi

# GCPの認証情報を確認する。
echo "GCPの認証情報を表示します。"
echo "******************************"
gcloud auth list
echo "******************************"

echo "デプロイしたいGCPプロジェクトに所属するアカウントでしたか？"
read -p "Y/n> " answer2
if [[ "${answer2}" == "n" ]]; then
    echo "gcloud auth loginでアカウントを切り替えてください。"
    echo "デプロイを中断します。"
    exit 1
fi

# GCPへデプロイする。
echo "デプロイを開始します。"
gcloud app deploy deploy.yaml --project="${PROJECT}"
echo "デプロイを完了しました。"
