CREATE DATABASE IF NOT EXISTS 'appointment' DEFAULT CHARACTER SET utf8;
USE `appointment`;

CREATE TABLE IF NOT EXISTS `appointment`.`appointment`(
    `id` CHAR(36) NOT NULL,
    `date_time` DATETIME NOT NULL,
    `provider` CHAR(36) NOT NULL,
    `patient` CHAR(36) NOT NULL,
    FOREIGN KEY (`provider`) REFERENCES `provider`.`provider`.(`id`),
    FOREIGN KEY (`patient`) REFERENCES `patient`.`patient`.(`id`)
)