# ArcNote

ArcNote is managed as a single repository containing frontend, backend, and infrastructure code.

## Local Development

Run the backend server from the repository root:

```bash
docker compose up --build
```

The API will be available at `http://localhost:8080`.

## Repository Layout

```text
/
  frontend/              # Next.js application
  backend/               # Go API and application modules
  infra/                 # IaC and deployment manifests
    terraform/
      modules/
      envs/
    kubernetes/
      base/
      overlays/
  docs/                  # Architecture and design documents
  scripts/               # Local automation scripts
```

## Intent

- `frontend/` owns the web UI and BFF-facing client code.
- `backend/` owns the Go API, domain, application, and adapters.
- `infra/` owns reproducible infrastructure and deployment definitions.
- `docs/` keeps architectural decisions close to the codebase.

## Backend Layout

```text
backend/
  Dockerfile
  cmd/
    api/                 # API サーバーの entrypoint
  internal/
    domain/              # ビジネスルールとエンティティ
    application/         # usecase とアプリケーションサービス
    port/                # application が依存する interface
    adapter/             # HTTP や persistence など外部入出力
    infrastructure/      # 依存組み立てや設定
  migrations/
  test/
```

## Backend Responsibilities

### `cmd/`

- 実行可能ファイルの入口を置く
- できるだけ薄く保ち、起動処理に専念させる
- 例: `cmd/api/main.go`

### `internal/domain/`

- ドメイン知識そのものを置く
- エンティティ、値オブジェクト、ドメインルールをここに集める
- 外部ライブラリや HTTP 都合を持ち込まない
- 例: `internal/domain/studyrecord`

### `internal/application/`

- ユースケースを置く
- 「何をしたいか」を表現し、domain を使って処理を進める
- 永続化や通知などは直接実装せず、必要なら port 経由で依存する
- 例: `internal/application/studyrecord`

### `internal/port/`

- application が必要とする interface を置く
- repository や notifier など、外部依存の抽象をここで定義する
- 実装は adapter 側に置く

### `internal/adapter/`

- 外部との接続部分を置く
- HTTP handler、DB 実装、外部サービス連携などをここで扱う
- application や domain を外部 I/O に接続する層
- 例:
  - `internal/adapter/http`
  - `internal/adapter/http/studyrecord`
  - `internal/adapter/persistence`

### `internal/infrastructure/`

- 依存関係の組み立てを置く
- どの handler にどの usecase を渡すか、どの実装を使うかを決める
- main や handler に具体実装の生成責務を持たせすぎないための層
- 例: `internal/infrastructure/bootstrap/http.go`

## Current Flow

現在の backend は、最小構成として次の流れで動きます。

```text
cmd/api
  -> infrastructure/bootstrap
  -> adapter/http
  -> application
  -> domain
```

今後 persistence を追加する場合は、次のように port と adapter が間に入る想定です。

```text
adapter/http
  -> application
  -> port
  -> adapter/persistence
```

## Infrastructure Layout

```text
infra/
  terraform/
    modules/
    envs/
      local/
      dev/
      prod/
  kubernetes/
    base/
    overlays/
      dev/
      prod/
```
