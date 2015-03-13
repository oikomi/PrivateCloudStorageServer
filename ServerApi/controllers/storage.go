//
// Copyright 2015-2099 Hong Miao. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/oikomi/PrivateCloudStorageServer/ServerApi/conf"
)

type StorageController struct {
	beego.Controller
}

func (this *StorageController) Get() {
	//var err error
	action := this.GetString(conf.KEY_ACTION)
	if action == "" {
		beego.Error("[para is null] | action ")
		return
	}
	
	fo := NewFileOperation()
	
	switch action {
	case conf.ACTION_GET_SERVER_FILE_LIST:
		path := this.GetString(conf.KEY_PATH)
		if path == "" {
			beego.Error("[para is null] | path ")
			return
		}
		res, err := fo.getServerFileList(path)
		if err != nil {
			return
		}
		this.Data["json"] = res
		this.ServeJson()
	default:
		
	}
}

func (this *StorageController) Post() {
	
}

type FileOperation struct {
	
}

func NewFileOperation() *FileOperation {
	return &FileOperation {
	
	}
}

func (this *FileOperation)getFileStat(path string) (os.FileInfo, error) {
	fileInfo ,err := os.Stat(path)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	
	return fileInfo ,err
}

func (this *FileOperation)listDir(path string) ([]FileStat, error) {
	fs := NewFileStat()
	fss := make([]FileStat, 0, 10)
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fs.Status = 1
		return nil, err
	}
	//PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		//if fi.IsDir() { 
		//	continue
		//}
		
		fi ,err := this.getFileStat(path + fi.Name())
		if err != nil {
			fs.Status = 1
			return nil, err
		}
		
		fs.Name = fi.Name()
		fs.Path = path + fi.Name()
		fs.Size = fi.Size()
		fs.Modify = fi.ModTime().Unix()
		if fi.IsDir() {
			fs.Type = "dir"
		} else {
			fs.Type = "file"
		}
		
		fss = append(fss, fs)
	}
	return fss, nil
}

func (this *FileOperation) getServerFileList(relativePath string) (string, error) {
	path := config.BaseDir + relativePath
	beego.Info(path)
	fs, err := this.listDir(path)
	if err != nil {
		beego.Error(err)
		return "", err
	}
	
	data, err := json.Marshal(fs)
	if err != nil {
		beego.Error(err)
		return "", err
	}
	beego.Info(string(data))
	
	return string(data), nil
}