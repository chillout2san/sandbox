#!/bin/bash

echo "GCPの認証情報を表示します。"
echo "******************************"
gcloud auth list
echo "******************************"

echo "デプロイしたいGCPプロジェクトのアカウントでしたか？"
read -p "Y/n> " answer1
if [[ "${answer1}" == "n" ]]; then
    echo "デプロイを中断します。"
    exit 1
fi

echo "デプロイを開始します。"
gcloud app deploy
echo "デプロイを完了しました。"