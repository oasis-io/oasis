# oasis

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/carina-io/carina/blob/main/LICENSE)

简体中文 | [English](./README_en.md)

### 介绍
    封装Operator, 基于Kubernetes来部署和维护数据库实例。它提供一下功能：

* Kubernetes deploy MySQL
* SQL 查询
* SQL 审核
* 数据迁移
* 数据库实例管理


### 编译

bash scripts/build.sh


### 部署

```bash
创建数据库
CREATE DATABASE IF NOT EXISTS oasis default character set utf8mb4;

创建用户与权限
CREATE USER  'oasis'@'%' IDENTIFIED BY 'eQWJjZGV1A(MjAxOQo';
GRANT ALL PRIVILEGES ON oasis.* TO 'oasis'@'%' ; 
FLUSH PRIVILEGES;

启动
nohup ./oasis -c oasis.yaml &
```

