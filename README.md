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

## 開発

### 前提条件

- Go 1.21以上
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
