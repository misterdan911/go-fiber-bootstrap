DROP TABLE IF EXISTS `users`;

CREATE TABLE `users`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NULL,
  `email` varchar(50) NULL,
  `password` varchar(255) NULL,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  PRIMARY KEY (`id`)
);
