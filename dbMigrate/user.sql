drop table if exists `user`;

create table `user`
(
    id                 bigint primary key auto_increment comment '自增id',
    user_id            bigint comment '用户id',
    username           varchar(20),
    encrypted_password varchar(100),
    salt               varchar(64) comment "盐",
    nickname           varchar(20),
    avatar             varchar(100),
    status             tinyint,
    created_at         datetime,
    updated_at         datetime default null,
    deleted_at         datetime default null,
    UNIQUE (username),
    unique key `idx_user_id` (`user_id`) using btree
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  AUTO_INCREMENT = 1
  COLLATE = utf8mb4_unicode_ci;