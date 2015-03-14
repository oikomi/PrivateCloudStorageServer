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
)

type FileStat struct {
	Status     uint32 `json:"status"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Size       int64  `json:"size"`
	Modify     int64  `json:"modify"`
	Type       string `json:"type"`
}

func NewFileStat() FileStat {
	return FileStat{}
}


type LoginData struct {
	Status     uint32 `json:"status"`
	Token      string `json:"token"`
}

func NewLoginData() LoginData {
	return LoginData{}
}


type UserInfoData struct {
	Name      string `json:"user_name"`
	Password  string `json:"password"`
}

func NewUserInfoData() UserInfoData {
	return UserInfoData{}
}

type CommResData struct {
	Status     uint32 `json:"status"`
}

func NewCommResData() CommResData {
	return CommResData{}
}