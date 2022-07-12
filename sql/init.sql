CREATE DATABASE IF NOT EXISTS chatdbnew;
USE chatdbnew;
-- DROP TABLE IF EXISTS `users`;
-- CREATE TABLE `users` (  
CREATE TABLE IF NOT EXISTS `users` ( 
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    `username` varchar(255) UNIQUE,
    `password` varchar(255)  NOT NULL,
    `fullname` varchar(255)  NOT NULL,
    `salary` bigint(20) DEFAULT NULL,
    `delstatus` TINYINT(3) DEFAULT 0,
    `active` TINYINT(3) DEFAULT 1,
    `permission` TINYINT(3) DEFAULT 0,
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

) DEFAULT CHARSET UTF8 COMMENT '';