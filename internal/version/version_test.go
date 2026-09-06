package version

import "testing"

func TestVersion(t *testing.T) {
	tests := []struct {
		name     string
		injected string
		want     string
	}{
		{name: "未設定なら dev を返す", injected: "", want: "dev"},
		{name: "注入された値をそのまま返す", injected: "1.2.3", want: "1.2.3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ldflags の -X はパッケージ変数を書き換えるので、テストでも同じ変数を差し替える。
			original := version
			t.Cleanup(func() { version = original })
			version = tt.injected

			if got := Version(); got != tt.want {
				t.Errorf("Version() = %q, want %q", got, tt.want)
			}
		})
	}
}
