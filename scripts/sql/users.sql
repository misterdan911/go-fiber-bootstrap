DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(50) DEFAULT NULL,
    `email` varchar(50) DEFAULT NULL,
    `password` varchar(255) DEFAULT NULL,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`) USING HASH,
    UNIQUE KEY `email` (`email`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci