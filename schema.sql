-- Membuat database jika belum ada
CREATE DATABASE IF NOT EXISTS xyz_db;
USE xyz_db;

-- Menghapus tabel jika sudah ada untuk eksekusi ulang yang bersih
DROP TABLE IF EXISTS `transactions`;
DROP TABLE IF EXISTS `customer_limits`;
DROP TABLE IF EXISTS `customers`;

-- Tabel untuk menyimpan data personal konsumen
CREATE TABLE `customers` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `nik` VARCHAR(16) NOT NULL,
  `full_name` VARCHAR(255) NOT NULL,
  `legal_name` VARCHAR(255) NOT NULL,
  `birth_place` VARCHAR(100) NOT NULL,
  `birth_date` DATE NOT NULL,
  `salary` DECIMAL(15, 2) NOT NULL,
  `ktp_photo_url` VARCHAR(255) NOT NULL,
  `selfie_photo_url` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_nik` (`nik`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Tabel untuk menyimpan limit kredit konsumen per tenor
CREATE TABLE `customer_limits` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `customer_id` BIGINT UNSIGNED NOT NULL,
  `tenor_in_months` INT NOT NULL,
  `amount` DECIMAL(15, 2) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_customer_tenor` (`customer_id`, `tenor_in_months`),
  FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Tabel untuk mencatat semua transaksi
CREATE TABLE `transactions` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `contract_number` VARCHAR(50) NOT NULL,
  `customer_id` BIGINT UNSIGNED NOT NULL,
  `otr` DECIMAL(15, 2) NOT NULL,
  `admin_fee` DECIMAL(15, 2) NOT NULL,
  `installment_amount` DECIMAL(15, 2) NOT NULL,
  `interest_amount` DECIMAL(15, 2) NOT NULL,
  `asset_name` VARCHAR(255) NOT NULL,
  `transaction_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_contract_number` (`contract_number`),
  FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Memasukkan data contoh
INSERT INTO `customers` (id, nik, full_name, legal_name, birth_place, birth_date, salary, ktp_photo_url, selfie_photo_url) VALUES
(1, '3273010101900001', 'Budi Santoso', 'Budi Santoso', 'Bandung', '1990-01-01', 8000000.00, '/images/ktp_budi.jpg', '/images/selfie_budi.jpg'),
(2, '3273020202920002', 'Annisa Fitriani', 'Annisa Fitriani', 'Jakarta', '1992-02-02', 15000000.00, '/images/ktp_annisa.jpg', '/images/selfie_annisa.jpg');

INSERT INTO `customer_limits` (customer_id, tenor_in_months, amount) VALUES
(1, 1, 100000.00), (1, 2, 200000.00), (1, 3, 500000.00), (1, 6, 700000.00),
(2, 1, 1000000.00), (2, 2, 1200000.00), (2, 3, 1500000.00), (2, 6, 2000000.00);
*/