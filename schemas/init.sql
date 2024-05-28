-- -- -- Provider Schema -- -- --
-- This should NEVER be used on a prod database
-- It is meant to set up a fresh MySQL server for the patient and provider APIs
CREATE DATABASE IF NOT EXISTS `provider` DEFAULT CHARACTER SET utf8;
USE `provider`;

CREATE TABLE IF NOT EXISTS `provider`.`provider` (
    `id` CHAR(36) NOT NULL,
    `firstname` VARCHAR(255) NOT NULL,
    `lastname` VARCHAR(255) NOT NULL,
    `suffix` VARCHAR(255) NOT NULL,
    `bio` TEXT NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(15),
    `password` VARCHAR(255) NOT NULL DEFAULT 'password',
    `image` VARCHAR(255),
    UNIQUE (`email`),
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `provider`.`service` (
    `id` INT AUTO_INCREMENT,
    `service` CHAR(255),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS `provider`.`language` (
    `id` INT AUTO_INCREMENT,
    `language` CHAR(50) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `provider`.`provider_service` (
    `provider_id` CHAR(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `service_id` INT NOT NULL,
    FOREIGN KEY (`provider_id`) REFERENCES `provider`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`service_id`) REFERENCES `service`(`id`),
    PRIMARY KEY (`provider_id`, `service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `provider`.`provider_language` (
    `provider_id` CHAR(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `language_id` INT NOT NULL,
    FOREIGN KEY (`provider_id`) REFERENCES `provider`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`language_id`) REFERENCES `language`(`id`),
    PRIMARY KEY (`provider_id`, `language_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `provider`.`provider` (`id`, `firstname`, `lastname`, `suffix`, `bio`, `email`, `phone`) VALUES
    ('00000000-0000-0000-0000-000000000000', 'John', 'Doe', 'MD', 'John Doe, MD. is a great doctor', 'drjohn@gmail.com', '8011234567');

INSERT INTO `provider`.`service` (`service`) VALUES ('Johnology');

INSERT INTO `provider`.`language` (`language`) VALUES ('english');

INSERT INTO `provider`.`provider_service` (`provider_id`, `service_id`) VALUES ('00000000-0000-0000-0000-000000000000', 1);

INSERT INTO `provider`.`provider_language` (`provider_id`, `language_id`) VALUES ('00000000-0000-0000-0000-000000000000', 1);

-- -- -- Patient Schema -- -- --
CREATE DATABASE IF NOT EXISTS `patient` DEFAULT CHARACTER SET utf8;
USE `patient`;

CREATE TABLE IF NOT EXISTS `patient`.`patient` (
    `id` CHAR(36) NOT NULL,
    `firstname` VARCHAR(255) NOT NULL,
    `lastname` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(15),
    `language` INT NOT NULL,
    `birth` DATE NOT NULL,
    `gender` ENUM('male', 'female') NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    FOREIGN KEY (`language`) REFERENCES `provider`.`language`(`id`),
    UNIQUE (`email`),
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `patient`.`allergy` (
    `patient_id` CHAR(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255),
    FOREIGN KEY (`patient_id`) REFERENCES `patient`(`id`) ON DELETE CASCADE,
    PRIMARY KEY (`patient_id`, `name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `patient`.`prescription` (
    `provider_id` CHAR(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    `patient_id` CHAR(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `dosage` VARCHAR(255) NOT NULL,
    `frequency` VARCHAR(255) NOT NULL,
    `start` DATE NOT NULL,
    `end` DATE,
    FOREIGN KEY (`patient_id`) REFERENCES `patient`.`patient`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`provider_id`) REFERENCES `provider`.`provider`(`id`) ON DELETE SET NULL,
    PRIMARY KEY (`patient_id`, `name`, `start`)
);

INSERT INTO `patient`.`patient` (`id`, `firstname`, `lastname`, `email`, `phone`, `language`, `birth`, `gender`, `password`)
VALUES ('11111111-1111-1111-1111-111111111111', 'Bob', 'Johnson', 'bob@example.com', '9876543210', 1, '1985-05-15', 'male', 'unhashed');

INSERT INTO `patient`.`allergy` (`patient_id`, `name`, `description`) VALUES ('11111111-1111-1111-1111-111111111111', 'Peanuts', 'Severe allergy to peanuts');

INSERT INTO `patient`.`prescription` (`provider_id`, `patient_id`, `name`, `dosage`, `frequency`, `start`, `end`) VALUES
    ('00000000-0000-0000-0000-000000000000', '11111111-1111-1111-1111-111111111111', 'Amoxicillin', '500mg', 'Twice daily', '2023-01-01', '2023-01-14'),
    ('00000000-0000-0000-0000-000000000000', '11111111-1111-1111-1111-111111111111', 'Ibuprofen', '200mg', 'As needed', '2023-02-01', NULL);

-- -- -- Appointment Schema -- -- --
CREATE DATABASE IF NOT EXISTS `appointment` DEFAULT CHARACTER SET utf8;
USE `appointment`;

CREATE TABLE IF NOT EXISTS `appointment`.`appointment`(
    `id` CHAR(36) NOT NULL,
    `date_time` DATETIME NOT NULL,
    `provider` CHAR(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `patient` CHAR(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `service` INT NOT NULL,
    `description` TEXT NOT NULL,
    FOREIGN KEY (`provider`) REFERENCES `provider`.`provider`(`id`),
    FOREIGN KEY (`patient`) REFERENCES `patient`.`patient`(`id`),
    FOREIGN KEY (`service`) REFERENCES `provider`.`service`(`id`)
)