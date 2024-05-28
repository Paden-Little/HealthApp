CREATE DATABASE IF NOT EXISTS 'appointment' DEFAULT CHARACTER SET utf8;
USE `appointment`;

CREATE TABLE IF NOT EXISTS `appointment`.`appointment`(
    `id` CHAR(36) NOT NULL,
    `date` DATE NOT NULL,
    `start_time` TIME NOT NULL,
    `end_time` TIME NOT NULL,
    `provider` CHAR(36) NOT NULL,
    `patient` CHAR(36) NOT NULL,
    `service` INT NOT NULL,
    `description` TEXT NOT NULL,
    FOREIGN KEY (`provider`) REFERENCES `provider`.`provider`(`id`),
    FOREIGN KEY (`patient`) REFERENCES `patient`.`patient`(`id`),
    FOREIGN KEY (`service`) REFERENCES `provider`.`service`(`id`)
)