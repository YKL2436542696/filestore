CREATE TABLE `tbl_file` (
                            `id` INT ( 11 ) NOT NULL AUTO_INCREMENT,
                            `file_sha1` CHAR ( 40 ) NOT NULL DEFAULT '' COMMENT '文件 HASH ',
                            `file_name` VARCHAR ( 256 ) NOT NULL DEFAULT '' COMMENT '文件名',
                            `file_size` BIGINT ( 20 ) DEFAULT '0' COMMENT '文件大小',
                            `file_addr` VARCHAR ( 1024 ) NOT NULL DEFAULT '' COMMENT '件存储位置',
                            `create_at` DATETIME DEFAULT NOw() COMMENT '创建日期 ',
                            `update_at` DATETIME DEFAULT NOw() ON UPDATE CURRENT_TIMESTAMP () COMMENT '更新日期',
                            `status` INT ( 11 ) NOT NULL DEFAULT '0' COMMENT '状态(可用/禁用/已删除等状态)',
                            `ext1` INT ( 11 ) DEFAULT '0' COMMENT '备用字段1',
                            `ext2` TEXT COMMENT '备用字段2',
                            PRIMARY KEY ( `id` ),
                            UNIQUE KEY `idx_file_hash` (`file_sha1`),
                            KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `tbl_user` (
                            `id` INT ( 11 ) NOT NULL AUTO_INCREMENT,
                            `user_name` VARCHAR ( 64 ) NOT NULL DEFAULT '' COMMENT '用户名',
                            `user_pwd` VARCHAR ( 256 ) NOT NULL DEFAULT '' COMMENT '用户encoded密码',
                            `email` VARCHAR ( 64 ) DEFAULT '' COMMENT '邮箱',
                            `phone` VARCHAR ( 128 ) DEFAULT '' COMMENT '手机号',
                            `email_validated` TINYINT ( 1 ) DEFAULT 0 COMMENT '邮箱是否已验证',
                            `phone_validated` TINYINT ( 1 ) DEFAULT 0 COMMENT '手机号是否已验证',
                            `signup_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '注册日期',
                            `last_active` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后活跃时间戳',
                            `profile` TEXT COMMENT '用户属性',
                            `status` INT ( 11 ) NOT NULL DEFAULT '0' COMMENT '账户状态 (启用/禁用/锁定/标记删除等) ',
                            PRIMARY KEY ( `id`),
                            KEY `idx_phone`(`phone`),
                            KEY `idx_status`(`status`)
)ENGINE=innoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;


