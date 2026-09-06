# CLAUDE.md

AI エージェント向けの入口。AGENTS.md は置かず、このファイルがその役割を担う。

## 構成

- `README.md`: プロジェクトの目的と位置づけ
- `LICENSE`: MIT ライセンス
- `CLAUDE.md`: このファイル
- `CONTEXT.md`: 用語集
- `docs/adr/`: 設計判断の結論
- `cmd/tsukumo/`: エントリポイント
- `internal/version/`: バージョン文字列
- `.devcontainer/`: 開発コンテナの定義
- `.github/workflows/`: CI のワークフロー
- `.golangci.yml`: lint の設定
- `.gitignore`: ビルド成果物の無視設定
- `go.mod`: module path、Go のバージョン、golangci-lint のバージョン
- `skills-lock.json`: skills CLI が導入したスキルの出所とハッシュ
- `.agents/skills/`: スキルの実体。skills CLI で導入する
- `.claude/skills/`: `.agents/skills/` 配下の各スキルへのシンボリックリンク

## 開発

devcontainer が唯一の公式な開発環境である。
ホストの Go は使わず、VS Code でコンテナを開いてから作業する。

```bash
go build ./...
go test ./...
go tool golangci-lint run
```

CI もこの 3 つを同じ順で実行する。

### ホストの LLM ランタイム

推論の相手はホストの macOS で動かす Ollama で、コンテナからは `host.docker.internal` へ HTTP で到達する。
コンテナ内の `localhost` はホストに届かないため、ホスト側で待ち受けを開いておく。

```bash
OLLAMA_HOST=0.0.0.0 ollama serve
```

`go test ./...` はランタイムを起動していなくても通る。
ホストに置く理由は `docs/adr/0001-llm-runtime-on-host.md` にある。
