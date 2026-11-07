-- Infectious-disease domain tables (MySQL).
-- Dev/test prefer GORM AutoMigrate via models.AllModels(); this file is for prod.
-- Usage: ./scripts/migrate.sh up

ALTER TABLE `users`
    ADD COLUMN IF NOT EXISTS `name` VARCHAR(100) NOT NULL DEFAULT '' AFTER `email`,
    ADD COLUMN IF NOT EXISTS `phone_number` VARCHAR(20) NOT NULL DEFAULT '' AFTER `name`;

-- Note: MySQL <8.0.12 may not support IF NOT EXISTS on ADD COLUMN; run manually if needed.
CREATE UNIQUE INDEX IF NOT EXISTS `uk_users_phone_number` ON `users` (`phone_number`);

CREATE TABLE IF NOT EXISTS `employee_profiles` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `gender` VARCHAR(20) NOT NULL DEFAULT '',
    `is_pregnant` TINYINT(1) NOT NULL DEFAULT 0,
    `pregnancy_weeks` INT NOT NULL DEFAULT 0,
    `id_number` VARCHAR(32) NOT NULL DEFAULT '',
    `birth_date` DATE NULL,
    `age` INT NOT NULL DEFAULT 0,
    `ethnicity` VARCHAR(50) NOT NULL DEFAULT '',
    `education_level` VARCHAR(50) NOT NULL DEFAULT '',
    `work_on_plateau_start_date` DATE NULL,
    `department` VARCHAR(100) NOT NULL DEFAULT '',
    `specific_occupation` VARCHAR(100) NOT NULL DEFAULT '',
    `medical_personnel_type` VARCHAR(100) NOT NULL DEFAULT '',
    `other_position_name` VARCHAR(100) NOT NULL DEFAULT '',
    `other_phone_number` VARCHAR(20) NOT NULL DEFAULT '',
    `emergency_contact_name` VARCHAR(100) NOT NULL DEFAULT '',
    `emergency_contact_phone_number` VARCHAR(20) NOT NULL DEFAULT '',
    `emergency_contact_relation` VARCHAR(50) NOT NULL DEFAULT '',
    `has_medical_history` TINYINT(1) NOT NULL DEFAULT 0,
    `has_hypertension` TINYINT(1) NOT NULL DEFAULT 0,
    `has_diabetes` TINYINT(1) NOT NULL DEFAULT 0,
    `has_hyperlipidemia` TINYINT(1) NOT NULL DEFAULT 0,
    `has_hyperuricemia` TINYINT(1) NOT NULL DEFAULT 0,
    `has_coronary_heart_disease` TINYINT(1) NOT NULL DEFAULT 0,
    `has_stroke` TINYINT(1) NOT NULL DEFAULT 0,
    `has_other_cardiovascular_diseases` TINYINT(1) NOT NULL DEFAULT 0,
    `has_asthma` TINYINT(1) NOT NULL DEFAULT 0,
    `has_copd` TINYINT(1) NOT NULL DEFAULT 0,
    `has_peptic_ulcer` TINYINT(1) NOT NULL DEFAULT 0,
    `has_malignant_tumor` TINYINT(1) NOT NULL DEFAULT 0,
    `has_lung_cancer` TINYINT(1) NOT NULL DEFAULT 0,
    `has_other_cancer` TINYINT(1) NOT NULL DEFAULT 0,
    `other_cancer_name` VARCHAR(100) NOT NULL DEFAULT '',
    `has_severe_mental_disorders` TINYINT(1) NOT NULL DEFAULT 0,
    `has_tuberculosis` TINYINT(1) NOT NULL DEFAULT 0,
    `has_hepatitis` TINYINT(1) NOT NULL DEFAULT 0,
    `has_occupational_disease` TINYINT(1) NOT NULL DEFAULT 0,
    `has_chronic_kidney_disease` TINYINT(1) NOT NULL DEFAULT 0,
    `has_chronic_liver_disease` TINYINT(1) NOT NULL DEFAULT 0,
    `has_immunodeficiency` TINYINT(1) NOT NULL DEFAULT 0,
    `has_typhus` TINYINT(1) NOT NULL DEFAULT 0,
    `is_postpartum_in_six_weeks` TINYINT(1) NOT NULL DEFAULT 0,
    `has_dust_exposure` TINYINT(1) NOT NULL DEFAULT 0,
    `has_other_diseases` TINYINT(1) NOT NULL DEFAULT 0,
    `other_diseases_name` VARCHAR(200) NOT NULL DEFAULT '',
    `smoking_status` VARCHAR(50) NOT NULL DEFAULT '',
    `drinking_status` VARCHAR(50) NOT NULL DEFAULT '',
    `height` INT NOT NULL DEFAULT 0,
    `weight` INT NOT NULL DEFAULT 0,
    `is_vaccinated_for_covid` TINYINT(1) NOT NULL DEFAULT 0,
    `is_vaccinated_for_flu` TINYINT(1) NOT NULL DEFAULT 0,
    `is_vaccinated_for_plague` TINYINT(1) NOT NULL DEFAULT 0,
    `is_vaccinated_for_bcg` TINYINT(1) NOT NULL DEFAULT 0,
    `is_vaccinated_for_hepatitis` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_employee_profiles_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `daily_health_statuses` (
    `status_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `check_in_date` DATE NOT NULL,
    `check_in_time` DATETIME NULL,
    `is_health` TINYINT(1) NULL,
    `disease_type_name` VARCHAR(100) NOT NULL DEFAULT '',
    `diagnosis_disease_type_name` VARCHAR(100) NOT NULL DEFAULT '',
    `longitude` DOUBLE NOT NULL DEFAULT 0,
    `latitude` DOUBLE NOT NULL DEFAULT 0,
    `province` VARCHAR(50) NOT NULL DEFAULT '',
    `city` VARCHAR(50) NOT NULL DEFAULT '',
    `county` VARCHAR(50) NOT NULL DEFAULT '',
    `location_name` VARCHAR(200) NOT NULL DEFAULT '',
    PRIMARY KEY (`status_id`),
    KEY `idx_daily_health_user_date` (`user_id`, `check_in_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `disease_types` (
    `disease_type_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `disease_type_name` VARCHAR(100) NOT NULL DEFAULT '',
    `description` VARCHAR(500) NOT NULL DEFAULT '',
    PRIMARY KEY (`disease_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `symptom_weighting` (
    `symptom_weighting_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `disease_type_id` BIGINT UNSIGNED NOT NULL DEFAULT 0,
    `disease_type_name` VARCHAR(100) NOT NULL DEFAULT '',
    `table_name` VARCHAR(100) NOT NULL DEFAULT '',
    `symptom_field_name` VARCHAR(100) NOT NULL DEFAULT '',
    `weight_score` DOUBLE NOT NULL DEFAULT 0,
    PRIMARY KEY (`symptom_weighting_id`),
    KEY `idx_symptom_weight_disease` (`disease_type_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `disease_scoring` (
    `scoring_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `status_id` BIGINT UNSIGNED NOT NULL,
    `symptom_time` DATE NULL,
    `disease_type_id` BIGINT UNSIGNED NOT NULL DEFAULT 0,
    `disease_type_name` VARCHAR(100) NOT NULL DEFAULT '',
    `score` DOUBLE NOT NULL DEFAULT 0,
    PRIMARY KEY (`scoring_id`),
    KEY `idx_disease_scoring_status` (`status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `system_settings` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `is_ai` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `hospitals` (
    `hospital_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `hospital_name` VARCHAR(200) NOT NULL DEFAULT '',
    `address` VARCHAR(500) NOT NULL DEFAULT '',
    `phone` VARCHAR(50) NOT NULL DEFAULT '',
    PRIMARY KEY (`hospital_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `user_feedbacks` (
    `feedback_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0,
    `name` VARCHAR(100) NOT NULL DEFAULT '',
    `phone_number` VARCHAR(20) NOT NULL DEFAULT '',
    `feedback_title` VARCHAR(200) NOT NULL DEFAULT '',
    `feedback_text` TEXT,
    `feedback_date` DATE NULL,
    PRIMARY KEY (`feedback_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `satisfaction_ratings` (
    `rating_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `survey_id` BIGINT UNSIGNED NOT NULL DEFAULT 0,
    `phone_number` VARCHAR(20) NOT NULL DEFAULT '',
    `satisfaction_level` VARCHAR(50) NOT NULL DEFAULT '',
    `rating_time` DATETIME NULL,
    PRIMARY KEY (`rating_id`),
    KEY `idx_satisfaction_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `disease_learning_materials` (
    `material_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `disease_type_name` VARCHAR(100) NOT NULL DEFAULT '',
    `title` VARCHAR(200) NOT NULL DEFAULT '',
    `content` TEXT,
    `file_name` VARCHAR(255) NOT NULL DEFAULT '',
    `sort_order` INT NOT NULL DEFAULT 0,
    `enabled` TINYINT(1) NOT NULL DEFAULT 1,
    PRIMARY KEY (`material_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Remaining symptom/diagnosis/lab tables are created by GORM AutoMigrate in dev,
-- or mirror internal/models/*.go when promoting to production.
