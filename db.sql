create database  if not exists tensir_blog;

use tensir_blog;

CREATE TABLE if not exists `tensir_blog_user`(
    `id` int(11) NOT NULL AUTO_INCREMENT comment '自增主键',
    `username` varchar(60) NOT NULL comment '用户名',
    `password` varchar(60) NOT NULL comment '密码',
    `role` varchar(50) DEFAULT NULL comment '角色',
    `ctime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '创建时间',
    `utime` datetime DEFAULT NULL comment '更新时间',
    `avatar` varchar(200) DEFAULT NULL comment '头像',
    `email` varchar(100) DEFAULT NULL comment '邮箱',
    `class_id` int(11) default null  comment '类别id',
    PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=1;

create table if not exists  `tensir_blog_article` (
    `id` int(11) not null auto_increment comment '自增主键',
    `title` varchar(100) not null comment '文章标题',
    `content` text not null comment '文章内容（富文本）',
    `ctime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '文章创建时间',
    `utime` datetime default null comment '文章更新时间',
    `author_id` int(11) not null comment '作者id',
    `status` tinyint comment '文章状态：0草稿，1已提交，2回收站',
    primary key (`id`)
) ENGINE =MyISAM AUTO_INCREMENT=1;

create table if not exists  `tensir_blog_tag` (
    `id` int(11) not null auto_increment comment '自增主键',
    `name` varchar(60) not null comment '标签名',
    `description` tinytext default null comment '标签描述',
    `ctime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '标签创建时间',
    `utime` datetime default null comment '标签更新时间',
    primary key (`id`)
) ENGINE =MyISAM AUTO_INCREMENT=1;

create table if not exists `tensir_blog_article_tag` (
    `article_id` int(11) not null comment '文章id主键',
    `tag_id` int(11) not null comment '标签id主键',
    primary key (`article_id`, `tag_id`)
) ENGINE =MyISAM;

create table if not exists  `tensir_blog_class` (
    `id` int(11) not null auto_increment comment '自增主键',
    `name` varchar(60) not null  comment '类别名称',
    `description` tinytext default null comment '类别描述',
    primary key (`id`)
) ENGINE =MyISAM;