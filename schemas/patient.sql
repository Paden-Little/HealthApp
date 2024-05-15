CREATE DATABASE IF NOT EXISTS `provider` DEFAULT CHARACTER SET utf8;
CREATE DATABASE IF NOT EXISTS `patient` DEFAULT CHARACTER SET utf8;
USE `patient`;

CREATE TABLE IF NOT EXISTS `provider`.`language` (
   `id` INT AUTO_INCREMENT,
   `language` CHAR(50) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `patient`.`patient` (
    `id` CHAR(36) NOT NULL,
    `firstname` VARCHAR(255) NOT NULL,
    `lastname` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255),
    `phone` VARCHAR(15),
    `language` INT,
    `birth` DATE NOT NULL,
    `gender` ENUM('male', 'female') NOT NULL,
    FOREIGN KEY (`language`) REFERENCES `provider`.`language`(`id`),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS `patient`.`allergy` (
    `patient_id` CHAR(36) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255),
    FOREIGN KEY (`patient_id`) REFERENCES `patient`(`id`) ON DELETE CASCADE,
    PRIMARY KEY (`patient_id`, `name`)
);

CREATE TABLE IF NOT EXISTS `patient`.`prescription` (
    `provider_id` CHAR(36),
    `patient_id` CHAR(36) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `dosage` VARCHAR(255) NOT NULL,
    `frequency` VARCHAR(255) NOT NULL,
    `start` DATE NOT NULL,
    `end` DATE,
    FOREIGN KEY (`patient_id`) REFERENCES `patient`.`patient`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`provider_id`) REFERENCES `provider`.`provider`(`id`) ON DELETE SET NULL,
    -- start is included  in the primary key to allow multiple prescriptions with the same name at different dates
    PRIMARY KEY (`patient_id`, `name`, `start`)
);