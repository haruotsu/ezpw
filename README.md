# ezpw (Easy Playwright)
**ezpw** (Easy Playwright) は、PlaywrightによるE2Eテストを、YAML形式で簡潔に記述できるツールです。

## 特徴

- シンプルなYAML記述 - JavaScriptやTypeScriptの知識不要
- マルチブラウザ対応 - Chromium、Firefox、WebKit
- モバイル対応 - デバイスエミュレーション

## クイックスタート

### インストール

```bash
# Go 1.21以上が必要
go install github.com/haruotsu/ezpw/cmd/ezpw@latest

# または、リリースページからバイナリをダウンロード
```

### 基本的な使い方

1. **テストシナリオの作成** (`test.yml`)

```yaml

```

2. **テストの実行**

```bash
# 基本実行
ezpw run test.yml

# 特定ブラウザで実行
ezpw run test.yml --browser firefox

# ヘッドレスモード無効化（ブラウザ表示）
ezpw run test.yml --no-headless

# 詳細ログ付きで実行
ezpw run test.yml --verbose
```

## CI/CD

### 自動化されているプロセス

このプロジェクトでは、GitHub Actionsを使用して以下のプロセスを自動化しています：

#### CI

プルリクエスト・プッシュ時に自動実行

1. Lint - コード品質チェック
   - `gofmt`によるフォーマットチェック
   - `golangci-lint`による静的解析（30種類以上のlinter）

2. Test - 自動テスト
   - Go 1.23と1.24でのマトリックステスト
   - レースコンディション検出付きテスト実行
   - カバレッジレポートの生成とアップロード

3. Build - ビルド検証
   - バイナリのビルド成功確認
   - CLIコマンドの動作確認

#### CD

mainブランチへのプッシュ時に自動実行

1. 自動バージョニング
   - セマンティックバージョニングによる自動タグ付け
   - コミットメッセージからバージョンを決定
     - `fix:` → パッチバージョンアップ (0.0.X)
     - `feat:` → マイナーバージョンアップ (0.X.0)
     - `BREAKING CHANGE:` → メジャーバージョンアップ (X.0.0)

2. リリース作成
   - GitHub Releaseの自動作成
   - 変更ログの自動生成
   - マルチプラットフォームバイナリの配布
     - Linux (amd64)
     - macOS (Intel/Apple Silicon)
     - Windows (amd64)

3. Dockerイメージ
   - GitHub Container Registry (ghcr.io)への自動プッシュ
   - 最新タグとバージョンタグの両方を付与

### Dockerサポート

```bash
# 最新版のイメージを取得
docker pull ghcr.io/haruotsu/ezpw:latest

# コンテナでテスト実行
docker run -v $(pwd):/workspace ghcr.io/haruotsu/ezpw:latest run test.yml
```

## 開発

### 前提条件

- Go 1.24以上
- Node.js (Playwright インストール用)

### セットアップ

```bash
# リポジトリクローン
git clone https://github.com/haruotsu/ezpw.git
cd ezpw

# 依存関係インストール
go mod download

# Playwright インストール
npm install -g playwright
playwright install

# ビルド
make build

# テスト実行
make test

# リント
make lint
```

### ディレクトリ構成

```
ezpw/
├── cmd/ezpw/           # CLI エントリーポイント
├── internal/           # 内部パッケージ
│   ├── cli/           # CLI 処理
│   ├── parser/        # YAML パーサー
│   ├── executor/      # 実行エンジン
│   ├── playwright/    # Playwright 統合
│   └── reporter/      # レポート生成
├── pkg/types/         # 公開型定義
├── testdata/          # テストデータ
└── docs/              # ドキュメント
```

## ライセンス

[MIT License](LICENSE) の下でライセンスされています。
