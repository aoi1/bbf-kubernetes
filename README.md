# bbf-kubernetes

Repository for the book **"Build, Breaking, Fixing: A Playful Way to Learn Kubernetes"**（書籍『Kubernetes 実践入門』ハンズオン用）。

## 構成

| ディレクトリ | 内容 |
| --- | --- |
| `hello-server/` | Go 1.21 サンプルサーバー（`/healthz`, `/metrics`） |
| `chapter-01/`〜`chapter-11/` | 章別 Kubernetes マニフェスト |
| `kind/` | kind クラスタ設定 |
| `kube-prometheus-stack/` | Prometheus Helm values |

## 前提ツール

- kubectl / kind / helm / Go 1.21+
- Docker（イメージビルド時）

## CI

tag push 時に `blux2/hello-server:*` を Docker Hub へマルチ arch ビルド（`.github/workflows/build.yaml`）。

---

## 改善点バックログ

> 監査日: 2026-05-31。Tech: Go hello-server + K8s マニフェスト + Kustomize + kind。`.gitignore`・LICENSE・Dependabot なし。

### 機能 (Functionality)

- [ ] `[P1]` probe パス不一致 — `hello-server` は `/healthz` だが chapter-08〜10 は `/health` を参照
- [ ] `[P1]` `chapter-01/hello-server/` がルート版と API 不一致 — 差分表を README に記載
- [ ] `[P2]` Prometheus scrape target が固定 hostname — namespace/Service 対応表を追加
- [ ] `[P2]` kind 設定の利用手順が README にない
- [ ] `[P3]` `*-destruction.yaml` 演習用 manifest の学習目的を章見出し下に記載

### デザイン/UX (Design)

- [ ] `[P2]` README が 1 行のみ — 書籍タイトル・対象読者・章構成・前提ツールを追加（本 README 冒頭で着手）
- [ ] `[P2]` 各章の期待結果（成功時 `kubectl get` 出力例）が不足
- [ ] `[P3]` マニフェストに初学者向けコメント追加

### セキュリティ (Security)

- [ ] `[P1]` `chapter-06/secret/nginx-sample.yaml` に学習用平文資格情報 — 「本番禁止」を明記し削除手順追加
- [ ] `[P2]` Docker イメージ `blux2/hello-server:*` 固定 — 自分でビルドする手順と digest pin 推奨
- [ ] `[P2]` CI が DockerHub secrets 依存 — Secrets 設定手順と最小権限 token 方針
- [ ] `[P3]` Dependabot（gomod / docker / github-actions）追加
- [ ] `[P3]` hello-server に TLS/認証なし — 本番 hardening チェックリストを別 doc で

### システム設計 (System Design)

- [ ] `[P1]` `.gitignore` なし — 標準 `.gitignore` を追加
- [ ] `[P2]` ルート `hello-server/` と `chapter-01/hello-server/` の二重管理 — sync 方針決定
- [ ] `[P2]` Kustomize overlay 適用手順が README 未記載
- [ ] `[P2]` CI が tag push のみ — PR 時に `kubectl kustomize build` / `go test` 追加
- [ ] `[P3]` GitHub Actions を v4/v6 世代へ更新
- [ ] `[P3]` Docker ビルド `no-cache: true` 固定 — キャッシュ戦略見直し

### ドキュメント/運用 (Docs & Ops)

- [ ] `[P1]` LICENSE ファイルなし
- [ ] `[P1]` 章一覧とファイル対応表がない — chapter-01〜11 トピック一覧
- [ ] `[P1]` 前提ソフトウェアのバージョン要件なし
- [ ] `[P2]` Dependabot 未設定
- [ ] `[P2]` イメージタグと CI の関係（tag push → `blux2/hello-server:${tag}`）を追記
- [ ] `[P2]` Prometheus スタック導入の helm install 例
- [ ] `[P3]` CONTRIBUTING / Issue テンプレート
