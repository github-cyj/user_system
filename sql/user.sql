CREATE TABLE `user`
(
    `id`          int(11) primary key auto_increment NOT NULL COMMENT '主键id',
    `avatar`      varchar(100)                       NOT NULL DEFAULT '' COMMENT '头像url',
    `username`    varchar(32)                        NOT NULL DEFAULT '' COMMENT '用户名',
    `sex`         tinyint(2)                         NOT NULL DEFAULT '1' COMMENT '性别 1男 2女',
    `tel`         varchar(11)                        NOT NULL DEFAULT '' COMMENT '电话号码',
    `email`       varchar(64)                        NOT NULL DEFAULT '' COMMENT '邮箱',
    `password`    char(32)                           NOT NULL DEFAULT '' COMMENT '密码',
    `create_time` datetime                           NOT NULL COMMENT '创建时间',
    `update_time` datetime                           NOT NULL COMMENT '更新时间',
    `delete_time` datetime                                    DEFAULT NULL COMMENT '删除时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT "用户表";