--
-- Table structure for table `schools`
--
CREATE TABLE `schools`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`    timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `head_master`          varchar(30) NOT NULL,
    `school` varchar(30) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `schools_head_master_uindex` (`head_master`),
    UNIQUE KEY `schools_school_uindex` (`school`)
    );
