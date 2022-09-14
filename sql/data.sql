-- init token map data
set global time_zone='+8:00';
set time_zone='+8:00';
flush privileges;

CREATE DATABASE IF NOT EXISTS heco_monitor DEFAULT CHARACTER SET utf8;
USE heco_monitor;

GRANT ALL PRIVILEGES ON *.* TO 'root'@'%'  WITH GRANT OPTION;
flush privileges;

CREATE TABLE IF NOT EXISTS  `bridge_token` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `token_name` varchar(10) COMMENT 'token_name',

  `eth_address` varchar(64) binary NOT NULL COMMENT 'eth_address',
  `eth_amount` DECIMAL(64, 0) COMMENT 'eth_amount',
  `eth_decimals` int COMMENT 'eth_decimals',

  `heco_address` varchar(64) binary NOT NULL COMMENT 'heco_address',
  `heco_amount` DECIMAL(64, 0) COMMENT 'heco_amount',
  `heco_decimals` int COMMENT 'heco_decimals',

  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;