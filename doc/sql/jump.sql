-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2023-05-20 16:55:11
-- 服务器版本： 8.0.12
-- PHP 版本： 8.0.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `jump`
--

-- --------------------------------------------------------

--
-- 表的结构 `api_log`
--

CREATE TABLE `api_log` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求方式',
  `request_ip` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求ip',
  `request_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求url',
  `query_params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '请求参数',
  `response_data` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '返回的数据 不包含data'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='请求日志表';

--
-- 转存表中的数据 `api_log`
--

INSERT INTO `api_log` (`id`, `create_at`, `update_at`, `method`, `request_ip`, `request_url`, `query_params`, `response_data`) VALUES
(1, '2021-06-25 00:28:59', '2021-06-25 00:28:59', 'POST', '[\"127.0.0.1\",\"127.0.0.1\"]', '//127.0.0.1:8787/api/getData//127.0.0.1:8787/api/getData', '[]', '{\"code\":200,\"data\":[]}'),
(2, '2021-06-25 00:29:06', '2021-06-25 00:29:06', 'POST', '[\"127.0.0.1\",\"127.0.0.1\"]', '//127.0.0.1:8787/api/getData//127.0.0.1:8787/api/getData?2222', '{\"2222\":\"\"}', '{\"code\":200,\"data\":[]}'),
(3, '2021-06-25 00:29:09', '2021-06-25 00:29:09', 'POST', '[\"127.0.0.1\",\"127.0.0.1\"]', '//127.0.0.1:8787/api/getData//127.0.0.1:8787/api/getData?2222', '{\"2222\":\"\"}', '{\"code\":200,\"data\":[]}'),
(4, '2021-06-25 00:29:51', '2021-06-25 00:29:51', 'POST', '[\"127.0.0.1\",\"127.0.0.1\"]', '/api/getData//127.0.0.1:8787/api/getData?2222', '[{\"2222\":\"\"},\"\"]', '{\"code\":200,\"data\":[]}'),
(5, '2021-06-25 00:31:21', '2021-06-25 00:31:21', 'GET', '[\"127.0.0.1\",\"127.0.0.1\"]', '/2', '[[],\"\"]', '');

-- --------------------------------------------------------

--
-- 表的结构 `api_user`
--

CREATE TABLE `api_user` (
  `id` bigint(20) NOT NULL,
  `api_key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口用户名',
  `api_secret` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口秘钥',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `ip_list` json DEFAULT NULL COMMENT 'api接口访问ip列表',
  `status` tinyint(1) NOT NULL DEFAULT '10' COMMENT '10默认未启用20启用30禁止',
  `end_date` datetime DEFAULT NULL COMMENT 'api终止使用日期',
  `deleted_at` datetime DEFAULT NULL COMMENT '是否删除'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='接口用户表';

--
-- 转存表中的数据 `api_user`
--

INSERT INTO `api_user` (`id`, `api_key`, `api_secret`, `create_time`, `update_time`, `ip_list`, `status`, `end_date`, `deleted_at`) VALUES
(1, 'key', 'secret', '2023-05-20 13:31:09', '2023-05-20 13:31:09', NULL, 10, NULL, NULL),
(2, 'key111', 'secret1111', '2023-05-20 13:31:09', '2023-05-20 13:31:09', NULL, 10, NULL, NULL);

-- --------------------------------------------------------

--
-- 表的结构 `jump_link`
--

CREATE TABLE `jump_link` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `jump_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '需要跳转到url',
  `hash_key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'id的哈希值',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '是否删除',
  `end_time` datetime DEFAULT NULL COMMENT '最后存活时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='链接跳转数据表';

--
-- 转存表中的数据 `jump_link`
--

INSERT INTO `jump_link` (`id`, `jump_url`, `hash_key`, `created_at`, `updated_at`, `deleted_at`, `end_time`) VALUES
(33, 'https%3A%2F%2Fwww.cnblogs.com%2Fchenqionghe%2Fp%2F13409556.html', '174200537', '2023-05-20 16:52:34', '2023-05-20 16:52:34', NULL, '2023-05-25 08:00:00');

--
-- 转储表的索引
--

--
-- 表的索引 `api_log`
--
ALTER TABLE `api_log`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `api_user`
--
ALTER TABLE `api_user`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `jump_link`
--
ALTER TABLE `jump_link`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `api_log`
--
ALTER TABLE `api_log`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `api_user`
--
ALTER TABLE `api_user`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `jump_link`
--
ALTER TABLE `jump_link`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
