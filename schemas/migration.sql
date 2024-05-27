-- Make emails unique

SET @dbName = 'patient';
SET @tableName = 'patient';
SET @preparedStatement = (SELECT IF(
     (
         SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS
         WHERE
             (table_name = @tableName)
           AND (table_schema = @dbName)
           AND (constraint_name = 'email')
     ) > 0,
     'SELECT 1',
     CONCAT('ALTER TABLE ', @dbName, '.',  @tableName, ' ADD CONSTRAINT UNIQUE (`email`);')
));
PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
-- Do the same for the other db
SET @dbName = 'provider';
SET @tableName = 'provider';

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
MODIFY suffix VARCHAR(255) NOT NULL;

ALTER TABLE `provider`.`provider`
MODIFY bio TEXT NOT NULL;

ALTER TABLE `provider`.`provider`
MODIFY email VARCHAR(255) NOT NULL;

-- Split provider name into firstname and lastname
SET @dbName = 'provider';
SET @tableName = 'provider';
SET @columnName = 'name';

SET @preparedStatement = (
    SELECT IF(
       (
           SELECT COUNT(*)
           FROM INFORMATION_SCHEMA.COLUMNS
           WHERE
               table_name = @tableName
             AND table_schema = @dbName
             AND column_name = @columnName
       ) > 0,
       CONCAT(
           'ALTER TABLE `', @dbName, '`.`', @tableName, '` ',
           'ADD COLUMN `firstname` VARCHAR(255) NOT NULL AFTER `id`, ',
           'ADD COLUMN `lastname` VARCHAR(255) NOT NULL AFTER `firstname`; '
       ),
       'SELECT 1'
   )
);

PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;

-- Update firstname and lastname columns with split values
SET @updateStatement = (
    SELECT IF(
           (
               SELECT COUNT(*)
               FROM INFORMATION_SCHEMA.COLUMNS
               WHERE
                   table_name = @tableName
                 AND table_schema = @dbName
                 AND column_name = 'firstname'
           ) > 0,
           CONCAT(
               'UPDATE `', @dbName, '`.`', @tableName, '` ',
               'SET `firstname` = SUBSTRING_INDEX(`name`, '' '', 1), ',
               '`lastname` = SUBSTRING_INDEX(`name`, '' '', -1); '
           ),
           'SELECT 1'
   )
);

PREPARE updateIfExists FROM @updateStatement;
EXECUTE updateIfExists;
DEALLOCATE PREPARE updateIfExists;

-- Drop the original 'name' column
SET @dropStatement = (
    SELECT IF(
       (
           SELECT COUNT(*)
           FROM INFORMATION_SCHEMA.COLUMNS
           WHERE
               table_name = @tableName
             AND table_schema = @dbName
             AND column_name = 'firstname'
       ) > 0,
       CONCAT(
           'ALTER TABLE `', @dbName, '`.`', @tableName, '` ',
           'DROP COLUMN `name`;'
       ),
       'SELECT 1'
   )
);

PREPARE dropIfExists FROM @dropStatement;
EXECUTE dropIfExists;
DEALLOCATE PREPARE dropIfExists;
