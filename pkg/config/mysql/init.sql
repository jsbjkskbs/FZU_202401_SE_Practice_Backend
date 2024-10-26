SET NAMES utf8mb4;
DROP TABLE IF EXISTS `Message`;
CREATE TABLE `Message` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '消息ID',
    `from_user_id` BIGINT NOT NULL COMMENT '发送者ID',
    `to_user_id` BIGINT NOT NULL COMMENT '接受者ID',
    `content` VARCHAR(255) NOT NULL COMMENT '内容',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `deleted_at` BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_from_to_user` (`from_user_id`, `to_user_id`) USING BTREE COMMENT '发送者与接收者索引'
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='消息表';

DROP TABLE IF EXISTS `User`;
CREATE TABLE `User` (
    `uid` BIGINT NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `user_name` VARCHAR(255) NOT NULL COMMENT '用户名称',
    `password` VARCHAR(255) NOT NULL COMMENT '用户密码',
    `email` VARCHAR(255) NOT NULL COMMENT '用户邮箱',
    `role` VARCHAR(255) NOT NULL COMMENT '权限',
    `avatar_url` VARCHAR(255) COMMENT '用户头像url',
    `created_at` BIGINT NOT NULL COMMENT '创建账号时间',
    `updated_at` BIGINT NOT NULL COMMENT '最近登录时间',
    `deleted_at` BIGINT DEFAULT NULL COMMENT '账号删除时间',
    `mfa_secret` VARCHAR(255) COMMENT 'mfa秘钥',
    `mfa_enable` BOOLEAN NOT NULL COMMENT '是否使用mfa',
    PRIMARY KEY (`uid`),
    KEY `idx_username` (`user_name`) USING BTREE COMMENT '用户名索引'
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

DROP TABLE IF EXISTS `Video`;
CREATE TABLE `Video` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '视频ID',
    `user_id` BIGINT NOT NULL COMMENT '作者ID',
    `video_url` VARCHAR(255) NOT NULL COMMENT '视频url',
    `cover_url` VARCHAR(255) NOT NULL COMMENT '封面url',
    `title` VARCHAR(255) NOT NULL COMMENT '标题',
    `description` VARCHAR(255) NOT NULL COMMENT '简介',
    `visit_count` BIGINT NOT NULL COMMENT '浏览量',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `updated_at` BIGINT NOT NULL COMMENT '修改时间',
    `deleted_at` BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_created_at` (`created_at`) USING BTREE COMMENT '时间查询索引',
    KEY `idx_user_id` (`user_id`) USING BTREE COMMENT '作者查询索引',
    KEY `idx_title` (`title`) USING BTREE COMMENT '标题查询索引'
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='视频表';


DROP TABLE IF EXISTS `VideoComment`;
CREATE TABLE `VideoComment` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '评论ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `video_id` BIGINT NOT NULL COMMENT '视频ID',
    `parent_id` BIGINT NOT NULL COMMENT '父评论ID',
    `content` VARCHAR(255) NOT NULL COMMENT '评论内容',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `updated_at` BIGINT NOT NULL COMMENT '修改时间',
    `deleted_at` BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) USING BTREE COMMENT '用户索引',
    KEY `idx_video_id` (`video_id`) USING BTREE COMMENT '视频索引'
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='视频评论表';


DROP TABLE IF EXISTS `Follow`;
CREATE TABLE `Follow` (
    `followed_id` BIGINT NOT NULL COMMENT '被关注者ID',
    `follower_id` BIGINT NOT NULL COMMENT '关注者ID',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `deleted_at` BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`followed_id`, `follower_id`),
    KEY `idx_follower_id` (`follower_id`) USING BTREE COMMENT '关注者索引',
    KEY `idx_followed_id` (`followed_id`) USING BTREE COMMENT '被关注者索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关注关系表';


DROP TABLE IF EXISTS `Activity`;
CREATE TABLE `Activity` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '动态ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `content` VARCHAR(255) NOT NULL COMMENT '动态内容',
    `media_url` VARCHAR(255) COMMENT '媒体URL',
    `visit_count` BIGINT NOT NULL COMMENT '浏览量',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `updated_at` BIGINT NOT NULL COMMENT '修改时间',
    `deleted_at` BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) USING BTREE COMMENT '用户索引',
    KEY `idx_created_at` (`created_at`) USING BTREE COMMENT '创建时间索引',
    KEY `idx_user_created` (`user_id`, `created_at`) USING BTREE COMMENT '用户与创建时间索引'
) ENGINE=InnoDB  AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='动态表';


DROP TABLE IF EXISTS `VideoReport`;
CREATE TABLE `VideoReport` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '举报ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `video_id` BIGINT NOT NULL COMMENT '视频ID',
    `reason` VARCHAR(255) NOT NULL COMMENT '举报原因',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `status` VARCHAR(255) NOT NULL COMMENT '举报状态',
    `resolved_at` BIGINT DEFAULT NULL COMMENT '解决时间',
    `admin_id` BIGINT NOT NULL COMMENT '管理员ID',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) USING BTREE COMMENT '用户索引',
    KEY `idx_video_id` (`video_id`) USING BTREE COMMENT '视频索引',
    KEY `idx_created_at` (`created_at`) USING BTREE COMMENT '创建时间索引',
    KEY `idx_status` (`status`) USING BTREE COMMENT '状态索引'
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='视频举报表';

DROP TABLE IF EXISTS `VideoLike`;
CREATE TABLE `VideoLike` (
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `video_id` BIGINT NOT NULL COMMENT '视频ID',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `deleted_at` BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`user_id`, `video_id`),
    KEY `idx_video_id` (`video_id`) USING BTREE COMMENT '视频索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频点赞表';

