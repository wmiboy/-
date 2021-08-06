/*
 Navicat Premium Data Transfer

 Source Server         : 99v66
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Host           : localhost:3306
 Source Schema         : ruan_jian_guan_li

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 07/08/2021 04:48:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for s_apply
-- ----------------------------
DROP TABLE IF EXISTS `s_apply`;
CREATE TABLE `s_apply`  (
  `aid` int(0) NOT NULL AUTO_INCREMENT,
  `i_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '应用名称',
  `i_mold` tinyint(0) NOT NULL COMMENT '收费类型 1扣时 2扣点',
  `i_login` tinyint(0) NOT NULL COMMENT '登录限制 1免登陆 2单点登录 3多点登录',
  `i_bin` tinyint(0) NOT NULL COMMENT '设备限制 1不绑定 2绑定机器码 3绑定IP',
  `i_msg` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '远程数据',
  `i_state` tinyint(0) NOT NULL COMMENT '应用状态 1=可用 2=禁止使用',
  `i_time` int(0) NOT NULL,
  PRIMARY KEY (`aid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10001 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for s_key
-- ----------------------------
DROP TABLE IF EXISTS `s_key`;
CREATE TABLE `s_key`  (
  `kid` int(0) NOT NULL AUTO_INCREMENT,
  `i_cdk` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `i_mold` tinyint(0) NOT NULL COMMENT '卡密类型 1扣时 2扣点',
  `i_point` int(0) NOT NULL COMMENT '点数',
  `i_day` int(0) NOT NULL COMMENT '小时数',
  `i_aid` int(0) NOT NULL COMMENT '对应软件',
  `i_uid` int(0) NOT NULL COMMENT '使用者UID',
  `i_cread_time` int(0) NOT NULL COMMENT '创建时间',
  `i_use_time` int(0) NOT NULL COMMENT '使用时间',
  `i_state` tinyint(0) NOT NULL COMMENT '1=可用 2=不可用',
  PRIMARY KEY (`kid`) USING BTREE,
  UNIQUE INDEX `unx_cdk`(`i_cdk`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10020 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for s_power
-- ----------------------------
DROP TABLE IF EXISTS `s_power`;
CREATE TABLE `s_power`  (
  `pid` int(0) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `i_uid` int(0) NOT NULL COMMENT '账户ID',
  `i_aid` int(0) NOT NULL COMMENT '应用ID',
  `i_kid` int(0) NOT NULL COMMENT '卡密ID',
  `i_day` int(0) NOT NULL COMMENT '到期时间',
  `i_point` int(0) NOT NULL COMMENT '剩余点数',
  `i_ip` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '首次登陆IP',
  `i_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '首次登陆机器码',
  `i_state` tinyint(0) NOT NULL COMMENT '1=可用 2=不可用',
  `i_time` int(0) NOT NULL COMMENT '登陆时间',
  PRIMARY KEY (`pid`) USING BTREE,
  INDEX `inx_uid`(`i_uid`, `i_aid`, `i_kid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for s_user
-- ----------------------------
DROP TABLE IF EXISTS `s_user`;
CREATE TABLE `s_user`  (
  `uid` int(0) NOT NULL AUTO_INCREMENT,
  `i_user` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `i_pwd` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `i_msg` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户信息',
  `i_type` tinyint(0) NOT NULL COMMENT '1管理员  1普通用户 ',
  `i_state` tinyint(0) NOT NULL COMMENT '账户状态 1=可用 2=不可用',
  `i_time` int(0) NOT NULL,
  `i_uptime` int(0) NOT NULL COMMENT '改密时间,强退时间 用于作废之前的token',
  PRIMARY KEY (`uid`) USING BTREE,
  UNIQUE INDEX `unx_user`(`i_user`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10025 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
