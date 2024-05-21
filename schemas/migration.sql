ALTER TABLE `provider`.`provider`
ADD CONSTRAINT `unique_email` UNIQUE (`email`);

ALTER TABLE `provider`.`provider`
ADD COLUMN `password` VARCHAR(255) NOT NULL DEFAULT '';

ALTER TABLE `patient`.`patient`
ADD CONSTRAINT `unique_email` UNIQUE (`email`);

ALTER TABLE `patient`.`patient`
ADD COLUMN `password` VARCHAR(255) NOT NULL DEFAULT '';
