-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2023-05-06 14:07:09
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
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求方式',
  `request_ip` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求ip',
  `request_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求url',
  `query_params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '请求参数',
  `response_data` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '返回的数据 不包含data',
  `exec_time` datetime DEFAULT NULL COMMENT '开始执行的时间',
  `response_time` datetime DEFAULT NULL COMMENT '返回结果的时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='请求日志表';

--
-- 转存表中的数据 `api_log`
--

INSERT INTO `api_log` (`id`, `create_time`, `update_time`, `method`, `request_ip`, `request_url`, `query_params`, `response_data`, `exec_time`, `response_time`) VALUES
(1, '2021-06-25 00:28:59', '2021-06-25 00:28:59', 'POST', '[\"127.0.0.1\",\"127.0.0.1\"]', '//127.0.0.1:8787/api/getData//127.0.0.1:8787/api/getData', '[]', '{\"code\":200,\"data\":[]}', '2021-06-25 00:28:59', '2021-06-25 00:28:59'),
(2, '2021-06-25 00:29:06', '2021-06-25 00:29:06', 'POST', '[\"127.0.0.1\",\"127.0.0.1\"]', '//127.0.0.1:8787/api/getData//127.0.0.1:8787/api/getData?2222', '{\"2222\":\"\"}', '{\"code\":200,\"data\":[]}', '2021-06-25 00:29:06', '2021-06-25 00:29:06'),
(3, '2021-06-25 00:29:09', '2021-06-25 00:29:09', 'POST', '[\"127.0.0.1\",\"127.0.0.1\"]', '//127.0.0.1:8787/api/getData//127.0.0.1:8787/api/getData?2222', '{\"2222\":\"\"}', '{\"code\":200,\"data\":[]}', '2021-06-25 00:29:09', '2021-06-25 00:29:09'),
(4, '2021-06-25 00:29:51', '2021-06-25 00:29:51', 'POST', '[\"127.0.0.1\",\"127.0.0.1\"]', '/api/getData//127.0.0.1:8787/api/getData?2222', '[{\"2222\":\"\"},\"\"]', '{\"code\":200,\"data\":[]}', '2021-06-25 00:29:51', '2021-06-25 00:29:51'),
(5, '2021-06-25 00:31:21', '2021-06-25 00:31:21', 'GET', '[\"127.0.0.1\",\"127.0.0.1\"]', '/2', '[[],\"\"]', '', '2021-06-25 00:31:21', '2021-06-25 00:31:21');

-- --------------------------------------------------------

--
-- 表的结构 `jump_link`
--

CREATE TABLE `jump_link` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `jump_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '需要跳转到url',
  `md5_data` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'url的md5值',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '是否删除',
  `end_time` datetime DEFAULT NULL COMMENT '最后存活时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='链接跳转数据表';

--
-- 转存表中的数据 `jump_link`
--

INSERT INTO `jump_link` (`id`, `jump_url`, `md5_data`, `created_at`, `updated_at`, `deleted_at`, `end_time`) VALUES
(1, 'https%3A%2F%2Fwww.baidu.com%2F', 'e81c1f5749545c5f7d247b3a100ffe62', '2021-06-24 22:41:04', '2021-06-24 22:41:04', NULL, '2021-06-30 00:00:00'),
(2, 'https%3A%2F%2Fwww.baidu.com%2F', 'e81c1f5749545c5f7d247b3a100ffe62', '2021-06-25 00:13:37', '2021-06-25 00:13:37', NULL, '2021-06-30 00:00:00');

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `avatar` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deleted_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='111';

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`id`, `name`, `avatar`, `deleted_at`) VALUES
(1, '1111', '2222', NULL),
(2, '333', '4444', NULL),
(3, '001111', '00222000', NULL),
(4, '30033000', '00000676', NULL),
(5, '1111', '2222', NULL),
(6, '333', '4444', NULL),
(7, '001111', '00222000', NULL),
(8, '30033000', '00000676', NULL);

--
-- 转储表的索引
--

--
-- 表的索引 `api_log`
--
ALTER TABLE `api_log`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `jump_link`
--
ALTER TABLE `jump_link`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user`
--
ALTER TABLE `user`
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
-- 使用表AUTO_INCREMENT `jump_link`
--
ALTER TABLE `jump_link`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
