# trial-nginx

nginx の挙動確認用のリポジトリ

## 試したこと

- `proxy_set_header Host $http_host;`
  - web コンテナや app コンテナで受けたリクエストの Host ヘッダが proxy コンテナでリクエストを受けたときの URL のホスト名(+:ポート番号)が渡るようになった
- `proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;`
  - X-Forwarded-For ヘッダーフィールドにアクセス元(\$remote_addr)をカンマ区切りで追加されるようになった
