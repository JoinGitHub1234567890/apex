// Code generated by go-bindata.
// sources:
// _boilerplate/functions/hello/index.js
// _boilerplate/infrastructure/.gitignore
// _boilerplate/infrastructure/_env/main.tf
// _boilerplate/infrastructure/modules/iam/iam.tf
// _boilerplate/infrastructure/modules/iam/outputs.tf
// DO NOT EDIT!

package boilerplate

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _functionsHelloIndexJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\xcc\x41\x0a\x83\x30\x10\x85\xe1\x7d\x4e\xf1\x36\x25\x06\x82\x07\x10\x7a\x18\x8d\x53\xb5\x0c\x33\x92\x19\x5b\x41\xbc\x7b\xeb\xa6\xd0\xcd\xdb\xbc\x9f\xaf\xa8\x98\x32\xb5\xac\x53\x13\xcd\xfb\xea\x8b\x4c\x78\x6c\x52\x7c\x51\x89\x29\xd0\xbe\x6a\x75\x6b\xe7\x5e\x46\x26\xdc\x7f\x5f\x43\x19\xc5\xf7\xef\x0c\x09\x47\x00\xfe\xa8\xb5\x6a\x21\xb3\x0b\xa3\x17\x89\x77\xb8\x3d\x63\x06\xa5\x2b\x1c\x1a\xd9\x98\x33\x0e\xcc\xc4\xac\x1d\xe2\x5b\x2b\x8f\x11\x67\x0a\x67\xf8\x04\x00\x00\xff\xff\x92\xa4\x17\xbd\x93\x00\x00\x00")

func functionsHelloIndexJsBytes() ([]byte, error) {
	return bindataRead(
		_functionsHelloIndexJs,
		"functions/hello/index.js",
	)
}

