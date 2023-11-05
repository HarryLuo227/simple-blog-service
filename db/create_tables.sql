-- 建立標籤表
CREATE TABLE `blog_tag` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) DEFAULT '' COMMENT '標籤名稱',
    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '建立時間',
    `created_by` VARCHAR(100) DEFAULT '' COMMENT '建立人',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改時間',
    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '刪除時間',
    `is_del` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT '是否刪除 0 為未刪除、1 為已刪除',
    `state` TINYINT(3) UNSIGNED DEFAULT '1' COMMENT '狀態 0 為禁用、1 為啟用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='標籤管理';

-- 建立文章表
CREATE TABLE `blog_article` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(100) DEFAULT '' COMMENT '文章標題',
    `desc` VARCHAR(255) DEFAULT '' COMMENT '文章簡述',
    `cover_image_url` VARCHAR(255) DEFAULT '' COMMENT '封面圖片位址',
    `content` LONGTEXT COMMENT '文章內容',
    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '建立時間',
    `created_by` VARCHAR(100) DEFAULT '' COMMENT '建立人',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改時間',
    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '刪除時間',
    `is_del` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT '是否刪除 0 為未刪除、1 為已刪除',
    `state` TINYINT(3) UNSIGNED DEFAULT '1' COMMENT '狀態 0 為禁用、1 為啟用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

-- 建立文章標籤連結表
CREATE TABLE `blog_article_tag` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `article_id` INT(11) NOT NULL COMMENT '文章 ID',
    `tag_id` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '標籤 ID',
    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '建立時間',
    `created_by` VARCHAR(100) DEFAULT '' COMMENT '建立人',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改時間',
    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '刪除時間',
    `is_del` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT '是否刪除 0 為未刪除、1 為已刪除',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章標籤連結';

-- 建立認證表
CREATE TABLE `blog_auth` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user` VARCHAR(20) DEFAULT '' COMMENT 'User',
	`password` VARCHAR(50) DEFAULT '' COMMENT 'Password',
    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '建立時間',
    `created_by` VARCHAR(100) DEFAULT '' COMMENT '建立人',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改時間',
    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '刪除時間',
    `is_del` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT '是否刪除 0 為未刪除、1 為已刪除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='認證管理';