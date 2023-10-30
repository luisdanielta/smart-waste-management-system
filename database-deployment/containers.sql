/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE TABLE `containers` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `controller_id` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `location` varchar(254) NOT NULL,
  `usecase` varchar(100) NOT NULL,
  `size` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_containers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

INSERT INTO `containers` (`id`, `created_at`, `updated_at`, `deleted_at`, `controller_id`, `name`, `location`, `usecase`, `size`) VALUES
(7, '2023-10-28 01:34:15.939', '2023-10-28 01:34:15.939', NULL, 'TEGY526D', '001C', '-75.46585943854299, 10.370504818304838', 'Others', 'L');
INSERT INTO `containers` (`id`, `created_at`, `updated_at`, `deleted_at`, `controller_id`, `name`, `location`, `usecase`, `size`) VALUES
(9, '2023-10-28 01:35:45.451', '2023-10-28 01:35:45.451', NULL, 'TEGY526D', '002C', '-75.46564692223424, 10.37027117007581', 'Others', 'L');
INSERT INTO `containers` (`id`, `created_at`, `updated_at`, `deleted_at`, `controller_id`, `name`, `location`, `usecase`, `size`) VALUES
(10, '2023-10-28 01:36:15.533', '2023-10-28 01:36:15.533', NULL, 'TEGY526D', '003C', '-75.4656889658942, 10.37000905678633', 'Others', 'L');
INSERT INTO `containers` (`id`, `created_at`, `updated_at`, `deleted_at`, `controller_id`, `name`, `location`, `usecase`, `size`) VALUES
(11, '2023-10-28 01:36:19.723', '2023-10-28 01:36:19.723', NULL, 'TEGY526D', '004C', '-75.46517385129899, 10.369995225099842', 'Others', 'L'),
(12, '2023-10-28 01:36:35.622', '2023-10-28 01:36:35.622', NULL, 'TEGY526D', '005C', '-75.46607059406848, 10.369727724692908', 'Others', 'L'),
(13, '2023-10-28 01:36:50.051', '2023-10-28 01:36:50.051', NULL, 'TEGY526D', '006C', '-75.46531671530121, 10.369609843115981', 'Others', 'L');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;