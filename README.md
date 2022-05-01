# Simple Object Storage
简易的对象存储系统

- 支持上传文件, 不会重复存储
- 支持分享文件, 可以设置下载密码
- 只支持单机，不支持集群
- 需要本地有redis

## 上传
POST /api/upload
- body为form-data
- 文件的key为file

## 查看文件列表
GET /api/list
- username: 文件的所有者, 精确匹配
- filename: 文件名, 模糊匹配
- is_public:
    - true: 从公开的资源中查询
    - false: 从username所属的资源中查询

## 分享文件
POST /api/share
- file_id: 文件的id
- need_pwd: 下载是否需要面
- password: 需要密码时可以自己指定密码, 不传则随机生成
- duration: 分享的有效时间, 单位为秒

## 下载文件
GET /api/download
- share_id: 分享id
- password: 分享密码

## 配置文件
conf/conf.yml
```yml
port: "6789"
file_root: ./file_root/ # Root directory for the storage, must be end of a '/'

mysql_host: localhost
mysql_port: 3306#数字,非字符串
mysql_user: root
mysql_password: rootroot
mysql_db: object_storage
```