### 概要

本アプリケーションは会員登録機能、ログイン機能、認証しているかどうか確認する機能の3つを持っており、他の機能は実装しない代わりに日々学んだことを活かしてブラッシュアップしていくことが役目です。

### 環境構築手順

1. Goのアプリケーションが使用するenvファイルを生成します。特に中身を書き換える必要はありません。

```bash
cd chi_authentication/api
cp .env.sample .env
```

2. MySQLのDockerコンテナが使用するenvファイルを生成します。こちらも中身を書き換える必要はありません。

```bash
cd chi_authentication/mysql
cp .env.sample .env
```

3. GoとMySQLのDockerコンテナを立ち上げます。初回はそれなりに時間がかかります。

```bash
cd chi_authentication
docker compose up -d
```

4. アプリケーションが立ち上がっているか確認します。STATUSカラムが`db`コンテナと`go`コンテナで`running`になっていることを確認してください。

```bash
docker compose ps
```

5. Postman等で`http://localhost:8080/account/check_auth`のエンドポイントを以下のボディをつけてPostで叩いてください。
```json
{
    "id": "id",
    "token": "token"
}
```

以下のようなエラーメッセージが返ってきたら、アプリケーションは正常に動作しています。
```json
{
    "hasAuth": false,
    "errMessage": "token contains an invalid number of segments"
}
```

### デプロイ手順
環境構築手順は完了しているものとみなして進めます。
なおDBはCloudSQL、ホストはGoogleAppEngineのflex環境で行う前提です。

1. `main`ブランチにいることを確認した上で、デプロイスクリプトを叩く。`deploy.yaml`ファイルが作成されることを確認。

```bash
cd chi_authentication
./deploy.sh
```
2. `deploy.yaml`の`env_variables`に環境変数を記載する。

```yaml
env_variables:
  ALLOW_ORIGIN: APIコールを許可するオリジンを記載
  DB_USER: CloudSQLのユーザー名
  DB_PASS: CloudSQLのパスワード名
  DB_NAME: CloudSQLに作成したデータベース名
  INSTANCE_CONNECTION_NAME: CloudSQLの接続名
  JWT_SECRET_KEY: 適当な文字列。テスト用ならsecret等で構わない
  DB_FLAG: GCP ここはそのまま触らない
```

3. `deploy.sh`の`PROJECT`定数をデプロイしたいプロジェクトの名前に変更。

4. デプロイスクリプト実行。

```bash
cd chi_authentication
./deploy.sh
```