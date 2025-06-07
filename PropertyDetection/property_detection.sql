/*
 Navicat Premium Dump SQL

 Source Server         : dockerMySQL
 Source Server Type    : MySQL
 Source Server Version : 90300 (9.3.0)
 Source Host           : 127.0.0.1:3306
 Source Schema         : property_detection

 Target Server Type    : MySQL
 Target Server Version : 90300 (9.3.0)
 File Encoding         : 65001

 Date: 06/06/2025 17:16:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for copyright
-- ----------------------------
DROP TABLE IF EXISTS `copyright`;
CREATE TABLE `copyright`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `vector` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `owner` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `register_date` date NULL DEFAULT NULL,
  `deleted` int NULL DEFAULT 0,
  `status` int NULL DEFAULT 1,
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of copyright
-- ----------------------------

-- ----------------------------
-- Table structure for detection_records
-- ----------------------------
DROP TABLE IF EXISTS `detection_records`;
CREATE TABLE `detection_records`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `deleted` int NULL DEFAULT 0,
  `status` int NULL DEFAULT 1,
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `score` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0.0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1748070321538 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of detection_records
-- ----------------------------

-- ----------------------------
-- Table structure for image
-- ----------------------------
DROP TABLE IF EXISTS `image`;
CREATE TABLE `image`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `vector` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `deleted` int NULL DEFAULT 0,
  `status` int NULL DEFAULT 1,
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `relation_patent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1749195594015 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of image
-- ----------------------------

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `deleted` int NULL DEFAULT 0 COMMENT '逻辑删除标志，0 表示未删除，1 表示已删除',
  `status` int NULL DEFAULT 1 COMMENT '状态，1 表示正常，其他值表示异常',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单标签',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单路径',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单组件',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'icon-caidan' COMMENT '菜单图标',
  `icon_bg_color` varchar(7) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '#fff' COMMENT '菜单图标背景颜色',
  `meta` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单元数据',
  `href` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单链接',
  `parent_id` bigint NULL DEFAULT NULL COMMENT '父菜单 ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1740496414534 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1749104617, 1, 1, '123', '123', '12324', '123', '', '', '', '', 1740496414531);
INSERT INTO `menu` VALUES (1740496414501, 0, 1, NULL, '工具', '/util', NULL, 'icon-caidan', '#fff', '{\"i18n\": \"util\"}', NULL, 0);
INSERT INTO `menu` VALUES (1740496414502, 0, 1, NULL, '缓冲', '/cache', 'views/util/cache', 'icon-caidan', '#fff', '{\"i18n\": \"cache\", \"keepAlive\": true}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414503, 0, 1, NULL, '参数', '/params', 'views/util/params', 'icon-caidan', '#fff', '{\"i18n\": \"params\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414504, 0, 1, NULL, '详情页', '/detail', 'views/util/detail', 'icon-caidan', '#fff', '{\"i18n\": \"detail\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414505, 0, 1, NULL, '标签', '/tags', 'views/util/tags', 'icon-caidan', '#fff', '{\"i18n\": \"tags\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414506, 0, 1, NULL, '存储', '/store', 'views/util/store', 'icon-caidan', '#fff', '{\"i18n\": \"store\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414507, 0, 1, NULL, '日志监控', '/logs', 'views/util/logs', 'icon-caidan', '#fff', '{\"i18n\": \"logs\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414508, 0, 1, NULL, '通用模板', '/crud', 'views/util/crud', 'icon-caidan', '#fff', '{\"i18n\": \"crud\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414509, 0, 1, NULL, '表格', '/table', 'views/util/table', 'icon-caidan', '#fff', '{\"i18n\": \"table\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414510, 0, 1, NULL, '表单', '/form', 'views/util/form', 'icon-caidan', '#fff', '{\"i18n\": \"form\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414511, 0, 1, NULL, '权限', '/permission', 'views/util/permission', 'icon-caidan', '#fff', '{\"i18n\": \"permission\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414512, 0, 1, NULL, '表格表单', '/crud-form', 'views/util/crud-form', 'icon-caidan', '#fff', '{\"i18n\": \"crudForm\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414513, 0, 1, NULL, '返回顶部', '/top', 'views/util/top', 'icon-caidan', '#fff', '{\"i18n\": \"top\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414514, 0, 1, NULL, '图钉', '/affix', 'views/util/affix', 'icon-caidan', '#fff', '{\"i18n\": \"affix\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414515, 0, 1, NULL, '多级菜单', '/deep', NULL, 'icon-caidan', '#fff', NULL, NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414516, 0, 1, NULL, '外部页面', '/out', NULL, 'icon-caidan', '#fff', '{\"i18n\": \"out\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414517, 0, 1, NULL, '异常页', '/error', NULL, 'icon-caidan', '#fff', '{\"i18n\": \"error\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414518, 0, 1, NULL, '关于', '/about', 'views/util/about', 'icon-caidan', '#fff', '{\"i18n\": \"about\"}', NULL, 1740496414501);
INSERT INTO `menu` VALUES (1740496414519, 0, 1, NULL, '多级菜单1-1', 'deep', NULL, 'icon-caidan', '#fff', NULL, NULL, 1740496414515);
INSERT INTO `menu` VALUES (1740496414520, 0, 1, NULL, '多级菜单2-1', 'deep', 'views/util/deep', 'icon-caidan', '#fff', NULL, NULL, 1740496414519);
INSERT INTO `menu` VALUES (1740496414521, 0, 1, NULL, '官方网站(内嵌页面)', 'website', NULL, 'icon-caidan', '#fff', '{\"i18n\": \"website\"}', 'https://avuejs.com', 1740496414516);
INSERT INTO `menu` VALUES (1740496414522, 0, 1, NULL, '全局函数(外链页面)', 'api', NULL, 'icon-caidan', '#fff', '{\"i18n\": \"api\", \"target\": \"_blank\"}', 'https://avuejs.com/docs/api?test1=1&test2=2', 1740496414516);
INSERT INTO `menu` VALUES (1740496414523, 0, 1, NULL, 'error403', 'error', 'components/error-page/403', 'icon-caidan', '#fff', NULL, NULL, 1740496414517);
INSERT INTO `menu` VALUES (1740496414524, 0, 1, NULL, 'error404', '404', 'components/error-page/404', 'icon-caidan', '#fff', NULL, NULL, 1740496414517);
INSERT INTO `menu` VALUES (1740496414525, 0, 1, NULL, 'error500', '500', 'components/error-page/500', 'icon-caidan', '#fff', NULL, NULL, 1740496414517);
INSERT INTO `menu` VALUES (1740496414526, 0, 1, NULL, '系统管理', '/system', NULL, 'icon-caidan', '#fff', '{\"i18n\": \"system\"}', NULL, 0);
INSERT INTO `menu` VALUES (1740496414528, 0, 1, NULL, '菜单管理', '/menu', 'views/system/menu', 'icon-caidan', '#fff', '{\"i18n\": \"menu\"}', NULL, 1740496414526);
INSERT INTO `menu` VALUES (1740496414529, 0, 1, NULL, '用户管理', '/user', 'views/system/user', 'icon-caidan', '#fff', '{\"i18n\":\"用户管理\"}', NULL, 1740496414526);
INSERT INTO `menu` VALUES (1740496414530, 0, 1, NULL, '角色管理', '/role', 'views/system/role', 'icon-caidan', '#fff', '{\"i18n\":\"角色管理\"}', NULL, 1740496414526);
INSERT INTO `menu` VALUES (1740496414531, 0, 1, NULL, '知产管理', '/image', NULL, 'icon-caidan', '#fff', '{\"i18n\": \"image\"}', NULL, 0);
INSERT INTO `menu` VALUES (1740496414532, 0, 1, NULL, '知产数据库', '/image', 'views/image/image', 'icon-caidan', '#fff', '{\"i18n\":\"知产数据库\"}', NULL, 1740496414531);
INSERT INTO `menu` VALUES (1740496414533, 0, 1, NULL, '知产检索', '/task', 'views/image/task', 'icon-caidan', '#fff', '{\"i18n\":\"知产检索\"}', NULL, 1740496414531);

-- ----------------------------
-- Table structure for menu_permission
-- ----------------------------
DROP TABLE IF EXISTS `menu_permission`;
CREATE TABLE `menu_permission`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `deleted` int NULL DEFAULT 0 COMMENT '逻辑删除标志，0 表示未删除，1 表示已删除',
  `status` int NULL DEFAULT 1 COMMENT '状态，1 表示正常，其他值表示异常',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `role_id` bigint NULL DEFAULT NULL COMMENT '角色 ID',
  `menu_id` bigint NULL DEFAULT NULL COMMENT '菜单 ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1740757119900 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '菜单权限关联表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menu_permission
-- ----------------------------
INSERT INTO `menu_permission` VALUES (1749104617, 1, 1, 'admin', 1, 1749104617);
INSERT INTO `menu_permission` VALUES (1740496414501, 0, 1, 'admin', 1, 1740496414501);
INSERT INTO `menu_permission` VALUES (1740496414502, 0, 1, 'admin', 1, 1740496414502);
INSERT INTO `menu_permission` VALUES (1740496414503, 0, 1, 'admin', 1, 1740496414503);
INSERT INTO `menu_permission` VALUES (1740496414504, 0, 1, 'admin', 1, 1740496414504);
INSERT INTO `menu_permission` VALUES (1740496414505, 0, 1, 'admin', 1, 1740496414505);
INSERT INTO `menu_permission` VALUES (1740496414506, 0, 1, 'admin', 1, 1740496414506);
INSERT INTO `menu_permission` VALUES (1740496414507, 0, 1, 'admin', 1, 1740496414507);
INSERT INTO `menu_permission` VALUES (1740496414508, 0, 1, 'admin', 1, 1740496414508);
INSERT INTO `menu_permission` VALUES (1740496414509, 0, 1, 'admin', 1, 1740496414509);
INSERT INTO `menu_permission` VALUES (1740496414510, 0, 1, 'admin', 1, 1740496414510);
INSERT INTO `menu_permission` VALUES (1740496414511, 0, 1, 'admin', 1, 1740496414511);
INSERT INTO `menu_permission` VALUES (1740496414512, 0, 1, 'admin', 1, 1740496414512);
INSERT INTO `menu_permission` VALUES (1740496414513, 0, 1, 'admin', 1, 1740496414513);
INSERT INTO `menu_permission` VALUES (1740496414514, 0, 1, 'admin', 1, 1740496414514);
INSERT INTO `menu_permission` VALUES (1740496414515, 0, 1, 'admin', 1, 1740496414515);
INSERT INTO `menu_permission` VALUES (1740496414516, 0, 1, 'admin', 1, 1740496414516);
INSERT INTO `menu_permission` VALUES (1740496414517, 0, 1, 'admin', 1, 1740496414517);
INSERT INTO `menu_permission` VALUES (1740496414518, 0, 1, 'admin', 1, 1740496414518);
INSERT INTO `menu_permission` VALUES (1740496414519, 0, 1, 'admin', 1, 1740496414519);
INSERT INTO `menu_permission` VALUES (1740496414520, 0, 1, 'admin', 1, 1740496414520);
INSERT INTO `menu_permission` VALUES (1740496414521, 0, 1, 'admin', 1, 1740496414521);
INSERT INTO `menu_permission` VALUES (1740496414522, 0, 1, 'admin', 1, 1740496414522);
INSERT INTO `menu_permission` VALUES (1740496414523, 0, 1, 'admin', 1, 1740496414523);
INSERT INTO `menu_permission` VALUES (1740496414524, 0, 1, 'admin', 1, 1740496414524);
INSERT INTO `menu_permission` VALUES (1740496414525, 0, 1, 'admin', 1, 1740496414525);
INSERT INTO `menu_permission` VALUES (1740496414526, 0, 1, 'admin', 1, 1740496414526);
INSERT INTO `menu_permission` VALUES (1740496414528, 0, 1, 'admin', 1, 1740496414528);
INSERT INTO `menu_permission` VALUES (1740496414529, 0, 1, 'admin', 1, 1740496414529);
INSERT INTO `menu_permission` VALUES (1740496414530, 0, 1, 'admin', 1, 1740496414530);
INSERT INTO `menu_permission` VALUES (1740496414531, 0, 1, 'admin', 1, 1740496414531);
INSERT INTO `menu_permission` VALUES (1740496414532, 0, 1, 'admin', 1, 1740496414532);
INSERT INTO `menu_permission` VALUES (1740496414533, 0, 1, 'admin', 1, 1740496414533);
INSERT INTO `menu_permission` VALUES (1740757119842, 0, 1, 'app_admin', 2, 1740496414529);
INSERT INTO `menu_permission` VALUES (1740757119851, 0, 1, 'app_admin', 2, 1740496414530);

-- ----------------------------
-- Table structure for patent
-- ----------------------------
DROP TABLE IF EXISTS `patent`;
CREATE TABLE `patent`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `vector` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `owner` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `register_date` date NULL DEFAULT NULL,
  `deleted` int NULL DEFAULT 0,
  `status` int NULL DEFAULT 1,
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1748027105091 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of patent
-- ----------------------------

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `deleted` int NULL DEFAULT 0 COMMENT '逻辑删除标志，0 表示未删除，1 表示已删除',
  `status` int NULL DEFAULT 1 COMMENT '状态，1 表示正常，其他值表示异常',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '角色名称',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name` ASC) USING BTREE,
  UNIQUE INDEX `name_2`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1748704926 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, 0, 1, '超级管理员', 'admin');
INSERT INTO `role` VALUES (2, 0, 1, '平台管理员', 'app_admin');
INSERT INTO `role` VALUES (3, 0, 1, 'VIP用户', 'vip_user');
INSERT INTO `role` VALUES (4, 0, 1, '普通用户', 'user');

-- ----------------------------
-- Table structure for trademark
-- ----------------------------
DROP TABLE IF EXISTS `trademark`;
CREATE TABLE `trademark`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `vector` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `owner` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `register_date` date NULL DEFAULT NULL,
  `deleted` int NULL DEFAULT 0,
  `status` int NULL DEFAULT 1,
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of trademark
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `deleted` int NULL DEFAULT 0 COMMENT '逻辑删除标志，0 表示未删除，1 表示已删除',
  `status` int NULL DEFAULT 1 COMMENT '状态，1 表示正常，0表示封号',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '密码',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `role_id` bigint NULL DEFAULT 0 COMMENT '角色，默认为无',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE,
  UNIQUE INDEX `email`(`email` ASC) USING BTREE,
  UNIQUE INDEX `username_2`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1740496761241 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 0, 1, NULL, 'admin', '123456', '213269@qq.com', 1);
INSERT INTO `user` VALUES (1740496761240, 0, 1, NULL, 'app_admin', '12345678', '2131231342@qq.com', 2);

SET FOREIGN_KEY_CHECKS = 1;
