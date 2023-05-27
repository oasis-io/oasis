# Oasis

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/carina-io/carina/blob/main/LICENSE)

简体中文 | [English](README_en.md)

### 介绍

  基于Kubernetes 来部署和维护数据库实例。它提供以下功能：

* Kubernetes 部署MySQL
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

启动
nohup ./oasis -c oasis.toml &
```

## 技术交流

Email：oasis_2022@126.com





