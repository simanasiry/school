--
-- Table structure for table `registration`
--
CREATE TABLE `registrations`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`    timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `student_id`    bigint unsigned NOT NULL,
    `teacher_id`    bigint unsigned NOT NULL,
    PRIMARY KEY      (`id`),
    FOREIGN KEY      (`student_id`) REFERENCES `students` (`id`) ON UPDATE CASCADE,
    FOREIGN KEY      (`teacher_id`) REFERENCES `teachers` (`id`) ON UPDATE CASCADE,
    UNIQUE KEY `registrations_student_id_teacher_id_unix` (`student_id`, `teacher_id`)

);