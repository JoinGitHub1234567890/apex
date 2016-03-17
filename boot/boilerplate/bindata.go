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

var _functionsHelloIndexJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\xcc\x51\x0a\x83\x30\x0c\x80\xe1\xf7\x9e\x22\x2f\xa3\x2d\x48\x0f\x20\xec\x30\x92\x66\xea\x08\x89\x34\x71\x2b\x88\x77\x1f\xbe\x0c\x7c\xfe\x7f\xbe\x80\x2a\xa6\x4c\x85\x75\x4e\xd1\x7c\x6a\xbe\xca\x0c\xaf\x5d\xd0\x57\x95\x98\x03\xf5\x4d\x9b\x5b\x59\x26\xa9\x4c\xf0\xfc\xb7\x44\x03\xa0\xf7\x0c\x47\x00\xb8\x31\x5b\x53\x24\xb3\x0b\xa2\x0f\x89\x8f\xf0\x78\xc7\x01\x28\x5f\xa3\xf7\x62\x3b\x22\x51\x4d\x07\x2c\xc4\xac\x23\xc4\xaf\x36\xae\x11\xce\x1c\xce\x5f\x00\x00\x00\xff\xff\x4b\xf0\xa8\x46\x92\x00\x00\x00")

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

	info := bindataFileInfo{name: "functions/hello/index.js", size: 146, mode: os.FileMode(420), modTime: time.Unix(1456944518, 0)}
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

var _infrastructure_envMainTf = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xcc\x41\x0e\x02\x21\x0c\x40\xd1\x7d\x4f\xd1\x34\xae\x3b\x27\xf0\x2c\xa4\x0e\x98\x90\x94\xa9\x01\xea\x86\x70\x77\x51\xd6\xae\x7f\xfe\x83\x62\xd1\x35\x21\x65\x29\x84\x03\x10\x9b\x79\x3d\x13\xde\x91\x98\x8f\x5d\xdb\xf1\xad\x30\x01\xcc\xfb\xcb\x3b\x92\x4a\x79\x44\x09\x4f\xbf\xce\x9e\xed\x0a\xd5\x34\x85\x1c\xb7\xf0\x16\xf5\x1f\x70\x1b\xfb\xe7\xb5\xf3\x9f\x65\x2e\xf7\x13\x00\x00\xff\xff\x52\x50\x48\x14\x85\x00\x00\x00")

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

	info := bindataFileInfo{name: "infrastructure/_env/main.tf", size: 133, mode: os.FileMode(420), modTime: time.Unix(1458142461, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _infrastructureModulesIamIamTf = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x91\xbd\x6e\xf2\x30\x14\x86\x77\x5f\x85\x65\x7d\x13\x82\x08\x58\x3e\x29\xa2\x43\x06\xba\xb6\xa2\x52\x17\x84\xa2\x53\xe7\xa4\xb5\xe4\x1f\x64\x3b\x8d\x5a\x94\x7b\xaf\x8f\x09\xf4\xbf\x5d\xca\xc0\xe0\xf7\x39\xc7\x7e\x9f\x30\x8f\xc1\x75\x5e\x22\x17\xd0\x87\x5a\x81\xa9\xbd\xd3\x28\xb8\xd0\x60\xee\x1a\xa8\xdb\xce\xca\xa8\x9c\x15\xfc\xc0\x38\xb7\x60\x90\x5f\x7c\x0e\x53\x04\x21\x74\x06\xf3\x74\xbd\x77\x5a\xc9\xa7\x04\xae\x56\xeb\xab\x4b\x46\x93\xe2\x16\x7d\x20\xb4\xe4\x62\x39\x5f\x2c\x67\x8b\xf9\x6c\xf1\x5f\x4c\x29\xba\x89\x10\xd1\xa0\x8d\x29\xdc\xa6\x03\x9e\xef\xa2\x9f\x58\xb7\x2d\x4a\x3a\x17\x95\xd6\xae\xcf\x7c\x0e\xae\xbd\xb2\x52\xed\x41\xa7\xec\x44\xd3\x2a\xf4\x8f\x4a\x22\x0d\x1c\xdf\x58\x80\x81\x67\x67\x53\xb9\x42\x3a\x23\x46\x72\x38\xef\xa9\x8e\x05\x12\x1f\x62\x28\xab\x5c\x62\x43\x06\x32\x30\xa4\xff\x1d\x1b\x18\xb5\x18\xd8\x37\xb2\xc6\xba\xc9\x99\xd4\xae\x6b\x7a\x88\xf2\x41\xbb\xfb\x90\xf4\x68\x5d\x83\x94\x18\xc2\x7b\x7d\x3f\x70\x89\xa2\x9d\x44\xfd\x3b\xbc\xbd\xa5\xf8\xe0\xbc\x50\xcd\x40\xf4\xdf\xb9\x3e\xab\xd8\xbe\xfa\xa4\xf7\x95\x93\x93\xb6\xdd\xf4\xd7\xef\xb2\x19\x15\x51\x34\xf9\x4a\xe2\x4b\x00\x00\x00\xff\xff\x17\xa1\xca\x4c\x72\x02\x00\x00")

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

	info := bindataFileInfo{name: "infrastructure/modules/iam/iam.tf", size: 626, mode: os.FileMode(420), modTime: time.Unix(1457019469, 0)}
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

