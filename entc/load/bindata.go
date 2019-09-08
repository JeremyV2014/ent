// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package load

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

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x51\x5d\x6b\xdb\x30\x14\x7d\xb6\x7e\xc5\x99\xe9\xa8\x5d\x52\xa5\xed\xdb\x06\x79\x28\x6d\x06\x19\x5b\x3b\x48\x61\x0f\x5d\x29\x8a\x7d\x9d\x88\x3a\x92\x77\xa5\x94\x05\xa1\xff\x3e\x24\x27\x61\x7b\xb2\xa5\x73\xee\xf9\xd0\x0d\x61\x7a\x21\xee\xec\xb0\x67\xbd\xde\x78\xdc\x5c\x5d\x7f\xba\x1c\x98\x1c\x19\x8f\x2f\xaa\xa1\x95\xb5\x6f\x58\x98\x46\xe2\xb6\xef\x91\x49\x0e\x09\xe7\x77\x6a\xa5\x78\xda\x68\x07\x67\x77\xdc\x10\x1a\xdb\x12\xb4\x43\xaf\x1b\x32\x8e\x5a\xec\x4c\x4b\x0c\xbf\x21\xdc\x0e\xaa\xd9\x10\x6e\xe4\xd5\x11\x45\x67\x77\xa6\x15\xda\x64\xfc\xdb\xe2\x6e\xfe\xb0\x9c\xa3\xd3\x3d\xe1\x70\xc7\xd6\x7a\xb4\x9a\xa9\xf1\x96\xf7\xb0\x1d\xfc\x3f\x66\x9e\x89\xa4\xb8\x98\xc6\x28\x44\x08\x68\xa9\xd3\x86\x50\x6e\x95\x36\x25\x62\x14\xd3\x29\xee\x52\x9e\x35\x19\x62\xe5\xa9\xc5\x6a\x8f\x73\x32\xbe\x39\x5d\x9d\x4b\xdc\x3f\xe2\xe1\xf1\x09\xf3\xfb\xc5\x93\x14\x83\x6a\xde\xd4\x9a\x90\x34\x84\xd0\xdb\xc1\xb2\x47\x25\x8a\xd2\xba\x52\x14\xe5\x6a\xef\x29\xfd\x84\x00\x4f\xdb\xa1\x57\x9e\x50\x8e\x2c\x97\x2d\x33\x34\xb0\x36\xbe\x43\xf9\xf1\x77\x09\xf9\xe3\xa0\x18\xa3\xa8\x73\xcc\xb3\x95\x72\x84\xcf\x33\xe4\xef\x11\x4f\xb3\xef\x8a\xe1\x9a\x0d\x6d\x95\xc3\x0c\xcf\x2f\x64\xbc\x5c\x18\x4f\xdc\xa9\x86\x42\x96\x66\x65\xd6\x84\xb3\xd7\x09\xce\x8c\xda\x66\x19\xf9\xa0\xb6\xe4\x92\x7e\x51\x84\x70\x79\xd0\x8f\x51\xa6\xc3\x29\x8a\x0b\xb1\x3c\xcc\xc4\x38\xc9\x5a\x64\x5a\x5c\xc6\x28\xa2\x10\xdd\xce\x34\xb9\x73\x55\x23\x88\x22\x05\xe9\xb5\x21\x87\xe7\x97\xe7\x97\x54\x5a\x14\x9d\x65\xbc\x4e\x0e\xf9\x92\xef\x18\xe5\x98\x37\x88\xa2\x58\x4d\x40\xcc\x09\xfb\xae\xd8\x6d\x54\xbf\xcc\x60\x35\x72\x6a\x51\x14\xba\xcb\x8c\x0f\x33\x18\xdd\xe7\x99\xa2\x53\xba\xaf\x88\x39\xc1\xa9\xc2\xe8\x3b\x83\x1a\x06\x32\x6d\x95\x8f\x13\xac\x6a\x91\x50\xeb\xe4\xd2\xb7\x76\xe7\xe5\x4f\xd6\x9e\xaa\xbc\x0f\xf9\xd5\x6a\x73\x24\x8e\x71\xab\xf2\x97\x29\xeb\xba\x3e\x75\x3b\xba\x24\x7b\xcb\xb9\xe4\xa8\x45\xcc\xa3\xd6\xd2\xb3\x36\xeb\xc4\x91\xf3\xc4\xa9\xea\x3a\x73\xe6\x7f\xb4\xaf\xae\xb3\xd2\x7f\x5b\x1f\x4b\x8d\x4b\x3f\x3c\x66\x8c\xe2\x6f\x00\x00\x00\xff\xff\xe4\x6e\x0c\x4d\x4b\x03\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 843, mode: os.FileMode(420), modTime: time.Unix(1567330508, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4b\x6f\xdb\x38\x10\x3e\x4b\xbf\x62\x6a\xa0\x81\x14\x78\xe5\x6e\x6f\xeb\xc2\x87\xa2\x4d\x81\xec\x23\x2d\x9a\xee\x5e\x82\x20\xa5\xc5\xa1\xcd\x46\x22\x55\x92\x76\x92\x06\xf9\xef\x0b\xbe\xf4\xb0\xec\xa0\xdb\x6e\x02\x04\x36\xe7\x45\xce\x37\x1f\x87\xe3\xd9\x0c\xde\xc8\xe6\x4e\xf1\xd5\xda\xc0\xcb\x17\xbf\xfe\xf6\x4b\xa3\x50\xa3\x30\xf0\x8e\x94\xb8\x94\xf2\x1a\x4e\x45\x59\xc0\xeb\xaa\x02\x67\xa4\xc1\xea\xd5\x16\x69\x91\xce\x66\xf0\x69\xcd\x35\x68\xb9\x51\x25\x42\x29\x29\x02\xd7\x50\xf1\x12\x85\x46\x0a\x1b\x41\x51\x81\x59\x23\xbc\x6e\x48\xb9\x46\x78\x59\xbc\x88\x5a\x60\x72\x23\xa8\x0d\xc1\x85\x33\xf9\xf3\xf4\xcd\xc9\xd9\xf9\x09\x30\x5e\x61\x94\x29\x29\x0d\x50\xae\xb0\x34\x52\xdd\x81\x64\x60\x7a\xfb\x19\x85\x58\xa4\x69\x43\xca\x6b\xb2\x42\xa8\x24\xa1\x69\xca\xeb\x46\x2a\x03\x59\x9a\x4c\x50\x94\x92\x72\xb1\x9a\x7d\xd1\x52\x4c\xd2\x64\xc2\x6a\x63\x3f\x14\xb2\x0a\x4b\x33\x49\xd3\x64\xb2\xe2\x66\xbd\x59\x16\xa5\xac\x67\x2c\x24\xcc\x45\xb9\x59\x12\x23\xd5\x0c\x85\x99\xe9\x72\x8d\x35\x99\x21\x5d\xe1\x77\x39\x4c\xfe\x43\x50\xc6\xb1\xa2\x93\x34\x4f\x2d\x0c\xe7\x4e\x06\x0a\x43\x01\x34\x10\x01\x28\x4c\x11\x14\x66\x4d\x0c\xdc\x10\xed\xf2\x44\x0a\x4c\xc9\x1a\x08\x94\xb2\x6e\x2a\x6e\xc1\xd6\xa8\x20\x60\x51\xa4\xe6\xae\xc1\x18\x52\x1b\xb5\x29\x0d\xdc\xa7\xc9\x19\xa9\x11\xe2\x9f\x36\x8a\x8b\x55\xbb\xfc\x6c\x51\x9a\x4f\x04\xa9\x71\x2a\x6b\x6e\xb0\x6e\xcc\xdd\xe4\x73\x9a\x9c\xd0\x15\xea\x68\x75\x71\x79\x6c\xd7\x3b\x4e\x16\x1d\x3d\xf4\x7a\x67\x73\xd3\x9d\x97\x5b\x0f\xbd\x5c\xfa\x3b\x6e\xa7\x82\xe2\x6d\xdc\xee\xe2\xf2\xd8\xad\x87\x6e\xdc\x9b\x0c\xfd\xce\x5d\x8e\x61\xd3\x8b\xcb\xe3\xde\x3a\xfa\x79\x18\xae\xf6\xec\xfa\xe0\x0a\xe0\x8d\xc7\xf8\x7b\xf9\x0f\xc0\xef\x1d\x0f\xa0\xdf\x87\xff\x20\xf4\x9f\x6c\x98\xf6\xcf\x9d\xbc\x70\xb2\xe0\x61\xb7\xd9\xf1\x20\x5d\x45\xf7\xed\x61\xc8\x6a\x07\x38\xfe\xad\xb7\xc5\x31\x17\x66\x88\xb6\xe6\xdf\x76\xb6\xf8\x5b\xf0\xaf\x9b\xd6\x67\x29\x65\x35\xf4\xd8\x38\xfd\xd0\xe7\x8c\x57\x15\x59\x56\x78\xc8\x47\x04\xfd\xd0\xeb\x7d\x63\xb8\x14\xa4\x3a\xe4\x25\x83\x7e\xe8\xf5\x16\x19\xd9\x54\xe6\xe0\xf9\xa8\xd7\xef\x24\xd5\x50\x62\x30\xba\xee\x49\xca\xe9\xaf\xf6\xfa\x9e\xd6\xf5\xc6\xb4\xd9\x8d\x7d\x79\xd4\x0f\xdd\xfe\x21\x15\xa7\xb6\x25\x38\xb6\x77\xc0\x47\xb7\x6d\xab\xdf\x43\xd6\x3e\xbf\x77\x28\x7b\x6b\x50\x59\xd0\x02\xf1\x1c\x6b\x80\x22\xe3\x02\x69\x6c\xac\xbe\x01\xc5\x36\xd1\x8b\xd5\xb1\xb5\x65\x52\x20\xd1\x21\x02\xb5\xac\x1e\xda\x3d\x42\xe6\x9d\x80\x23\x0e\xbf\x91\x75\x6d\x5f\xa0\x1d\xc3\xd2\x8b\x87\xb6\x1f\xae\x57\x1f\x88\x59\xef\xda\x36\xd7\xab\xab\x86\x98\xf5\x4e\x27\xab\x97\x48\xed\xe5\x0d\x35\x8a\xdd\x2b\x88\xf7\xc0\xec\x7a\xdd\xb8\x25\x38\xf1\x0f\x74\x04\xe7\xb7\xa7\x21\xfc\x6f\xd0\x7d\x6f\xd1\x3e\x22\xf3\x9b\x0f\xed\x14\xb2\xab\xf1\xee\x1f\x91\x85\xee\xe0\x5b\x7f\x67\x7c\xa0\x31\x0c\xe1\xdd\xd7\x0f\x4e\xc5\x16\x95\xc6\x5d\x53\xee\xc5\xbb\xdb\x7f\xdd\x70\x35\xaa\x9a\x0a\xe2\x3d\x55\xf3\x8f\xc6\xb8\x6c\x5e\xfe\x03\x75\xf3\x8e\x5d\xe1\x42\xa6\xed\x55\x7f\x24\xd3\xf0\x7a\x5e\x5c\x0e\x91\x3e\xfc\x62\xee\x5a\x1e\x7c\xaf\xce\xf0\xc6\xd5\xa3\x54\x48\x0c\xba\x24\x43\x46\x36\xb8\x4f\xcb\x7d\xa3\xa8\x4b\xc5\x1b\x23\x55\x91\xb2\x8d\x28\xa3\x67\x86\x14\x8e\xad\x45\xf1\xb6\xb5\xc8\x43\x91\xef\xd3\x44\x20\xcc\x17\x70\x64\x97\xf7\x69\x62\xa9\x35\xf7\x34\x40\x5a\x7c\x22\xab\xa9\x95\xdd\x35\x38\x6f\x65\x96\x8d\x69\xe2\x58\xdd\x0a\xed\xc2\x0a\x3d\x62\x73\x2f\xf4\x0b\x2b\x0e\x3c\x98\x3b\x71\x58\x58\x79\xac\xf9\xdc\xca\xe3\xc2\x2b\x58\x88\xef\x14\x2c\xc4\x7f\x48\x13\xce\x40\x21\xb3\x47\xf6\x9a\x57\x6e\xf9\x6c\x01\x82\x57\x36\x9d\x44\xa0\x15\xc3\xa2\x4d\x5f\x21\xcb\x9d\xab\x42\xb3\x51\x02\x04\x06\x64\xff\x22\x4a\xaf\x49\x15\xc6\x27\x37\x46\xa2\x6b\x9b\xbd\x71\x8c\x0b\x83\xca\x4e\x77\xf6\x9b\x04\x02\xbf\x9f\xbf\x3f\xb3\xce\x8e\x5e\x25\x11\xb0\xb4\xc8\x5b\x57\xea\x4d\x6c\x80\xe0\x2c\x97\x5f\xb0\x34\xe1\x23\x14\x65\xb0\x69\xa6\xe3\xde\x96\xb5\x61\xa7\x1c\xb2\x25\x5c\x5c\x2e\xef\x0c\x4e\x01\x95\xb2\xff\xb6\x62\xf7\x69\xa2\x5d\xa9\xbc\xef\xbd\x07\x88\x0b\x3f\x38\x67\x61\xdc\x75\xf5\x79\xcf\x42\xe4\x3c\x77\xa5\xc9\xf2\x87\x34\x09\x0c\x73\x21\xe7\x0b\xd0\x84\xa1\xe7\x62\xb4\x75\xe0\x5a\x6d\x0f\xcd\x88\x19\xaf\xa6\xc0\x6a\x53\x9c\xd8\xb3\xb0\x6c\x12\x0e\xfe\xfc\xeb\x1c\x9e\x6f\x27\x53\xd0\x9e\x02\xd6\xdd\x83\xcd\xa4\x82\xab\x29\xb8\x4a\x29\x22\x2c\x53\x3d\xf1\x6d\x54\x46\xad\x98\xf5\x08\x99\xe5\x69\x92\x68\x67\x7d\xe4\x4e\x65\xcd\x7a\x1c\xf3\x93\x51\x47\xb4\x1e\x27\xa3\x2a\x12\xb3\x47\xe1\x56\xe5\x79\xdc\xa3\x67\xd4\x74\x1c\x6d\x87\x97\x79\xb7\x59\x1c\x57\xac\x3a\x4e\x29\x9d\x3a\x4a\x9c\xba\x9d\x0e\xe6\x51\xdd\x4a\x9c\xbe\x1b\x03\x9c\x41\x85\x22\x63\xb4\xe8\xa4\xb9\xb3\x0a\x93\xc9\xbc\x3b\x60\x9c\x55\x7c\x4d\x7c\x16\xfd\x21\x66\xee\xb2\x18\x8c\x35\x9d\xe9\x43\x9a\xd8\x9a\x32\x5a\xb8\xf9\xef\xd9\x02\x5e\x38\xfc\x13\xcd\xbc\x64\x01\x47\x41\x19\xac\x75\x11\xfa\xd3\x02\x48\xd3\xa0\xa0\x59\x94\x4c\x41\x87\x7b\xe4\x9b\x5a\x9f\x47\xae\xfb\x3d\x25\x8d\xb0\xa3\x91\xdb\xdd\x05\xd5\x85\xef\xba\xbd\xa3\x9e\xf8\xa3\xb5\x9d\x6f\xc0\xb0\xdc\x87\x8c\x3f\x2c\xfa\x09\x84\xdf\x23\x4f\x99\x02\xa7\xb7\x5d\x12\xe1\x0c\x2e\x70\x50\x70\x7a\x3b\xba\x0f\x45\xfc\x9d\xd4\x4b\xf1\x34\x1e\xff\xc8\x7d\x73\xe5\x74\x69\xcf\xc1\xc5\xf0\x10\x58\xa9\xaf\xdb\xdc\x49\x43\x0d\xfb\x97\xc0\x8a\x3b\xfa\x3f\x0c\x7a\xa4\x7d\x93\x8a\xd0\xaa\x32\x9d\x87\x86\xd9\xb5\x0c\xb8\x51\xa4\xd1\xfd\x19\x33\xc8\x6b\x34\x6b\x49\xe1\x86\x9b\x35\x28\x2c\xe5\x16\x15\x18\x09\x28\xf4\x46\x21\x08\x09\x0d\x11\xbc\xd4\x76\x40\xad\x7d\x78\x2e\x56\xa1\x35\x8e\x3a\xd2\xa8\x2f\xb2\xf8\x76\xb6\x3f\xd8\x76\x3b\x24\x45\x86\x0a\x6c\xb8\x2c\xf7\xe8\x32\xd8\x3a\xdc\xfd\x61\xb2\xfc\x15\x6c\xfb\x65\x4d\xac\xff\x62\x4f\x45\x63\x46\xfe\xc0\xa1\xb8\x5b\x5b\x96\xd0\x49\xc1\x05\xf1\xf7\xe6\xc1\xd6\x2b\x60\x37\x70\xcf\xf2\xa9\xb3\xea\x00\xf4\x9c\x1d\xe1\xe7\xc5\x3f\x0b\x5f\xff\x22\x8e\xd0\xf3\x37\xc7\x83\x67\x0d\x9f\x10\x3b\x9f\xcd\x1e\xe8\x30\xdc\xd8\xc7\x90\xf3\x49\x8c\x80\x8b\x77\x61\x04\x5d\x54\xfc\x2c\x78\xc3\x26\x30\x82\x2f\xde\x59\x0f\xa0\x33\x7e\x42\x04\x63\x52\x7b\x30\xe4\x6d\x53\x78\x0c\xc5\x98\x4d\x87\xa3\x4b\xb4\x9d\x17\x0c\xf4\x27\x86\x7c\xb0\xb2\x67\xb3\x6d\xcb\x14\x7f\x70\x41\xb3\x1c\x16\x8b\x56\xff\xc1\x28\x77\x74\x03\x0b\x30\xc5\x49\x85\x75\x36\x68\x1d\x26\x7d\x48\xff\x0d\x00\x00\xff\xff\xb1\x82\x4a\x61\x6f\x14\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 5231, mode: os.FileMode(420), modTime: time.Unix(1567952289, 0)}
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
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
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
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
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
