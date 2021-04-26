
为 Promethues 提供基于 mysql 的 规则管理 操作

# 使用
## 环境准备

golang 1.13
源码下载

## DB初始化

用 rules.sql 初始化数据库

## 配置

config/config.toml 
- 配置端口
- 配置mysql

## 运行

go run main.go

# 说明
## 加载外部规则内容
参看 servies/load_test.go 中的示例
- TestParseFromByte：直接从字节流中解析
- TestParseFromHttpFile：从网络规则文件中解析