/*
 Navicat Premium Data Transfer

 Source Server         : 本地存储
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : shopping

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 25/04/2025 21:03:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for base_information
-- ----------------------------
DROP TABLE IF EXISTS `base_information`;
CREATE TABLE `base_information`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `home_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '首页标题',
  `home_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '首页描述',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号码',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱地址',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '地址',
  `wechat_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '微信图片地址',
  `is_deleted` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否删除 0：删除 1：正常',
  `is_enable` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否启用 0：未启用 1：启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of base_information
-- ----------------------------
INSERT INTO `base_information` VALUES (1, '測試用戶', NULL, '1234567890', '123@test.com', '測試地址', 'https://i.ebayimg.com/images/g/EcIAAOSwropmIhRN/s-l1200.jpg', 0, 0);

-- ----------------------------
-- Table structure for carousel
-- ----------------------------
DROP TABLE IF EXISTS `carousel`;
CREATE TABLE `carousel`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `order` int NOT NULL DEFAULT 0 COMMENT '排序',
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图片URL地址',
  `redirect_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图片URL地址',
  `is_enabled` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否启用 0：未启用 1：启用',
  `is_deleted` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否删除 0：删除 1：正常',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of carousel
-- ----------------------------
INSERT INTO `carousel` VALUES (1, '測試', 1, 'https://i.ebayimg.com/images/g/EcIAAOSwropmIhRN/s-l1200.jpg', 'https://i.ebayimg.com/images/g/EcIAAOSwropmIhRN/s-l1200.jpg', 0, 1, '2025-04-25 00:51:09', '2025-04-25 00:51:13');

-- ----------------------------
-- Table structure for certificates
-- ----------------------------
DROP TABLE IF EXISTS `certificates`;
CREATE TABLE `certificates`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '证书名称',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '证书描述',
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '证书图片',
  `is_deleted` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否删除 0：删除 1：正常',
  `is_available` tinyint(1) NOT NULL DEFAULT 1 COMMENT '上下架状态 0:下架 1:上架',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of certificates
-- ----------------------------
INSERT INTO `certificates` VALUES (1, '证书测试', '11', 'http://127.0.0.1:8088/v1/getLatestImage?fileName=shopping_1745585749758469.jfif', 1, 1, '2025-04-25 18:16:18', '2025-04-25 20:55:51');
INSERT INTO `certificates` VALUES (2, '11', '333', 'http://127.0.0.1:8088/v1/getLatestImage?fileName=shopping_1745578806650419.jfif', 1, 1, '2025-04-25 19:05:33', '2025-04-25 20:50:31');

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品名称',
  `cover_image` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品封面图',
  `detail_images` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商品详情图片',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品描述',
  `sale_price` float(10, 2) NOT NULL DEFAULT 0.00 COMMENT '销售价格',
  `cost_price` float(10, 2) NOT NULL DEFAULT 0.00 COMMENT '成本价格',
  `stock_quantity` int NOT NULL DEFAULT 0 COMMENT '库存数量',
  `brand` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商品品牌',
  `category_id` int NOT NULL DEFAULT 0 COMMENT '商品分类ID',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `is_available` tinyint(1) NOT NULL DEFAULT 1 COMMENT '上下架状态 0:下架 1:上架',
  `is_deleted` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否删除 0：删除 1：正常',
  `type` int NOT NULL DEFAULT 0 COMMENT '默认类型：0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product
-- ----------------------------
INSERT INTO `product` VALUES (1, 'ipad 10', 'https://i.ebayimg.com/images/g/EcIAAOSwropmIhRN/s-l1200.jpg', NULL, 'ipad 10 商品', 1000.00, 500.00, 0, 'Apple', 0, '2025-04-25 01:34:05', '2025-04-25 01:34:08', 1, 1, 0);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `role` int NOT NULL DEFAULT 1 COMMENT '角色 1：普通用户',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ones_username`(`username`) USING BTREE COMMENT '用户名唯一索引',
  UNIQUE INDEX `ones_email`(`email`) USING BTREE COMMENT '邮箱号唯一索引'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, NULL, 'fondy', '$2a$10$jX7.U2p083a3f5t7yL/LFuk1ez8TveoSENiv0XN4cS6LKeVqdccc2', '2025-04-24 15:38:14', '2025-04-24 15:38:14', 1, NULL);
INSERT INTO `user` VALUES (2, NULL, 'ss', '123', '2025-04-24 22:52:07', '2025-04-24 22:52:43', 1, NULL);

SET FOREIGN_KEY_CHECKS = 1;
