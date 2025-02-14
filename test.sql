/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80029 (8.0.29)
 Source Host           : localhost:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 80029 (8.0.29)
 File Encoding         : 65001

 Date: 14/02/2025 16:47:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ad_items
-- ----------------------------
DROP TABLE IF EXISTS `ad_items`;
CREATE TABLE `ad_items`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '广告名称',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '广告标题',
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '广告图片',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '广告内容',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '广告链接',
  `start_time` timestamp NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` timestamp NULL DEFAULT NULL COMMENT '结束时间',
  `status` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'N' COMMENT '状态 N 隐藏-Y显示',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ad_items
-- ----------------------------

-- ----------------------------
-- Table structure for ad_space_item
-- ----------------------------
DROP TABLE IF EXISTS `ad_space_item`;
CREATE TABLE `ad_space_item`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `item_id` bigint NULL DEFAULT NULL COMMENT '广告内容ID',
  `space_id` bigint NULL DEFAULT NULL COMMENT '广告位置ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 57 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ad_space_item
-- ----------------------------

-- ----------------------------
-- Table structure for ad_spaces
-- ----------------------------
DROP TABLE IF EXISTS `ad_spaces`;
CREATE TABLE `ad_spaces`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '广告位置',
  `key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '广告标识',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注说明',
  `status` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'N' COMMENT '状态 N 隐藏-Y显示',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `key`(`key` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ad_spaces
-- ----------------------------

-- ----------------------------
-- Table structure for app_tools
-- ----------------------------
DROP TABLE IF EXISTS `app_tools`;
CREATE TABLE `app_tools`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `taxonomy_id` int NULL DEFAULT NULL COMMENT '归类ID',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '工具标识代码',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '工具图标',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'URL地址',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '描述',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '详情',
  `status` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'N' COMMENT '是否上线 Y 禁用 N 否',
  `sort` int NULL DEFAULT 0 COMMENT '排序',
  `price` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '价格',
  `market_price` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '市场价',
  `recommend` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'N' COMMENT '是否推荐 Y-是 N-否',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `taxonomy_id`(`taxonomy_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 97 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app_tools
-- ----------------------------
INSERT INTO `app_tools` VALUES (1, 66, 'tool:nianzhongzongjie', '/assets/icons/text-2.svg', '/pages/tool/text', '年终总结', '年终总结写的好，升值加薪少不了', '无', 'Y', 0, 0.00, NULL, 'Y', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (2, 66, 'tool:shuzhibaogao', '/assets/icons/text-2.svg', '/pages/tool/text', '述职报告', '工作做的好不好，述职报告要搞好', '无', 'Y', 0, 0.00, NULL, 'Y', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (3, 66, 'tool:gongzuojihua', '/assets/icons/text-2.svg', '/pages/tool/text', '工作计划', '今天工作计划做的好，明天换你当领导', '无', 'Y', 0, 0.00, NULL, 'Y', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (4, 66, 'tool:xindetihui', '/assets/icons/text-2.svg', '/pages/tool/text', '心得体会', '工作心得，学习心得，根据模板快速生成', '无', 'Y', 0, 0.00, NULL, 'Y', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (5, 66, 'tool:zhuanyewenzhang', '/assets/icons/text-2.svg', '/pages/tool/text', '专业文章', '行业领域专业内容写作，规范格式', '无', 'Y', 0, 0.00, NULL, 'Y', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (6, 66, 'tool:shiyanbaogao', '/assets/icons/text-2.svg', '/pages/tool/text', '实验报告', '实验的结果数据分析和报告', '无', 'Y', 0, 0.00, NULL, 'Y', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (7, 66, 'tool:shixunbaogao', '/assets/icons/text-2.svg', '/pages/tool/text', '实训报告', '职业教育，技能培训，岗前培训等记录总结和报告', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (8, 66, 'tool:shijianbaogao', '/assets/icons/text-2.svg', '/pages/tool/text', '实践报告', '做个总结，让这次实践更有意义', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (9, 66, 'tool:shixibaogao', '/assets/icons/text-2.svg', '/pages/tool/text', '实习报告', '仅需1分钟就能完成一篇优质的实习报告', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (10, 66, 'tool:sixianghuibao', '/assets/icons/text-2.svg', '/pages/tool/text', '思想汇报', '思想汇报', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (11, 66, 'tool:jingxuangao', '/assets/icons/text-2.svg', '/pages/tool/text', '竞选稿', '竞选稿', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (12, 66, 'tool:tuijianxin', '/assets/icons/text-2.svg', '/pages/tool/text', '推荐信', '推荐信', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (13, 66, 'tool:gongzuozongjie', '/assets/icons/text-2.svg', '/pages/tool/text', '工作总结', '工作总结', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (14, 66, 'tool:tongzhi', '/assets/icons/text-2.svg', '/pages/tool/text', '通知', '自动生成各类通知', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (15, 66, 'tool:gongzuobaogao', '/assets/icons/text-2.svg', '/pages/tool/text', '工作报告', '工作日报，周报，部门汇报，项目汇报等', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (16, 66, 'tool:huiyijiyao', '/assets/icons/text-2.svg', '/pages/tool/text', '会议纪要', '会议纪要', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (17, 66, 'tool:diaoyanbaogao', '/assets/icons/text-2.svg', '/pages/tool/text', '调研报告', '调研报告', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (18, 66, 'tool:jinshengzongjie', '/assets/icons/text-2.svg', '/pages/tool/text', '晋升总结', '晋升总结', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (19, 66, 'tool:zhaopinxuqiu', '/assets/icons/text-2.svg', '/pages/tool/text', '招聘需求', '招聘需求', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (20, 66, 'tool:gongwen', '/assets/icons/text-2.svg', '/pages/tool/text', '公文', '按照公文格式标准一键生成', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (21, 66, 'tool:guizhangzhidu', '/assets/icons/text-2.svg', '/pages/tool/text', '规章制度', '规章制度', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (22, 67, 'tool:manfenzuowen', '/assets/icons/text-3.svg', '/pages/tool/text', '满分作文', '通用文章写作', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (23, 67, 'tool:duilian', '/assets/icons/text-3.svg', '/pages/tool/text', '春节对联', '春节对联', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (24, 67, 'tool:yanjianggao', '/assets/icons/text-3.svg', '/pages/tool/text', '演讲稿', '演讲稿', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (25, 67, 'tool:bianlun', '/assets/icons/text-3.svg', '/pages/tool/text', '辩论稿', '辩论稿', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (26, 67, 'tool:zhuchi', '/assets/icons/text-3.svg', '/pages/tool/text', '主持稿', '主持稿', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (27, 67, 'tool:gushichi', '/assets/icons/text-3.svg', '/pages/tool/text', '古诗词', '诗词歌赋，琴棋书画', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (28, 67, 'tool:zhengwen', '/assets/icons/text-3.svg', '/pages/tool/text', '征文', '征文', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (29, 67, 'tool:duhougan', '/assets/icons/text-3.svg', '/pages/tool/text', '读后感', '读完一部名作后的感悟及评价', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (30, 67, 'tool:guanhougan', '/assets/icons/text-3.svg', '/pages/tool/text', '观后感', '观看影视、表演、文化景点后的观后感', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (31, 67, 'tool:weiwenxin', '/assets/icons/text-3.svg', '/pages/tool/text', '慰问信', '慰问信', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (32, 67, 'tool:banjiangci', '/assets/icons/text-3.svg', '/pages/tool/text', '颁奖词', '颁奖词', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (33, 67, 'tool:zhici', '/assets/icons/text-3.svg', '/pages/tool/text', '致辞', '会议致辞、活动致辞...', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (34, 67, 'tool:xiandaishi', '/assets/icons/text-3.svg', '/pages/tool/text', '现代诗', '现代诗', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (35, 67, 'tool:tuokouxiu', '/assets/icons/text-3.svg', '/pages/tool/text', '脱口秀', '脱口秀', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (36, 67, 'tool:xiaopinjuben', '/assets/icons/text-3.svg', '/pages/tool/text', '小品剧本', '小品剧本', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (37, 67, 'tool:wenyipinglun', '/assets/icons/text-3.svg', '/pages/tool/text', '文艺评论', '文学、影视、音乐、绘画、节目、戏剧的评论、推荐、解说', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (38, 67, 'tool:weixiaoshuo', '/assets/icons/text-3.svg', '/pages/tool/text', '微小说', '微小说', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (39, 67, 'tool:geci', '/assets/icons/text-3.svg', '/pages/tool/text', '歌词', '歌词', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (40, 67, 'tool:sanwen', '/assets/icons/text-3.svg', '/pages/tool/text', '散文', '散文', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (41, 67, 'tool:minshiqisushu', '/assets/icons/text-3.svg', '/pages/tool/text', '民事起诉书', '民事起诉书', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (42, 67, 'tool:qingshugaobai', '/assets/icons/text-3.svg', '/pages/tool/text', '情书告白', '情书告白', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (43, 68, 'tool:lvxingjihua', '/assets/icons/text-4.svg', '/pages/tool/text', '旅行计划', '旅行计划', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (44, 68, 'tool:jiantaoshu', '/assets/icons/text-4.svg', '/pages/tool/text', '检讨书', '检讨书', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (45, 68, 'tool:ziwojieshao', '/assets/icons/text-4.svg', '/pages/tool/text', '自我介绍', '自我介绍', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (46, 68, 'tool:shenqingshu', '/assets/icons/text-4.svg', '/pages/tool/text', '申请书', '申请书', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (47, 68, 'tool:jianyishu', '/assets/icons/text-4.svg', '/pages/tool/text', '建议书', '建议书', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (48, 68, 'tool:ziwopingjia', '/assets/icons/text-4.svg', '/pages/tool/text', '自我评价', '自我评价', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (49, 68, 'tool:zhuhexin', '/assets/icons/text-4.svg', '/pages/tool/text', '祝贺信', '祝贺信', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (50, 68, 'tool:jianli', '/assets/icons/text-4.svg', '/pages/tool/text', '简历', '简历制作', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (51, 69, 'tool:pengyouquan', '/assets/icons/wechat-pyq.svg', '/pages/tool/text', '微信朋友圈文案', '微信朋友圈文案', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (52, 69, 'tool:xiaohongshu', '/assets/icons/xiaohongshu.svg', '/pages/tool/text', '小红书文案', '小红书文案', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (53, 69, 'tool:jinritoutiao', '/assets/icons/toutiao.svg', '/pages/tool/text', '今日头条文章', '今日头条文章', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (54, 69, 'tool:gongzhonghao', '/assets/icons/weichat.svg', '/pages/tool/text', '公众号文章', '公众号文章', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (55, 69, 'tool:zhihu', '/assets/icons/zhihu.svg', '/pages/tool/text', '知乎问答', '知乎问答', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (56, 69, 'tool:xinwen', '/assets/icons/text-4.svg', '/pages/tool/text', '新闻稿', '新闻稿', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (57, 69, 'tool:kepu', '/assets/icons/text-4.svg', '/pages/tool/text', '科普文案', '科普文案', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (58, 71, 'tool:yingxiaocehua', '/assets/icons/text-5.svg', '/pages/tool/text', '营销策划', '营销策划', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (59, 71, 'tool:huodongcehua', '/assets/icons/text-5.svg', '/pages/tool/text', '活动策划', '活动策划', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (60, 71, 'tool:anlifenxibaogao', '/assets/icons/text-5.svg', '/pages/tool/text', '案例分析报告', '案例分析报告', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (61, 71, 'tool:chanpincuxiao', '/assets/icons/text-5.svg', '/pages/tool/text', '产品促销文案', '产品促销文案', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (62, 71, 'tool:gongsijieshao', '/assets/icons/text-5.svg', '/pages/tool/text', '公司介绍', '市场营销推广，官网宣传，媒体广告投放，需要一份公司介绍', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (63, 71, 'tool:chanpinjieshao', '/assets/icons/text-5.svg', '/pages/tool/text', '产品介绍', '产品介绍', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (64, 71, 'tool:hetongxieyi', '/assets/icons/text-5.svg', '/pages/tool/text', '合同协议', '合同协议', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (65, 71, 'tool:diaochawenjuan', '/assets/icons/text-5.svg', '/pages/tool/text', '调查问卷', '调查问卷', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (66, 71, 'tool:toubiaoshu', '/assets/icons/text-5.svg', '/pages/tool/text', '投标书', '投标书', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (67, 71, 'tool:hangyeyanjiubaogao', '/assets/icons/text-5.svg', '/pages/tool/text', '行业研究报告', '行业研究报告', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (68, 71, 'tool:chuangyejihuashu', '/assets/icons/text-5.svg', '/pages/tool/text', '创业计划书', '创业计划书', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (69, 71, 'tool:zhiyeguihua', '/assets/icons/text-5.svg', '/pages/tool/text', '职业规划', '职业规划', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (70, 71, 'tool:yaoqinghan', '/assets/icons/text-5.svg', '/pages/tool/text', '邀请函', '邀请函', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (71, 71, 'tool:pinggubaogao', '/assets/icons/text-5.svg', '/pages/tool/text', '评估报告', '评估报告', '无', 'N', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (72, 71, 'tool:chanpcuxiaofangan', '/assets/icons/text-5.svg', '/pages/tool/text', '产品促销方案', '产品促销方案', '无', 'N', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (73, 71, 'tool:okr', '/assets/icons/text-5.svg', '/pages/tool/text', 'OKR制定', 'OKR制定', '无', 'N', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (74, 71, 'tool:swot', '/assets/icons/text-5.svg', '/pages/tool/text', 'SWOT分析', 'SWOT分析', '无', 'N', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (75, 71, 'tool:pest', '/assets/icons/text-5.svg', '/pages/tool/text', 'PEST分析', 'PEST分析', '无', 'N', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (76, 71, 'tool:shejishuomingshu', '/assets/icons/text-5.svg', '/pages/tool/text', '设计说明书', '设计说明书', '无', 'N', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (77, 70, 'tool:jiaoxuesheji', '/assets/icons/text-6.svg', '/pages/tool/text', '教学设计', '教学设计', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (78, 70, 'tool:banhuizhuti', '/assets/icons/text-6.svg', '/pages/tool/text', '主题班会', '主题班会', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (79, 70, 'tool:shitishengcheng', '/assets/icons/text-6.svg', '/pages/tool/text', '试题生成', '试题生成', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (80, 70, 'tool:monidabian', '/assets/icons/text-6.svg', '/pages/tool/text', '模拟答辩', '模拟答辩', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (81, 70, 'tool:zhichengpingxuan', '/assets/icons/text-6.svg', '/pages/tool/text', '职称评选报告', '职称评选报告', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (82, 70, 'tool:zhutiyanjiang', '/assets/icons/text-6.svg', '/pages/tool/text', '主题演讲', '主题演讲', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (83, 70, 'tool:xuekejiaoan', '/assets/icons/text-6.svg', '/pages/tool/text', '学科教案', '学科教案', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (84, 70, 'tool:jiaoxuehuodong', '/assets/icons/text-6.svg', '/pages/tool/text', '教学活动安排', '教学活动安排', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (85, 70, 'tool:jiaoshipeixunxinde', '/assets/icons/text-6.svg', '/pages/tool/text', '教师培训心得', '教师培训心得', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (86, 70, 'tool:jiaoyangongzuozongjie', '/assets/icons/text-6.svg', '/pages/tool/text', '教研工作总结', '教研工作总结', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (87, 70, 'tool:jiaoshishuzhibaogao', '/assets/icons/text-6.svg', '/pages/tool/text', '教师述职报告', '教师述职报告', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (88, 70, 'tool:jiaoxuemubiao', '/assets/icons/text-6.svg', '/pages/tool/text', '教学目标', '教学目标', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (89, 70, 'tool:jiaoxuejianyi', '/assets/icons/text-6.svg', '/pages/tool/text', '教学建议与意见', '教学建议与意见', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (90, 70, 'tool:zhaoshengjihuawenan', '/assets/icons/text-6.svg', '/pages/tool/text', '招生计划文案', '招生计划文案', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (91, 70, 'tool:youxiujiaoshidaibiaofayan', '/assets/icons/text-6.svg', '/pages/tool/text', '优秀教师代表发言', '优秀教师代表发言', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (92, 70, 'tool:jiafangbaogao', '/assets/icons/text-6.svg', '/pages/tool/text', '家访报告', '家访报告', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (93, 70, 'tool:xinlishudaofangan', '/assets/icons/text-6.svg', '/pages/tool/text', '心理疏导方案', '青少年问题，心理疏导方案', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (94, 70, 'tool:banjiguanlizhidu', '/assets/icons/text-6.svg', '/pages/tool/text', '班级管理制度', '班级管理制度', '无', 'Y', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (95, 69, 'tool:abc', '/assets/icons/text-4.svg', '/pages/tool/text', '一键去背景', '自选写作类型，满足多样化的写作需求', '无', 'N', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');
INSERT INTO `app_tools` VALUES (96, 69, 'tool:abc', '/assets/icons/text-4.svg', '/pages/tool/text', '一键去背景', '自选写作类型，满足多样化的写作需求', '无', 'N', 0, 0.00, NULL, 'N', '2025-01-01 16:49:52', '2025-01-01 16:49:55');

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '图片名称',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'URL',
  `user_id` bigint NULL DEFAULT 0 COMMENT '归属用户ID',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '存储路径',
  `size` bigint NULL DEFAULT NULL COMMENT '大小',
  `hash` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图片hash',
  `uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'UUID',
  `extension` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图片后缀',
  `mime` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `disk` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '上传驱动',
  `download_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '下载链接地址',
  `expired_at` timestamp NULL DEFAULT NULL COMMENT '过期时间',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uuid`(`uuid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12481 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of files
-- ----------------------------

-- ----------------------------
-- Table structure for options
-- ----------------------------
DROP TABLE IF EXISTS `options`;
CREATE TABLE `options`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '配置标题',
  `key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '键名',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '键值',
  PRIMARY KEY (`id`, `key`) USING BTREE,
  UNIQUE INDEX `key`(`key` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of options
-- ----------------------------
INSERT INTO `options` VALUES (1, '系统版本号', 'version', '当前版本 24.12.25');
INSERT INTO `options` VALUES (2, '网站标题', 'site_title', '实用在线工具箱');
INSERT INTO `options` VALUES (3, '网站关键词', 'site_keywords', '在线工具,js代码格式化,css代码格式化,js代码格式化,加密解密工具,图片处理,图片裁剪,在线二维码生成');
INSERT INTO `options` VALUES (4, '网站描述', 'site_description', '实用在线工具箱提供在线工具,js代码格式化,css代码格式化,js代码格式化,加密解密工具,图片处理,图片裁剪,在线二维码生成等各类在线工具');
INSERT INTO `options` VALUES (5, '站点状态', 'site_status', '0');
INSERT INTO `options` VALUES (6, '版权信息', 'copyright', 'Copyright © 2024');
INSERT INTO `options` VALUES (7, '统计代码', 'statistics', '<script>\nvar _hmt = _hmt || [];\n(function() {\n  var hm = document.createElement(\"script\");\n  hm.src = \"https://hm.baidu.com/hm.js?8ea2e85debcb174a1cd3b509d8715045\";\n  var s = document.getElementsByTagName(\"script\")[0]; \n  s.parentNode.insertBefore(hm, s);\n})();\n</script>');
INSERT INTO `options` VALUES (8, '公司名称', 'company', '');
INSERT INTO `options` VALUES (9, '电话', 'telephone', '40069925');
INSERT INTO `options` VALUES (10, 'email', 'email', 'kefu@kktt.cn');
INSERT INTO `options` VALUES (11, '备案号', 'icp_beian', '蜀ICP备19009987号-2');
INSERT INTO `options` VALUES (12, '公安备案', 'gongan_beian', '川公网安备 51010702001950号');
INSERT INTO `options` VALUES (13, '公安备案链接', 'gongan_beian_url', 'http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=51010702001950');

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  `client_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户端ClientID',
  `order_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '订单编号',
  `order_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '订单类型',
  `pay_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '支付方式 alipay,wechat',
  `pay_order_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '支付订单号',
  `pay_time` timestamp NULL DEFAULT NULL COMMENT '支付时间',
  `product_id` bigint NULL DEFAULT NULL COMMENT '商品编号',
  `product_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商品编码',
  `product_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商品名称',
  `product_price` decimal(10, 2) NULL DEFAULT NULL COMMENT '商品价格',
  `quantity` int NOT NULL COMMENT '购买数量',
  `amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '订单总金额',
  `pay_amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '支付金额',
  `pay_qrcode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '二维码字符串',
  `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'pending' COMMENT '订单状态\r\n\'pending\' => \'待付款\',\r\n\'paid\' => \'已付款\',\r\n\'completed\' => \'完成\',\r\n\'canceled\' => \'已取消\',\r\n\'refunded\' => \'已退款\'',
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `extra` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '额外信息',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `expired_at` timestamp NULL DEFAULT NULL COMMENT '订单过期时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_order_no`(`order_no` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 170 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of orders
-- ----------------------------
INSERT INTO `orders` VALUES (14, 2, NULL, '78203239315888', 'recharge', 'wechat', NULL, NULL, NULL, 'recharge', '余额充值', 0.01, 1, 0.01, 0.01, NULL, 'paid', NULL, NULL, NULL, '2023-07-09 23:29:48', '2023-07-09 23:29:48', '2023-07-09 23:59:48');
INSERT INTO `orders` VALUES (16, 2, NULL, '130914471832693', 'recharge', 'wechat', '4200002062202312053773110389', '2023-12-05 22:12:42', NULL, 'recharge', '余额充值', 0.02, 1, 0.02, 0.02, 'weixin://wxpay/bizpayurl?pr=dL2nkAizz', 'completed', NULL, NULL, NULL, '2023-12-05 22:12:21', '2023-12-05 22:12:42', '2023-12-05 22:42:20');
INSERT INTO `orders` VALUES (20, 2, NULL, '144273036933920', 'recharge', 'wechat', '4200002058202401129146354416', '2024-01-12 16:08:43', NULL, 'recharge', '余额充值', 0.10, 1, 0.10, 0.10, 'weixin://wxpay/bizpayurl?pr=Qsj2p1mzz', 'completed', NULL, NULL, NULL, '2024-01-12 16:08:29', '2024-01-12 16:08:43', '2024-01-12 16:38:29');
INSERT INTO `orders` VALUES (21, 2, NULL, '144274535034529', 'recharge', 'alipay', '2024011222001430241418519453', '2024-01-12 16:14:50', NULL, 'recharge', '余额充值', 0.10, 1, 0.10, 0.10, 'https://qr.alipay.com/bax08375rawwtpmzrw9i0041', 'completed', NULL, NULL, NULL, '2024-01-12 16:14:35', '2024-01-12 16:14:50', '2024-01-12 16:44:35');
INSERT INTO `orders` VALUES (22, 2, NULL, '165584422602635', 'recharge', 'wechat', '4200002163202403123058701398', '2024-03-12 21:25:08', NULL, 'recharge', '余额充值', 0.10, 1, 0.10, 0.10, 'weixin://wxpay/bizpayurl?pr=K3gtgPozz', 'completed', NULL, NULL, NULL, '2024-03-12 21:24:44', '2024-03-21 09:49:15', '2024-03-12 21:54:44');
INSERT INTO `orders` VALUES (23, 52, NULL, '168228889392250', 'recharge', 'wechat', '4200002164202403206625135752', '2024-03-20 08:45:25', NULL, 'recharge', '余额充值', 1.00, 1, 1.00, 1.00, 'weixin://wxpay/bizpayurl?pr=gP9fvxfzz', 'completed', NULL, NULL, NULL, '2024-03-20 08:45:06', '2024-03-21 09:49:16', '2024-03-20 09:15:06');
INSERT INTO `orders` VALUES (24, 295, NULL, '168596188013887', 'recharge', 'wechat', '4200002167202403211159534290', '2024-03-21 09:40:26', NULL, 'recharge', '余额充值', 0.10, 1, 0.10, 0.10, 'weixin://wxpay/bizpayurl?pr=z3uRLYHzz', 'completed', NULL, NULL, NULL, '2024-03-21 09:39:39', '2024-03-21 09:49:16', '2024-03-21 10:09:38');
INSERT INTO `orders` VALUES (25, 295, NULL, '168652055776861', 'recharge', 'alipay', '2024032122001430241425743864', '2024-03-21 13:27:16', NULL, 'recharge', '余额充值', 0.10, 1, 0.10, 0.10, 'https://qr.alipay.com/bax031009kilzgklruyb3099', 'completed', NULL, NULL, NULL, '2024-03-21 13:26:58', '2024-03-22 21:53:10', '2024-03-21 13:56:58');
INSERT INTO `orders` VALUES (26, 65, NULL, '183253761416622', 'recharge', 'alipay', '2024050122001409701450591811', '2024-05-01 19:41:54', NULL, 'recharge', '余额充值', 10.00, 1, 10.00, 10.00, 'https://qr.alipay.com/bax091746idonq62yaeb253f', 'completed', NULL, NULL, NULL, '2024-05-01 19:41:28', '2024-05-01 19:41:54', '2024-05-01 20:11:27');
INSERT INTO `orders` VALUES (27, 74, NULL, '194191484345232', 'recharge', 'alipay', '2024060122001400651414012888', '2024-06-01 17:27:33', NULL, 'recharge', '余额充值', 10.00, 1, 10.00, 10.00, 'https://qr.alipay.com/bax06484umgdl2aig3nl25fa', 'completed', NULL, NULL, NULL, '2024-06-01 17:27:10', '2024-06-01 17:27:33', '2024-06-01 17:57:10');
INSERT INTO `orders` VALUES (94, 104, NULL, '209742947846374', 'recharge', 'alipay', '2024071522001495801439463074', '2024-07-15 16:06:50', NULL, 'recharge', '余额充值', 2.01, 1, 2.01, 2.01, 'https://qr.alipay.com/bax05721yankpprqslsh30c4', 'completed', NULL, NULL, NULL, '2024-07-15 16:06:14', '2024-07-15 16:06:50', '2024-07-15 16:36:14');
INSERT INTO `orders` VALUES (97, 2, NULL, '221443654660305', 'recharge', 'wechat', '4200002327202408177056371993', '2024-08-17 17:36:47', NULL, 'recharge', '余额充值', 0.01, 1, 0.01, 0.01, 'weixin://wxpay/bizpayurl?pr=tQ0gsmiz3', 'completed', NULL, NULL, NULL, '2024-08-17 17:36:32', '2024-08-17 17:36:47', '2024-08-17 18:06:32');
INSERT INTO `orders` VALUES (98, 295, NULL, '245403329274193', 'recharge', 'wechat', '4200002348202410242824980793', '2024-10-24 10:28:57', NULL, 'recharge', '余额充值', 0.10, 1, 0.10, 0.10, 'weixin://wxpay/bizpayurl?pr=l54E2Ayz1', 'completed', NULL, NULL, NULL, '2024-10-24 10:28:42', '2024-10-24 10:28:57', '2024-10-24 10:58:42');
INSERT INTO `orders` VALUES (101, 2, NULL, '255517345120639', 'recharge', 'alipay', '2024112222001430241458616236', '2024-11-22 00:26:44', NULL, 'recharge', '余额充值', 0.20, 1, 0.20, 0.20, 'https://qr.alipay.com/bax01902be3zyvw5undp0073', 'completed', NULL, NULL, NULL, '2024-11-22 00:22:44', '2024-11-22 00:26:44', '2024-11-22 00:52:44');
INSERT INTO `orders` VALUES (102, 2, NULL, '255518410634294', 'recharge', 'wechat', '4200002514202411220997754154', '2024-11-22 00:27:22', NULL, 'recharge', '余额充值', 0.30, 1, 0.30, 0.30, 'weixin://wxpay/bizpayurl?pr=veksa9vz1', 'completed', NULL, NULL, NULL, '2024-11-22 00:27:05', '2024-11-22 00:27:22', '2024-11-22 00:57:04');
INSERT INTO `orders` VALUES (105, 295, NULL, '266847776140729', 'recharge', 'wechat', '4200002499202412248199364391', '2024-12-24 00:46:46', NULL, 'recharge', '余额充值', 1.00, 1, 1.00, 1.00, 'weixin://wxpay/bizpayurl?pr=WwawSqmz3', 'completed', NULL, NULL, NULL, '2024-12-24 00:46:23', '2024-12-24 00:46:46', '2024-12-24 01:16:22');
INSERT INTO `orders` VALUES (113, 1, NULL, '271043626552826', 'recharge', 'alipay', '2025010422001430241458956295', '2025-01-04 21:19:56', NULL, 'recharge', '余额充值', 1.00, 1, 1.00, 1.00, 'https://qr.alipay.com/bax01938eu9slf0czoc92505', 'completed', NULL, NULL, NULL, '2025-01-04 21:19:21', '2025-01-04 21:19:56', '2025-01-04 21:49:20');
INSERT INTO `orders` VALUES (131, 307, '', '1876294259846418432', 'recharge', 'wechat', '', NULL, NULL, '', '', 0.00, 0, 0.10, 0.10, '', 'paid', '充值订单', '充值订单', NULL, '2025-01-06 23:46:26', '2025-01-07 12:25:13', '2025-01-07 00:16:26');
INSERT INTO `orders` VALUES (132, 307, '', '1876510048432164864', 'recharge', 'wechat', '', NULL, NULL, '', '', 0.00, 0, 0.10, 0.10, '', 'completed', '充值订单', '充值订单', NULL, '2025-01-07 14:03:54', '2025-01-07 14:18:06', '2025-01-07 14:33:54');
INSERT INTO `orders` VALUES (133, 307, '', '1876675404970856448', 'product', 'wechat', '', NULL, 2, 'vip:month:1', '1个月', 19.00, 1, 19.00, 19.00, '', 'closed', '1个月', '', '{\"month\":1}', '2025-01-08 01:00:58', '2025-01-08 01:31:01', '2025-01-08 01:30:58');
INSERT INTO `orders` VALUES (134, 307, '', '1876675450030264320', 'product', 'wechat', '', NULL, 3, 'vip:month:3', '3个月', 50.00, 1, 50.00, 50.00, '', 'closed', '3个月', '', '{\"month\":3}', '2025-01-08 01:01:09', '2025-01-08 01:32:02', '2025-01-08 01:31:09');
INSERT INTO `orders` VALUES (135, 307, '', '1876675685221666816', 'recharge', 'wechat', '', NULL, 0, '', '', 0.00, 0, 5.00, 5.00, '', 'closed', '充值订单', '充值订单', '', '2025-01-08 01:02:05', '2025-01-08 01:33:02', '2025-01-08 01:32:05');
INSERT INTO `orders` VALUES (136, 307, '', '1876675724811702272', 'product', 'wechat', '', NULL, 2, 'vip:month:1', '1个月', 19.00, 1, 19.00, 19.00, '', 'closed', '1个月', '', '{\"month\":1}', '2025-01-08 01:02:14', '2025-01-08 01:33:02', '2025-01-08 01:32:14');
INSERT INTO `orders` VALUES (137, 307, '', '1876676052000968704', 'product', 'wechat', '', NULL, 2, 'vip:month:1', '1个月', 19.00, 1, 19.00, 19.00, '', 'closed', '1个月', '', '{\"month\":1}', '2025-01-08 01:03:32', '2025-01-08 01:34:02', '2025-01-08 01:33:32');
INSERT INTO `orders` VALUES (138, 307, '', '1876676176081063936', 'product', 'wechat', '', NULL, 2, 'vip:month:1', '1个月', 19.00, 1, 19.00, 19.00, '', 'closed', '1个月', '', '{\"month\":1}', '2025-01-08 01:04:02', '2025-01-08 01:35:03', '2025-01-08 01:34:02');
INSERT INTO `orders` VALUES (139, 307, '', '1876676602767609856', 'product', 'wechat', '', NULL, 2, 'vip:month:1', '1个月', 19.00, 1, 19.00, 19.00, '', 'closed', '1个月', '', '{\"month\":1}', '2025-01-08 01:05:44', '2025-01-08 01:36:01', '2025-01-08 01:35:44');
INSERT INTO `orders` VALUES (140, 307, '', '1876677869002821632', 'recharge', 'wechat', '', NULL, 0, '', '', 0.00, 0, 5.00, 5.00, '', 'closed', '充值订单', '充值订单', '', '2025-01-08 01:10:46', '2025-01-08 01:41:02', '2025-01-08 01:40:46');
INSERT INTO `orders` VALUES (141, 307, '', '1876677958823841792', 'product', 'wechat', '', NULL, 2, 'vip:month:1', '1个月', 19.00, 1, 19.00, 19.00, '', 'closed', '1个月', '', '{\"month\":1}', '2025-01-08 01:11:07', '2025-01-08 01:42:02', '2025-01-08 01:41:07');
INSERT INTO `orders` VALUES (142, 307, '', '1876679318025801728', 'product', 'wechat', '', NULL, 2, 'vip:month:1', '1个月', 19.00, 1, 19.00, 19.00, '', 'closed', '1个月', '1个月', '{\"month\":1}', '2025-01-08 01:16:31', '2025-01-08 01:47:01', '2025-01-08 01:46:31');
INSERT INTO `orders` VALUES (143, 307, '', '1876680082567729152', 'product', 'wechat', '', NULL, 2, 'vip:month:1', '1个月', 0.01, 1, 0.01, 0.01, '', 'completed', '1个月', '1个月', '{\"month\":1}', '2025-01-08 01:19:33', '2025-01-08 01:33:54', '2025-01-08 01:49:33');
INSERT INTO `orders` VALUES (144, 307, '', '1876684101386244096', 'product', 'wechat', '', NULL, 2, 'vip:month:1', '1个月', 0.01, 1, 0.01, 0.01, '', 'completed', '1个月', '1个月', '{\"month\":1}', '2025-01-08 01:35:32', '2025-01-08 01:35:38', '2025-01-08 02:05:32');
INSERT INTO `orders` VALUES (145, 313, NULL, '272406920891419', 'recharge', 'wechat', NULL, NULL, NULL, 'recharge', '余额充值', 1.00, 1, 1.00, 1.00, 'weixin://wxpay/bizpayurl?pr=Jb0GcHGz1', 'closed', NULL, NULL, NULL, '2025-01-08 17:46:36', '2025-01-08 18:17:02', '2025-01-08 18:16:35');
INSERT INTO `orders` VALUES (146, 313, NULL, '272407049507219', 'recharge', 'alipay', NULL, NULL, NULL, 'recharge', '余额充值', 1.00, 1, 1.00, 1.00, 'https://qr.alipay.com/bax038308ss1ouwv0itj5523', 'closed', NULL, NULL, NULL, '2025-01-08 17:47:08', '2025-01-08 18:18:02', '2025-01-08 18:17:07');
INSERT INTO `orders` VALUES (147, 307, '', '1877349418068348928', 'product', 'wechat', '', NULL, 2, 'vip:month:1', '1个月', 0.01, 1, 0.01, 0.01, '', 'completed', '1个月', '1个月', '{\"month\":1}', '2025-01-09 21:39:15', '2025-01-09 21:39:24', '2025-01-09 22:09:15');
INSERT INTO `orders` VALUES (148, 307, '', '1878495787944513536', 'product', 'wechat', '', NULL, 4, 'vip:month:6', '6个月', 99.00, 1, 99.00, 99.00, '', 'closed', '6个月', '6个月', '{\"month\":6}', '2025-01-13 01:34:31', '2025-01-13 02:05:02', '2025-01-13 02:04:31');
INSERT INTO `orders` VALUES (149, 307, '', '1878862155504160768', 'product', 'wechat', '', NULL, 2, 'vip:month:1', '1个月', 0.01, 1, 0.01, 0.01, '', 'completed', '1个月', '1个月', '{\"month\":1}', '2025-01-14 01:50:20', '2025-01-14 01:50:28', '2025-01-14 02:20:20');
INSERT INTO `orders` VALUES (150, 319, '', '1878980492628332544', 'recharge', 'wechat', '', NULL, 0, '', '', 0.00, 0, 5.00, 5.00, '', 'closed', '充值订单', '充值订单', '', '2025-01-14 09:40:34', '2025-01-14 10:11:02', '2025-01-14 10:10:34');
INSERT INTO `orders` VALUES (151, 323, '', '1879094829107908608', 'product', 'wechat', '', NULL, 2, 'vip:month:1', '1个月', 0.01, 1, 0.01, 0.01, '', 'completed', '1个月', '1个月', '{\"month\":1}', '2025-01-14 17:14:54', '2025-01-14 17:15:06', '2025-01-14 17:44:54');
INSERT INTO `orders` VALUES (152, 307, '', '1882714515129241600', 'point', 'wechat', '', NULL, 7, 'point:100', '100点数', 10.00, 1, 10.00, 10.00, '', 'closed', '100点数(赠10点)', '100点数', '{\"point\":100,\"gift\":10}', '2025-01-24 16:58:14', '2025-01-24 17:29:02', '2025-01-24 17:28:14');
INSERT INTO `orders` VALUES (153, 307, '', '1882715354283642880', 'point', 'wechat', '', NULL, 6, 'point:10', '10点数', 0.01, 1, 0.01, 0.01, '', 'completed', '10点数(赠1点)', '10点数(赠1点)', '{\"point\":10,\"gift\":10}', '2025-01-24 17:01:34', '2025-01-24 17:01:52', '2025-01-24 17:31:34');
INSERT INTO `orders` VALUES (154, 307, '', '1882717215703175168', 'point', 'wechat', '', NULL, 6, 'point:10', '10点数', 0.01, 1, 0.01, 0.01, '', 'completed', '10点数(赠1点)', '10点数(赠1点)', '{\"point\":10,\"gift\":1}', '2025-01-24 17:08:58', '2025-01-24 17:09:28', '2025-01-24 17:38:58');
INSERT INTO `orders` VALUES (155, 307, '', '1882772598916386816', 'vip', 'wechat', '', NULL, 2, 'vip:month:1', '1个月', 0.01, 1, 0.01, 0.01, '', 'completed', '1个月(赠10点)', '1个月', '{\"month\":1,\"point\":10}', '2025-01-24 20:49:02', '2025-01-24 20:49:14', '2025-01-24 21:19:02');
INSERT INTO `orders` VALUES (156, 307, '', '1882779392032116736', 'point', 'wechat', '', NULL, 9, 'point:1000', '1000点数', 100.00, 1, 100.00, 100.00, '', 'closed', '1000点数(赠100点)', '1000点数(赠100点)', '{\"point\":1000,,\"gift\":100}', '2025-01-24 21:16:02', '2025-01-24 21:47:02', '2025-01-24 21:46:02');
INSERT INTO `orders` VALUES (157, 307, '', '1883157395786436608', 'product', 'wechat', '', NULL, 5, 'vip:month:12', '12个月', 159.00, 1, 159.00, 159.00, '', 'closed', '12个月(赠120点)', '12个月', '{\"month\":12\",point\":120}', '2025-01-25 22:18:05', '2025-01-25 22:49:02', '2025-01-25 22:48:05');
INSERT INTO `orders` VALUES (158, 345, '', '1887164865328451584', 'vip', 'wechat', '', NULL, 5, 'vip:month:12', '12个月', 159.00, 1, 159.00, 159.00, '', 'closed', '12个月(赠120点)', '12个月', '{\"month\":12\",point\":120}', '2025-02-05 23:42:20', '2025-02-06 00:13:02', '2025-02-06 00:12:20');
INSERT INTO `orders` VALUES (159, 346, '', '1887164870806212608', 'vip', 'wechat', '', NULL, 5, 'vip:month:12', '12个月', 159.00, 1, 159.00, 159.00, '', 'closed', '12个月(赠120点)', '12个月', '{\"month\":12\",point\":120}', '2025-02-05 23:42:22', '2025-02-06 00:13:02', '2025-02-06 00:12:22');
INSERT INTO `orders` VALUES (160, 344, '', '1887164910664683520', 'point', 'wechat', '', NULL, 9, 'point:1000', '1000点数', 100.00, 1, 100.00, 100.00, '', 'closed', '1000点数(赠100点)', '1000点数(赠100点)', '{\"point\":1000,,\"gift\":100}', '2025-02-05 23:42:31', '2025-02-06 00:13:02', '2025-02-06 00:12:31');
INSERT INTO `orders` VALUES (161, 345, '', '1887165013727121408', 'point', 'wechat', '', NULL, 9, 'point:1000', '1000点数', 100.00, 1, 100.00, 100.00, '', 'closed', '1000点数(赠100点)', '1000点数(赠100点)', '{\"point\":1000,,\"gift\":100}', '2025-02-05 23:42:56', '2025-02-06 00:13:02', '2025-02-06 00:12:56');
INSERT INTO `orders` VALUES (162, 346, '', '1887165018760286208', 'point', 'wechat', '', NULL, 9, 'point:1000', '1000点数', 100.00, 1, 100.00, 100.00, '', 'closed', '1000点数(赠100点)', '1000点数(赠100点)', '{\"point\":1000,,\"gift\":100}', '2025-02-05 23:42:57', '2025-02-06 00:13:02', '2025-02-06 00:12:57');
INSERT INTO `orders` VALUES (163, 344, '', '1887165024338710528', 'vip', 'wechat', '', NULL, 5, 'vip:month:12', '12个月', 159.00, 1, 159.00, 159.00, '', 'closed', '12个月(赠120点)', '12个月', '{\"month\":12\",point\":120}', '2025-02-05 23:42:58', '2025-02-06 00:13:02', '2025-02-06 00:12:58');
INSERT INTO `orders` VALUES (164, 347, '', '1887309780783992832', 'vip', 'wechat', '', NULL, 5, 'vip:month:12', '12个月', 159.00, 1, 159.00, 159.00, '', 'closed', '12个月(赠120点)', '12个月', '{\"month\":12\",point\":120}', '2025-02-06 09:18:11', '2025-02-06 09:49:02', '2025-02-06 09:48:11');
INSERT INTO `orders` VALUES (165, 345, '', '1887309782503657472', 'vip', 'wechat', '', NULL, 5, 'vip:month:12', '12个月', 159.00, 1, 159.00, 159.00, '', 'closed', '12个月(赠120点)', '12个月', '{\"month\":12\",point\":120}', '2025-02-06 09:18:11', '2025-02-06 09:49:02', '2025-02-06 09:48:11');
INSERT INTO `orders` VALUES (166, 347, '', '1887309826954891264', 'point', 'wechat', '', NULL, 9, 'point:1000', '1000点数', 100.00, 1, 100.00, 100.00, '', 'closed', '1000点数(赠100点)', '1000点数(赠100点)', '{\"point\":1000,,\"gift\":100}', '2025-02-06 09:18:22', '2025-02-06 09:49:02', '2025-02-06 09:48:22');
INSERT INTO `orders` VALUES (167, 347, '', '1887309929006501888', 'point', 'wechat', '', NULL, 9, 'point:1000', '1000点数', 100.00, 1, 100.00, 100.00, '', 'closed', '1000点数(赠100点)', '1000点数(赠100点)', '{\"point\":1000,,\"gift\":100}', '2025-02-06 09:18:46', '2025-02-06 09:49:02', '2025-02-06 09:48:46');
INSERT INTO `orders` VALUES (168, 345, '', '1887309930797469696', 'point', 'wechat', '', NULL, 9, 'point:1000', '1000点数', 100.00, 1, 100.00, 100.00, '', 'closed', '1000点数(赠100点)', '1000点数(赠100点)', '{\"point\":1000,,\"gift\":100}', '2025-02-06 09:18:47', '2025-02-06 09:49:02', '2025-02-06 09:48:47');
INSERT INTO `orders` VALUES (169, 347, '', '1887309940633112576', 'vip', 'wechat', '', NULL, 5, 'vip:month:12', '12个月', 159.00, 1, 159.00, 159.00, '', 'closed', '12个月(赠120点)', '12个月', '{\"month\":12\",point\":120}', '2025-02-06 09:18:49', '2025-02-06 09:49:02', '2025-02-06 09:48:49');

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产品名称',
  `type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产品类型',
  `price` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '单价',
  `market_price` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '市场价',
  `status` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Y' COMMENT '状态 Y上架 N下架',
  `order` bigint NULL DEFAULT NULL COMMENT '排序',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产品代码',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '关键词',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '描述',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '详情',
  `extra` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '额外信息',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of products
-- ----------------------------
INSERT INTO `products` VALUES (1, '首页广告A', 'ad', 100.00, 200.00, 'Y', 0, 'home-a', NULL, '', NULL, '<p>首页广告A</p>', NULL, '2023-05-14 01:15:35', '2023-05-14 01:15:52');
INSERT INTO `products` VALUES (2, '1个月', 'vip', 19.99, 19.99, 'Y', NULL, 'vip:month:1', '1个月(赠10点)', '', '1个月', NULL, '{\"month\":1,\"point\":10}', '2025-01-07 22:33:31', '2025-01-07 22:33:31');
INSERT INTO `products` VALUES (3, '3个月', 'vip', 50.00, 59.00, 'Y', NULL, 'vip:month:3', '3个月(赠30点)', '', '3个月', NULL, '{\"month\":3\",point\":30}', '2025-01-07 22:33:31', '2025-01-07 22:33:31');
INSERT INTO `products` VALUES (4, '6个月', 'vip', 99.00, 119.00, 'Y', NULL, 'vip:month:6', '6个月(赠60点)', '', '6个月', NULL, '{\"month\":6\",point\":60}', '2025-01-07 22:33:31', '2025-01-07 22:33:31');
INSERT INTO `products` VALUES (5, '12个月', 'vip', 159.00, 239.00, 'Y', NULL, 'vip:month:12', '12个月(赠120点)', '', '12个月', NULL, '{\"month\":12\",point\":120}', '2025-01-07 22:33:31', '2025-01-07 22:33:31');
INSERT INTO `products` VALUES (6, '10点数', 'point', 0.01, 1.00, 'Y', NULL, 'point:10', '10点数(赠1点)', '', '10点数(赠1点)', NULL, '{\"point\":10,\"gift\":1}', '2025-01-24 15:06:14', '2025-01-24 15:06:16');
INSERT INTO `products` VALUES (7, '100点数', 'point', 10.00, 10.00, 'Y', NULL, 'point:100', '100点数(赠10点)', '', '100点数(赠10点)', NULL, '{\"point\":100,\"gift\":10}', '2025-01-24 15:06:14', '2025-01-24 15:06:16');
INSERT INTO `products` VALUES (8, '500点数', 'point', 50.00, 50.00, 'Y', NULL, 'point:500', '500点数(赠50点)', '', '500点数(赠50点)', NULL, '{\"point\":500,\"gift\":50}', '2025-01-24 15:06:14', '2025-01-24 15:06:16');
INSERT INTO `products` VALUES (9, '1000点数', 'point', 100.00, 100.00, 'Y', NULL, 'point:1000', '1000点数(赠100点)', '', '1000点数(赠100点)', NULL, '{\"point\":1000,,\"gift\":100}', '2025-01-24 15:06:14', '2025-01-24 15:06:16');

-- ----------------------------
-- Table structure for search_logs
-- ----------------------------
DROP TABLE IF EXISTS `search_logs`;
CREATE TABLE `search_logs`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '搜索关键词',
  `count` bigint NULL DEFAULT NULL COMMENT '搜索次数',
  `sort` bigint NULL DEFAULT 0 COMMENT '排序',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of search_logs
-- ----------------------------
INSERT INTO `search_logs` VALUES (1, '请假事由', 11, 0, '2025-01-20 18:10:48', '2025-02-14 13:43:26');
INSERT INTO `search_logs` VALUES (2, '融资计划书', 2, 0, '2025-01-30 03:26:58', '2025-02-05 23:50:55');
INSERT INTO `search_logs` VALUES (3, '音频合成', 1, 0, '2025-02-05 23:50:05', '2025-02-05 23:50:05');
INSERT INTO `search_logs` VALUES (4, '个人简历', 1, 0, '2025-02-05 23:50:28', '2025-02-05 23:50:28');
INSERT INTO `search_logs` VALUES (5, '简历', 3, 0, '2025-02-05 23:50:32', '2025-02-11 15:38:22');

-- ----------------------------
-- Table structure for task_stat
-- ----------------------------
DROP TABLE IF EXISTS `task_stat`;
CREATE TABLE `task_stat`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `task_id` bigint NOT NULL COMMENT '任务ID',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务唯一编号',
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务类型 none-一次性,daily,weekly,monthly,none-无限制',
  `today_date` int NULL DEFAULT NULL COMMENT '今日',
  `week_date` int NULL DEFAULT NULL COMMENT '本周',
  `month_date` int NULL DEFAULT NULL COMMENT '本月',
  `year_date` int NULL DEFAULT NULL COMMENT '今年',
  `today_count` int NULL DEFAULT 0 COMMENT '今日完成次数',
  `week_count` int NULL DEFAULT 0 COMMENT '本周完成次数',
  `month_count` int NULL DEFAULT 0 COMMENT '本月完成次数',
  `year_count` int NULL DEFAULT 0 COMMENT '今年完成次数',
  `count` int NULL DEFAULT 0 COMMENT '总共完成次数',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_user_id_code`(`user_id` ASC, `code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 133 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of task_stat
-- ----------------------------

-- ----------------------------
-- Table structure for tasks
-- ----------------------------
DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务唯一编号',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务标题',
  `launch_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '启动按钮标题',
  `launch_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '启动链接地址',
  `launch_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '启动方式 navigateTo,redirectTo,reLaunch,switchTab,request',
  `completed_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务完成后显示标题',
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务类型 none-一次性,daily,weekly,monthly,none-无限制',
  `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Y' COMMENT '任务状态 Y-显示 N-隐藏',
  `app_status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Y' COMMENT 'APP任务启用 Y-显示 N-隐藏',
  `reward_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '奖励类型 money:余额 gold: 金币 score:积分',
  `reward_amount` decimal(10, 4) NULL DEFAULT NULL COMMENT '任务奖励数量',
  `reward_count` int NULL DEFAULT 0 COMMENT '任务次数限制',
  `reward_icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '奖品图标',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务图标',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '详情',
  `payload` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '任务载荷数据',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_code`(`code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of tasks
-- ----------------------------
INSERT INTO `tasks` VALUES (1, 'check_in:daily', '每日签到', '签到', '/pages/task/check-in', 'navigateTo', '已签到', 'daily', 'Y', 'Y', 'score', 150.0000, 1, '/assets/icons/score.svg', '/assets/task-icons/check-in.svg', '每日签到奖励', NULL, NULL, '2025-01-02 10:22:47', '2025-01-02 10:22:49');
INSERT INTO `tasks` VALUES (2, 'post:release', '发布文章', '去写作', NULL, NULL, NULL, 'none', 'Y', 'N', 'score', 500.0000, 9999, '/assets/icons/score.svg', '/assets/task-icons/post.svg', '发布文章奖励', NULL, NULL, '2025-01-02 16:00:20', '2025-01-02 16:00:22');
INSERT INTO `tasks` VALUES (3, 'wechat:subscribe', '关注微信公众号', '去关注', NULL, NULL, NULL, 'once', 'Y', 'N', 'score', 3000.0000, 1, '/assets/icons/score.svg', '/assets/task-icons/follow.svg', '关注微信公众号奖励', NULL, NULL, '2025-01-02 16:00:27', '2025-01-02 16:00:29');
INSERT INTO `tasks` VALUES (4, 'setting:email', '绑定email账号', '去设置', NULL, NULL, NULL, 'once', 'Y', 'N', 'score', 3000.0000, 1, '/assets/icons/score.svg', '/assets/task-icons/email.svg', '设置email账号奖励', NULL, NULL, '2025-01-02 16:05:04', '2025-01-02 16:05:06');
INSERT INTO `tasks` VALUES (5, 'user:register', '新人红包', '领取', '/task/handler', 'request', '已获得', 'once', 'Y', 'Y', 'point', 10.0000, 1, '/assets/icons/ticket.svg', '/assets/task-icons/hongbao.svg', '新用户赠送10次点数', NULL, NULL, '2025-01-12 15:28:38', '2025-01-12 15:28:40');
INSERT INTO `tasks` VALUES (6, 'tool:used-text:once', '体验文章创作', '去体验', '/pages/index/index', 'switchTab', '已体验', 'once', 'Y', 'Y', 'score', 1000.0000, 1, '/assets/icons/score.svg', '/assets/task-icons/text.svg', '体验文章创作', NULL, NULL, '2025-01-12 16:13:58', '2025-01-12 16:14:00');
INSERT INTO `tasks` VALUES (7, 'user:invite', '分享邀请', '去邀请', '/pages/user/invite', 'navigateTo', '去邀请', 'none', 'Y', 'Y', 'score', 5000.0000, 9999, '/assets/icons/score.svg', '/assets/task-icons/invite.svg', '邀请好友注册奖励', NULL, NULL, '2025-01-12 22:39:31', '2025-01-12 22:39:33');
INSERT INTO `tasks` VALUES (8, 'vip:gift:daily', '会员每日礼包', '领取', '/task/handler', 'request', '已领取', 'daily', 'Y', 'Y', 'point', 10.0000, 1, '/assets/icons/ticket.svg', '/assets/task-icons/vip-gift.svg', '会员每日礼包', NULL, NULL, '2025-01-24 19:32:38', '2025-01-24 19:32:41');

-- ----------------------------
-- Table structure for taxonomy
-- ----------------------------
DROP TABLE IF EXISTS `taxonomy`;
CREATE TABLE `taxonomy`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级ID',
  `taxonomy` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'category,navigation,doc',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `slug` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Slug',
  `order` int NULL DEFAULT 0 COMMENT '排序',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '关键词',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '描述',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '详情',
  `status` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Y' COMMENT '显示状态 Y-显示 N-隐藏',
  `icon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `taxonomy`(`taxonomy` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 72 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of taxonomy
-- ----------------------------
INSERT INTO `taxonomy` VALUES (1, 0, 'tool', '查询工具', 'cx', 0, NULL, '', '', NULL, 'Y', NULL, '2023-04-15 17:40:08', '2023-04-15 17:40:10');
INSERT INTO `taxonomy` VALUES (12, 0, 'tool', '代码格式化', 'gs', 0, NULL, '', '', NULL, 'Y', NULL, '2023-04-15 17:40:08', '2023-04-15 17:40:10');
INSERT INTO `taxonomy` VALUES (18, 0, 'tool', '加密/解密', 'jm', 0, NULL, '', '', NULL, 'Y', NULL, '2023-04-15 17:40:08', '2023-04-15 17:40:10');
INSERT INTO `taxonomy` VALUES (30, 0, 'tool', '在线测试工具', 'test', 0, NULL, '', '', NULL, 'Y', NULL, '2023-04-15 17:40:08', '2023-04-15 17:40:10');
INSERT INTO `taxonomy` VALUES (36, 0, 'tool', '二维码&条形码', 'qrcode', 0, NULL, '', '', NULL, 'Y', NULL, '2023-04-15 17:40:08', '2023-04-15 17:40:10');
INSERT INTO `taxonomy` VALUES (39, 0, 'tool', '图像处理', 'image', 0, NULL, '', '', NULL, 'Y', NULL, '2023-04-15 17:40:08', '2023-04-15 17:40:10');
INSERT INTO `taxonomy` VALUES (46, 0, 'tool', '文本处理', 'txt', 0, NULL, '', '', NULL, 'Y', NULL, '2023-04-15 17:40:08', '2023-04-15 17:40:10');
INSERT INTO `taxonomy` VALUES (50, 0, 'tool', '常用词典', 'cd', 0, NULL, '', '', NULL, 'Y', NULL, '2023-04-15 17:40:08', '2023-04-15 17:40:10');
INSERT INTO `taxonomy` VALUES (51, 0, 'tag', 'php', 'php', 0, 'php', 'php', 'php', 'PHP是世界上最好的语言', 'Y', 'icon-php', '2023-05-13 20:10:01', '2024-12-27 10:40:08');
INSERT INTO `taxonomy` VALUES (52, 0, 'tag', 'Vue', 'vue', 2, 'Vue', 'Vue', 'Vue', 'Vue', 'Y', 'icon-Vue', '2023-06-12 21:11:08', '2024-12-27 10:43:09');
INSERT INTO `taxonomy` VALUES (53, 0, 'navigation', '工具', 'tool.index', 1, '工具箱', '工具箱', '工具箱', '工具箱', 'Y', NULL, '2023-11-18 00:18:43', '2023-11-18 00:43:37');
INSERT INTO `taxonomy` VALUES (54, 0, 'navigation', '标签', 'web.tag', 2, '标签', '标签', '标签', '标签', 'N', NULL, '2023-11-18 00:19:54', '2024-03-03 17:11:59');
INSERT INTO `taxonomy` VALUES (55, 0, 'navigation', '手册', 'web.manual', 3, '参考手册', '参考手册', '参考手册', '参考手册', 'Y', NULL, '2023-11-18 00:44:42', '2023-11-18 00:44:42');
INSERT INTO `taxonomy` VALUES (59, 0, 'tool', '编码转换', 'convert', 0, '编码转换', '编码转换', '编码转换', '<p>编码转换</p>', 'Y', NULL, '2023-11-30 22:58:04', '2023-11-30 22:58:31');
INSERT INTO `taxonomy` VALUES (60, 0, 'tool', '网站信息', 'website', 0, '网站信息', '网站信息', '网站信息', '<p>网站信息</p>', 'Y', NULL, '2023-11-30 23:09:11', '2023-11-30 23:10:11');
INSERT INTO `taxonomy` VALUES (61, 0, 'tool', '办公管理', 'office', NULL, '办公管理', '办公管理', '办公管理', NULL, 'Y', NULL, '2023-12-03 00:04:47', '2023-12-03 00:04:47');
INSERT INTO `taxonomy` VALUES (62, 0, 'navigation', '图标', 'web.icon', 4, '字体图标', '字体图标', '字体图标', '字体图标', 'Y', NULL, '2023-12-30 21:17:31', '2023-12-30 21:18:38');
INSERT INTO `taxonomy` VALUES (63, 0, 'navigation', '开发文档', 'web.doc', 0, NULL, NULL, NULL, NULL, 'N', NULL, '2024-03-13 22:52:28', '2024-03-13 22:52:58');
INSERT INTO `taxonomy` VALUES (64, 0, 'navigation', '数据API', 'web.open-api', 5, '开放API', 'Open API，免费数据API', '免费数据API，提供大量免费 API 接口服务，接口稳定响应快，无调用次数限制，速率60次/分钟，请合理利用资源。目前提供 IP查询API，IP批量查询API，万年历查询API，福利彩票开奖查询API，全国行政区域查询API，成语查询API，成语接龙查询API，域名whois信息查询API，域名ssl证书信息查询API，网址安全检测API，国家法定节日放假安排查询API...', NULL, 'Y', NULL, '2024-05-26 16:09:54', '2024-06-19 01:12:03');
INSERT INTO `taxonomy` VALUES (65, 0, 'navigation', '排行榜', 'home.ranking', 6, '排行榜', '排行榜', '排行榜', NULL, 'Y', NULL, '2024-12-26 01:46:21', '2024-12-26 01:46:23');
INSERT INTO `taxonomy` VALUES (66, 0, 'mini-app', '工作', 'work', 0, '工作', '', '', NULL, 'Y', NULL, '2025-01-01 16:37:11', '2025-01-01 16:37:14');
INSERT INTO `taxonomy` VALUES (67, 0, 'mini-app', '写作', 'write', 0, '写作', '', '', NULL, 'Y', NULL, '2025-01-01 16:37:11', '2025-01-01 16:37:14');
INSERT INTO `taxonomy` VALUES (68, 0, 'mini-app', '个人', 'person', 0, '个人', '', '', NULL, 'Y', NULL, '2025-01-01 16:37:11', '2025-01-01 16:37:14');
INSERT INTO `taxonomy` VALUES (69, 0, 'mini-app', '媒体', 'media', 0, '媒体', '', '', NULL, 'Y', NULL, '2025-01-01 16:37:11', '2025-01-01 16:37:14');
INSERT INTO `taxonomy` VALUES (70, 0, 'mini-app', '教育', 'edu', 0, '教育', '', '', NULL, 'Y', NULL, '2025-01-01 16:37:11', '2025-01-01 16:37:14');
INSERT INTO `taxonomy` VALUES (71, 0, 'mini-app', '商业', 'biz', 0, '商业', '', '', NULL, 'Y', NULL, '2025-01-01 16:37:11', '2025-01-01 16:37:14');

-- ----------------------------
-- Table structure for taxonomy_meta
-- ----------------------------
DROP TABLE IF EXISTS `taxonomy_meta`;
CREATE TABLE `taxonomy_meta`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `taxonomy_id` int NULL DEFAULT NULL COMMENT '归类ID',
  `color` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字体颜色',
  `bg_color` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '背景色彩',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of taxonomy_meta
-- ----------------------------
INSERT INTO `taxonomy_meta` VALUES (1, 66, NULL, '#d1ebff');
INSERT INTO `taxonomy_meta` VALUES (2, 67, NULL, '#ffd9c9');
INSERT INTO `taxonomy_meta` VALUES (3, 68, NULL, '#c5efdd');
INSERT INTO `taxonomy_meta` VALUES (4, 69, NULL, '#d7daf7');
INSERT INTO `taxonomy_meta` VALUES (5, 70, NULL, '#ffedca');
INSERT INTO `taxonomy_meta` VALUES (6, 71, NULL, '#ffdee4');

-- ----------------------------
-- Table structure for user_activity_logs
-- ----------------------------
DROP TABLE IF EXISTS `user_activity_logs`;
CREATE TABLE `user_activity_logs`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `uuid` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'UUID',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作代码',
  `amount` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '扣费金额',
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `content_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '响应内容类型 text,image,audio,video',
  `content` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '用户展示结果内容',
  `form_schemas` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '表单信息',
  `request_body` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求内容',
  `response_body` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应内容',
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'pending' COMMENT '状态: pending-已提交, fail-处理失败, success-成功, refunded-已退费',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 108 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_activity_logs
-- ----------------------------

-- ----------------------------
-- Table structure for user_balance_logs
-- ----------------------------
DROP TABLE IF EXISTS `user_balance_logs`;
CREATE TABLE `user_balance_logs`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '余额代码',
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'increase-增加 decrease-减少\n\n',
  `balance` decimal(10, 2) NULL DEFAULT NULL COMMENT '当前余额',
  `amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '消费金额',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 91 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_balance_logs
-- ----------------------------

-- ----------------------------
-- Table structure for user_check_in
-- ----------------------------
DROP TABLE IF EXISTS `user_check_in`;
CREATE TABLE `user_check_in`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  `count` int UNSIGNED NULL DEFAULT 0 COMMENT '连续签到次数',
  `date` int NULL DEFAULT 0 COMMENT '签到日期 格式20250101',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_user_id_date`(`user_id` ASC, `date` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 118 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_check_in
-- ----------------------------

-- ----------------------------
-- Table structure for user_invite_logs
-- ----------------------------
DROP TABLE IF EXISTS `user_invite_logs`;
CREATE TABLE `user_invite_logs`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint NULL DEFAULT NULL COMMENT '邀请人UID',
  `invite_user_id` bigint NULL DEFAULT NULL COMMENT '被邀人UID',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_invite_logs
-- ----------------------------

-- ----------------------------
-- Table structure for user_login_logs
-- ----------------------------
DROP TABLE IF EXISTS `user_login_logs`;
CREATE TABLE `user_login_logs`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  `user_agent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'user-agent',
  `client_ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户端IP',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_login_logs
-- ----------------------------

-- ----------------------------
-- Table structure for user_meta
-- ----------------------------
DROP TABLE IF EXISTS `user_meta`;
CREATE TABLE `user_meta`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  `invite_id` bigint NULL DEFAULT NULL COMMENT '邀请码MMID',
  `invite_qrcode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邀请二维码',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `identity_card` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '身份证号码',
  `alipay_account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '支付宝账号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 60 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_meta
-- ----------------------------

-- ----------------------------
-- Table structure for user_payment_account
-- ----------------------------
DROP TABLE IF EXISTS `user_payment_account`;
CREATE TABLE `user_payment_account`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `pay_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收款账号类型',
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '账号',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户姓名',
  `bank_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '银行名称',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_user_id_pay_type`(`user_id` ASC, `pay_type` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_payment_account
-- ----------------------------

-- ----------------------------
-- Table structure for user_point_logs
-- ----------------------------
DROP TABLE IF EXISTS `user_point_logs`;
CREATE TABLE `user_point_logs`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作类型 increase-增加 decrease-减少\n',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '点数来源代码',
  `amount` bigint NULL DEFAULT 0 COMMENT '点数数量',
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_id_source`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 185 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_point_logs
-- ----------------------------

-- ----------------------------
-- Table structure for user_reward_logs
-- ----------------------------
DROP TABLE IF EXISTS `user_reward_logs`;
CREATE TABLE `user_reward_logs`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作类型 increase-增加 decrease-减少\n',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '奖金来源代码',
  `amount` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '奖金数量',
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_reward_logs
-- ----------------------------

-- ----------------------------
-- Table structure for user_score_logs
-- ----------------------------
DROP TABLE IF EXISTS `user_score_logs`;
CREATE TABLE `user_score_logs`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作类型 increase-增加 decrease-减少\n',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '积分来源代码',
  `amount` bigint NULL DEFAULT 0 COMMENT '积分数量',
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_id_source`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 209 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_score_logs
-- ----------------------------

-- ----------------------------
-- Table structure for user_stat
-- ----------------------------
DROP TABLE IF EXISTS `user_stat`;
CREATE TABLE `user_stat`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务唯一编号',
  `today_date` int NULL DEFAULT NULL COMMENT '今日',
  `week_date` int NULL DEFAULT NULL COMMENT '本周',
  `month_date` int NULL DEFAULT NULL COMMENT '本月',
  `year_date` int NULL DEFAULT NULL COMMENT '今年',
  `today_count` int NULL DEFAULT 0 COMMENT '今日完成次数',
  `week_count` int NULL DEFAULT 0 COMMENT '本周完成次数',
  `month_count` int NULL DEFAULT 0 COMMENT '本月完成次数',
  `year_count` int NULL DEFAULT 0 COMMENT '今年完成次数',
  `count` int NULL DEFAULT 0 COMMENT '总共完成次数',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_stat
-- ----------------------------

-- ----------------------------
-- Table structure for user_transfer
-- ----------------------------
DROP TABLE IF EXISTS `user_transfer`;
CREATE TABLE `user_transfer`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  `amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '提现转账金额',
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收款账号',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `order_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '转账单号',
  `pay_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收款方式 wechat;alipay',
  `pay_order_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '支付单号',
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'pending' COMMENT '转账状态 pending,paid,fail',
  `pay_time` timestamp NULL DEFAULT NULL COMMENT '转账时间',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户提现转账' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_transfer
-- ----------------------------

-- ----------------------------
-- Table structure for user_vip
-- ----------------------------
DROP TABLE IF EXISTS `user_vip`;
CREATE TABLE `user_vip`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  `active` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'N' COMMENT '是否有效 Y-有效 N-过期',
  `expire_time` timestamp NULL DEFAULT NULL COMMENT '会员到期时间',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_vip
-- ----------------------------

-- ----------------------------
-- Table structure for user_wechat
-- ----------------------------
DROP TABLE IF EXISTS `user_wechat`;
CREATE TABLE `user_wechat`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NULL DEFAULT 0 COMMENT '关联用户ID',
  `openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '微信公众号openID',
  `subscribe` tinyint UNSIGNED NULL DEFAULT 0 COMMENT '是否关注 0-否 1-是',
  `is_account` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'N' COMMENT '是否绑定账号 Y-是 N-否',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `openid`(`openid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 80 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_wechat
-- ----------------------------

-- ----------------------------
-- Table structure for user_wechat_app
-- ----------------------------
DROP TABLE IF EXISTS `user_wechat_app`;
CREATE TABLE `user_wechat_app`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NULL DEFAULT 0 COMMENT '关联用户ID',
  `openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '微信小程序openid',
  `unionid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '微信开放平台unionid',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `openid`(`openid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 65 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_wechat_app
-- ----------------------------

-- ----------------------------
-- Table structure for user_wechat_scan
-- ----------------------------
DROP TABLE IF EXISTS `user_wechat_scan`;
CREATE TABLE `user_wechat_scan`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '微信公众号openID',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '场景KEY，qrscene_为前缀，后面为二维码的参数值',
  `ticket` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '二维码的ticket，可用来换取二维码图片',
  `status` tinyint NULL DEFAULT 0 COMMENT '是否可用 0 可用 1已经',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `openid`(`openid` ASC) USING BTREE,
  INDEX `ticket`(`ticket` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 243 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_wechat_scan
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `mmid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '唯一ID编号',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `remember_token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `balance` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '账户余额',
  `score` bigint NULL DEFAULT 0 COMMENT '积分',
  `reward` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '奖励金额',
  `point` bigint NULL DEFAULT 0 COMMENT '点数',
  `gender` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'female' COMMENT '性别 Female-女 Male-男性',
  `is_admin` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'N' COMMENT '是否管理员 Y-是 N-否',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `website` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '个人网站',
  `bio` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '个人介绍',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `users_email_unique`(`email` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 403 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
