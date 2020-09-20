/*
 Navicat Premium Data Transfer

 Source Server         : book
 Source Server Type    : MySQL
 Source Server Version : 80018
 Source Host           : 47.95.239.29:3306
 Source Schema         : novel

 Target Server Type    : MySQL
 Target Server Version : 80018
 File Encoding         : 65001

 Date: 19/09/2020 17:33:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_account
-- ----------------------------
DROP TABLE IF EXISTS `t_account`;
CREATE TABLE `t_account`  (
  `id` bigint(20) NOT NULL,
  `uid` int(11) NOT NULL,
  `account_balance` bigint(11) NOT NULL,
  `monthly_balance` bigint(11) NOT NULL,
  `recommend_balance` bigint(11) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_admin
-- ----------------------------
DROP TABLE IF EXISTS `t_admin`;
CREATE TABLE `t_admin`  (
  `id` bigint(20) NOT NULL,
  `nickname` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `username` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `roles` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_banner
-- ----------------------------
DROP TABLE IF EXISTS `t_banner`;
CREATE TABLE `t_banner`  (
  `id` bigint(20) NOT NULL,
  `title` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `enable` tinyint(1) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NULL DEFAULT NULL,
  `update_time` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_book
-- ----------------------------
DROP TABLE IF EXISTS `t_book`;
CREATE TABLE `t_book`  (
  `id` bigint(20) NOT NULL,
  `type_id` bigint(20) NOT NULL,
  `title` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `synopsis` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `cover` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `recommend` int(11) NULL DEFAULT NULL,
  `click` int(11) NULL DEFAULT NULL,
  `collection` int(11) NULL DEFAULT NULL,
  `instalments` tinyint(1) NOT NULL,
  `word_num` bigint(20) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  `third_id` bigint(20) NOT NULL,
  `platform_id` int(11) NOT NULL,
  `is_serial` tinyint(1) NOT NULL COMMENT '1完结，0连载中',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `third_id`(`third_id`, `platform_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_book_author
-- ----------------------------
DROP TABLE IF EXISTS `t_book_author`;
CREATE TABLE `t_book_author`  (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `book_id` bigint(20) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`, `book_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_chapter
-- ----------------------------
DROP TABLE IF EXISTS `t_chapter`;
CREATE TABLE `t_chapter`  (
  `id` bigint(20) NOT NULL,
  `book_id` bigint(20) NOT NULL,
  `sorted` int(11) NOT NULL,
  `name` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  `third_id` bigint(20) NOT NULL,
  `platform_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `third_id`(`third_id`, `platform_id`) USING BTREE,
  UNIQUE INDEX `book_id`(`book_id`, `sorted`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_circle
-- ----------------------------
DROP TABLE IF EXISTS `t_circle`;
CREATE TABLE `t_circle`  (
  `id` bigint(20) NOT NULL,
  `uid` int(11) NOT NULL,
  `book_id` bigint(20) NOT NULL,
  `content` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` tinyint(1) NULL DEFAULT 1,
  `create_time` int(11) NULL DEFAULT NULL,
  `update_time` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_circle_comment
-- ----------------------------
DROP TABLE IF EXISTS `t_circle_comment`;
CREATE TABLE `t_circle_comment`  (
  `id` bigint(20) NOT NULL,
  `circle_id` bigint(20) NOT NULL,
  `uid` int(11) NOT NULL,
  `mast` bigint(20) NOT NULL,
  `reply_id` bigint(20) NOT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_circle_like
-- ----------------------------
DROP TABLE IF EXISTS `t_circle_like`;
CREATE TABLE `t_circle_like`  (
  `id` bigint(20) NOT NULL,
  `uid` int(11) NOT NULL,
  `type` tinyint(1) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  `circle_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_collection
-- ----------------------------
DROP TABLE IF EXISTS `t_collection`;
CREATE TABLE `t_collection`  (
  `id` bigint(20) NOT NULL,
  `uid` int(11) NOT NULL,
  `type` tinyint(1) NOT NULL,
  `book_id` bigint(20) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_comment
-- ----------------------------
DROP TABLE IF EXISTS `t_comment`;
CREATE TABLE `t_comment`  (
  `id` bigint(20) NOT NULL,
  `uid` int(11) NOT NULL,
  `content` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `book_id` bigint(20) NOT NULL,
  `reply_id` bigint(20) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_follow
-- ----------------------------
DROP TABLE IF EXISTS `t_follow`;
CREATE TABLE `t_follow`  (
  `id` bigint(20) NOT NULL,
  `follow_id` int(11) NOT NULL,
  `fans_id` int(11) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_label
-- ----------------------------
DROP TABLE IF EXISTS `t_label`;
CREATE TABLE `t_label`  (
  `id` bigint(20) NOT NULL,
  `name` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_message
-- ----------------------------
DROP TABLE IF EXISTS `t_message`;
CREATE TABLE `t_message`  (
  `id` bigint(20) NOT NULL,
  `cmd` tinyint(4) NOT NULL,
  `len` int(11) NOT NULL,
  `form_id` int(11) NOT NULL,
  `to_id` int(11) NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_message_read
-- ----------------------------
DROP TABLE IF EXISTS `t_message_read`;
CREATE TABLE `t_message_read`  (
  `id` bigint(20) NOT NULL,
  `uid` int(11) NOT NULL,
  `msg_id` bigint(20) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_recharge
-- ----------------------------
DROP TABLE IF EXISTS `t_recharge`;
CREATE TABLE `t_recharge`  (
  `id` bigint(20) NOT NULL,
  `uid` int(11) NOT NULL,
  `amount` tinyint(4) NOT NULL,
  `type` bigint(11) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_record
-- ----------------------------
DROP TABLE IF EXISTS `t_record`;
CREATE TABLE `t_record`  (
  `id` bigint(20) NOT NULL,
  `uid` int(11) NOT NULL,
  `book_id` bigint(20) NOT NULL,
  `chapter_id` bigint(20) NOT NULL,
  `page` int(11) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_type
-- ----------------------------
DROP TABLE IF EXISTS `t_type`;
CREATE TABLE `t_type`  (
  `id` bigint(20) NOT NULL,
  `pid` bigint(20) NOT NULL,
  `pic` varchar(155) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `type_name_index`(`name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_user_details
-- ----------------------------
DROP TABLE IF EXISTS `t_user_details`;
CREATE TABLE `t_user_details`  (
  `id` bigint(20) NOT NULL,
  `uid` int(11) NOT NULL,
  `head_img` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `phone` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `login_type` tinyint(4) NULL DEFAULT NULL,
  `birthday` int(11) NULL DEFAULT NULL,
  `region` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `nation` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `autograph` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_user_info
-- ----------------------------
DROP TABLE IF EXISTS `t_user_info`;
CREATE TABLE `t_user_info`  (
  `id` bigint(20) NOT NULL,
  `uid` int(11) NOT NULL,
  `username` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `nickname` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` char(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `type` tinyint(4) NOT NULL COMMENT '1管理员，2普通用户，3作家',
  `public_key` char(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uid`(`uid`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  INDEX `password`(`password`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
