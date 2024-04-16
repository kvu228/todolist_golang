-- MySQL dump 10.13  Distrib 8.0.35, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: to_do_list
-- ------------------------------------------------------
-- Server version	8.0.36

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `images`
--

DROP TABLE IF EXISTS `images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `images` (
  `id` varchar(36) NOT NULL,
  `title` varchar(150) DEFAULT NULL,
  `file_name` varchar(36) DEFAULT NULL,
  `file_size` int DEFAULT NULL,
  `file_type` varchar(20) DEFAULT NULL,
  `storage_provider` enum('aws_s3','local') DEFAULT NULL,
  `status` enum('activated','uploaded','deleted') DEFAULT 'uploaded',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `uploaded_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `images_id_index` (`id`),
  KEY `images_title_index` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `images`
--

LOCK TABLES `images` WRITE;
/*!40000 ALTER TABLE `images` DISABLE KEYS */;
INSERT INTO `images` VALUES ('018edc0a-57c1-7947-9746-20a47951819f','sample','1713088626201670600_download.png',56136,'image/png','aws_s3','activated','2024-04-14 16:57:29','2024-04-14 16:57:29');
/*!40000 ALTER TABLE `images` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `posts`
--

DROP TABLE IF EXISTS `posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `posts` (
  `id` varchar(36) NOT NULL,
  `title` varchar(50) NOT NULL,
  `body` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` enum('doing','done','delete') DEFAULT 'doing',
  `owner_id` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `posts_id_index` (`id`),
  KEY `posts_status_index` (`status`),
  KEY `posts_title_index` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;
INSERT INTO `posts` VALUES ('018eb46e-70c7-7b2e-af30-3f7468c7eda0','Test 1','lorem ipsum','2024-04-06 17:22:55','2024-04-08 02:01:11','doing','018ebb6f-ddb3-70e5-871a-332e9d84c6f3'),('018eb46e-f26e-7d34-b4ed-f0ecbbbc301b','Test 2','lorem ipsum','2024-04-06 17:22:56','2024-04-08 02:01:11','doing','018ebb6f-ddb3-70e5-871a-332e9d84c6f3'),('018edaef-de3c-7118-8513-62da9f719b2b','test title','lorem ipsum','2024-04-14 11:48:57','2024-04-14 11:48:57','doing','018ed68c-963a-7e1c-9305-3f49f85ed05d'),('018edc57-83b6-70a4-92f7-b3c3ac0aeca9','test title 2','lorem ipsum','2024-04-14 18:21:47','2024-04-14 18:21:47','doing','018ed68c-963a-7e1c-9305-3f49f85ed05d'),('018edc7b-149a-7980-affb-c3c48a6ec96e','test title 3','lorem ipsum','2024-04-14 19:00:38','2024-04-14 19:00:38','doing','018ed68c-963a-7e1c-9305-3f49f85ed05d');
/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_sessions`
--

DROP TABLE IF EXISTS `user_sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_sessions` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36) NOT NULL,
  `refresh_token` varchar(80) NOT NULL,
  `refresh_token_exp_at` timestamp NOT NULL,
  `access_token_exp_at` timestamp NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_sessions`
--

LOCK TABLES `user_sessions` WRITE;
/*!40000 ALTER TABLE `user_sessions` DISABLE KEYS */;
INSERT INTO `user_sessions` VALUES ('018ebecf-5ef7-7cd0-aeae-ca5cbbf744cc','018ebb6f-ddb3-70e5-871a-332e9d84c6f3','c1aded913a36b6d5f4e4b3f8eb00d4527a1b0f10be4f17e12cfb3ccb0bc0','2024-04-23 00:44:05','2024-04-16 00:44:05','2024-04-08 17:44:05','2024-04-08 17:44:05'),('018ec3dc-1f7b-7da3-ba54-968c210a02f4','018ebb6f-ddb3-70e5-871a-332e9d84c6f3','75f4376ea739b07f8af3ccc432db205f8596b2d2af7943d1bea5afc0bfa1','2024-04-24 00:16:07','2024-04-17 00:16:07','2024-04-09 17:16:06','2024-04-09 17:16:06'),('018ed68d-79d1-7f2b-83eb-7a13356158d5','018ed68c-963a-7e1c-9305-3f49f85ed05d','06499502831933af2ae051a8b6ab5d91e88f92689b738280f1450675f1dc','2024-04-27 15:23:00','2024-04-20 15:23:00','2024-04-13 08:22:59','2024-04-13 08:22:59'),('018edae3-5fd2-76f2-9697-20da3a17e111','018ed68c-963a-7e1c-9305-3f49f85ed05d','749fdcf726efac46c7c0486c9833d4e4b07daf75847d46883044ff2b3613','2024-04-28 11:35:18','2024-04-21 11:35:18','2024-04-14 04:35:18','2024-04-14 04:35:18');
/*!40000 ALTER TABLE `user_sessions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `first_name` varchar(50) NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `email` varchar(150) NOT NULL,
  `password` varchar(100) NOT NULL,
  `salt` varchar(80) NOT NULL,
  `avatar` varchar(150) DEFAULT NULL,
  `status` enum('activated','banned') DEFAULT 'activated',
  `role` enum('user','admin') DEFAULT 'user',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_pk_2` (`email`),
  KEY `users_email_index` (`email`),
  KEY `users_first_name_index` (`first_name`),
  KEY `users_id_index` (`id`),
  KEY `users_last_name_index` (`last_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('018ebb6f-ddb3-70e5-871a-332e9d84c6f3','kiet','vu','vuhardyz@gmail.com','$2a$04$L52pWn8oGM/HK3iTVf4pHuQylVnPjahkFDHo3ylaCXnuNfzuFoNwC','1e9c9472db1e928e38dfd5dec5d19522842fe7e8ca5f8ed67cb4da1be292','','activated','user','2024-04-08 02:00:54','2024-04-08 02:00:54'),('018ed68c-963a-7e1c-9305-3f49f85ed05d','john','doe','johndoe@test.com','$2a$04$ZhS9H/1v9Gnteh8DMKSUd.JBgN1qqZPL53ZX9ePeKSv1l6VOFtF0K','60661d39029a58256bf54bdb7faf483eb43ece2b94e83ffa99abd8d25149','1713088626201670600_download.png','activated','user','2024-04-13 08:22:01','2024-04-15 23:59:38');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-04-16  8:22:53
