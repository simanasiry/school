--
-- Table structure for table `students`
--
CREATE TABLE `students`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`    timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `school_id`     bigint unsigned NOT NULL,
    `student`       varchar(30) NOT NULL,
    PRIMARY KEY      (`id`),
    FOREIGN KEY      (`school_id`) REFERENCES `schools` (`id`) ON UPDATE CASCADE,
    UNIQUE KEY       `students_student_uindex` (`student`)


);