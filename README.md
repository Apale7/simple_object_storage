# Simple Object Storage
简易的对象存储系统

- 支持上传文件, 不会重复存储
- 支持下载文件
- 支持分享文件, 可以设置下载密码
- 只支持单机，不支持集群
- 需要本地有redis

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