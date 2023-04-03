CREATE DATABASE IF NOT EXISTS `ndv`;
GRANT ALL ON `ndv`.* TO 'user'@'%';

USE ndv;

create table ndv.partners
(
    id int not null,
    name varchar(255),
    address_id int not null,
    image varchar(255),
    phone varchar(255),
    mail varchar(255),
    rating tinyint,
    created_at datetime not null,
    last_updated tinyint default 1 not null,
    is_deleted bool
);

create unique index partners_id_uindex
    on ndv.partners(id);

alter table ndv.partners modify id int auto_increment;


create table ndv.address
(
    id int not null,
    name varchar(255),
    created_at datetime not null,
    last_updated tinyint default 1 not null,
    is_deleted bool default false
);