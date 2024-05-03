-- Assumed to be using mysql bc its the goat
CREATE DATABASE IF NOT EXISTS records;

-- All the ids are VARCHAR(36) because we're using UUIDs

-- `gender` refers to the gender assigned at birth for the patient
CREATE TABLE IF NOT EXISTS `patient` (
    `id` VARCHAR(36) NOT NULL, 
    `birth_date` DATE NOT NULL,
    `death_date` DATE,
    `first` VARCHAR(255) NOT NULL,
    `middle` VARCHAR(255),
    `last` VARCHAR(255) NOT NULL,
    `gender` ENUM('male', 'female') NOT NULL,
    `race` VARCHAR(50),
    `ethnicity` VARCHAR(50),
    `address` VARCHAR(255) NOT NULL,
    `city` VARCHAR(255) NOT NULL,
    `state` VARCHAR(255) NOT NULL,
    `zip` VARCHAR(10) NOT NULL,
    PRIMARY KEY (`id`)
);

-- We only use `name` for provider, instead of 
-- `first`, `middle`, and `last` for patient
-- `gender` refers to the gender assigned at birth for the provider
-- `organization` is the name of the provider's organization, might be a separate table in the future
CREATE TABLE IF NOT EXISTS `provider` (
    `id` VARCHAR(36) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `organization` VARCHAR(255) NOT NULL,
    `gender` ENUM('male', 'female') NOT NULL,
    `specialty` VARCHAR(255) NOT NULL,
    `address` VARCHAR(255) NOT NULL,
    `city` VARCHAR(255) NOT NULL,
    `state` VARCHAR(2) NOT NULL,
    `zip` VARCHAR(10) NOT NULL,
    PRIMARY KEY (`id`)
);

-- Encounter is a visit to a provider by a patient
-- 'urgentcare' visits are basic checkups, like sore throat, achy joints, etc.
-- 'wellness' visits are routine checkups, like physicals, vaccinations, etc.
-- 'inpatient' visits are for patients who are admitted to the hospital
-- `description` is a description of what happened during the encounter
-- `reason` is a human-readable description of the reason for the encounter
-- We could add SNOMED or ICD codes but that would require a lot of work so... yeah
CREATE TABLE IF NOT EXISTS `encounter` (
    `id` VARCHAR(36) NOT NULL,
    `date` DATE NOT NULL,
    `patient_id` VARCHAR(36) NOT NULL,
    `provider_id` VARCHAR(36) NOT NULL,
    `type` ENUM('emergency', 'ambulatory', 'wellness', 'urgentcare', 'inpatient') NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `reason` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`patient_id`) REFERENCES `patient`(`id`),
    FOREIGN KEY (`provider_id`) REFERENCES `provider`(`id`)
);

-- `expiration` is the expiration date of the vaccine, if applicable
CREATE TABLE IF NOT EXISTS `immunization` (
    `id` VARCHAR(36) NOT NULL,
    `date` DATE NOT NULL,
    `patient_id` VARCHAR(36) NOT NULL,
    `provider_id` VARCHAR(36) NOT NULL,
    `vaccine` VARCHAR(255) NOT NULL,
    `lot` VARCHAR(255) NOT NULL,
    `expiration` DATE,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`patient_id`) REFERENCES `patient`(`id`),
    FOREIGN KEY (`provider_id`) REFERENCES `provider`(`id`)
);

-- `encounter_id` is the encounter that the medication was prescribed in
CREATE TABLE IF NOT EXISTS `medication` (
    `id` VARCHAR(36) NOT NULL,
    `start` DATE NOT NULL,
    `end` DATE NOT NULL,
    `patient_id` VARCHAR(36) NOT NULL,
    `provider_id` VARCHAR(36) NOT NULL,
    `encounter_id` VARCHAR(36) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `dose` VARCHAR(255) NOT NULL,
    `frequency` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`patient_id`) REFERENCES `patient`(`id`),
    FOREIGN KEY (`provider_id`) REFERENCES `provider`(`id`),
    FOREIGN KEY (`encounter_id`) REFERENCES `encounter`(`id`)
);

-- condition is a diagnosis
-- `start` is when the condition was first diagnosed
-- `end` is when the condition was resolved, if applicable
CREATE TABLE IF NOT EXISTS `condition` (
    `id` VARCHAR(36) NOT NULL,
    `start` DATE NOT NULL,
    `end` DATE,
    `patient_id` VARCHAR(36) NOT NULL,
    `provider_id` VARCHAR(36) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`patient_id`) REFERENCES `patient`(`id`),
    FOREIGN KEY (`provider_id`) REFERENCES `provider`(`id`)
);