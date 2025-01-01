ALTER TABLE `pythons`
  ADD CONSTRAINT `fk_pythons_users`
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `pythons_stats`
  ADD CONSTRAINT `fk_pythons_stats_python`
  FOREIGN KEY (`python_id`) REFERENCES `pythons` (`id`);

ALTER TABLE `pythons_stats`
  ADD CONSTRAINT `fk_pythons_stats_skill1`
  FOREIGN KEY (`skill1`) REFERENCES `pythons_skills` (`id`);

ALTER TABLE `pythons_stats`
  ADD CONSTRAINT `fk_pythons_stats_skill2`
  FOREIGN KEY (`skill2`) REFERENCES `pythons_skills` (`id`);

ALTER TABLE `pythons_stats`
  ADD CONSTRAINT `fk_pythons_stats_skill3`
  FOREIGN KEY (`skill3`) REFERENCES `pythons_skills` (`id`);

ALTER TABLE `pythons_stats`
  ADD CONSTRAINT `fk_pythons_stats_skill4`
  FOREIGN KEY (`skill4`) REFERENCES `pythons_skills` (`id`);
