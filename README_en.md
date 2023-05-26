# oasis

English | [简体中文](./README.md)


## build
bash scripts/build.sh


## Install

CREATE DATABASE IF NOT EXISTS oasis default character set utf8mb4;

CREATE USER  'oasis'@'%' IDENTIFIED BY 'eQWJjZGV1A(MjAxOQo';
GRANT ALL PRIVILEGES ON oasis.* TO 'oasis'@'%' ; 
FLUSH PRIVILEGES;

Start
nohup ./oasis -c config.yaml &


