CREATE DATABASE IF NOT EXISTS social_go;

USE social_go;

DROP TABLE IF EXISTS user;

CREATE TABLE user(
    id int auto_increment primary key,
    name varchar(50) not null,
    username varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(50) not null,
    created_at timestamp default current_timestamp()
) ENGINE=INNODB;