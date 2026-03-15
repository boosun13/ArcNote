# ArcNote リポジトリ構成

## 目的

frontend、backend、infra を 1 つのリポジトリで管理しつつ、責務を混在させないことを目的とする。

## 推奨構成

```text
/
  frontend/
    app/
    components/
    features/
    lib/
    public/
    styles/
    tests/
    package.json

  backend/
    cmd/
      api/
        main.go
    internal/
      domain/
      application/
      port/
      adapter/
        http/
        persistence/
        notification/
      infrastructure/
    migrations/
    test/
    go.mod
    go.sum

  infra/
    terraform/
      modules/
        network/
        database/
        app_runtime/
      envs/
        local/
        dev/
        prod/
    kubernetes/
      base/
        namespace.yaml
        api-deployment.yaml
        api-service.yaml
      overlays/
        dev/
        prod/

  docs/
  scripts/
  .github/
    workflows/
```

## この構成にする理由

### Frontend

- Next.js に関する責務を Go や IaC から分離できる
- Node.js の依存管理や CI を独立して扱いやすい

### Backend

- 既存のアーキテクチャ方針である modular monolith と layered/hexagonal の境界に合わせやすい
- `cmd` をエントリーポイント、`internal` をアプリケーション本体として明確に分けられる

### Infrastructure

- クラウド資源の作成は `terraform`、デプロイ定義は `kubernetes` と責務を分けられる
- Terraform は再利用可能な module と環境ごとの root を分けて管理できる
- Kubernetes は base と overlay を分けることで環境差分を局所化できる

## 境界ルール

- `frontend/` から IaC や backend の内部実装には依存しない
- `backend/` から frontend の資産には依存しない
- `infra/terraform/` はインフラの作成だけを担う
- `infra/kubernetes/` はビルド済みアーティファクトのデプロイだけを担う
- 共有契約はまず `docs/` に置き、コード共有が本当に必要になったら専用ディレクトリを追加する

## 運用メモ

- CI は frontend、backend、infra で責務ごとに分ける
- Dockerfile は deployable unit ごとに持たせ、リポジトリ直下に責務の異なる build ロジックを増やしすぎない
- 将来的に ArgoCD などの GitOps を導入する場合は `infra/kubernetes/overlays/<env>` を参照先にする
- Terraform の state は環境ごとに分離する
