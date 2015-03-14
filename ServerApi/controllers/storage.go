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
	"net/url"
	"io/ioutil"
	"crypto/md5"
	"encoding/hex"
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
		this.Abort("400")
		return
	}
	
	fo := NewFileOperation()
	
	switch action {
	case conf.ACTION_GET_SERVER_FILE_LIST:
		path := this.GetString(conf.KEY_PATH)
		if path == "" {
			beego.Error("[para is null] | path ")
			this.Abort("400")
			return
		}
		
		rawPath, err := url.QueryUnescape(path)
		if err != nil {
			beego.Error(err)
			this.Abort("400")
			return
		}
		
		res, err := fo.getServerFileList(rawPath)
		if err != nil {
			beego.Error(err)
			this.Abort("400")
			return
		}
		this.Data["json"] = res
		this.ServeJson()
	case conf.ACTION_MKDIR:
		dir := this.GetString(conf.KEY_DIR)
		if dir == "" {
			beego.Error("[para is null] | dir ")
			this.Abort("400")
			return
		}
		
		//rawDir, err := url.QueryUnescape(dir)
		//if err != nil {
		//	beego.Error(err)
		//	this.Abort("400")
		//	return
		//}
		
		res, err := fo.mkdir(dir)
		if err != nil {
			beego.Error(err)
			this.Abort("400")
			return
		}
		this.Data["json"] = res
		this.ServeJson()
		
	case conf.ACTION_RMFILE:
		dir := this.GetString(conf.KEY_DIR)
		if dir == "" {
			beego.Error("[para is null] | dir ")
			this.Abort("400")
			return
		}
		
		//rawdir, err := url.QueryUnescape(dir)
		//if err != nil {
		//	beego.Error(err)
		//	this.Abort("400")
		//	return
		//}
		
		res, err := fo.rmFile(dir)
		if err != nil {
			beego.Error(err)
			this.Abort("400")
			return
		}
		this.Data["json"] = res
		this.ServeJson()
	case conf.ACTION_RENAME:
		oldDir := this.GetString(conf.KEY_OLD_DIR)
		newDir := this.GetString(conf.KEY_NEW_DIR)
		if oldDir == ""  || newDir == "" {
			beego.Error("[para is null] | oldDir | newDir ")
			this.Abort("400")
			return
		}
		
		//rawoldDir, err := url.QueryUnescape(oldDir)
		//if err != nil {
		//	beego.Error(err)
		//	this.Abort("400")
		//	return
		//}
		//rawnewDir, err := url.QueryUnescape(newDir)
		//if err != nil {
		//	beego.Error(err)
		//	this.Abort("400")
		//	return
		//}
		
		res, err := fo.rename(oldDir, newDir)
		if err != nil {
			beego.Error(err)
			this.Abort("400")
			return
		}
		this.Data["json"] = res
		this.ServeJson()
	default:
		
	}
}

func (this *StorageController) Post() {
	action := this.GetString(conf.KEY_ACTION)
	if action == "" {
		beego.Error("[para is null] | action ")
		this.Abort("400")
		return
	}

	fo := NewFileOperation()
	
	switch action {
	case conf.ACTION_LOGIN:
		name := this.Input().Get("user_name")
		password := this.Input().Get("password")
		beego.Info(name)
		beego.Info(password)
		if name == "" || password == "" {
			beego.Error("[para is null] | user_name |  password ")
			this.Abort("400")
			return
		}
		
		res, err := fo.login(name, password)
		if err != nil {
			beego.Error(err)
			this.Abort("400")
			return
		}
		this.Data["json"] = res
		this.ServeJson()
	default:
		
	}
}

type FileOperation struct {
	userInfoData   *UserInfoData
}

func NewFileOperation() *FileOperation {
	return &FileOperation {
		userInfoData : new(UserInfoData),
	}
}

func (this *FileOperation)getUserInfo()  error {
	beego.Info(config.UserInfoDir)
	file, err := os.Open(config.UserInfoDir)
	if err != nil {
		beego.Error(err.Error())
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	err = dec.Decode(this.userInfoData)
	if err != nil {
		beego.Error(err.Error())
		return err
	}
	
	return  err
}

func (this *FileOperation) login(name string, password string) (*LoginData, error) {
	ld := NewLoginData()
	h := md5.New()
	err := this.getUserInfo()
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	
	ld.Status = 1
	ld.Token = ""
	
	if this.userInfoData.Name == name && this.userInfoData.Password == password {
		ld.Status = 0
		h.Write([]byte(config.Salt + name + password))
		ld.Token = hex.EncodeToString(h.Sum(nil))
	}
	
	//data, err := json.Marshal(ld)
	//if err != nil {
	//	beego.Error(err)
	//	return "", err
	//}
	
	return &ld, nil
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
	dir, err := ioutil.ReadDir(config.BaseDir + path)
	if err != nil {
		fs.Status = 1
		return nil, err
	}
	//PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		//if fi.IsDir() { 
		//	continue
		//}
		
		fi ,err := this.getFileStat(config.BaseDir + path + fi.Name())
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

func (this *FileOperation) getServerFileList(relativePath string) ([]FileStat, error) {
	path := config.BaseDir + relativePath
	beego.Info(path)
	fs, err := this.listDir(relativePath)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	
	//data, err := json.Marshal(fs)
	//if err != nil {
	//	beego.Error(err)
	//	return "", err
	//}
	//beego.Info(string(data))
	
	return fs, nil
}

func (this *FileOperation) mkdir(relativePath string) (*CommResData, error) {
	path := config.BaseDir + relativePath
	crd := NewCommResData()
	crd.Status = 1
	err := os.MkdirAll(path, 0777)
	if err != nil {
		crd.Status = 1
		beego.Error(err)
		return nil, err
	}
	
	crd.Status = 0
	
	//data, err := json.Marshal(crd)
	//if err != nil {
	//	crd.Status = 1
	//	beego.Error(err)
	//	return "", err
	//}
	
	//beego.Info(string(data))

	return &crd, nil
}

func (this *FileOperation) rmFile(relativePath string) (*CommResData, error) {
	path := config.BaseDir + relativePath
	crd := NewCommResData()
	crd.Status = 1
	err := os.RemoveAll(path)
	if err != nil {
		crd.Status = 1
		beego.Error(err)
		return nil, err
	}
	
	crd.Status = 0
	
	//data, err := json.Marshal(crd)
	//if err != nil {
	//	crd.Status = 1
	//	beego.Error(err)
	//	return "", err
	//}
	
	//beego.Info(string(data))

	return &crd, nil
}

func (this *FileOperation) rename(oldDir, newDir string) (*CommResData, error) {
	oldPath := config.BaseDir + oldDir
	newPath := config.BaseDir + "/" + newDir
	beego.Info(oldPath)
	beego.Info(newPath)
	crd := NewCommResData()
	crd.Status = 1
	err := os.Rename(oldPath, newPath)
	if err != nil {
		crd.Status = 1
		beego.Error(err)
		return nil, err
	}
	
	crd.Status = 0
	
	//data, err := json.Marshal(crd)
	//if err != nil {
	//	crd.Status = 1
	//	beego.Error(err)
	//	return "", err
	//}
	
	//beego.Info(string(data))

	return &crd, nil
}
