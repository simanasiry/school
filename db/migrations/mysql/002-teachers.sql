--
-- Table structure for table `teachers`
--
CREATE TABLE `teachers`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`    timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `school_id`     bigint unsigned NOT NULL,
    `teacher`       varchar(30) NOT NULL,
    PRIMARY KEY      (`id`),
    FOREIGN KEY      (`school_id`) REFERENCES `schools` (`id`) ON UPDATE CASCADE,
    UNIQUE KEY       `teachers_teacher_uindex` (`teacher`)


);