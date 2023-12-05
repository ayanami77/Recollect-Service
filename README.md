# 🐬Recollect-Service

自分史作成アプリ「Recollect」のバックエンドです。

## 🛠️技術スタック

**Recollect-Service**は主に以下の技術スタックで構成されています。

- Go
- Gin
- Gorm
- postgresql
- docker

## 🚧環境構築

1 `.env.example`をコピーして、`.env`をルート直下に置いてください。

```sh
cp .env.example .env
```


## 🐢作業時の注意

- コミットメッセージには prefix を付けよう。
  - `feat:` .. 何か機能を実装した時
  - `update:` .. 機能やスタイルは変わらず、実装を更新した時
  - `wip:` .. 作業は途中だが一旦 push しておきたい時
  - `refac:` .. リファクタリング時
  - `fix:` .. 機能のバグの修正時
  - `chore:` .. ライブラリや補助ツールを導入したい時など
  - `docs:` .. ドキュメントの更新時
- PR作成時
  - 基本、PRのテンプレートに沿って記入してください。（場合によっては内容の省略ok）
  - また、本プロジェクトではgithub projectsを利用してタスク管理を行っているので、それをもとにissueを立ててくれると嬉しいです。