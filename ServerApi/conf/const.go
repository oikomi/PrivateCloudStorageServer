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

package conf

const (
	KEY_ACTION = "action"
	KEY_PATH = "path"
	KEY_TOKEN = "token"
	KEY_DIR = "dir"
	KEY_OLD_DIR = "old_dir"
	KEY_NEW_DIR = "new_dir"
	KEY_FROM_PATH = "from_path"
	KEY_TO_PATH = "to_path"
)

const (
	ACTION_LOGIN = "login"
	ACTION_GET_SERVER_FILE_LIST = "get_server_file_list"
	ACTION_MKDIR = "mkdir"
	ACTION_RMFILE = "rm_file"
	ACTION_RENAME = "rename"
	ACTION_MOVEFILE = "move_file"
	ACTION_UPLOAD = "upload"
	ACTION_GET_STORAGE_INFO = "get_storage_info"
	ACTION_IS_FILE_EXIST = "is_file_exist"
)