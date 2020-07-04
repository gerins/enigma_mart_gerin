-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema enigma_mart
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema enigma_mart
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `enigma_mart` DEFAULT CHARACTER SET utf8 ;
USE `enigma_mart` ;

-- -----------------------------------------------------
-- Table `enigma_mart`.`produk_category`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `enigma_mart`.`produk_category` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nama` VARCHAR(45) NULL DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `Status` VARCHAR(5) NOT NULL DEFAULT 'A',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 13
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `enigma_mart`.`produk`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `enigma_mart`.`produk` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `produk_category_id` INT NOT NULL,
  `nama` VARCHAR(45) NOT NULL,
  `harga` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `Status` VARCHAR(5) NOT NULL DEFAULT 'A',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  INDEX `fk_produk_produk_category_idx` (`produk_category_id` ASC) VISIBLE,
  CONSTRAINT `fk_produk_produk_category`
    FOREIGN KEY (`produk_category_id`)
    REFERENCES `enigma_mart`.`produk_category` (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 40
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `enigma_mart`.`transaksi_penjualan`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `enigma_mart`.`transaksi_penjualan` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `total_penjualan` INT NOT NULL,
  `Status` VARCHAR(5) NOT NULL DEFAULT 'A',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 51
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `enigma_mart`.`transaksi_produk`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `enigma_mart`.`transaksi_produk` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `transaksi_penjualan_id` INT NOT NULL,
  `produk_id` INT UNSIGNED NOT NULL,
  `kuantiti` INT NOT NULL,
  `total` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `Status` VARCHAR(5) NOT NULL DEFAULT 'A',
  PRIMARY KEY (`id`),
  INDEX `fk_transaksi_produk_transaksi_penjualan1_idx` (`transaksi_penjualan_id` ASC) VISIBLE,
  INDEX `fk_transaksi_produk_produk1_idx` (`produk_id` ASC) VISIBLE,
  CONSTRAINT `fk_transaksi_produk_produk1`
    FOREIGN KEY (`produk_id`)
    REFERENCES `enigma_mart`.`produk` (`id`),
  CONSTRAINT `fk_transaksi_produk_transaksi_penjualan1`
    FOREIGN KEY (`transaksi_penjualan_id`)
    REFERENCES `enigma_mart`.`transaksi_penjualan` (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 105
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `enigma_mart`.`user_data`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `enigma_mart`.`user_data` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nama` VARCHAR(45) NOT NULL,
  `password` VARCHAR(256) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `nama_UNIQUE` (`nama` ASC) VISIBLE)
ENGINE = InnoDB
AUTO_INCREMENT = 19
DEFAULT CHARACTER SET = utf8;

USE `enigma_mart` ;

-- -----------------------------------------------------
-- Placeholder table for view `enigma_mart`.`produk_idx`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `enigma_mart`.`produk_idx` (`id` INT, `p_name` INT, `cat_id` INT, `cat_name` INT, `harga` INT, `created_at` INT, `updated_at` INT, `status` INT);

-- -----------------------------------------------------
-- Placeholder table for view `enigma_mart`.`transaksi_produk_idx`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `enigma_mart`.`transaksi_produk_idx` (`trans_id` INT, `nama` INT, `kategori` INT, `harga` INT, `kuantiti` INT, `total` INT);

-- -----------------------------------------------------
-- View `enigma_mart`.`produk_idx`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `enigma_mart`.`produk_idx`;
USE `enigma_mart`;
CREATE  OR REPLACE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `enigma_mart`.`produk_idx` AS select `p`.`id` AS `id`,`p`.`nama` AS `p_name`,`pc`.`id` AS `cat_id`,`pc`.`nama` AS `cat_name`,`p`.`harga` AS `harga`,`p`.`created_at` AS `created_at`,`p`.`updated_at` AS `updated_at`,`p`.`Status` AS `status` from (`enigma_mart`.`produk` `p` join `enigma_mart`.`produk_category` `pc` on((`p`.`produk_category_id` = `pc`.`id`)));

-- -----------------------------------------------------
-- View `enigma_mart`.`transaksi_produk_idx`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `enigma_mart`.`transaksi_produk_idx`;
USE `enigma_mart`;
CREATE  OR REPLACE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `enigma_mart`.`transaksi_produk_idx` AS select `tp`.`transaksi_penjualan_id` AS `trans_id`,`p`.`nama` AS `nama`,`pc`.`nama` AS `kategori`,`p`.`harga` AS `harga`,`tp`.`kuantiti` AS `kuantiti`,`tp`.`total` AS `total` from ((`enigma_mart`.`transaksi_produk` `tp` join `enigma_mart`.`produk` `p` on((`p`.`id` = `tp`.`produk_id`))) join `enigma_mart`.`produk_category` `pc` on((`pc`.`id` = `p`.`produk_category_id`)));

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
