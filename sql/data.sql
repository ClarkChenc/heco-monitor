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
  `assert_type` int COMMENT 'assert_type',

  `heco_address` varchar(64) binary NOT NULL COMMENT 'heco_address',
  `heco_amount` DECIMAL(64, 0) COMMENT 'heco_amount',
  `heco_decimals` int COMMENT 'heco_decimals',

  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (`id`),
  UNIQUE KEY `heco_address` (`heco_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE IF NOT EXISTS  `root_token` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `token_name` varchar(10) COMMENT 'token_name',

  `root_amount` DECIMAL(64, 0) COMMENT 'root_amount',
  `root_decimals` int COMMENT 'root_decimals',

  `eth_address` varchar(64) binary NOT NULL COMMENT 'eth_address',
  `eth_amount` DECIMAL(64, 0) COMMENT 'eth_amount',
  `eth_decimals` int COMMENT 'eth_decimals',

  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (`id`),
  UNIQUE KEY `eth_address` (`eth_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE IF NOT EXISTS  `root_account_token` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `token_name` varchar(10) COMMENT 'token_name',

  `account` varchar(64) binary NOT NULL COMMENT 'account',
  `account_balance` double COMMENT 'account_balance',

  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (`id`),
  UNIQUE KEY `token_name__account` (`token_name`, `account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;