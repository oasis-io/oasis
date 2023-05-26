# oasis

简体中文 | [English](./README_en.md)


### 编译
bash scripts/build.sh


### 部署
```bash
数据库
CREATE DATABASE IF NOT EXISTS oasis default character set utf8mb4;

用户与权限
CREATE USER  'oasis'@'%' IDENTIFIED BY 'eQWJjZGV1A(MjAxOQo';
GRANT ALL PRIVILEGES ON oasis.* TO 'oasis'@'%' ; 
FLUSH PRIVILEGES;

启动
nohup ./oasis -c config.yaml &
```