DROP TABLE IF EXISTS `ActivityComment`;
CREATE TABLE `ActivityComment` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '评论ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `activity_id` BIGINT NOT NULL COMMENT '动态ID',
    `parent_id` BIGINT NOT NULL COMMENT '父评论ID',
    `content` VARCHAR(255) NOT NULL COMMENT '评论内容',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `updated_at` BIGINT NOT NULL COMMENT '修改时间',
    `deleted_at` BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) USING BTREE COMMENT '用户索引',
    KEY `idx_activity_id` (`activity_id`) USING BTREE COMMENT '动态索引'
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='动态评论表';

DROP TABLE IF EXISTS `ActivityCommentLike`;
CREATE TABLE `ActivityCommentLike` (
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `comment_id` BIGINT NOT NULL COMMENT '动态评论ID',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `deleted_at` BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`user_id`, `comment_id`),
    KEY `idx_comment_id` (`comment_id`) USING BTREE COMMENT '动态评论索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='动态评论点赞表';

DROP TABLE IF EXISTS `ActivityLike`;
CREATE TABLE `ActivityLike` (
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `activity_id` BIGINT NOT NULL COMMENT '动态ID',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `deleted_at` BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`user_id`, `activity_id`),
    KEY `idx_activity_id` (`activity_id`) USING BTREE COMMENT '动态索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='动态点赞表';

DROP TABLE IF EXISTS `ActivityCommentReport`;
CREATE TABLE `ActivityCommentReport` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '举报ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `comment_id` BIGINT NOT NULL COMMENT '动态评论ID',
    `reason` VARCHAR(255) NOT NULL COMMENT '举报原因',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `status` VARCHAR(255) NOT NULL COMMENT '举报状态',
    `resolved_at` BIGINT DEFAULT NULL COMMENT '解决时间',
    `admin_id` BIGINT NOT NULL COMMENT '管理员ID',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) USING BTREE COMMENT '用户索引',
    KEY `idx_comment_id` (`comment_id`) USING BTREE COMMENT '动态评论索引',
    KEY `idx_status` (`status`) USING BTREE COMMENT '状态索引'
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='动态评论举报表';

DROP TABLE IF EXISTS `ActivityReport`;
CREATE TABLE `ActivityReport` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '举报ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `activity_id` BIGINT NOT NULL COMMENT '动态ID',
    `reason` VARCHAR(255) NOT NULL COMMENT '举报原因',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `status` VARCHAR(255) NOT NULL COMMENT '举报状态',
    `resolved_at` BIGINT DEFAULT NULL COMMENT '解决时间',
    `admin_id` BIGINT NOT NULL COMMENT '管理员ID',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) USING BTREE COMMENT '用户索引',
    KEY `idx_activity_id` (`activity_id`) USING BTREE COMMENT '动态索引',
    KEY `idx_status` (`status`) USING BTREE COMMENT '状态索引'
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='动态举报表';

DROP TABLE IF EXISTS `Tag`;
CREATE TABLE `Tag` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '标签ID',
    `tag_name` VARCHAR(255) NOT NULL COMMENT '标签名称',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `deleted_at` BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_tag_name` (`tag_name`) USING BTREE COMMENT '标签名称唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='标签静态库表';

DROP TABLE IF EXISTS `VideoTag`;
CREATE TABLE `VideoTag` (
    `video_id` BIGINT NOT NULL COMMENT '视频ID',
    `tag_id` BIGINT NOT NULL COMMENT '标签ID',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `deleted_at` BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`video_id`,`tag_id`),
    KEY `idx_video_id` (`video_id`) USING BTREE COMMENT '视频索引',
    KEY `idx_tag_id` (`tag_id`) USING BTREE COMMENT '标签索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频标签关联表';

DROP TABLE IF EXISTS `Review`;
CREATE TABLE `Review` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '审核ID',
    `video_id` BIGINT NOT NULL COMMENT '视频ID',
    `reviewer_id` BIGINT NOT NULL COMMENT '审核者ID',
    `submitted_at` BIGINT NOT NULL COMMENT '提交时间',
    `reviewed_at` BIGINT NOT NULL COMMENT '审核时间',
    `review_result` VARCHAR(255) NOT NULL COMMENT '审核结果',
    PRIMARY KEY (`id`),
    KEY `idx_video_id` (`video_id`) USING BTREE COMMENT '视频索引',
    KEY `idx_reviewer_id` (`reviewer_id`) USING BTREE COMMENT '审核者索引'
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='视频审核表';

DROP TABLE IF EXISTS `VideoCommentReport`;
CREATE TABLE `VideoCommentReport` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '举报ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `comment_id` BIGINT NOT NULL COMMENT '视频评论ID',
    `reason` VARCHAR(255) NOT NULL COMMENT '举报原因',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `status` VARCHAR(255) NOT NULL COMMENT '举报状态',
    `resolved_at` BIGINT DEFAULT NULL COMMENT '解决时间',
    `admin_id` BIGINT NOT NULL COMMENT '管理员ID',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) USING BTREE COMMENT '用户索引',
    KEY `idx_comment_id` (`comment_id`) USING BTREE COMMENT '视频评论索引',
    KEY `idx_status` (`status`) USING BTREE COMMENT '状态索引'
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='视频评论举报表';

DROP TABLE IF EXISTS `VideoCommentLike`;
CREATE TABLE `VideoCommentLike` (
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `comment_id` BIGINT NOT NULL COMMENT '视频评论ID',
    `created_at` BIGINT NOT NULL COMMENT '创建时间',
    `deleted_at` BIGINT DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`user_id`, `comment_id`),
    KEY `idx_comment_id` (`comment_id`) USING BTREE COMMENT '视频评论索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频评论点赞表';