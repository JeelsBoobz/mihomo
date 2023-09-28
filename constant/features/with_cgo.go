//go:build with_cgo
package features

func init() {
	TAGS = append(TAGS, "with_cgo")
}
