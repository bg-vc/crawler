
CREATE DATABASE `crawler` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `douban_movie` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(128) DEFAULT '' COMMENT '标题',
  `subtitle` varchar(256) DEFAULT '' COMMENT '副标题',
  `other` varchar(512) DEFAULT '' COMMENT '其他',
  `desc_one` varchar(512) DEFAULT '' COMMENT '描述1',
  `desc_two` varchar(512) DEFAULT '' COMMENT '描述2',
  `score` decimal(10,2) NOT NULL DEFAULT '0' COMMENT '评分',
  `comment` int(11) unsigned DEFAULT '0' COMMENT '评价',
  `quote` varchar(512) DEFAULT '' COMMENT '引用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='豆瓣电影';


CREATE TABLE `qczj_car` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `city` varchar(255) DEFAULT '' COMMENT '城市',
  `title` varchar(255) DEFAULT '' COMMENT '标题',
  `price` decimal(10,2) NOT NULL DEFAULT '0'  COMMENT '价格',
  `kilometer` decimal(10,2) NOT NULL DEFAULT '0' COMMENT '公里',
  `date` int(10) NOT NULL DEFAULT '0' COMMENT '日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='汽车之家二手车';
