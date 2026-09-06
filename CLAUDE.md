# CLAUDE.md

AI エージェント向けの入口。AGENTS.md は置かず、このファイルがその役割を担う。

## 構成

- `README.md`: プロジェクトの目的と位置づけ
- `LICENSE`: MIT ライセンス
- `CLAUDE.md`: このファイル
- `skills-lock.json`: skills CLI が導入したスキルの出所とハッシュ
- `.agents/skills/`: スキルの実体。skills CLI で導入する
- `.claude/skills/`: `.agents/skills/` 配下の各スキルへのシンボリックリンク
