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

 Date: 24/04/2025 23:39:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '用户名',
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
