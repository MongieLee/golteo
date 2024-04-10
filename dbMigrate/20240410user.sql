alter table user
    add column deleted_at datetime default null;

alter table user
    add column salt varchar(64) default null;