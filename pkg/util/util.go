package util

import (
	"fmt"

	"github.com/xuewenG/webdav/pkg/config"
)

func WithPrefix(path string) string {
	prefix := config.Config.Prefix

	if prefix == "" || prefix == "/" {
		return path
	}

	return fmt.Sprintf("%s%s", prefix, path)
}
