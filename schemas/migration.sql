-- Make emails unique
ALTER TABLE `provider`.`provider`
ADD UNIQUE (`email`);

ALTER TABLE `patient`.`patient`
ADD UNIQUE (`email`);

-- Add password field
ALTER TABLE `provider`.`provider`
ADD COLUMN `password` VARCHAR(255) NOT NULL DEFAULT '';

ALTER TABLE `patient`.`patient`
ADD COLUMN `password` VARCHAR(255) NOT NULL DEFAULT '';

-- Make fields not null
ALTER TABLE `patient`.`patient`
MODIFY COLUMN `email` VARCHAR(255) NOT NULL;

ALTER TABLE `patient`.`patient`
MODIFY COLUMN `language` INT NOT NULL;

ALTER TABLE `provider`.`provider`
MODIFY name VARCHAR(255) NOT NULL;

ALTER TABLE `provider`.`provider`
MODIFY suffix VARCHAR(255) NOT NULL;

ALTER TABLE `provider`.`provider`
MODIFY bio TEXT NOT NULL;

ALTER TABLE `provider`.`provider`
MODIFY email VARCHAR(255) NOT NULL;