func functionsHelloIndexJs() (*asset, error) {
	bytes, err := functionsHelloIndexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "functions/hello/index.js", size: 147, mode: os.FileMode(420), modTime: time.Unix(1462198236, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _infrastructureGitignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x2b\x49\x2d\x2a\x4a\x4c\xcb\x2f\xca\x05\x04\x00\x00\xff\xff\x75\x76\xd9\x6d\x0a\x00\x00\x00")

func infrastructureGitignoreBytes() ([]byte, error) {
	return bindataRead(
		_infrastructureGitignore,
		"infrastructure/.gitignore",
	)
}

func infrastructureGitignore() (*asset, error) {
	bytes, err := infrastructureGitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "infrastructure/.gitignore", size: 10, mode: os.FileMode(420), modTime: time.Unix(1457019469, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _infrastructure_envMainTf = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8c\x31\x0e\xc3\x20\x0c\x45\x77\x4e\x61\xa1\xce\xe4\x04\x3d\x0b\x72\x02\xad\x90\x4c\x5c\x19\x9c\x0e\x51\xee\x5e\x53\xe6\xac\xff\xfd\xf7\x0e\x94\x82\x2b\x65\xf0\xf8\x6d\x51\xf2\xbb\xf0\xee\xe1\xbc\x9c\xab\x9c\x74\xec\x05\xab\x0d\x0e\xa0\xb1\xca\x96\xe1\x09\x3e\x84\x65\xd2\xb6\x0c\xea\xec\xcd\xda\x3f\xda\xc1\x13\xd6\x35\x61\x7c\xe9\xbe\x75\x4b\x45\x61\xca\xb1\xa4\x59\x38\x90\xf4\x1f\x78\x9c\xd3\x0f\xa6\x87\x1b\xe5\x1a\xdd\x5f\x00\x00\x00\xff\xff\xc6\xc7\xbc\xc7\x9f\x00\x00\x00")

func infrastructure_envMainTfBytes() ([]byte, error) {
	return bindataRead(
		_infrastructure_envMainTf,
		"infrastructure/_env/main.tf",
	)
}

func infrastructure_envMainTf() (*asset, error) {
	bytes, err := infrastructure_envMainTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "infrastructure/_env/main.tf", size: 159, mode: os.FileMode(420), modTime: time.Unix(1459260005, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _infrastructureModulesIamIamTf = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x91\x4d\x4f\x02\x31\x10\x86\xef\xfb\x2b\x9a\xc6\x13\x81\x0d\x70\x31\x21\x78\xe0\x80\x57\x0d\x26\x5e\x08\x69\xc6\x32\x68\x93\x7e\x90\xb6\x2b\x2a\xd9\xff\xee\x4c\x59\x50\x89\x1f\x17\xf7\xb0\x87\xbe\x4f\xa7\x7d\x9f\x56\x11\x53\x68\xa2\x46\x21\x61\x97\x94\x01\xa7\x62\xb0\x28\x85\xb4\xe0\x1e\xd6\xa0\x36\x8d\xd7\xd9\x04\x2f\xc5\xbe\x12\xc2\x83\x43\x71\x45\xec\x16\x5f\xd4\x39\x41\x39\xa4\xd4\x38\x2c\x23\xd4\x36\x58\xa3\x5f\x89\x9e\x4e\xe7\x37\xd7\x15\x6f\x97\xf7\x18\x13\xa3\x13\x21\xc7\xc3\xd1\x78\x30\x1a\x0e\x46\x97\xb2\xcf\xd1\x5d\x86\x8c\x0e\x7d\xa6\x70\x49\x0b\xa2\x1c\xc8\x9f\x9c\x6f\x36\xa8\x79\x5d\xce\xac\x0d\xbb\xc2\x97\xe0\x36\x1a\xaf\xcd\x16\x2c\x65\x47\x9a\x47\x61\x7c\x36\x1a\x79\xc3\xe1\x8e\x35\x38\x78\x0b\x9e\x1a\xd6\x3a\x38\xd9\x91\xed\x69\xce\xec\x50\x80\xf8\x94\xd3\x64\x56\x4a\x2c\x58\x43\x01\x5a\xfa\xaf\xaa\xb6\xe2\x16\x6d\xf5\x83\xb1\xae\x2e\x89\xd3\x36\x34\xeb\x1d\x64\xfd\x64\xc3\x63\x22\x3d\xd6\x2a\xd0\x1a\x53\xfa\xea\xf0\x17\x8e\x28\x9e\xc9\xd4\xc5\xfe\xf3\x29\xf5\x99\xf3\xda\xac\x5b\xa6\xff\xcf\xf5\x49\xc5\xf2\xc3\x27\xdf\x6f\xd2\x3b\x6a\x5b\xf5\xff\x7c\x97\x45\xa7\x88\xa3\xde\x77\x12\xdf\x03\x00\x00\xff\xff\xb5\x16\x6f\x72\x77\x02\x00\x00")

func infrastructureModulesIamIamTfBytes() ([]byte, error) {
	return bindataRead(
		_infrastructureModulesIamIamTf,
		"infrastructure/modules/iam/iam.tf",
	)
}

func infrastructureModulesIamIamTf() (*asset, error) {
	bytes, err := infrastructureModulesIamIamTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "infrastructure/modules/iam/iam.tf", size: 631, mode: os.FileMode(420), modTime: time.Unix(1459260005, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _infrastructureModulesIamOutputsTf = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\xca\x2f\x2d\x29\x28\x2d\x51\x50\xca\x49\xcc\x4d\x4a\x49\x8c\x4f\x2b\xcd\x4b\x2e\xc9\xcc\xcf\x8b\x2f\xca\xcf\x49\x8d\xcf\x4c\x51\x52\xa8\xe6\x52\x50\x28\x4b\xcc\x29\x4d\x55\xb0\x55\x50\x52\xa9\x4e\x2c\x2f\x8e\xcf\x4c\xcc\x05\xcb\xeb\xa1\x69\xd2\x4b\x2c\xca\xab\x55\xe2\xaa\x05\x04\x00\x00\xff\xff\x60\xb0\xec\x67\x55\x00\x00\x00")

func infrastructureModulesIamOutputsTfBytes() ([]byte, error) {
	return bindataRead(
		_infrastructureModulesIamOutputsTf,
		"infrastructure/modules/iam/outputs.tf",
	)
}

func infrastructureModulesIamOutputsTf() (*asset, error) {
	bytes, err := infrastructureModulesIamOutputsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "infrastructure/modules/iam/outputs.tf", size: 85, mode: os.FileMode(420), modTime: time.Unix(1457019469, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"functions/hello/index.js": functionsHelloIndexJs,
	"infrastructure/.gitignore": infrastructureGitignore,
	"infrastructure/_env/main.tf": infrastructure_envMainTf,
	"infrastructure/modules/iam/iam.tf": infrastructureModulesIamIamTf,
	"infrastructure/modules/iam/outputs.tf": infrastructureModulesIamOutputsTf,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"functions": &bintree{nil, map[string]*bintree{
		"hello": &bintree{nil, map[string]*bintree{
			"index.js": &bintree{functionsHelloIndexJs, map[string]*bintree{}},
		}},
	}},
	"infrastructure": &bintree{nil, map[string]*bintree{
		".gitignore": &bintree{infrastructureGitignore, map[string]*bintree{}},
		"_env": &bintree{nil, map[string]*bintree{
			"main.tf": &bintree{infrastructure_envMainTf, map[string]*bintree{}},
		}},
		"modules": &bintree{nil, map[string]*bintree{
			"iam": &bintree{nil, map[string]*bintree{
				"iam.tf": &bintree{infrastructureModulesIamIamTf, map[string]*bintree{}},
				"outputs.tf": &bintree{infrastructureModulesIamOutputsTf, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

