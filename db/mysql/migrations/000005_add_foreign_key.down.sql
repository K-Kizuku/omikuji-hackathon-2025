ALTER TABLE `pythons_stats`
  DROP FOREIGN KEY `fk_pythons_stats_skill4`,
  DROP FOREIGN KEY `fk_pythons_stats_skill3`,
  DROP FOREIGN KEY `fk_pythons_stats_skill2`,
  DROP FOREIGN KEY `fk_pythons_stats_skill1`,
  DROP FOREIGN KEY `fk_pythons_stats_python`;

ALTER TABLE `pythons`
  DROP FOREIGN KEY `fk_pythons_users`;
