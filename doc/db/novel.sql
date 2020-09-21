/*
 Navicat Premium Data Transfer

 Source Server         : my_node1
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : 119.29.117.244:3306
 Source Schema         : novel

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 21/09/2020 10:15:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_account
-- ----------------------------
DROP TABLE IF EXISTS `t_account`;
CREATE TABLE `t_account`  (
  `id` bigint(0) NOT NULL,
  `uid` int(0) NOT NULL,
  `account_balance` bigint(0) NOT NULL,
  `monthly_balance` bigint(0) NOT NULL,
  `recommend_balance` bigint(0) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_auth
-- ----------------------------
DROP TABLE IF EXISTS `t_auth`;
CREATE TABLE `t_auth`  (
  `id` bigint(0) NOT NULL,
  `role_id` bigint(0) NOT NULL,
  `user_id` bigint(0) NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_banner
-- ----------------------------
DROP TABLE IF EXISTS `t_banner`;
CREATE TABLE `t_banner`  (
  `id` bigint(0) NOT NULL,
  `title` varchar(55) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `link` varchar(255) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `img` varchar(255) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `enable` tinyint(1) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NULL DEFAULT NULL,
  `update_time` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_book
-- ----------------------------
DROP TABLE IF EXISTS `t_book`;
CREATE TABLE `t_book`  (
  `id` bigint(0) NOT NULL,
  `type_id` bigint(0) NOT NULL,
  `title` varchar(55) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `synopsis` varchar(4096) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `cover` varchar(1024) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `recommend` int(0) NULL DEFAULT NULL,
  `click` int(0) NULL DEFAULT NULL,
  `collection` int(0) NULL DEFAULT NULL,
  `instalments` tinyint(1) NOT NULL,
  `word_num` bigint(0) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  `third_id` bigint(0) NOT NULL,
  `platform_id` int(0) NOT NULL,
  `is_serial` tinyint(1) NOT NULL COMMENT '1完结，0连载中',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `third_id`(`third_id`, `platform_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_book_author
-- ----------------------------
DROP TABLE IF EXISTS `t_book_author`;
CREATE TABLE `t_book_author`  (
  `id` bigint(0) NOT NULL,
  `user_id` bigint(0) NOT NULL,
  `book_id` bigint(0) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`, `book_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_chapter
-- ----------------------------
DROP TABLE IF EXISTS `t_chapter`;
CREATE TABLE `t_chapter`  (
  `id` bigint(0) NOT NULL,
  `book_id` bigint(0) NOT NULL,
  `sorted` int(0) NOT NULL,
  `name` varchar(55) CHARACTER SET utf8mb4 NOT NULL,
  `content` mediumtext CHARACTER SET utf8mb4 NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  `third_id` bigint(0) NOT NULL,
  `platform_id` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `third_id`(`third_id`, `platform_id`) USING BTREE,
  UNIQUE INDEX `book_id`(`book_id`, `sorted`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_circle
-- ----------------------------
DROP TABLE IF EXISTS `t_circle`;
CREATE TABLE `t_circle`  (
  `id` bigint(0) NOT NULL,
  `uid` int(0) NOT NULL,
  `book_id` bigint(0) NOT NULL,
  `content` varchar(2048) CHARACTER SET utf8mb4 NOT NULL,
  `status` tinyint(1) NULL DEFAULT 1,
  `create_time` int(0) NULL DEFAULT NULL,
  `update_time` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_circle_comment
-- ----------------------------
DROP TABLE IF EXISTS `t_circle_comment`;
CREATE TABLE `t_circle_comment`  (
  `id` bigint(0) NOT NULL,
  `circle_id` bigint(0) NOT NULL,
  `uid` int(0) NOT NULL,
  `mast` bigint(0) NOT NULL,
  `reply_id` bigint(0) NOT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_circle_like
-- ----------------------------
DROP TABLE IF EXISTS `t_circle_like`;
CREATE TABLE `t_circle_like`  (
  `id` bigint(0) NOT NULL,
  `uid` int(0) NOT NULL,
  `type` tinyint(1) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  `circle_id` bigint(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_collection
-- ----------------------------
DROP TABLE IF EXISTS `t_collection`;
CREATE TABLE `t_collection`  (
  `id` bigint(0) NOT NULL,
  `uid` int(0) NOT NULL,
  `type` tinyint(1) NOT NULL,
  `book_id` bigint(0) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_comment
-- ----------------------------
DROP TABLE IF EXISTS `t_comment`;
CREATE TABLE `t_comment`  (
  `id` bigint(0) NOT NULL,
  `uid` int(0) NOT NULL,
  `content` varchar(55) CHARACTER SET utf8mb4 NOT NULL,
  `book_id` bigint(0) NOT NULL,
  `reply_id` bigint(0) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_follow
-- ----------------------------
DROP TABLE IF EXISTS `t_follow`;
CREATE TABLE `t_follow`  (
  `id` bigint(0) NOT NULL,
  `follow_id` int(0) NOT NULL,
  `fans_id` int(0) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_label
-- ----------------------------
DROP TABLE IF EXISTS `t_label`;
CREATE TABLE `t_label`  (
  `id` bigint(0) NOT NULL,
  `name` varchar(55) CHARACTER SET utf8mb4 NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_message
-- ----------------------------
DROP TABLE IF EXISTS `t_message`;
CREATE TABLE `t_message`  (
  `id` bigint(0) NOT NULL,
  `cmd` tinyint(0) NOT NULL,
  `len` int(0) NOT NULL,
  `form_id` int(0) NOT NULL,
  `to_id` int(0) NOT NULL,
  `content` text CHARACTER SET utf8mb4 NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_message_read
-- ----------------------------
DROP TABLE IF EXISTS `t_message_read`;
CREATE TABLE `t_message_read`  (
  `id` bigint(0) NOT NULL,
  `uid` int(0) NOT NULL,
  `msg_id` bigint(0) NOT NULL,
  `type` tinyint(0) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_recharge
-- ----------------------------
DROP TABLE IF EXISTS `t_recharge`;
CREATE TABLE `t_recharge`  (
  `id` bigint(0) NOT NULL,
  `uid` int(0) NOT NULL,
  `amount` tinyint(0) NOT NULL,
  `type` bigint(0) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_record
-- ----------------------------
DROP TABLE IF EXISTS `t_record`;
CREATE TABLE `t_record`  (
  `id` bigint(0) NOT NULL,
  `uid` int(0) NOT NULL,
  `book_id` bigint(0) NOT NULL,
  `chapter_id` bigint(0) NOT NULL,
  `page` int(0) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_role
-- ----------------------------
DROP TABLE IF EXISTS `t_role`;
CREATE TABLE `t_role`  (
  `id` bigint(0) NOT NULL,
  `role_name` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
  `remakes` varchar(255) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_type
-- ----------------------------
DROP TABLE IF EXISTS `t_type`;
CREATE TABLE `t_type`  (
  `id` bigint(0) NOT NULL,
  `pid` bigint(0) NOT NULL,
  `pic` varchar(155) CHARACTER SET utf8mb4 NOT NULL,
  `name` varchar(55) CHARACTER SET utf8mb4 NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `type_name_index`(`name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_user_details
-- ----------------------------
DROP TABLE IF EXISTS `t_user_details`;
CREATE TABLE `t_user_details`  (
  `id` bigint(0) NOT NULL,
  `uid` int(0) NOT NULL,
  `head_img` varchar(512) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `phone` char(11) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `login_type` tinyint(0) NULL DEFAULT NULL,
  `birthday` int(0) NULL DEFAULT NULL,
  `region` varchar(255) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `nation` varchar(55) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `autograph` varchar(255) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_user_info
-- ----------------------------
DROP TABLE IF EXISTS `t_user_info`;
CREATE TABLE `t_user_info`  (
  `id` bigint(0) NOT NULL,
  `uid` int(0) NOT NULL,
  `username` varchar(55) CHARACTER SET utf8mb4 NOT NULL,
  `nickname` varchar(55) CHARACTER SET utf8mb4 NOT NULL,
  `password` char(128) CHARACTER SET utf8mb4 NOT NULL,
  `salt` char(10) CHARACTER SET utf8mb4 NOT NULL,
  `type` tinyint(0) NOT NULL COMMENT '1管理员，2普通用户，3作家',
  `public_key` char(5) CHARACTER SET utf8mb4 NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(0) NOT NULL,
  `update_time` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uid`(`uid`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  INDEX `password`(`password`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
