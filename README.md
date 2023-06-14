# Oasis

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/carina-io/carina/blob/main/LICENSE)

简体中文 | [English](README_en.md)

### 介绍
  Oasis 提供对MySQL 数据库全生命周期的管理。基于Kubernetes 来部署数据库实例，并提供对数据库的维护。它提供以下功能：

* Kubernetes 部署MySQL
* MySQL 高可用管理
* SQL 审核和查询
* 数据迁移
* 数据库实例管理


### 编译

```bash
git clone https://github.com/oasis-io/oasis.git
cd oasis
bash scripts/build.sh
```


### 部署

```bash
创建数据库
CREATE DATABASE IF NOT EXISTS oasis default character set utf8mb4;


创建用户与权限
CREATE USER  'oasis'@'%' IDENTIFIED BY 'eQWJjZGV1A(MjAxOQo';
GRANT ALL PRIVILEGES ON oasis.* TO 'oasis'@'%' ; 
FLUSH PRIVILEGES;


修改配置文件的帐号与密码
cat oasis.toml
[mysql]
user = "oasis"
host = "127.0.0.1" 
port = "3306"
password = "eQWJjZGV1A(MjAxOQo"
database = "oasis"


启动
nohup ./oasis -c oasis.toml &
```

## 技术交流

Email：oasis_2022@126.com





