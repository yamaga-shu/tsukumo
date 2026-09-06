// Package version は tsukumo のバージョン文字列を提供する。
package version

// version はビルド時に -ldflags "-X ...version.version=<値>" で差し込まれる。
var version string

// Version はバージョン文字列を返す。ldflags で差されていなければ "dev" を返す。
func Version() string {
	if version == "" {
		return "dev"
	}
	return version
}
