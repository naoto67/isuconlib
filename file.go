package isuconlib

import (
	"os"
	"path"

	"github.com/pkg/errors"
)

func WriteFile(p string, data []byte) error {
	dir := path.Dir(p)
	if err := os.MkdirAll(dir, os.FileMode(0o755)); err != nil && !os.IsExist(err) {
		return errors.Wrap(err, "dirの作成に失敗しました")
	}
	if err := os.WriteFile(p, data, os.FileMode(0o755)); err != nil {
		return errors.Wrap(err, "fileの書き込みに失敗しました。")
	}
	return nil
}

func WriteFileWithFileMode(p string, data []byte, perm os.FileMode) error {
	dir := path.Dir(p)
	if err := os.MkdirAll(dir, os.FileMode(0o755)); err != nil && !os.IsExist(err) {
		return errors.Wrap(err, "dirの作成に失敗しました")
	}
	if err := os.WriteFile(p, data, perm); err != nil {
		return errors.Wrap(err, "fileの書き込みに失敗しました。")
	}
	return nil
}
