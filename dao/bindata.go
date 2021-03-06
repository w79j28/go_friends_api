// Code generated by go-bindata.
// sources:
// db.sql
// DO NOT EDIT!

package dao

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

var _dbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xcd\x4a\xc4\x30\x14\x46\xd7\x53\xe8\x3b\x7c\x0f\x20\x81\x66\x2b\x2e\x4a\xb9\x03\x23\x75\x2a\x6d\xd1\xa5\x14\x9a\xca\x05\x67\x26\xe6\xc7\xe7\x97\xb4\x18\x33\x10\x7f\x96\xf7\x7e\x1c\x38\x9c\xa6\xa7\x7a\x24\x74\x3d\x7a\x7a\x6c\xeb\x86\xf0\x74\xa0\x67\x7c\xbc\x2c\x86\xd5\x79\xb6\xa8\x87\xb2\xc0\x40\x2d\x35\x23\x7c\x25\x78\x46\x3d\xc0\x5b\x65\x2a\x9e\x6f\xca\x02\x40\x78\xab\xd3\xc4\x6f\x71\x59\xaf\xaf\x51\x26\x8c\xfc\x66\xe4\x35\x23\xd7\x6b\xdd\xf6\x7d\xf7\x80\x49\x73\x54\x58\x36\x04\x2d\xed\x47\xdc\x77\x87\xe3\xba\x06\x0a\xbe\x42\x77\xc4\x22\xc2\xc1\x73\x85\xbb\xcd\xf1\x17\x40\xa6\x80\x0c\x40\x10\xbc\x2d\x8b\x5d\x59\xfc\x18\x43\x2b\x73\x62\x6b\xf9\x72\xce\xf7\x30\xea\xdd\x2b\xeb\x2e\x26\x53\x24\x6e\xd9\x2a\x6e\x32\xaf\xca\x65\xa2\x6c\x43\xca\x68\x61\xdd\xe4\xbc\xc5\x75\xa5\xc4\x4d\xff\x15\x4a\x8b\x68\xf3\xdf\x54\x5a\x6c\x26\xb1\xd4\xee\x33\x00\x00\xff\xff\x96\x12\x37\x7e\x33\x02\x00\x00")

func dbSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbSql,
		"db.sql",
	)
}

func dbSql() (*asset, error) {
	bytes, err := dbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db.sql", size: 563, mode: os.FileMode(438), modTime: time.Unix(1497510442, 0)}
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
	"db.sql": dbSql,
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
	"db.sql": &bintree{dbSql, map[string]*bintree{}},
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

