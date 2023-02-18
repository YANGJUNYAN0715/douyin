// Copyright 2022 CloudWeGo Authors
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
//

package consts

const (
	UserTableName     = "user"
	RelationTableName = "relations"
	//数据库表名
	VideoTableName = "video"

	SecretKey           = "secret key"
	IdentityKey         = "id"
	Total               = "total"
	Notes               = "notes"
	ApiServiceName      = "api"
	UserServiceName     = "user"
	RelationServiceName = "relation"
	FavorteServiceName  = "favorite"

	MySQLDefaultDSN     = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP                 = "tcp"
	UserServiceAddr     = ":9000"
	RelationServiceAddr = ":8087"

	FavoriteServiceAddr = ":8090"

	NoteServiceAddr = ":10000"
	ExportEndpoint  = ":4317"
	ETCDAddress     = "127.0.0.1:2379"
	DefaultLimit    = 10

	StatusOK = 200

	//favorite actiontype,1是点赞，2是取消点赞
	Like   = 1
	Unlike = 2

	OssEndPoint        = "oss-cn-shenzhen.aliyuncs.com" //Oss
	OssAccessKeyId     = "oss"
	OssAccessKeySecret = "oss"
	OssBucket          = "dousheng1"
)
