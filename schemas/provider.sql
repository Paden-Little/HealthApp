
CREATE TABLE IF NOT EXISTS `provider` (
    `id` CHAR(36) NOT NULL,
    `name` VARCHAR(255),
    `suffix` VARCHAR(255),
    `bio` TEXT,
    `email` VARCHAR(255),
    `phone` VARCHAR(15),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS `service` (
    `id` INT AUTO_INCREMENT,
    `service` CHAR(255),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS `language` (
    `id` INT AUTO_INCREMENT,
    `language` CHAR(50),
    PRIMARY KEY (`id`)
);


CREATE TABLE IF NOT EXISTS `provider_service` (
    `provider_id` CHAR(36) NOT NULL,
    `service_id` INT NOT NULL,
    FOREIGN KEY (`provider_id`) REFERENCES `provider`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`service_id`) REFERENCES `service`(`id`),
    PRIMARY KEY (`provider_id`, `service_id`)
);

CREATE TABLE IF NOT EXISTS `provider_language` (
    `provider_id` CHAR(36) NOT NULL,
    `language_id` INT NOT NULL,
    FOREIGN KEY (`provider_id`) REFERENCES  `provider`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`language_id`) REFERENCES  `language`(`id`),
    PRIMARY KEY (`provider_id`, `language_id`)
);
