drop table if exists `user`;

create table `user`
(
    id                 int primary key auto_increment comment 'id',
    username           varchar(20),
    encrypted_password varchar(100),
    nickname           varchar(20),
    avatar             varchar(100),
    status             bit,
    created_at         datetime,
    updated_at         datetime,
    UNIQUE (username)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  AUTO_INCREMENT = 1
  COLLATE = utf8mb4_unicode_ci;