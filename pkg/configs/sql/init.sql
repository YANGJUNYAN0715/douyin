-- 用户信息表
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `username`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Username',
  `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
  `name`   varchar(128) NOT NULL DEFAULT 'xiaodouyin' COMMENT 'Name',
  
  `follow_count` bigint NOT NULL DEFAULT 0,
  `follower_count` bigint NOT NULL DEFAULT 0,
  `is_follow`      boolean  NOT NULL DEFAULT 1 COMMENT 'IsFollow',
  `avatar`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Avatar',
  `background_image`   varchar(128) NOT NULL DEFAULT '' COMMENT 'BackgroundImage',
  `signature`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Signature',
  `total_favorited`   varchar(128) NOT NULL DEFAULT '' COMMENT 'TotalFavorited',
  `work_count` bigint NOT NULL DEFAULT 0,
  `favorite_count` bigint NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
  PRIMARY KEY (`id`),
  KEY          `idx_username` (`username`) COMMENT 'Username index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';
-- 视频表
CREATE TABLE `video` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `author_id` bigint NOT NULL,
  `publish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `play_url` varchar(128) NOT NULL,
  `cover_url` varchar(128) NOT NULL,
  `favorite_count` bigint DEFAULT 0,
  `comment_count` bigint DEFAULT 0,
  `title` varchar(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'video create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'video update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'video delete time',
  PRIMARY KEY (`id`),
  KEY          `idx_video_of_user_id` (`author_id`) COMMENT 'VideoOfUserId index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video table';
-- 用户聊天表
CREATE TABLE `message`
(
    `id`         bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `from_user_id` bigint NOT NULL COMMENT 'FromUserID',
    `to_user_id` bigint NOT NULL COMMENT 'ToUserID',
    `content`    TEXT NULL COMMENT 'Content',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Message create_time',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Message create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Message update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Message delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id` (`from_user_id`) COMMENT 'UserId index'
    
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Message table';
-- 用户关系表
CREATE TABLE `relation` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `from_user_id` bigint NOT NULL,
  `to_user_id` bigint NOT NULL,
  -- `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Relation create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Relation update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Relation delete time',
  PRIMARY KEY (`id`),
  KEY          `idx_from_user_id` (`from_user_id`) COMMENT 'Relation index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Relation table';

-- 点赞表
CREATE TABLE `favorite` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `user_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'follow create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'follow update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'follow delete time',
  PRIMARY KEY (`id`),
  KEY          `idx_from_user_id` (`user_id`) COMMENT 'Favorite index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Favorite table';

-- 评论表

CREATE TABLE `comment` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `user_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  `parent_id` bigint NOT NULL,
	`is_valid`   boolean  NOT NULL DEFAULT 1 COMMENT 'IsValid',
  `content`    TEXT NULL COMMENT 'Content',
  `create_date` varchar(128) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'comment create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'comment update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'comment delete time',
  PRIMARY KEY (`id`),
  KEY          `idx_comment_user_id` (`user_id`) COMMENT 'Comment index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Comment table';


ALTER TABLE `video` ADD FOREIGN KEY (`author_id`) REFERENCES `user` (`id`);
ALTER TABLE `message` ADD FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`);
ALTER TABLE `message` ADD FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`);
ALTER TABLE `relation` ADD FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`);
ALTER TABLE `relation` ADD FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`);
ALTER TABLE `favorite` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `favorite` ADD FOREIGN KEY (`video_id`) REFERENCES `video` (`id`);
ALTER TABLE `comment` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `comment` ADD FOREIGN KEY (`video_id`) REFERENCES `video` (`id`);