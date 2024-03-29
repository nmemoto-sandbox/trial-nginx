# trial-nginx

nginx の挙動確認用のリポジトリ

## 試したこと

- `proxy_set_header Host $http_host;`
  - web コンテナや app コンテナで受けたリクエストの Host ヘッダが proxy コンテナでリクエストを受けたときの URL のホスト名(+:ポート番号)が渡るようになった
- `proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;`
  - X-Forwarded-For ヘッダーフィールドにアクセス元(\$remote_addr)をカンマ区切りで追加されるようになった
  - 参考ページ: http://nginx.org/en/docs/http/ngx_http_proxy_module.html
- Forwarded ヘッダーを追加
  - RFC 7239 で策定された比較的新しいヘッダー。前述の X-Forwarded-For ヘッダーの標準化されたもの。
  - 参考ページ
    - https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/Forwarded
    - https://www.nginx.com/resources/wiki/start/topics/examples/forwarded/
- `proxy_http_version 1.1;`
  - プロキシ先への接続(proxy->web, web->app)を HTTP/1.1 で行うように設定した。
- X-Forwarded-Proto ヘッダーを追加
  - proxy でリクエストを受けたときの\$scheme を付加し、web では proxy でつけたヘッダーをそのまま app へのリクエストに付加した(\$http_x_forwarded_proto を利用した)
- X-Real-Ip ヘッダーを追加
  - proxy でリクエストを受けたときの$remote_addr(リクエスト元のIP)を付加し、web では proxy でつけたヘッダーをそのまま app へのリクエストに付加した($http_x_real_ip を利用した)
- web コンテナに app コンテナのレスポンスのキャッシュを作成するようにした
  - expires でクライアント側のキャッシュの有効期限を 30 秒に設定
  - proxy_cache_valid で app コンテナのレスポンスをキャッシュする時間を設定
    - http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_cache_valid
