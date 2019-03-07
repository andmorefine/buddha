# goa2-sample

### 使用技術
- goa v2
- mysql
- swagger-ui
- viron

## 環境構築
golangおよびdockerは既にインストール済みであることを想定

ライブラリのインストール
```
$ make install
```

## 開発
- designディレクトリ以下に新しいAPIデザインを追加
- ```$ make goagen```でコード自動生成
- 新しく作られたコントローラの雛形を元にビジネスロジックを追加
- ```$ make run```でビルド&実行

## 関連サーバ起動
mysql、swagger-ui、vironはそれぞれ別個のdockerコンテナで起動する

```
$ make docker-build
$ make docker-up
```

コンテナの停止
```
$ make docker-rm
```
