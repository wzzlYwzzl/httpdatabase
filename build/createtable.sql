DROP DATABASE IF EXISTS `httpdb`;
CREATE DATABASE `httpdb`;
USE `httpdb`;

DROP TABLE IF EXISTS `userpasswd`;
CREATE TABLE `userpasswd` (
    `user_id` int NOT NULL AUTO_INCREMENT,
    `name` char(40) NOT NULL,
    `password` char(100) NOT NULL,
    `cpus` int DEFAULT 2,
    `memory` int DEFAULT 500,
    PRIMARY KEY (`user_id`),
    UNIQUE KEY (`name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `userns`;
CREATE TABLE `userns` (
    `userns_id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `namespace` varchar(50) NOT NULL,
    PRIMARY KEY (`userns_id`),
    UNIQUE KEY (`namespace`),
    FOREIGN KEY (`user_id`) REFERENCES userpasswd(`user_id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
    `res_id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `cpus_use` int DEFAULT NULL,
    `mem_use` int DEFAULT NULL,
    PRIMARY KEY (`res_id`),
    UNIQUE KEY (`user_id`)
    FOREIGN KEY (`user_id`) REFERENCES userpasswd(`user_id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `deployment`;
CREATE TABLE `deployment` (
    `deploy_id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `app_name` varchar(100) NOT NULL,
    `cpus_use` int DEFAULT NULL,
    `mem_use` int DEFAULT NULL,
    PRIMARY KEY (`deploy_id`),
    UNIQUE KEY (`user_id`)
    FOREIGN KEY (`user_id`) REFERENCES userpasswd(`user_id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
