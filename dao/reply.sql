DROP TABLE IF EXISTS `reply_00`;
CREATE TABLE `reply_00` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `source_id` int(11) unsigned NOT NULL COMMENT '资源id',
  `type_id` tinyint(4) unsigned NOT NULL COMMENT '类型id',
  `mid` int(11) unsigned NOT NULL COMMENT '用户id',
  `comment` text NOT NULL COMMENT '内容',
  `parent_id` int(11) unsigned NOT NULL COMMENT '父级id',
  `path` text NOT NULL COMMENT '楼层关系',
  `state` tinyint(4) unsigned DEFAULT 0 NOT NULL COMMENT '状态',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '评论表';

DROP TABLE IF EXISTS `reply_01`;
CREATE TABLE `reply_01` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `source_id` int(11) unsigned NOT NULL COMMENT '资源id',
  `type_id` tinyint(4) unsigned NOT NULL COMMENT '类型id',
  `mid` int(11) unsigned NOT NULL COMMENT '用户id',
  `comment` text NOT NULL COMMENT '内容',
  `parent_id` int(11) unsigned NOT NULL COMMENT '父级id',
  `path` text NOT NULL COMMENT '楼层关系',
  `state` tinyint(4) unsigned DEFAULT 0 NOT NULL COMMENT '状态',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '评论表';

DROP TABLE IF EXISTS `reply_02`;
CREATE TABLE `reply_02` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `source_id` int(11) unsigned NOT NULL COMMENT '资源id',
  `type_id` tinyint(4) unsigned NOT NULL COMMENT '类型id',
  `mid` int(11) unsigned NOT NULL COMMENT '用户id',
  `comment` text NOT NULL COMMENT '内容',
  `parent_id` int(11) unsigned NOT NULL COMMENT '父级id',
  `path` text NOT NULL COMMENT '楼层关系',
  `state` tinyint(4) unsigned DEFAULT 0 NOT NULL COMMENT '状态',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '评论表';

DROP TABLE IF EXISTS `reply_03`;
CREATE TABLE `reply_03` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `source_id` int(11) unsigned NOT NULL COMMENT '资源id',
  `type_id` tinyint(4) unsigned NOT NULL COMMENT '类型id',
  `mid` int(11) unsigned NOT NULL COMMENT '用户id',
  `comment` text NOT NULL COMMENT '内容',
  `parent_id` int(11) unsigned NOT NULL COMMENT '父级id',
  `path` text NOT NULL COMMENT '楼层关系',
  `state` tinyint(4) unsigned DEFAULT 0 NOT NULL COMMENT '状态',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '评论表';


DROP TABLE IF EXISTS `reply_04`;
CREATE TABLE `reply_04` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `source_id` int(11) unsigned NOT NULL COMMENT '资源id',
  `type_id` tinyint(4) unsigned NOT NULL COMMENT '类型id',
  `mid` int(11) unsigned NOT NULL COMMENT '用户id',
  `comment` text NOT NULL COMMENT '内容',
  `parent_id` int(11) unsigned NOT NULL COMMENT '父级id',
  `path` text NOT NULL COMMENT '楼层关系',
  `state` tinyint(4) unsigned DEFAULT 0 NOT NULL COMMENT '状态',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '评论表';

DROP TABLE IF EXISTS `reply_05`;
CREATE TABLE `reply_05` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `source_id` int(11) unsigned NOT NULL COMMENT '资源id',
  `type_id` tinyint(4) unsigned NOT NULL COMMENT '类型id',
  `mid` int(11) unsigned NOT NULL COMMENT '用户id',
  `comment` text NOT NULL COMMENT '内容',
  `parent_id` int(11) unsigned NOT NULL COMMENT '父级id',
  `path` text NOT NULL COMMENT '楼层关系',
  `state` tinyint(4) unsigned DEFAULT 0 NOT NULL COMMENT '状态',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '评论表';

DROP TABLE IF EXISTS `reply_06`;
CREATE TABLE `reply_06` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `source_id` int(11) unsigned NOT NULL COMMENT '资源id',
  `type_id` tinyint(4) unsigned NOT NULL COMMENT '类型id',
  `mid` int(11) unsigned NOT NULL COMMENT '用户id',
  `comment` text NOT NULL COMMENT '内容',
  `parent_id` int(11) unsigned NOT NULL COMMENT '父级id',
  `path` text NOT NULL COMMENT '楼层关系',
  `state` tinyint(4) unsigned DEFAULT 0 NOT NULL COMMENT '状态',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '评论表';

DROP TABLE IF EXISTS `reply_07`;
CREATE TABLE `reply_07` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `source_id` int(11) unsigned NOT NULL COMMENT '资源id',
  `type_id` tinyint(4) unsigned NOT NULL COMMENT '类型id',
  `mid` int(11) unsigned NOT NULL COMMENT '用户id',
  `comment` text NOT NULL COMMENT '内容',
  `parent_id` int(11) unsigned NOT NULL COMMENT '父级id',
  `path` text NOT NULL COMMENT '楼层关系',
  `state` tinyint(4) unsigned DEFAULT 0 NOT NULL COMMENT '状态',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '评论表';

DROP TABLE IF EXISTS `reply_08`;
CREATE TABLE `reply_08` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `source_id` int(11) unsigned NOT NULL COMMENT '资源id',
  `type_id` tinyint(4) unsigned NOT NULL COMMENT '类型id',
  `mid` int(11) unsigned NOT NULL COMMENT '用户id',
  `comment` text NOT NULL COMMENT '内容',
  `parent_id` int(11) unsigned NOT NULL COMMENT '父级id',
  `path` text NOT NULL COMMENT '楼层关系',
  `state` tinyint(4) unsigned DEFAULT 0 NOT NULL COMMENT '状态',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '评论表';

DROP TABLE IF EXISTS `reply_09`;
CREATE TABLE `reply_09` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `source_id` int(11) unsigned NOT NULL COMMENT '资源id',
  `type_id` tinyint(4) unsigned NOT NULL COMMENT '类型id',
  `mid` int(11) unsigned NOT NULL COMMENT '用户id',
  `comment` text NOT NULL COMMENT '内容',
  `parent_id` int(11) unsigned NOT NULL COMMENT '父级id',
  `path` text NOT NULL COMMENT '楼层关系',
  `state` tinyint(4) unsigned DEFAULT 0 NOT NULL COMMENT '状态',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '评论表';
