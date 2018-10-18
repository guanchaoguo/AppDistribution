## Mysql脚本 #

>脚本说明：app热点分发平台mysql初始化脚本
>以下为具体脚本：
-- Target Server Type    : MYSQL
-- Target Server Version : 50718
-- File Encoding         : 65001

CREATE DATABASE IF NOT EXISTS `AppDistribution` /*!40100 DEFAULT CHARACTER SET utf8mb4 */

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `admin_user`
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) NOT NULL DEFAULT '' COMMENT '后台用户登录名',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '用户密码（加密后的密码，算法为：md5(sha1(value + 盐值))）',
  `salt` varchar(20) NOT NULL DEFAULT '' COMMENT '8-20位密码加密盐值',
  `account` varchar(50) NOT NULL DEFAULT '' COMMENT '用户真实姓名',
  `last_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '账号最后修改时间',
  `last_login_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登录时间',
  `last_login_ip` char(15) DEFAULT '' COMMENT '最后登录IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=26 DEFAULT CHARSET=utf8 COMMENT='后台管理员账户表';

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES ('1', 'admin', 'b78f498e97e6caad5388bed340e98cf4', 'RxFS57aGMD', '开发用户', '2017-11-27 06:14:43', '2018-01-24 16:19:25', '127.0.0.1');

-- ----------------------------
-- Table structure for `app_info`
-- ----------------------------
DROP TABLE IF EXISTS `app_info`;
CREATE TABLE `app_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Id',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '上传者用户Id',
  `logo` varchar(225) NOT NULL DEFAULT '' COMMENT '应用logo',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '应用名称',
  `app_url` varchar(255) NOT NULL DEFAULT '' COMMENT 'app下载地址',
  `type` tinyint(3) NOT NULL DEFAULT '0' COMMENT 'app类型 0 安卓  1 ios 默认安卓 ',
  `versions` varchar(20) NOT NULL DEFAULT '' COMMENT '应用版本号',
  `is_password` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否启动下载密码 0 不启动 1启用',
  `password` char(4) NOT NULL DEFAULT '' COMMENT '四位下载密码',
  `size` int(11) NOT NULL DEFAULT '0' COMMENT '应用文件大小',
  `shot_url` char(5) NOT NULL DEFAULT '' COMMENT '短连接',
  `app_id` varchar(225) NOT NULL DEFAULT '' COMMENT '应用唯一ID码',
  `is_merger` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否合并应用',
  `apk_id` int(11) NOT NULL COMMENT '被合并的安卓id',
  `ipa_id` int(11) NOT NULL COMMENT '被合并的苹果id',
  `ipa_name` varchar(255) NOT NULL COMMENT '被合并的苹果名称',
  `apk_name` varchar(225) NOT NULL DEFAULT '' COMMENT '被合并的安卓名称',
  `desc` varchar(500) NOT NULL DEFAULT '' COMMENT '应用描述',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '应用状态 0 正常 1 删除 2 冻结',
  `updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'app信息更新时间',
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'app 发布时间',
  `last_ip` char(15) NOT NULL DEFAULT '' COMMENT '发布者IP',
  `allow_count` int(11) DEFAULT '0' COMMENT 'app 允许下载的次数',
  `plist` varchar(255) DEFAULT '''''' COMMENT '苹果app文件plist',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=86 DEFAULT CHARSET=utf8mb4 COMMENT='热点app发布;列表';

-- ----------------------------
-- Table structure for `app_statistics`
-- ----------------------------
DROP TABLE IF EXISTS `app_statistics`;
CREATE TABLE `app_statistics` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Id',
  `upolader_id` int(11) NOT NULL DEFAULT '0' COMMENT 'app上传者Id',
  `app_id` int(11) NOT NULL DEFAULT '0' COMMENT '发布appId',
  `app_name` varchar(50) NOT NULL DEFAULT '' COMMENT '发布appId',
  `app_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '发布app类型  安卓 1 苹果2   默认0',
  `scan_count` int(11) DEFAULT '0' COMMENT '浏览次数',
  `down_count` int(11) DEFAULT '1' COMMENT '下载次数',
  `upload_count` int(1) DEFAULT '0' COMMENT '上传次数',
  `created` date NOT NULL DEFAULT '0000-00-00' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=379 DEFAULT CHARSET=utf8mb4 COMMENT='app数据统计';



## Mysql数据库时区配置 （my.cnf文件配置） #

>配置说明：把mysql数据库的时区修改为美国东部时区--西五区
>以下为具体配置内容：

`    [mysqld]
`    default-time_zone = '-5:00'
