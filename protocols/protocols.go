// Code generated by go-bindata.
// sources:
// .protocols/Tweets.proto
// DO NOT EDIT!

package protocols

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _tweetsProto = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x41\x4b\xc3\x40\x10\x85\xef\xf3\x2b\x1e\x3d\xe6\x52\x6c\x6b\x55\x42\x0e\x45\x50\x04\xf1\x20\xbd\xcb\x92\x8c\x69\x30\xd9\x89\x3b\x13\xdb\x52\xfa\xdf\x25\x69\x53\xd6\xde\xbc\xbe\x97\x99\x2f\xf3\xad\xee\xbd\xb9\x1d\x32\x4c\xda\x20\x26\xf3\x49\x4a\xad\xcb\xbf\x5c\xc9\x58\x6f\x99\x4d\x53\x22\x69\xad\x12\x8f\x52\x3e\xc6\x6a\xfc\x3c\x97\x7a\x92\x12\x35\xac\x7a\x99\xc0\x81\x80\xca\xdb\x72\x01\xc0\xfa\xe4\xa5\x40\x86\x9b\x94\x00\xb5\x50\xf9\x12\x30\xde\x19\x32\xcc\xe2\x2c\x0f\xec\x8c\x8b\x55\x5f\xcc\xd3\xd3\x8e\xf9\x0c\xc0\xa7\xfb\x91\x50\x19\x3f\x4a\xe7\xfb\x72\x11\x97\x81\x07\xc4\xd8\xdd\xc6\x1b\x6b\xe7\x4b\x64\x58\xc6\x59\xa7\x1c\xde\x5c\xd3\x9f\x70\x77\x9d\xbf\x4a\xee\x86\x4b\x33\xdc\xc7\x8c\xbe\x7b\x92\xba\x96\x2d\x07\x1d\x49\x0f\x29\x1d\x89\xa6\x49\x42\x48\xb0\xc2\x3b\x7f\x77\xac\x06\xdb\x38\xc3\x5e\x3a\xe8\x46\xba\xba\x40\x21\x79\xd7\xb0\x37\x42\x32\xbd\x78\x7a\x66\x1b\x54\x8d\x43\x87\xe8\x47\x94\x5d\xc8\x37\x6b\x0e\xcd\xb5\xb4\xba\x6a\xaa\xb3\xb5\xbf\x64\x6d\xc5\x2b\xff\x17\x7d\x9e\xea\xd9\x81\xdb\xc1\xfd\xf9\x01\x07\xa3\x7a\xc2\x1f\xe9\x37\x00\x00\xff\xff\x64\x7a\x49\x1b\x22\x02\x00\x00"

func tweetsProtoBytes() ([]byte, error) {
	return bindataRead(
		_tweetsProto,
		"Tweets.proto",
	)
}

func tweetsProto() (*asset, error) {
	bytes, err := tweetsProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Tweets.proto", size: 546, mode: os.FileMode(420), modTime: time.Unix(1516769341, 0)}
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
	"Tweets.proto": tweetsProto,
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
	"Tweets.proto": {tweetsProto, map[string]*bintree{}},
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
