// Code generated by go-bindata.
// sources:
// models/templates/app.tmpl
// DO NOT EDIT!

package models

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

var _templatesAppTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x3c\xfd\x73\xdb\x38\x76\xbf\xef\x5f\x81\x41\xd3\xca\xde\x95\x65\xd9\xdb\x6b\x7b\xdc\xa6\x33\x8e\xec\xec\xfa\xce\x4e\x54\xc9\x9b\x9d\x5e\x92\xe9\xd0\x24\x24\xf1\x4c\x01\x5c\x00\x74\xec\xd5\xe8\x7f\xef\xe0\x83\x24\x3e\x29\xd9\x4e\xe6\x9a\x4e\xf7\x2c\xf2\xe1\xe1\xe1\xe1\x7d\xe3\x81\x9b\x0d\xc8\xd1\xa2\xc0\x08\xc0\xb4\xaa\x20\xd8\x6e\xbf\x03\x60\xb3\x01\xaf\xd2\xaa\x02\xc9\x6b\x30\x3a\xab\xaa\xee\xe1\x3a\xc5\xc5\x02\x31\x2e\xdf\x5c\x37\x3f\xd4\xeb\xef\x00\x00\x00\x9e\xfd\x36\xbf\x41\xeb\xaa\x4c\x39\x7a\x4b\xe8\x3a\xe5\x1f\x10\x65\x05\xc1\x10\x24\x00\x9e\x8e\x4f\xc6\x47\xe3\x3f\x1f\x8d\xff\x0c\x87\x0a\x7c\x42\x70\x5e\xf0\x82\x60\x06\x13\x8d\x42\xce\xc4\x35\x0e\x00\x6f\xd3\x32\xc5\x19\xa2\x47\x59\x07\xea\xce\xed\x0d\xaa\x28\xc9\x10\x63\xbb\xc6\xc0\x4b\xcc\x11\xc5\x69\x29\x26\x07\xf0\x2d\x4e\x92\x8b\xdf\xeb\xb4\x14\xc4\x7c\x14\x4f\x66\x68\x01\x13\x03\x0c\x6c\x87\x00\xfe\x0f\x62\x10\x7c\x06\xdb\x61\x83\x65\x4a\x8b\xfb\x94\xa3\x1d\x48\x1a\xa8\x30\x8e\x37\x65\x8a\xef\xe6\x28\xab\x69\xc1\x1f\x7f\xa6\xa4\xae\x04\xc7\x36\x26\x3a\x90\x80\x8f\x1b\x89\x4d\xf0\xd2\x86\x15\x38\xe1\x67\xb5\x2e\x8d\x14\x4e\x53\x9a\xae\x11\x47\x54\x0e\xed\x67\x6e\x25\x60\x9f\xc0\xd8\x20\x7c\xb3\x96\x49\x59\x33\x8e\xa8\xb1\xa3\x00\xc0\x9b\xc7\x0a\x29\xc2\x39\x2d\xf0\x12\x0e\xbb\x57\xe7\x68\x91\xd6\x25\x97\x6f\xed\xe7\x2c\xa3\x45\xc5\x1b\xf1\x81\xfa\x55\xc7\xb5\x49\xcd\x38\x59\xdf\x90\xaa\xc8\x66\x35\xe6\xc5\x1a\x05\x26\xed\x9d\x33\x01\x10\x93\x1c\xfd\x9d\xc5\x26\x0e\xce\x7b\x8e\xaa\x92\x3c\xae\x11\xe6\xd7\xe9\x43\xb1\xae\xd7\xe1\x69\xdf\xd5\xeb\x5b\x44\x23\xd3\x9e\x8e\xc7\x3d\x73\x6a\xbc\xa0\x42\x34\x43\x98\xa7\x4b\x04\xc8\x02\x68\xf6\x23\x06\x38\x01\x77\x08\x55\x80\xd6\x18\x17\x78\x09\xbe\xac\x8a\x12\x81\x5c\xd2\x25\x96\xda\x47\x72\x81\x9f\x49\xf2\x9f\x7a\x29\x56\x68\xbf\x1e\xc5\x17\xf8\xbe\xa0\x04\x0b\x92\x9f\xb1\xab\x4f\xda\x4f\xd3\x0e\xec\x37\x8f\x85\xf0\x3d\x2e\x1f\x41\x5a\x96\xe4\x0b\x48\x33\xb1\x5c\xb1\x58\xbe\x2a\x18\x10\x56\x74\x41\xc9\x1a\x14\x98\x15\x39\x02\x7c\x85\xc0\x87\xe9\x24\x42\xf3\x3b\x62\xbe\x38\x13\x08\x51\xfe\x21\x2d\x6b\xa4\x8c\x89\x34\x1b\x43\x09\x07\x3e\x7b\x8b\xf8\x2b\x7a\xfc\xd6\x7c\x32\x2c\xdd\x33\xd8\xf4\x2b\x43\x60\x5e\xdf\x62\xc4\x99\x46\x24\xf8\xc4\x2a\x94\x15\x8b\x47\xc1\x96\x23\xc9\xa3\x92\xa4\x39\x68\x2c\x13\x40\x38\xaf\x48\x81\x39\xfb\x26\x3c\x9b\xa1\x12\xa5\x2c\xb4\xa0\xaf\x6d\xaa\x66\xa8\x22\xac\xe0\x84\x86\x36\xe9\x65\x93\xcd\x49\x4d\x33\x04\x32\x92\x23\x40\xbb\x69\x3c\x12\x6c\x97\xf1\xb5\xa9\xb8\x59\x21\x70\x65\x6d\x1d\xd3\xf3\x81\xa5\x98\x10\x2c\x08\x6d\x95\x22\x40\x9c\x12\x8c\x08\x59\x57\x05\xe3\xff\x79\xf6\xdb\x3c\x49\x2e\x26\xa7\x49\xa2\x80\x93\xe4\x32\xff\xaf\xe7\x90\xfa\x61\x3a\x01\x4c\xcd\xb7\x1f\x55\x71\xb9\xff\x36\xc4\x55\x5a\x3d\xf6\x23\xf2\x26\x65\x77\x33\x52\x7e\x7d\x29\xbe\x3c\xbb\x06\x02\xb1\x50\xd3\xb4\xaa\xca\x47\xf1\xc7\xc5\x64\x0e\xc4\x8c\x4c\x98\xf7\x38\x51\x4d\xdc\x67\xd1\xe4\x18\x84\x83\xd9\xc5\x7f\xff\x7a\x39\xbb\x38\x3f\x04\x57\xe9\xfa\x36\x4f\x81\xe1\xd0\xc1\x2f\x29\xce\x4b\x44\x81\xd6\x51\xd0\x60\x34\x08\xbe\x2e\xf0\x15\xc2\x4b\xbe\x92\xe4\x9e\x98\xaf\x1c\xab\xe4\xd3\x37\x9d\x44\xf8\xd5\xed\xe4\x87\xe9\x44\x6c\xe3\x73\x77\xb1\x7f\xd7\x3e\x4c\x27\x93\xcb\xf3\xd9\x57\xdf\x34\x31\xb3\x40\x1c\x9e\xde\x8a\x10\xaf\xd3\xaa\x2a\xf0\xd2\x54\x3a\x38\x25\x94\x4f\x29\xe1\x24\x23\x8e\x3b\x5c\x71\x5e\xa9\x18\x57\x08\x3c\xc2\x88\x1a\x70\xf0\x97\x9b\x9b\xa9\xb0\xb3\x97\x98\x71\xa1\xfe\xa1\x77\xd2\x00\xa1\x18\xc4\x1c\x76\xdc\xd1\xd3\xb1\xfe\xf9\xe6\x2f\x9e\xd0\x9a\x91\x67\x3d\xeb\xbb\x99\x44\x97\xa7\x5f\xc5\x27\x9b\xcf\xaf\xdc\xa9\xca\x9e\xa5\x09\xf0\x97\x4d\x05\xb6\xc1\xfd\x9e\x21\x26\x5d\x85\xb5\xe1\x66\x0c\x1d\x36\x22\x8d\x4e\x5c\x9e\x5d\x27\x89\x84\x31\x56\x32\xa5\xa4\x42\x94\x17\xc8\x36\xdd\xc2\x17\x33\x56\xaf\x91\x80\x9f\x92\xb2\xc8\x1e\xcf\x49\x56\x7b\xc1\x9c\x63\x2b\x44\x8a\x78\x7a\x74\x32\x3e\x3a\xf9\x77\x63\x12\x09\x34\xe7\x29\x47\x7a\xfc\x47\xeb\x15\x70\xf0\x49\xf0\x8b\xc5\x02\x65\x32\x42\x90\x31\x81\x83\x4d\x93\x5e\xe0\xac\xa8\x9a\xf4\x6f\x8e\xe8\x7d\x91\x21\x15\x35\x94\xd2\x1e\x8d\xd2\x75\xfa\x07\xc1\xe9\x17\x36\xca\xc8\xda\xca\xd8\xcc\x85\x66\xda\xa0\x7d\x04\x90\x71\x96\x74\x0b\xef\x42\x8e\xe6\xdf\xd6\xfa\x6d\xbe\xb5\x30\xc3\x69\xca\x57\x82\xf8\xe3\x8c\xe0\x7b\xf2\x70\x0c\xed\xb7\x82\xa1\x8a\xe5\x36\x2b\x5c\x46\x28\xc8\xc7\x77\xe9\x5a\x6d\x63\xbe\x2e\x70\xc1\x38\x4d\x39\xa1\x1e\x4b\xe0\x8e\x7d\x02\xfb\xee\x15\xf0\xf6\x4b\xf0\xd7\xdb\x11\x83\x73\xf0\x7b\xf1\xb3\x91\x4f\xf5\x00\x6c\x77\x70\xcf\xfc\xd5\x41\x6e\xfb\xb2\xc4\x1e\xe9\x56\x1e\x28\x49\xde\xd6\x58\x51\xb5\x97\x90\x4f\x48\x8e\x7c\x81\x9e\xff\xf8\xa6\xce\xee\x10\xef\x4a\x02\x7f\x21\x85\x96\x90\x23\x38\x14\xff\xa3\xf6\x15\x0e\x8d\x0a\x81\x24\x63\x86\x96\xd2\x92\x6f\xc1\x67\x5f\xdc\xe0\xfc\x47\x1d\xe5\xbb\x58\x15\x52\xaa\x5c\xe5\xb1\x85\xb6\xad\xc0\x6c\x87\x00\x1e\x2b\xc1\x3e\x5e\xc8\xe2\x4c\x41\xf0\xe8\x8f\xa2\x82\x6a\xae\xa8\x30\x6a\x4f\x2c\x90\x15\x38\x47\x0f\x23\xf4\xa0\xf3\x25\x0b\xec\x1a\xad\x09\x7d\x9c\x17\x7f\x48\xa6\x9e\x9c\xfe\x87\xfd\xba\xb1\x2e\x8a\xf4\x9f\x11\x3f\xe3\x4a\x36\x3c\x13\x24\x24\x83\x62\x4f\xdd\xa0\x91\xe3\xb7\xab\x0b\x94\x00\x9c\x51\x37\xc5\x1a\x91\x5a\x0a\xde\x8f\xe3\x31\x8c\x0b\x4a\xb8\x34\x42\x5b\xa3\x09\x46\x91\xaa\x48\x46\x09\xfe\x3b\xb9\xdd\x07\xb4\x29\xa0\x98\xa0\x7b\xd6\x5c\x98\xb2\x4f\x3d\xc8\x29\x5a\x0a\xe5\x7e\x8c\x62\x0f\x0d\x6a\xa2\x74\x18\x41\xca\xb8\xaa\x5a\xd9\xae\xe4\x7d\xcd\xab\x9a\xef\xae\xda\x11\x0d\x07\x46\xfd\x8b\xeb\xe0\x76\x70\xa3\x5d\x63\x78\x44\x97\xeb\x70\xee\x84\x36\xc2\x78\x89\xbc\xd0\x12\x9f\x16\xce\x75\x99\xdf\x89\xff\xdf\x6c\x44\xfe\x29\xf1\x1a\x85\xd2\x50\x75\xb1\x29\x91\xd2\x14\x2f\x11\x78\x75\x27\x2b\xa4\x17\x98\x53\x69\x7b\x59\xb3\x18\x78\x81\xd3\xdb\x12\xe5\x9b\x0d\xa8\xab\x0a\x51\x01\xb9\xdd\x76\x5a\xf1\x8e\x48\x95\x08\xd6\x11\xc5\x93\x39\x2a\x95\x0d\xfd\x08\xc6\xa6\x8e\xdb\xf8\xde\x36\xca\xad\xcc\x88\xd0\xfb\xa3\x13\xa9\x4e\x5a\xa3\xba\x75\xf5\xaf\xb0\x29\xf3\x39\xab\x43\xb1\xd5\x75\x64\x20\x8b\x0c\x23\xdc\x68\x6c\xee\x84\xac\xd7\xe9\x39\x2a\x8b\x75\xc1\x51\x2e\xc2\x20\x68\xd4\xaa\xda\xf4\xfe\x64\x38\x1e\x9e\xfe\xe9\xdf\xcc\x77\x56\x0a\xa1\xea\x55\x5e\xa5\x89\xd6\x78\x08\x26\xd3\x5f\x41\x8d\x0b\xae\x9e\x20\xa1\x3f\x68\x08\x52\x9c\x83\xeb\x37\x62\xc4\xec\xec\xda\x78\x03\x3b\xf9\xde\x97\x3d\xad\x08\xca\xf5\xc3\x2b\xb2\xb4\x53\xeb\x80\xbc\xb5\x30\x4a\xc2\x86\x3b\x66\x30\x14\x39\x36\x87\xed\xc4\xc8\x92\xc9\xff\x2a\xa0\x7d\xa6\xe8\xcc\xca\x5e\x65\xfe\xc8\xd1\x40\xb1\xe8\x86\x8d\x7e\x49\xd9\xb4\xdd\x0d\x2d\x1b\x8e\xf4\x74\xc0\x3a\xec\x62\x46\x59\xde\x10\xa3\x91\x10\x30\xb0\xdd\x5e\x4c\xe6\x22\xe9\x3c\x17\xc4\x17\x3c\x90\x58\x56\x08\xe7\xec\xbd\xf4\x86\x96\xc3\x1f\xb6\x81\x9d\x74\x2d\x9f\x03\x29\xa2\x02\x17\x39\x9f\x3b\x87\x01\x6c\xc4\x3d\x27\xa3\xf1\x7e\xc1\x81\x99\x98\xb7\x12\xd0\x3e\x74\xdc\x94\xa6\xf2\x86\xdc\x21\xbc\xd3\x4d\x46\x5d\xa4\x8e\xf4\x22\x51\x87\x13\x6b\xcc\x79\x9a\xdd\xc9\x11\xd2\x46\x88\xbd\x6d\x19\x0e\xfd\xf8\xc3\xac\x96\xb5\x88\x9a\x67\x0e\xa8\x53\xbc\x6d\xc1\xcd\xe7\xce\x90\x36\xb2\xd1\xa0\xe2\xb7\xeb\xca\x53\x76\xb7\x47\xd0\xdb\x84\xbb\xf6\x82\xbc\x70\xf7\x72\x9d\x2e\x0d\x38\xf9\x33\x04\xb8\xd9\x08\xe9\x46\x23\x69\xb2\x70\x3e\x3a\xa3\x34\x7d\xdc\x6e\xfd\x90\x57\x03\x04\x12\x14\x60\x69\x80\x0c\xa2\x86\xe0\x15\x2a\x65\x80\x2c\xf5\x61\x37\x7a\x93\x18\x89\x61\xbb\x1d\x6e\x36\xa8\x64\x68\xbb\xdd\x6c\x10\xce\xa3\x63\xe0\x66\xd3\xcc\xb5\xdd\xc2\x20\x69\xe1\xe1\x9f\x7d\x56\x88\xf9\x84\xb6\x63\x64\xd2\xac\xca\x15\x00\xc2\x7e\xb6\x6c\x36\xe0\x5e\x98\xc4\xc0\xd0\xad\x97\x59\x85\x89\x82\x93\xaa\xee\x04\xdc\xf0\x87\x27\x61\x7f\xd8\xee\xbf\xe7\x14\x5d\xc4\x2a\x7c\x0d\xe2\x3e\x7d\x29\xee\xd8\x59\x46\x0b\x70\x36\x9d\x36\x92\x28\xec\x6a\x54\x68\x85\x16\x9e\x4d\xfe\xaa\x61\x11\xbe\xd7\xbf\x23\xb0\x67\xbf\xcd\xff\x77\x76\xf1\xf3\xe5\xfb\x77\xe6\x08\xe3\x69\x78\x9c\x61\xac\x47\x06\xe9\x20\x28\x61\xca\xfa\x63\x04\x46\xd2\xdd\x01\x08\xc3\x70\xca\xb0\xb7\x2b\x33\xa4\x41\x0f\xf4\x45\x40\xa3\xd7\xde\xeb\x49\x6f\x9a\x40\x0c\x3d\x0e\xc1\x2b\x35\x8d\xf0\x59\x57\x05\xbe\xfb\x90\x52\x16\x26\x51\xf2\xff\x0e\x3d\xb6\xf4\xe9\x91\x21\xca\xe2\xb3\xc3\xe9\xec\xfd\xe4\x62\x3e\xf7\x4d\x90\x9b\xce\x7a\x72\xf2\x81\x94\xf5\x3a\x90\xd7\x03\x67\x53\xae\x49\x8d\xb9\x88\x24\xf5\x80\xbe\x9d\x19\x5d\xb2\xf9\x23\xe3\x68\xdd\xb3\x2d\xa3\x5f\x08\xe3\xdb\x6d\xb2\xd9\x8c\x26\x04\xf3\xb4\xc0\x88\x06\x25\x43\xad\x5b\xd8\x80\x08\xb2\x48\x66\x7a\x7c\xaf\x08\x3d\xf6\x33\x5e\xc7\x0b\x1d\x0b\x63\x25\x39\x26\xcc\x5a\x84\xb0\x50\x72\xbc\x6b\x5b\x7a\xde\x80\xf6\x58\x5d\x52\xf4\x8e\xa8\xc0\x0d\xb8\xa0\x9e\x3d\x84\x17\x0f\x9c\xa6\x82\xc6\x5d\x7b\xe6\x08\xa2\x50\xac\x76\xe8\x75\x5a\x45\x36\x30\xbc\x5f\x62\x90\xe9\xe3\xb4\xc4\x86\xd8\x21\xdc\x5c\x75\x96\xe7\x14\x31\xd6\x80\x37\x32\x1d\xf2\x04\x4f\x12\xf4\x17\xf0\xad\x89\xfa\xc2\x5c\x7b\x3e\xde\x29\xa1\xdc\x28\x6b\xf7\xec\xc8\x48\x80\xf6\x2a\x8e\xc8\x14\x0e\xd0\xef\x60\xd4\x14\x58\x55\x89\xf8\x10\x1c\xbc\x42\x22\xc6\x7d\xa3\x93\xdc\xc3\xbd\x75\x21\x11\xca\x10\x73\xec\x7d\x0e\x46\xd0\x2a\x8c\x67\x33\x25\xd8\x6e\x85\x10\xc0\xb0\x0e\x80\xc6\xd6\xb6\x2a\x03\xb6\xdb\x63\xf1\xa0\x5d\x49\x78\xf7\x41\xbf\x5e\xf5\xa8\x3d\x74\xa8\x4b\x76\x4e\xff\xff\x40\x79\xa7\xb4\xb8\x2f\x4a\xb4\x44\x79\x67\xaa\xbb\x67\xc1\x98\xf1\x8a\x2c\x27\x04\x2f\x8a\x65\x4d\xdd\xd4\xd6\x86\x3a\xa7\xc5\xbd\x2a\x99\xa5\x5f\x58\x49\x96\x2c\xe4\x9e\xdf\x57\x6e\xef\x93\xf5\x5a\x8f\x3c\xa2\xaa\x26\x98\x44\x6b\x85\x21\x5e\xb6\xa3\x97\x4d\xbe\x18\x4a\x42\x7b\x47\x32\x4e\x51\xba\x3e\xaa\x28\x5a\x14\x0f\x62\xa0\x2e\x5a\x7a\x63\x5c\x56\x3f\xaf\x4e\xab\x6d\x42\x40\xfc\xed\x64\xaf\x6d\x1a\x53\x09\x85\x53\x50\x09\x05\xfc\x76\x7e\x68\x90\xa3\x33\x57\x4b\x9d\x5d\x21\x83\x66\x4a\xdc\xca\x78\x53\xab\x96\x93\x45\xb2\x8c\x90\xd0\xda\x59\x5c\x20\x01\x94\xb9\x61\xf7\x3c\x94\xb0\xea\x13\x4a\xe1\x3b\xf5\x99\xc5\x7e\xb5\xea\xae\x2f\xab\x2b\x9f\xea\x67\x4e\xa2\xd5\x75\x0b\xf5\x4b\x7b\xd3\xf5\xf3\x0b\x4a\x4b\xbe\x7a\x9c\xaa\xde\x9f\x4e\x9b\xbc\xae\x23\x5f\xa9\x9a\x56\xa7\xbe\xb1\xba\x19\xca\x36\x5a\x2e\xc5\xac\xa0\x28\x9f\x88\xc0\x28\x18\xc3\x47\xea\x65\x7b\xc5\xf0\x7b\x89\xc9\x15\x49\xf3\xe6\x65\xc8\xf7\x04\xec\x44\x6b\x20\xf7\xcb\x58\xcd\x11\xc2\x23\xe8\x11\x07\x32\x1b\x94\xc6\x4b\xba\xb4\xf1\xa1\x65\x79\x43\x68\x4c\x5a\xbb\x92\x41\xc7\x9c\xfd\xe5\x7d\x87\xee\x3b\x46\xb7\xe7\x10\xc0\x14\xff\x48\x75\x23\xa8\x4f\x7e\x59\xa8\x6f\x97\xfd\x1a\x8f\x41\xb0\x63\x92\xcc\xe9\x76\x95\x04\x83\x6d\xaa\x76\xd9\xb4\x65\xa5\x59\x13\x7b\xa5\xcb\x70\x92\xbc\xe4\xb5\xa6\x77\x34\x35\x9e\x1a\xc0\xcd\x2c\x53\x69\x91\x05\x7c\x45\x0b\xcc\x17\x00\x36\xb8\xff\x99\x41\x1b\xa7\x5b\x7e\x1b\x99\x11\x92\x51\x73\x93\x0d\xa8\x81\x39\x82\xb1\xc7\x44\x18\x98\x45\x91\x79\x3d\x31\xd1\xee\x57\x77\xa9\x3b\xd1\xca\x5c\xc0\x6b\xd9\x7a\xd6\x96\x84\xab\xd8\xe1\xed\x68\x9b\x97\x44\xb2\xb8\x37\xf3\x3a\x41\x6b\xc6\x3b\x3b\xf8\x14\x1e\x7e\x93\xf6\xb3\xe7\x50\x28\x23\xcc\xe7\x90\x26\xcc\xa5\x32\x49\xed\x64\xb3\x14\xe7\x64\xcd\xc0\x41\xc1\x49\xda\xcd\x72\xe8\xf9\xe9\xde\x85\x3c\x6b\xfb\xed\x2a\x7d\xac\x80\xad\x37\xf8\xda\xb5\x7b\xbb\xa5\xa3\xd5\xbd\x96\xc7\x0e\x6b\x1d\x3e\xf6\xc7\x2f\xce\xd8\xee\xe0\xc3\x38\x4b\x70\x4d\xa7\xd8\x37\xcb\x3e\x8b\x71\x00\x9e\xbf\x9b\xab\xb4\xfa\xb3\xdd\x07\xf2\x4d\xc4\xb9\xf9\xf3\x29\xa1\x5a\x04\xbb\x55\x76\xd7\xab\x76\xf3\x8b\xaf\x23\xe1\xae\x0b\xfc\x06\x84\x9b\x62\x33\x72\xdd\x2e\xe0\xb4\x46\x52\x1e\x47\xa6\xb1\x7e\x99\xbc\xbb\x67\x46\xdf\x40\xe2\x03\x02\x17\x6b\x2e\x7d\x21\x27\xdd\x98\xf7\x54\xc4\x74\xe6\x4c\x46\x6f\x72\x30\xee\x85\x12\xcc\x3e\x35\xf4\xf2\x72\xb0\xc7\xe9\xc8\x51\x43\xaa\x57\x7c\xb2\x1b\x6b\x2f\xf1\x52\x17\x5b\x9c\x44\xa3\x57\xe7\x34\x94\x1b\x32\xaa\x02\xde\xb4\xbe\x2d\x8b\xcc\x4f\x7e\xe1\xa4\xc8\xe9\xa5\xe0\x36\x1c\x8f\xe4\xff\x1d\x8f\x03\xa7\x17\x91\xc4\xbd\x1b\x6d\xf4\x8a\xe8\xa6\x44\x3f\x31\x8c\x25\xe0\xf0\xb2\x32\xfb\xcf\x76\x65\xf9\xf0\x2d\x25\x6b\x23\x76\xb5\x74\xda\x03\xbe\x21\x31\x50\x3b\xc1\xdc\x15\x24\x3a\x3b\x1b\x28\x11\x98\x69\xd6\x87\x2a\xbb\xcc\x5d\xb6\x78\x0d\x02\xc3\xa8\x2a\x84\x8e\xbb\x95\xf8\x96\x29\xe3\x45\xd6\x59\x81\x02\x2f\x93\xc4\x34\x0a\x9d\x38\x3f\xcf\x49\x58\x79\xee\x1e\x7a\xda\xad\x3b\xa6\x3f\x4a\x04\xd1\xef\x60\x34\xcf\x56\x68\x8d\x00\x2c\xba\xdb\x50\x56\x20\xae\xde\xab\x26\xa1\x50\x7b\x90\xd1\xe1\xad\xf4\xef\x72\xa1\xa8\x6c\xba\xab\xed\xed\x37\x7a\x34\xec\x26\x6c\x57\x36\x3d\x40\x3b\x27\xb1\x54\x35\xa8\x0c\x1d\xe5\x0e\x61\xed\xb5\x90\xa1\xb9\xa6\xb8\x34\x79\x07\xa6\xd1\x25\x5f\x86\xb0\xf9\xeb\x0c\xae\xcd\x5f\x91\x2d\xee\x42\x74\x30\x92\x3d\x6c\xe7\x34\x2d\x70\x81\x97\xaa\xb1\x4f\x91\xa1\x65\x09\x26\xd2\xf9\x0c\xad\xfe\x28\x21\x2f\xcd\x18\xfd\x58\xe5\x78\xc3\x10\x76\xb3\xc7\x06\xc0\xcb\xbc\x44\x0e\x2a\xe3\x91\x8f\x86\x12\xc6\xfe\x46\x30\x6a\x08\xe9\x5e\xa9\x62\xc2\x64\x85\xb2\x3b\xb7\x84\xa1\xeb\x0c\x37\x2b\x8a\xd8\x8a\x94\xb2\x6e\x77\x6a\x8b\x99\x64\xed\x7d\xda\x5a\x23\x35\xa4\x79\xea\x9a\x19\x78\x93\xd2\x65\xb8\x51\xcf\xab\xd5\x1a\xe8\x0c\x13\x97\x44\xe5\x36\xa6\xae\x4d\x20\xa2\x51\x11\xca\x63\xe5\x5c\x73\xc6\x94\xaf\x1c\xc3\xe7\x1f\xd0\x3b\xfc\x57\x23\x8d\x1d\xb0\x80\x7f\xc5\xab\x20\x37\xbb\x74\xd8\xd8\x93\xa6\x4f\xf9\x6b\xfa\x35\xcb\xf9\x2b\x76\x8e\xc2\xe7\x69\xb6\x83\x31\xe2\x29\xa7\x7b\x5a\x8e\xdf\xdf\x03\xda\xa8\x1d\x15\x95\x09\xb1\x17\xda\x3f\x33\xc1\x1b\x76\xad\xdb\xf3\xab\x60\x1f\x71\xd4\xbf\x9a\xee\x61\x6f\xc7\x19\x6a\x0d\xb7\x38\xe7\x02\x84\x39\xd7\xe1\x51\x13\x87\x4a\x2b\x4f\xcc\x27\x03\x87\x9e\xf3\xf9\x95\xc1\xab\xc6\xf5\x7e\xbb\xbd\xf0\xa4\x20\x6a\xd0\xfb\x40\x5f\x4a\x86\x7f\x3e\xe1\x96\xcd\xbf\x6e\x54\x13\xe9\x19\xdf\x53\x81\x7d\x85\x7d\x78\xec\xd3\xda\x40\xb1\xd3\x6e\x45\x57\x6e\xc8\xc2\x13\xec\xd1\x97\x83\x9a\x28\xca\x02\x37\x5e\x85\x3a\x1f\x38\xa7\xc5\x6d\xcd\xd5\x82\x23\x07\xa9\x0d\x31\xbb\xc8\x00\x56\x2a\x2a\xdc\x55\xe0\x0c\xc4\x3b\x5f\x72\xf4\x87\xe9\xce\xce\x97\x6b\x90\xd7\x16\x3f\x74\x37\xcb\x97\x95\x17\xcb\xcf\xd5\x9b\x09\x21\x77\x05\x9a\xf3\x22\xbb\x2b\x30\x62\xac\x8d\x2a\xc4\xaa\xec\xdd\x4d\x17\xb2\xbe\xfa\x08\x2d\xb6\x04\xcb\xce\x1b\xb0\x47\x5a\x1c\x4b\xb6\xf4\xad\xf8\xd6\x5a\x80\x4e\xb8\x43\x57\xea\xdb\x4e\xef\xf6\xc8\x70\x67\x84\xbc\xf5\xc7\x38\x00\x1d\xb7\xda\x8d\x31\x52\x84\x5d\x69\x7b\xa0\x77\xdc\xe8\xa9\x94\xbd\x40\x13\x4a\xf0\x5f\xc8\x2d\xf3\x7b\xa3\x45\x14\x85\x9d\x5b\x3b\xbb\xee\xec\x44\x13\xe5\x3d\xef\xeb\xec\x71\x03\xa4\xe7\xae\x8e\xd7\xab\xb7\xeb\x9e\xce\xd7\xb9\xa5\xf3\x84\x3b\x3a\x91\xd3\x4a\xd3\x92\xc6\xef\xe6\x44\xad\xac\x1d\xd6\xd9\xca\xa2\xf7\xd7\x3d\x17\xdb\x79\x1b\x67\xcf\xbb\x38\xbd\x37\xa7\xc2\x3d\x26\x7b\xdc\x9e\x32\x79\x0a\x51\xc6\x92\x59\x8d\x6f\x52\x76\x17\x06\xb5\x6f\xf6\x04\x41\xcc\x84\x37\x62\xae\xcf\x28\x6e\x0f\x17\xc2\x20\x40\xd1\x92\x99\x27\x9d\x3b\xa2\x7a\x6b\x70\x4a\x71\x92\x7e\x61\x89\x40\x12\xf1\x03\xc0\xb7\x9c\xbd\xe7\xf0\x1a\xf3\x13\xd0\x9d\x65\x19\xa9\x31\xbf\xcc\x77\x60\xd4\xab\x3c\xee\xc1\xdc\xb6\xfc\x4d\xae\x7e\x9d\xdf\x5c\xcc\x60\xb4\xf1\x03\x78\x77\x80\xba\x7f\xa1\xa7\xfe\xb3\x50\xb4\xf4\x72\xd9\x0a\x6f\x16\x2c\xc9\x92\x25\x13\x8a\x52\x8e\xda\x6e\x86\x88\xdf\xb6\x41\xe7\xb2\xa7\xa1\x17\x76\x5a\xf3\x2b\xb2\xbc\xb8\x47\x98\xb3\x10\xb3\x02\x8e\xde\x16\xf1\x48\x0f\x5c\x23\x5c\x72\x92\x9e\x9b\x5f\x43\xd5\x28\xd4\x27\x15\x00\x0a\x2c\xaa\xb5\x23\x39\x4e\xbf\xb0\xe6\x5e\xd7\xf7\xfe\x5d\x2e\xf5\xef\x1f\xb8\x3b\xff\x38\x96\x07\xce\x4a\x3a\x69\x31\xce\x98\x01\x4c\xe2\x8c\x73\xc3\xf4\xa8\xb7\x30\x22\x01\x3f\x0e\xd0\xce\xba\xbd\x5d\x18\x73\xd8\xd1\x6b\x88\x6e\x6d\xaf\x75\xfd\xbb\x6b\x78\xfe\x05\xbe\x95\x7e\x60\x78\xae\x9e\xeb\x79\xcd\x4c\xc1\x63\x79\xe3\x56\x9e\xfe\x94\xce\xbf\x8e\x7e\xb4\x6e\xe2\xb7\x45\x89\x3f\x8d\xad\x3a\x92\x77\x73\x12\xfe\xad\xa8\xde\x16\x65\x60\x3f\xe1\x27\xec\x97\x63\x06\x35\x43\x80\x71\x5a\x64\x7c\xf0\x93\xeb\x3d\xef\x53\x0a\xd2\x2f\x0c\xbc\x06\x14\xfd\x5e\x17\x14\x1d\x0c\xd2\x2f\xec\x88\xe5\x77\x83\xc3\x20\x30\xca\x04\x30\x46\x5f\xc4\xb0\xd1\xc5\x64\x7e\xb0\x59\xa7\x0f\x33\xc4\x69\x81\x58\x72\x32\xde\x86\x87\x09\xf1\x35\xc6\x4d\x4a\x52\xe7\xbf\xa5\x3c\x5b\x5d\x91\x25\x3b\x08\x8f\xd1\x86\x1b\xbc\x06\x83\x80\x7d\xf6\xd6\x12\x31\x27\x7a\x76\x29\xcd\x02\x95\x65\x32\xcc\x16\x2f\x00\x07\x3f\x05\x3b\x66\x7b\x10\xeb\x9b\xa4\x1e\x5e\xe3\xba\x47\x14\xad\x44\xc0\xad\x2e\x0b\xc1\xa2\x8d\xbb\x2c\xff\x08\xca\xbf\x0b\xb4\x83\x54\xf1\x6a\x60\x74\x95\x0f\x12\x87\xde\xae\x06\xd7\xd3\x00\x22\x96\x32\x0c\x73\x28\x58\x95\x51\xd3\x0e\x92\xc1\xc0\xdd\x5d\xaf\xb9\x09\x3d\x54\x22\xf5\x6b\x14\x0e\xbc\x06\x0b\xad\xd8\x07\x48\x58\xbb\x21\xc8\x08\xe6\xe8\x81\x1f\x7a\xfc\x91\xb3\x48\x71\x51\x37\x25\xc0\x6b\x20\x87\x8c\xf4\xef\x11\x45\x55\x99\x66\xe8\xe0\xf8\x5f\xfe\xe9\xe0\xd3\xa7\xfc\x87\xc3\x9f\x8e\x97\xc3\x0e\xff\x5a\x48\xe1\x10\xe4\x28\x8b\xe0\x06\x80\x22\x5e\x53\x0c\xd4\x69\xfe\x68\x41\xc9\x7a\xb2\x4a\xa9\xd0\xcc\x03\x31\xcc\x13\x5e\xf1\x9f\x80\x1e\x34\x84\xaa\x46\x8b\xc0\x56\xc3\xe6\x0f\xc6\x53\xca\x51\xfe\xe6\x31\x01\x03\x91\xf9\x0c\x86\x31\x48\x5b\x7e\x12\x57\x9e\x3e\x2a\x56\xe8\x96\x92\xcf\x51\x34\x5a\xd5\x92\xe6\x8f\x38\xa0\x70\xae\x09\x38\x89\x02\x90\x7b\x44\x69\x91\x23\x96\xc4\x97\xa7\x10\xe9\xd6\xab\xf7\xdd\x80\x8f\x7d\x03\x80\x14\x6f\x9c\xae\x51\x62\x2d\x6a\xd8\x6c\x7c\xf2\x11\x0c\xd8\x6a\x30\x04\x83\xa3\x6c\xd0\x3e\x15\xc2\xda\x87\xf6\x73\xec\x65\x70\xd4\x36\xba\xa9\xec\x0e\x7d\x01\xaf\xc1\x75\xca\x57\xa3\x45\x49\x08\x3d\x90\x7f\x52\xd9\xd0\x71\x70\xf8\xfd\xc9\x78\x3c\x1e\x87\x65\xa2\x24\xcb\x03\x6b\x49\xe0\x07\x30\x48\x06\xe0\x87\xd6\xbc\xfc\x00\x06\xc7\x42\x0e\xe4\x2c\xaf\xc5\x1b\x39\xdd\x0f\x60\xb0\x66\xcd\x42\xe5\x63\x4b\xf2\x0d\x21\x47\x94\x46\xa5\x5b\xee\x05\x23\x25\x1a\x09\x42\x06\x88\xd2\xd3\xc1\x10\x88\x11\x41\x6a\xc5\x3f\x86\xb8\x76\x57\x07\xed\x14\x87\x60\x23\x9c\xc3\x88\xaa\x04\xe7\x40\x49\x79\xab\xb8\xa3\x9c\x60\x74\x28\x8c\x88\x20\x7d\x6f\x9d\xf1\x19\xde\x4c\x28\xd9\xb6\x46\x8c\xa5\x4b\x34\x04\xd9\x6d\x8f\x65\x50\x0d\xba\xc2\x48\x0b\x26\x1e\x09\x46\x1d\x08\x4f\x74\x9e\x72\x74\x70\x78\x38\x5a\xaa\xe5\x04\xdc\x10\xd8\x5b\x65\x1b\x17\x23\xec\x67\xd2\xfe\x8a\xaa\x49\xd9\xc4\x7b\x0a\x9e\x85\x62\x3f\xc5\x93\x88\xc4\xb0\x51\x66\x07\x8e\x2d\xc3\x9f\xb9\xe9\xbb\xf6\x7c\x3f\x36\x68\xea\x54\x88\xba\x97\x46\xeb\x2d\x4c\x40\xbb\x97\x22\x4c\x62\x3c\x5d\x57\x49\x64\x9b\x76\x68\x74\x94\xe9\xe0\xe9\xfb\x04\x9e\xb2\x57\x20\xb6\x5f\xc0\xe1\xb7\x5e\x6a\x9c\xd9\x72\x83\x2b\x23\xda\xef\xd4\xe9\x36\xa6\x3b\xbe\x93\x35\x4b\x34\x56\xd0\x1e\x38\xa3\x77\x7a\x68\x74\xc8\xb1\xa3\xae\x26\x42\x8a\xf9\x8a\x50\xae\xa3\x86\x59\xdd\x53\x63\xd3\x32\x91\x48\xa0\x9d\xb1\xb8\xd1\xea\x3b\xba\x22\x78\xa9\x67\x38\x62\xd9\x0a\xe5\xb5\xfd\x71\x9d\xb9\x7e\x76\xf1\x50\x51\xc4\x9a\x5a\x8f\x24\x4e\xbf\x71\xba\x8e\xd4\x69\xa6\x57\xf2\x96\x61\x7b\x34\xb6\xef\x52\x8d\xc8\xa5\xe2\xcb\x3c\x40\xb0\x3e\x38\x75\xce\x5e\x2b\x7d\xfa\xf8\xa9\xb9\x5f\xfe\x09\x26\xe0\x53\xd3\x16\x22\x7d\xc0\x76\xfb\x09\x0e\xc1\x27\xa8\x8d\x79\x07\xa0\x6f\x84\x4a\x00\x63\x8f\x43\x55\xd5\xc0\x16\xa9\xc4\x69\x8a\xe8\xba\x60\x2c\x94\x61\x01\x37\xc5\x32\x60\x43\xbb\x06\xec\xea\x68\xd6\x36\x25\xaa\x54\x3b\xb9\xc4\xf7\xe4\x0e\x85\x3e\x17\xd3\x3c\x53\xcd\x43\xcf\xe4\xbb\x51\xf8\x14\x93\x4a\x07\xc8\x9c\x52\xa7\x29\x2a\x32\x13\x96\x68\xa2\x3d\x82\x9e\x44\x1b\x13\xf7\x29\x4e\xb8\x72\x1d\xfa\xda\x88\xfa\x48\xc1\x4c\xbf\x09\x7c\xf3\x30\x70\xe9\x7e\x66\x80\xe9\xbd\x0d\x5e\xb6\x0f\x6b\xd4\xcb\xef\xcd\x1b\x9f\x66\xf4\x5a\xe1\xbd\x46\xb0\xef\x1a\x26\x45\xbf\xab\xe0\x7f\xa0\xc4\x66\x8a\x3a\xdf\xd4\xeb\xf4\xbe\x0f\xe1\x95\x7a\xbe\xd3\xbb\xd1\xcb\xd4\x9e\xde\xd0\xc0\xb0\xa1\xb7\x64\x2d\x00\xd1\x35\x39\xdf\x88\xb1\x7a\xf5\xc3\x1f\xa2\x70\xcf\x1d\x22\x9b\xb7\xd7\x99\x43\xb4\x80\xed\x54\xd5\xa3\x55\x29\xb7\xe2\x6f\xbd\x76\x6f\x28\xf4\x17\xc0\xec\xe3\x08\x77\x1e\xe3\x70\xc2\xab\x23\x41\x11\x43\xda\xea\xfb\x9c\xca\x92\xa1\xf2\xe1\xf2\x7f\xa3\xc8\x9d\x49\x8b\x1d\x57\x84\x0f\x2b\x2c\xbf\x61\x1f\x54\x58\x57\x34\x7c\xb0\xf8\x07\xde\xbe\xfe\xb7\xdb\x7a\x2b\x90\x10\xa9\x46\xbb\x92\xa4\xf9\x6d\xdb\x68\xa7\x7a\x40\x6f\x51\xe4\xfc\x21\x32\x46\xa9\x33\xa2\xcd\x49\x2e\x7b\x4b\xc9\x3a\xd8\xb2\xb7\x1b\xdb\xcc\xc5\xf5\x5b\xc1\x57\x7b\xe0\xca\x4e\x77\x12\x9f\x9d\x26\x67\x35\x5f\x11\x5a\xfc\x81\x82\xed\xa8\xde\xa8\xd0\x71\xb5\x51\x50\x0d\xf2\xf5\xfb\x00\x1a\xe7\x89\x73\xfd\x27\x28\xc4\xcd\x5f\x9f\x77\x5b\x53\xf3\xeb\x53\xfe\x47\x9d\x6c\x9b\x33\xff\x31\x49\xf4\x77\xd7\xb4\xd1\x39\x47\x25\x12\x72\xd2\x9e\x54\xc3\x19\x12\xc9\xf9\x0e\xa3\x24\xbf\xe1\x3c\x21\x98\x53\xd5\x39\xe3\x36\x28\xc2\x9b\xd4\xb9\x5d\xbc\x69\x3e\x5f\x02\x99\xbc\x6a\x2f\x6c\x6c\xdb\x2f\xa0\x6f\x4d\x5a\x5d\x7b\x2d\x7c\x5a\x55\x26\x70\x8f\xeb\x09\xb1\xcd\xe0\xda\xff\x05\x00\x00\xff\xff\x19\x98\x42\x00\xca\x5f\x00\x00")

func templatesAppTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAppTmpl,
		"templates/app.tmpl",
	)
}

func templatesAppTmpl() (*asset, error) {
	bytes, err := templatesAppTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/app.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/app.tmpl": templatesAppTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"app.tmpl": &bintree{templatesAppTmpl, map[string]*bintree{}},
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

