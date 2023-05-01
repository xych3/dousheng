DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
    `id` bigint NOT NULL COMMENT '主键ID',
    `user_id` bigint DEFAULT NULL COMMENT '用户ID',
    `video_id` bigint DEFAULT NULL COMMENT '视频ID',
    `comment_text` text COMMENT '评论内容',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites` (
     `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
     `user_id` bigint DEFAULT NULL COMMENT '用户ID',
     `video_id` bigint DEFAULT NULL COMMENT '视频ID',
     `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
     `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `deleted_at` timestamp NULL DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `relations`;
CREATE TABLE `relations` (
     `id` bigint NOT NULL COMMENT '主键ID',
     `to_user_id` bigint DEFAULT NULL COMMENT '被关注用户ID',
     `follower_id` bigint DEFAULT NULL COMMENT '关注用户ID',
     `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
     `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `deleted_at` timestamp NULL DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
     `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
     `username` varchar(40) DEFAULT NULL COMMENT '用户名',
     `password` varchar(40) DEFAULT NULL COMMENT '密码',
     `follow_count` bigint DEFAULT NULL COMMENT '关注总数',
     `follower_count` bigint DEFAULT NULL COMMENT '粉丝总数',
     `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
     `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `deleted_at` timestamp NULL DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
      `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
      `title` varchar(255) DEFAULT NULL COMMENT '视频标题',
      `user_id` bigint DEFAULT NULL COMMENT '视频作者ID',
      `play_url` varchar(255) DEFAULT NULL COMMENT '视频播放地址',
      `cover_url` varchar(255) DEFAULT NULL COMMENT '视频封面地址',
      `favorite_count` bigint DEFAULT NULL COMMENT '视频点赞总数',
      `comment_count` bigint DEFAULT NULL COMMENT '视频评论总数',
      `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
      `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      `deleted_at` timestamp NULL DEFAULT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `chat`;
CREATE TABLE `chat` (
      `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
      `from_user_id` bigint DEFAULT NULL COMMENT '发送消息用户ID',
      `to_user_id` bigint DEFAULT NULL COMMENT '接收消息用户ID',
      `content` varchar(250) DEFAULT NULL COMMENT '消息内容',
      `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
      `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      `deleted_at` timestamp NULL DEFAULT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;