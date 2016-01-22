package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Sha256 returns a base64 encoded SHA256 hash of `b`.
func Sha256(b []byte) string {
	h := sha256.New()
	h.Write(b)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// LoadFiles return filtered map of files; for filtering it uses shell file name pattern matching
func LoadFiles(root string, ignoredPatterns []string) (map[string]*os.File, error) {
	files := make(map[string]*os.File)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		mode := info.Mode()
		if !(mode.IsRegular() || mode&os.ModeSymlink == os.ModeSymlink) {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		for _, pattern := range ignoredPatterns {
			matched, err := filepath.Match(pattern, rel)
			if err != nil {
				return err
			}

			if matched {
				return nil
			}
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		files[rel] = file

		return nil
	})

	return files, err
}

// ReadIgnoreFile reads .apexignore in `dir` when present and returns a list of patterns.
func ReadIgnoreFile(dir string) ([]string, error) {
	path := filepath.Join(dir, ".apexignore")

	b, err := ioutil.ReadFile(path)

	if os.IsNotExist(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return strings.Split(string(b), "\n"), nil
}
