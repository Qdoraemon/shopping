# shopping

购物平台 api

## 项目介绍

本项目是一个购物平台 api，提供商品浏览、购物车、订单等功能。

## 功能介绍

- 商品浏览：用户可以浏览商品列表，包括商品名称、价格、图片等信息。
- 购物车：用户可以将商品加入购物车，并可以修改商品数量、删除商品等操作。
- 订单：用户可以查看订单列表，包括订单状态、商品信息、总价等信息。

## 技术栈

- 后端：golang , gin, gorm, mysql, redis
- 前端：Vue.js、Element UI

## 启动命令

```bash
# 进入项目目录
cd shopping

# 安装依赖
go mod tidy

# 启动项目
go run cmd/main.go
```
