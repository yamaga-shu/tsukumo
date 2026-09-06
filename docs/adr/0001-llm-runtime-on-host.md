# LLM ランタイムはホスト側に置く

tsukumo の開発は devcontainer 内の Linux で行うが、検証に使う LLM ランタイムはホストの macOS に各自が用意し、コンテナからは `host.docker.internal` 経由の HTTP で到達する。
コンテナに同梱すると macOS では CPU 推論に落ち、tsukumo が本来相手にする Metal 実行と性能特性が乖離するためである。
この結果、接続先が設定可能であることが tsukumo の設計要件になる。

検討した選択肢と経緯は #5 にある。
