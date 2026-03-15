# 学習記録アプリのアーキテクチャ方針

## 目的

このアプリは、学習記録を日々残し、集計し、振り返りできる Web アプリとして作る。

同時に、個人開発のポートフォリオとして以下を示すことを目的とする。

- 実装を最後まで作り切る力
- 設計意図を説明できる力
- 保守しやすい構成を選べる力
- Go でバックエンドを整理して実装する力

## 採用するアーキテクチャ

`モジュラモノリス + レイヤードアーキテクチャ + 軽量な Clean/Hexagonal Architecture`

## この方針を選ぶ理由

- 個人開発でも過剰に複雑になりにくい
- マイクロサービスより実装コストが低い
- レイヤー分離の基本を学べる
- Clean Architecture や Hexagonal Architecture の考え方を実践しやすい
- ドメインロジックをフレームワークや DB から分離しやすい
- テストしやすい構造にしやすい
- README や面接で設計意図を説明しやすい

## 採用しない方針

### マイクロサービス

現時点では採用しない。

理由:

- 個人開発には運用負荷が高い
- ネットワーク分割やデプロイ分離の複雑さに対して得られる価値が小さい
- ポートフォリオ段階では単一アプリ内で責務分離できれば十分

### フルスケール DDD

現時点では採用しない。

理由:

- 学習コストが高い
- アプリ規模に対して設計が重くなりやすい
- まずは Entity、UseCase、Repository、Adapter の責務分離を優先したい

## 全体構成

アプリケーションは 1 つのデプロイ単位として作る。
ただし内部では責務ごとにモジュールを分ける。

想定モジュール:

- auth
- study-record
- goal
- report
- notification

## レイヤー構成

### 1. Presentation

責務:

- HTTP リクエストの受付
- リクエストのバリデーション
- 認証情報の受け取り
- レスポンスの整形

例:

- HTTP handler
- router
- middleware

この層に書かないもの:

- 学習時間の集計ロジック
- 継続日数の計算
- DB アクセスの詳細

### 2. Application

責務:

- ユースケースの実行
- ドメインオブジェクトを組み合わせた処理の流れの制御
- トランザクション境界の調整

例:

- 学習記録を作成する
- 週次レポートを取得する
- 目標進捗を更新する
- 月次サマリーを生成する

この層に書かないもの:

- HTTP 固有の処理
- SQL の詳細

### 3. Domain

責務:

- ビジネスルールの表現
- エンティティや値オブジェクトの定義
- 学習記録アプリに固有の制約の管理

例:

- StudyRecord
- Goal
- StudySummary
- 連続学習日数の計算ルール
- 学習時間の妥当性チェック

原則:

- domain は DB や Web フレームワークを知らない
- domain は外部 API や ORM に依存しない

### 4. Infrastructure / Adapters

責務:

- PostgreSQL への永続化
- 外部サービス連携
- 時刻取得や設定読み込み
- domain/application が必要とする interface の実装

例:

- repository の実装
- DB 接続
- メール通知
- cron 実行

## Clean/Hexagonal の取り入れ方

完全な理論実装を目指すのではなく、以下を守る。

- 依存関係はできるだけ内側に向ける
- domain は外側の詳細を知らない
- repository や notifier の interface は内側で定義する
- DB や HTTP は adapter として外側に置く

重要な判断基準:

`I/O を行うものは外側に置く`

## ディレクトリ構成案

```text
/cmd
  /api
    main.go

/internal
  /domain
    /studyrecord
    /goal
    /report

  /application
    /usecase
      create_study_record.go
      get_weekly_report.go
      update_goal_progress.go

  /port
    study_record_repository.go
    goal_repository.go
    report_repository.go
    notifier.go
    clock.go

  /adapter
    /http
      handler.go
      middleware.go
      request.go
      response.go
    /persistence
      postgres_study_record_repository.go
      postgres_goal_repository.go
    /notification
      email_notifier.go

  /infrastructure
    db.go
    config.go
    logger.go
```

## 依存関係の原則

- presentation は application を呼ぶ
- application は domain と port を使う
- adapter は port を実装する
- domain は他レイヤーに依存しない

依存の向き:

`presentation -> application -> domain`

`adapter -> port <- application`

## 学習記録アプリで domain に置きたいもの

- 学習時間は 0 以下にできない
- 学習日が不正な未来日にならない
- 目標には開始日と終了日が必要
- 日次記録から週次・月次集計を計算する
- 継続日数を算出する

こうしたルールを handler や SQL まわりに散らさず、domain/application に集める。

## MVP 時点での現実的な構成

最初から全モジュールを作り込まない。
まずは以下を優先する。

1. auth
2. study-record
3. goal
4. report

notification は後から追加する。

## 技術スタック案

- Frontend: Next.js
- Backend: Go
- API: REST
- Database: PostgreSQL
- Container: Docker
- CI/CD: GitHub Actions

## テスト方針

### Domain

- 純粋なユニットテストを重視する
- 外部依存なしで高速に実行できるようにする

### Application

- port の fake/mock を使ってユースケースを検証する

### Adapter / Infrastructure

- DB や HTTP を含む結合テストを必要最小限で書く

## 設計上の注意

- handler に業務ロジックを書きすぎない
- repository を単なる便利箱にしない
- domain を ORM モデルと同一視しない
- 小さいうちは抽象化しすぎない
- ただし責務の境界は崩さない

## 将来の拡張余地

将来的に必要なら以下を追加できる。

- バッチによる週次レポート自動生成
- 通知機能
- 外部カレンダー連携
- 学習カテゴリ分析
- AI による振り返り支援

この時も、モジュラモノリスのまま十分拡張可能とする。

## 結論

このアプリでは、`モジュラモノリスをベースに、レイヤードアーキテクチャを採用し、Clean/Hexagonal の考え方を必要な範囲で取り入れる`。

目的は、過剰設計を避けつつ、保守性・説明可能性・テスト容易性の高い構成を実践することにある。
