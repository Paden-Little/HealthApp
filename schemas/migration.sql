-- Make emails unique
ALTER TABLE `provider`.`provider`
ADD UNIQUE (`email`);

ALTER TABLE `patient`.`patient`
ADD UNIQUE (`email`);

-- Add password field
SET @dbname = 'patient';
SET @tablename = 'patient';
SET @columnname = 'password';
SET @preparedStatement = (SELECT IF(
     (
         SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
         WHERE
             (table_name = @tablename)
           AND (table_schema = @dbname)
           AND (column_name = @columnname)
     ) > 0,
     'SELECT 1',
     CONCAT('ALTER TABLE ', @tablename, ' ADD ', @columnname, ' VARCHAR(255) NOT NULL;')
));
PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
-- Do the same for the other db
set @dbname = 'provider';
SET @tablename = 'provider';
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;

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

