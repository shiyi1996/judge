-- MySQL dump 10.13  Distrib 5.5.53, for Linux (x86_64)
--
-- Host: localhost    Database: fightcoder
-- ------------------------------------------------------
-- Server version	5.5.53

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


DROP DATABASE IF EXISTS `fightcoder`;
CREATE DATABASE `fightcoder`;
USE `fightcoder`;

--
-- Table structure for table `account`
--

DROP TABLE IF EXISTS `account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `email` varchar(50) NOT NULL COMMENT '邮箱',
  `password` varchar(30) NOT NULL COMMENT '密码',
  `phone` varchar(20) NOT NULL COMMENT '手机号',
  `qq_number` varchar(20) DEFAULT NULL COMMENT 'QQ号',
  `qq_id` int(11) DEFAULT NULL COMMENT '用于第三方登录',
  PRIMARY KEY (`id`),
  KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account`
--

LOCK TABLES `account` WRITE;
/*!40000 ALTER TABLE `account` DISABLE KEYS */;
/*!40000 ALTER TABLE `account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `contest`
--

DROP TABLE IF EXISTS `contest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `contest` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '比赛id',
  `user_id` bigint(20) NOT NULL COMMENT '举办者id',
  `type` int(11) NOT NULL COMMENT '比赛类型',
  `name` varchar(50) NOT NULL COMMENT '比赛名称',
  `description` varchar(200) DEFAULT NULL COMMENT '比赛描述',
  `problem_list` varchar(500) NOT NULL COMMENT '题目列表',
  `start_time` bigint(20) NOT NULL COMMENT '开始时间',
  `frozen_time` bigint(20) NOT NULL COMMENT '封榜时间',
  `end_time` bigint(20) NOT NULL COMMENT '结束时间',
  `password` varchar(20) DEFAULT NULL COMMENT '密码，为空表示公开',
  PRIMARY KEY (`id`),
  KEY `id` (`id`),
  KEY `user_id` (`user_id`),
  KEY `type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contest`
--

LOCK TABLES `contest` WRITE;
/*!40000 ALTER TABLE `contest` DISABLE KEYS */;
/*!40000 ALTER TABLE `contest` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group`
--

DROP TABLE IF EXISTS `group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '团队表id',
  `picture` varchar(50) DEFAULT NULL COMMENT '头像',
  `name` varchar(20) NOT NULL COMMENT '团队名称',
  `description` varchar(200) DEFAULT NULL COMMENT '团队描述',
  `set_up` varchar(500) DEFAULT NULL COMMENT '设置',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group`
--

LOCK TABLES `group` WRITE;
/*!40000 ALTER TABLE `group` DISABLE KEYS */;
/*!40000 ALTER TABLE `group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rank_group`
--

DROP TABLE IF EXISTS `rank_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rank_group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '团队排行id',
  `group_id` bigint(20) NOT NULL COMMENT '团队id',
  `ac_number` float DEFAULT NULL COMMENT 'AC数量',
  `grade` float DEFAULT NULL COMMENT '得分',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rank_group`
--

LOCK TABLES `rank_group` WRITE;
/*!40000 ALTER TABLE `rank_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `rank_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `letter`
--

DROP TABLE IF EXISTS `letter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `letter` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '站内信id',
  `content` varchar(1000) NOT NULL COMMENT '内容',
  `send_user_id` bigint(20) NOT NULL COMMENT '发送者id',
  `receive_user_id` bigint(20) NOT NULL COMMENT '接收者id',
  `send_time` bigint(20) NOT NULL COMMENT '发送时间',
  `is_read` tinyint(1) DEFAULT NULL COMMENT '是否已读',
  PRIMARY KEY (`id`),
  KEY `send_user_id` (`send_user_id`),
  KEY `receive_user_id` (`receive_user_id`),
  KEY `is_read` (`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `letter`
--

LOCK TABLES `letter` WRITE;
/*!40000 ALTER TABLE `letter` DISABLE KEYS */;
/*!40000 ALTER TABLE `letter` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `member`
--

DROP TABLE IF EXISTS `member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `member` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '成员表id',
  `group_id` bigint(20) NOT NULL COMMENT '团队Id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `role` varchar(255) DEFAULT NULL COMMENT '角色',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `group_id` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `member`
--

LOCK TABLES `member` WRITE;
/*!40000 ALTER TABLE `member` DISABLE KEYS */;
/*!40000 ALTER TABLE `member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `notice`
--

DROP TABLE IF EXISTS `notice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `notice` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '通知id',
  `user_id` bigint(20) NOT NULL COMMENT '用户Id',
  `time` bigint(20) NOT NULL COMMENT '时间',
  `content` varchar(255) NOT NULL COMMENT '内容',
  `is_read` tinyint(1) DEFAULT NULL COMMENT '是否已读',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `is_read` (`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `notice`
--

LOCK TABLES `notice` WRITE;
/*!40000 ALTER TABLE `notice` DISABLE KEYS */;
/*!40000 ALTER TABLE `notice` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `player`
--

DROP TABLE IF EXISTS `player`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '参与比赛者表id',
  `contest_id` bigint(20) NOT NULL COMMENT '竞赛id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `contest_id` (`contest_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `player`
--

LOCK TABLES `player` WRITE;
/*!40000 ALTER TABLE `player` DISABLE KEYS */;
/*!40000 ALTER TABLE `player` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `problem`
--

DROP TABLE IF EXISTS `problem`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `problem` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) NOT NULL COMMENT '题目提供者',
  `test_data` varchar(2000) COMMENT '测试数据',
  `titile` varchar(50) NOT NULL COMMENT '题目标题',
  `description` varchar(500) NOT NULL COMMENT '题目描述',
  `input_des` varchar(300) NOT NULL COMMENT '输入描述',
  `output_des` varchar(300) NOT NULL COMMENT '输出描述',
  `case` varchar(200) NOT NULL COMMENT '测试样例',
  `hint` varchar(300) DEFAULT NULL COMMENT '题目提示(可以为对样例输入输出的解释)',
  `time_limit` int(11) DEFAULT NULL COMMENT '时间限制',
  `memory_limit` int(11) DEFAULT NULL COMMENT '内存限制',
  `tag` varchar(200) DEFAULT NULL COMMENT '题目标签',
  `is_special_judge` tinyint(1) DEFAULT NULL COMMENT '是否特判',
  `special_judge_source` varchar(50) DEFAULT NULL COMMENT '特判程序源代码',
  `code` varchar(50) DEFAULT NULL COMMENT '标准程序',
  `language_limit` varchar(100) DEFAULT NULL COMMENT '语言限制',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `problem`
--

LOCK TABLES `problem` WRITE;
/*!40000 ALTER TABLE `problem` DISABLE KEYS */;
/*!40000 ALTER TABLE `problem` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `problem_check`
--

DROP TABLE IF EXISTS `problem_check`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `problem_check` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) NOT NULL COMMENT '题目提供者',
  `test_data` varchar(2000) COMMENT '测试数据',
  `titile` varchar(50) NOT NULL COMMENT '题目标题',
  `description` varchar(500) NOT NULL COMMENT '题目描述',
  `input_des` varchar(300) NOT NULL COMMENT '输入描述',
  `output_des` varchar(300) NOT NULL COMMENT '输出描述',
  `case` varchar(200) NOT NULL COMMENT '测试样例',
  `hint` varchar(300) DEFAULT NULL COMMENT '题目提示(可以为对样例输入输出的解释)',
  `time_limit` int(11) DEFAULT NULL COMMENT '时间限制',
  `memory_limit` int(11) DEFAULT NULL COMMENT '内存限制',
  `tag` varchar(200) DEFAULT NULL COMMENT '题目标签',
  `is_special_judge` tinyint(1) DEFAULT NULL COMMENT '是否特判',
  `special_judge_source` varchar(50) DEFAULT NULL COMMENT '特判程序源代码',
  `code` varchar(50) DEFAULT NULL COMMENT '标准程序',
  `language_limit` varchar(100) DEFAULT NULL COMMENT '语言限制',
  `check_status` varchar(100) DEFAULT NULL COMMENT '审核状态',
  `problem_id` bigint(20) DEFAULT NULL COMMENT '所在正式题库的Id',
  `problem_user_id` bigint(20) NOT NULL COMMENT '所在私人题库的Id',
  PRIMARY KEY (`id`),
  KEY `check_status` (`check_status`),
  KEY `problem_user_id` (`problem_user_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `problem_check`
--

LOCK TABLES `problem_check` WRITE;
/*!40000 ALTER TABLE `problem_check` DISABLE KEYS */;
/*!40000 ALTER TABLE `problem_check` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `problem_user`
--

DROP TABLE IF EXISTS `problem_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `problem_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) NOT NULL COMMENT '题目提供者',
  `test_data` varchar(2000) COMMENT '测试数据',
  `titile` varchar(50) NOT NULL COMMENT '题目标题',
  `description` varchar(500) NOT NULL COMMENT '题目描述',
  `input_des` varchar(300) NOT NULL COMMENT '输入描述',
  `output_des` varchar(300) NOT NULL COMMENT '输出描述',
  `case` varchar(200) NOT NULL COMMENT '测试样例',
  `hint` varchar(300) DEFAULT NULL COMMENT '题目提示(可以为对样例输入输出的解释)',
  `time_limit` int(11) DEFAULT NULL COMMENT '时间限制',
  `memory_limit` int(11) DEFAULT NULL COMMENT '内存限制',
  `tag` varchar(200) DEFAULT NULL COMMENT '题目标签',
  `is_special_judge` tinyint(1) DEFAULT NULL COMMENT '是否特判',
  `special_judge_source` varchar(50) DEFAULT NULL COMMENT '特判程序源代码',
  `code` varchar(50) DEFAULT NULL COMMENT '标准程序',
  `language_limit` varchar(100) DEFAULT NULL COMMENT '语言限制',
  `note` varchar(255) DEFAULT NULL COMMENT '备注(用户不可见)',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `problem_user`
--

LOCK TABLES `problem_user` WRITE;
/*!40000 ALTER TABLE `problem_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `problem_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rank_person`
--

DROP TABLE IF EXISTS `rank_person`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rank_person` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '个人排行id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `ac_number` int(11) DEFAULT NULL COMMENT 'AC数量',
  `grade` int(11) DEFAULT NULL COMMENT '得分',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rank_person`
--

LOCK TABLES `rank_person` WRITE;
/*!40000 ALTER TABLE `rank_person` DISABLE KEYS */;
/*!40000 ALTER TABLE `rank_person` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `save_code`
--

DROP TABLE IF EXISTS `save_code`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `save_code` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `problem_id` bigint(20) NOT NULL COMMENT '问题id',
  `user_id` bigint(20) NOT NULL COMMENT '提交用户Id',
  `code` varchar(50) NOT NULL COMMENT '代码',
  PRIMARY KEY (`id`),
  KEY `problem_id` (`problem_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `save_code`
--

LOCK TABLES `save_code` WRITE;
/*!40000 ALTER TABLE `save_code` DISABLE KEYS */;
/*!40000 ALTER TABLE `save_code` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `submit`
--

DROP TABLE IF EXISTS `submit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `submit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `problem_id` bigint(20) NOT NULL COMMENT '题目ID',
  `problem_type` varchar(20) NOT NULL COMMENT '题库类型',
  `user_id` bigint(20) NOT NULL COMMENT '提交用户ID',
  `language` varchar(20) NOT NULL COMMENT '提交语言',
  `submit_time` bigint(20) NOT NULL COMMENT '提交时间',
  `running_time` int(11) DEFAULT NULL COMMENT '耗时(ms)',
  `running_memory` int(11) DEFAULT NULL COMMENT '所占空间',
  `result` int DEFAULT NULL COMMENT '运行结果',
  `result_des` varchar(300) DEFAULT NULL COMMENT '结果描述',
  `code` varchar(50) NOT NULL COMMENT '提交代码',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `result` (`result`),
  KEY `problem_id` (`problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `submit`
--

LOCK TABLES `submit` WRITE;
/*!40000 ALTER TABLE `submit` DISABLE KEYS */;
/*!40000 ALTER TABLE `submit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `submit_contest`
--

DROP TABLE IF EXISTS `submit_contest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `submit_contest` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `problem_id` bigint(20) NOT NULL COMMENT '题目ID',
  `problem_type` varchar(20) NOT NULL COMMENT '题库类型',
  `user_id` bigint(20) NOT NULL COMMENT '提交用户ID',
  `language` varchar(20) NOT NULL COMMENT '提交语言',
  `submit_time` bigint(20) NOT NULL COMMENT '提交时间',
  `running_time` int(11) DEFAULT NULL COMMENT '耗时(ms)',
  `running_memory` int(11) DEFAULT NULL COMMENT '所占空间',
  `result` int DEFAULT NULL COMMENT '运行结果',
  `result_des` varchar(300) DEFAULT NULL COMMENT '结果描述',
  `code` varchar(50) NOT NULL COMMENT '提交代码',
  `contest_id` bigint(20) NOT NULL COMMENT '比赛Id',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `result` (`result`),
  KEY `contest_id` (`contest_id`),
  KEY `problem_id` (`problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `submit_contest`
--

LOCK TABLES `submit_contest` WRITE;
/*!40000 ALTER TABLE `submit_contest` DISABLE KEYS */;
/*!40000 ALTER TABLE `submit_contest` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `submit_user`
--

DROP TABLE IF EXISTS `submit_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `submit_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `problem_id` bigint(20) NOT NULL COMMENT '题目ID',
  `problem_type` varchar(20) NOT NULL COMMENT '题库类型',
  `user_id` bigint(20) NOT NULL COMMENT '提交用户ID',
  `language` varchar(20) NOT NULL COMMENT '提交语言',
  `submit_time` bigint(20) NOT NULL COMMENT '提交时间',
  `running_time` int(11) DEFAULT NULL COMMENT '耗时(ms)',
  `running_memory` int(11) DEFAULT NULL COMMENT '所占空间',
  `result` int DEFAULT NULL COMMENT '运行结果',
  `result_des` varchar(300) DEFAULT NULL COMMENT '结果描述',
  `code` varchar(50) NOT NULL COMMENT '提交代码',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `result` (`result`),
  KEY `problem_id` (`problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `submit_user`
--

LOCK TABLES `submit_user` WRITE;
/*!40000 ALTER TABLE `submit_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `submit_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `account_id` bigint(20) NOT NULL COMMENT '账号Id',
  `nick_name` varchar(20) NOT NULL COMMENT '昵称',
  `description` varchar(200) DEFAULT NULL COMMENT '个人描述',
  `sex` int(11) DEFAULT NULL COMMENT '性别',
  `birthday` bigint(20) DEFAULT NULL COMMENT '生日',
  `daily_address` varchar(100) DEFAULT NULL COMMENT '日常所在地：省、市',
  `recv_address` varchar(100) DEFAULT NULL COMMENT '收件地址，仅自己可见',
  `t_shirt_size` varchar(50) DEFAULT NULL COMMENT 'T-恤尺码(S、M、L、XL、XXL、XXL)',
  `stat_school` int(11) DEFAULT NULL COMMENT '当前就学状态(小学及以下、中学学生、大学学生、非在校生)',
  `blog` varchar(100) DEFAULT NULL COMMENT '博客地址',
  `git` varchar(100) DEFAULT NULL COMMENT 'Git地址',
  `avator` varchar(50) DEFAULT NULL COMMENT '头像',
  PRIMARY KEY (`id`),
  KEY `account_id` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_relation`
--

DROP TABLE IF EXISTS `user_relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_relation` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `leader_id` bigint(20) DEFAULT NULL COMMENT '领导者：即被关注者',
  `follower_id` bigint(20) DEFAULT NULL COMMENT '追随者：即关注者',
  PRIMARY KEY (`id`),
  KEY `leader_id` (`leader_id`),
  KEY `follower_id` (`follower_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_relation`
--

LOCK TABLES `user_relation` WRITE;
/*!40000 ALTER TABLE `user_relation` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_relation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_set`
--

DROP TABLE IF EXISTS `user_set`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_set` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户Id',
  `editor_set` varchar(255) DEFAULT NULL COMMENT '编辑器设置(JSON)',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_set`
--

LOCK TABLES `user_set` WRITE;
/*!40000 ALTER TABLE `user_set` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_set` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-11-18  1:22:43