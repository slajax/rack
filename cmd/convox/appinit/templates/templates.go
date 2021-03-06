// Code generated by go-bindata.
// sources:
// appinit/templates/Dockerfile
// appinit/templates/dockerignore
// DO NOT EDIT!

package templates

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

var _appinitTemplatesDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x41\x4f\xc3\x20\x18\x86\xef\xfc\x8a\xf7\xb0\xec\xe4\xca\x7d\x49\x4f\x3a\xa3\x31\x5b\x4d\x13\x35\x1e\x29\xfd\x6c\x49\x5b\x20\x08\x4b\x0c\xe1\xbf\x1b\xca\x76\x50\x13\x4f\x7c\x90\xe7\xe1\x85\xf7\xbe\x6d\x8e\x18\xc9\x99\x29\x70\x49\xbd\x70\x8c\xb5\x2f\x27\xc8\x1e\xdc\x2f\x16\xdb\x2d\x06\xe5\x21\x67\xa3\x09\xa3\xf7\xf6\x73\xcf\xf9\xa0\xfc\x18\xba\x4a\x9a\x85\x5f\xcc\xb2\xec\xba\xa0\xe6\xde\x0a\x39\xed\x62\x44\x35\x29\xdd\x23\x25\x76\x38\xbd\xe2\xa1\x39\x1e\x6a\x2e\xac\x65\x6f\x4d\xfb\x74\xf7\xd8\x62\xdd\xc4\x08\xf5\x81\x8a\xf4\x59\x39\xa3\x17\xd2\xfe\x2a\xc4\x08\x27\xf4\x40\xd8\x4c\xf4\x75\x83\xcd\x59\xcc\x81\xb0\xaf\x7f\xc3\x31\xae\x04\x52\xaa\xf3\x58\xb0\x94\xb2\x4f\x25\xfe\xc7\x94\xd3\xac\x23\x69\x16\xab\x66\x2a\xfe\x9f\x83\x0b\xce\x6e\x9b\xe7\x77\x54\xe5\xa5\x6b\x2b\xb9\x92\x7f\xff\xca\x3b\xa5\xf9\xf5\xae\xec\x15\x45\x0a\x39\x12\xfb\x0e\x00\x00\xff\xff\xdd\x1d\x1b\xfe\x6b\x01\x00\x00")

func appinitTemplatesDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_appinitTemplatesDockerfile,
		"appinit/templates/Dockerfile",
	)
}

func appinitTemplatesDockerfile() (*asset, error) {
	bytes, err := appinitTemplatesDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "appinit/templates/Dockerfile", size: 363, mode: os.FileMode(420), modTime: time.Unix(1475420665, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _appinitTemplatesDockerignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x4b\xce\xcf\x2b\xcb\xaf\xe0\xd2\x4b\xcd\x2b\xe3\xd2\x4b\xcf\x2c\x01\x13\x99\xe9\x79\xf9\x45\xa9\x5c\x5c\xd5\xd5\x0a\x99\x69\x0a\x7a\x10\xae\x5b\x66\x4e\x6a\xb1\x42\x6d\x2d\x48\xb4\x28\x31\x2f\x3d\x55\x41\x25\x2d\x33\x27\x55\xc1\xca\x16\x5d\x45\x75\x35\x54\xaa\xb6\x56\xa1\xba\x5a\x21\x35\x2f\x05\xaa\x0d\xca\x02\x04\x00\x00\xff\xff\xc9\x75\xf6\x46\x75\x00\x00\x00")

func appinitTemplatesDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_appinitTemplatesDockerignore,
		"appinit/templates/dockerignore",
	)
}

func appinitTemplatesDockerignore() (*asset, error) {
	bytes, err := appinitTemplatesDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "appinit/templates/dockerignore", size: 117, mode: os.FileMode(420), modTime: time.Unix(1475420665, 0)}
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
	"appinit/templates/Dockerfile": appinitTemplatesDockerfile,
	"appinit/templates/dockerignore": appinitTemplatesDockerignore,
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
	"appinit": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"Dockerfile": &bintree{appinitTemplatesDockerfile, map[string]*bintree{}},
			"dockerignore": &bintree{appinitTemplatesDockerignore, map[string]*bintree{}},
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

