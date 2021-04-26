/*
Navicat MySQL Data Transfer

Source Server         : dev
Source Server Version : 50732
Source Host           : 8.130.28.97:33106
Source Database       : db_promethues

Target Server Type    : MYSQL
Target Server Version : 50732
File Encoding         : 65001

Date: 2021-04-21 11:22:46
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for rules
-- ----------------------------
DROP TABLE IF EXISTS `t_rules`;
CREATE TABLE `t_rules` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '规则ID',
  `rule_name` varchar(255) DEFAULT NULL COMMENT '规则名称',
  `rule_type` varchar(20) DEFAULT NULL COMMENT '规则类型，{record,alert}',
  `rule_fn` varchar(255) DEFAULT NULL COMMENT '规则文件名(为兼容Prometheus配置文件管理模式)',
  `rule_gn` varchar(255) DEFAULT NULL COMMENT '规则组名(为兼容Prometheus配置文件管理模式)',
  `rule_interval` int(10) DEFAULT NULL COMMENT '规则运算间隔',
  `rule_alert` varchar(60) DEFAULT NULL COMMENT '规则告警',
  `rule_expr` text DEFAULT NULL COMMENT '规则表达式',
  `rule_for` varchar(20) DEFAULT NULL COMMENT '规则持续时长阈值',
  `note` varchar(255) DEFAULT '' COMMENT '规则说明',
  `state` int(4) NOT NULL DEFAULT '1' COMMENT '状态, 1-有效;2-无效',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近一次更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `t_rule_labels`;
CREATE TABLE `t_rule_labels` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '规则标签ID',
  `rule_id` int(10) NOT NULL COMMENT '规则ID',
  `label_key` varchar(255) DEFAULT NULL COMMENT '规则标签名称',
  `label_value` varchar(255) DEFAULT NULL COMMENT '规则标签值',
  `state` int(4) NOT NULL DEFAULT '1' COMMENT '状态, 1-有效;2-无效',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近一次更新时间',  
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `t_rule_annotations`;
CREATE TABLE `t_rule_annotations` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '规则评注ID',
  `rule_id` int(10) NOT NULL COMMENT '规则ID',
  `annotation_key` varchar(255) DEFAULT NULL COMMENT '规则评注名称',
  `annotation_value` varchar(255) DEFAULT NULL COMMENT '规则评注值',
  `state` int(4) NOT NULL DEFAULT '1' COMMENT '状态, 1-有效;2-无效',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近一次更新时间',  
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


-- 用于存储全局性配置项
--  -- rule_lastreloadedat:规则上次reload时间
--      -- prom自动reload并更新改值
--      -- reload时根据 rule_lastreloadedat > t_rules.updated_at 来提取本次要更新的rules
--      -- 启动时会将忽略该值，并强制加载全部有效 rules
DROP TABLE IF EXISTS `t_configs`;
CREATE TABLE `t_configs` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '配置项ID',
  `config_key` varchar(255) DEFAULT NULL COMMENT '配置项名称',
  `config_value` varchar(255) DEFAULT NULL COMMENT '配置项值',
  `state` int(4) NOT NULL DEFAULT '1' COMMENT '状态, 1-有效;2-无效',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近一次更新时间',  
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


