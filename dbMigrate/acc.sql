drop table if exists `sys_user`;

create table `sys_user`
(
    id                 bigint primary key auto_increment comment '自增id',
    user_id            bigint comment '用户id',
    username           varchar(20) comment '用户名',
    encrypted_password varchar(100) comment '密码',
    salt               varchar(64) comment '盐',
    nickname           varchar(20) comment '用户昵称',
    avatar             varchar(100) comment '头像地址',
    status             tinyint comment '用户状态，1正常，0封禁',
    created_at         datetime comment '创建时间',
    updated_at         datetime default null comment '更新时间',
    deleted_at         datetime default null comment '是否被软删除',
    UNIQUE (username),
    unique key `idx_user_id` (`user_id`) using btree
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  AUTO_INCREMENT = 1
  COLLATE = utf8mb4_unicode_ci comment '用户表';

