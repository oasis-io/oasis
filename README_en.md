# Oasis

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/carina-io/carina/blob/main/LICENSE)

English | [简体中文](README.md)

### Introduction

  Oasis provides full lifecycle management of MySQL databases. Deploy database instances based on Kubernetes and provide maintenance for the database. It provides the following functions:

* Kubernetes deploy MySQL
* MySQL HA Management
* SQL Audit and SQL Query
* Data migration
* Database instance management

### build

```bash
git clone https://github.com/oasis-io/oasis.git
cd oasis
bash scripts/build.sh
```

### Install

```bash
CREATE DATABASE IF NOT EXISTS oasis default character set utf8mb4;

CREATE USER  'oasis'@'%' IDENTIFIED BY 'eQWJjZGV1A(MjAxOQo';
GRANT ALL PRIVILEGES ON oasis.* TO 'oasis'@'%' ; 
FLUSH PRIVILEGES;

Modify the account and password of [mysql] in the oasis.toml configuration file


nohup ./oasis -c oasis.toml &
```

## Community

Email：oasis_2022@126.com