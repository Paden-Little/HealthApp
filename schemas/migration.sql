-- Make emails unique
ALTER TABLE `provider`.`provider`
ADD UNIQUE (`email`);

ALTER TABLE `patient`.`patient`
ADD UNIQUE (`email`);

-- Add password field
SET @dbname = 'patient';
SET @tablename = 'patient';
SET @columnname = 'password';
SET @columndefinition = 'VARCHAR(255) NOT NULL';
SET @preparedStatement = (SELECT IF(
     (
         SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
         WHERE
             (table_name = @tablename)
           AND (table_schema = @dbname)
           AND (column_name = @columnname)
     ) > 0,
     'SELECT 1',
     CONCAT('ALTER TABLE ', @dbname, '.',  @tablename, ' ADD ', @columnname, ' ', @columndefinition, ';')
));
PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
-- Do the same for the other db
SET @dbname = 'provider';
SET @tablename = 'provider';
EXECUTE alterIfNotExists;
-- Then add the image field
SET @columnname = 'image';
SET @columndefinition = 'VARCHAR(255)';
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

-- Add image field
ALTER TABLE `provider`.`provider`
