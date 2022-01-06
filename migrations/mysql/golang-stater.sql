-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 02, 2022 at 12:37 AM
-- Server version: 5.7.33
-- PHP Version: 7.2.34

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang-stater`
--

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `product_id` int(11) NOT NULL,
  `product_category_fkid` int(11) DEFAULT NULL,
  `admin_fkid` int(11) DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `price` int(11) NOT NULL,
  `description` text NOT NULL,
  `qty` int(11) NOT NULL DEFAULT '0',
  `image` text NOT NULL,
  `label` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`product_id`, `product_category_fkid`, `admin_fkid`, `name`, `price`, `description`, `qty`, `image`, `label`) VALUES
(1, 0, 0, 'Bayam Segar', 7999, 'Bayam segar /ikat +-300gr  Note : Tersedia bumbu dan sayuran packing siap olah lainya di etalase kami.  ORDER MAKSIMAL MASUK JAM 09.00 AKAN DIKIRIM DI HARI YANG SAMA. selebih jam 09.00 dikirim esok hari. dikarenakan untuk menjaga kualitas kesegaran sayur dan bumbu.  DISARANKAN MENGGUNAKAN KURIR INSTAN UNTUK MENJAGA KESEGARAN PRODUK (KARENA SEMUA PRODUK DIKEMAS PAGI HARI)   MOHON DIPAHAMI KURIR SAMEDAY AKAN MEMERLUKAN WAKTU YG LAMA DALAM PICKUP DAN PENGIRIMAN S/D,6JAM DAPAT MERUSAK KESEGARAN PRODUK SAYUR /IKAN/AYAM/DAGING', 100, 'test.com/image.png', 'sayuran daun'),
(2, 0, 0, 'Daun Bawang Segarikat', 4999, 'Daun bawang segar fresh.. harga untuk 1ikat.kurleb 70gr  DISARANKAN MENGGUNAKAN KURIR INSTAN UNTUK MENJAGA KESEGARAN PRODUK (KARENA SEMUA PRODUK DIKEMAS PAGI HARI)   MOHON DIPAHAMI KURIR SAMEDAY AKAN MEMERLUKAN WAKTU YG LAMA DALAM PICKUP DAN PENGIRIMAN S/D,6JAM DAPAT MERUSAK KESEGARAN PRODUK SAYUR /IKAN/AYAM/DAGING', 100, 'test.com/image.png', 'garnish'),
(3, 0, 0, 'Tomat Pack 100gr Sayuran Segar', 1999, '(Fast respon) Tanya hanya via diskusi produk  Untuk vitur chat eror overload  KHUSUS GOJEK   Sayuran segar yang kami sajikan untuk setiap packnya, di siapkan dengan pilihan bahan2 segar. untuk memudahkan anda yang ingin mendapatkan sayuran segar sehat dan higienis.  PAKET TOMAT SEGAR ISI :  - TOMAT SEGAR (100gr)  Note : Tersedia bumbu dan sayuran packing siap olah lainya di etalase kami.  ORDER MAKSIMAL MASUK JAM 09.00 AKAN DIKIRIM DI HARI YANG SAMA. selebih jam 09.00 dikirim esok hari. dikarenakan untuk menjaga kualitas kesegaran sayur dan bumbu.  HIDUP SEHAT BERSAMA KELUARGA.', 100, 'test.com/image.png', 'sayuran buah'),
(4, 0, 0, 'Sayurhd Bayam Segar 1 Ikat', 4700, '', 100, 'test.com/image.png', 'sayuran daun'),
(5, 0, 0, 'Sayurhd Sayur Segar Tempe 1papan', 7000, 'SayurHD  Sayur segar tempe dijual per papan  Pembayaran hari ini = masuk pengiriman h+1 stlh pembayaran ', 100, 'test.com/image.png', 'lauk nabati'),
(6, 0, 0, 'Sayurhd Sayur Segar Toge 100gr', 2000, 'SayurHD  Sayur segar toge dijual per 100gr  Pembayaran hari ini = pengiriman h+1 stlh pembayaran', 100, 'test.com/image.png', 'sayuran akar'),
(7, 0, 0, 'Sayur Daun Bayam Segar Per Ikat', 3500, 'Sayur Daun Bayam Segar Per Ikat  Berat 150 gram', 100, 'test.com/image.png', 'sayuran daun'),
(8, 0, 0, 'Tempe Per 1 Papan', 5000, 'Tempe  per 1 Papan', 100, 'test.com/image.png', 'lauk nabati'),
(9, 0, 0, 'Tahu Putih Kecil Per Pack Si 10', 6000, 'Tahu Putih Kecil Per Pack (isi 10)', 100, 'test.com/image.png', 'lauk nabati'),
(10, 0, 0, 'Terlaris Bayam Segar', 6900, 'TERLARIS BAYAM SEGAR 1 IKAT (90 - 110 GRAM) TERGANTUNG HASIL PANEN PRODUK KAMI SELALU FRESH  PENGIRIMAN SETIAP H+1 SETELAH TRANFER  PEMESANAN DALAM JUMLAH BANYAK LANGSUNG SAJA HUBUNGI NOMOR WHATSAPP KAMI  081210348245 081382628395', 100, 'test.com/image.png', 'sayuran daun');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `user_id` int(11) NOT NULL,
  `photo` text NOT NULL,
  `username` varchar(60) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `name` varchar(100) NOT NULL,
  `created_at` bigint(20) NOT NULL,
  `updated_at` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`user_id`, `photo`, `username`, `email`, `password`, `name`, `created_at`, `updated_at`) VALUES
(1, 'https://lh3.googleusercontent.com/-7NbJWT5zZVc/WomnR_q5wgI/AAAAAAAAFRY/yuW6Wd-7B0ocg9CSKlxqSuozeLkrZnw-gCEwYBhgLKs4DAMBZVoCN1JDyv2YpvnbbzVfPsZFncgL64Slg3_BJmvNihMXlHTUsmtq8bCZQrwNF3XUr4FhBoptnFnnx3TPECerUhZ1uzT8glAr46UWPynXpUNkaS4VWL6glsqdPoAec2NKdfcAgNbYW7O9UouasNXaFMS3EFOGnaWbVauMn_YpIR2v0pyyBhtAediPS-3zVzDz7txilDCh5_Fd0TlmtP5HDcWFunIUKCurQoY1tYggGE-3DC4oeu7JZwOYwyfjR2Z7wqyQ0diVcX9R-ayUhZf4zmiXHAaXPWb2_yj5Gs3P6ZD14H43nmtNHgmeoDkXPy01YPY7oUl2QLip4EN12vZbE-z4fOlNM69r4ODaW6xu5ko1BjdlHRL1Q2GSBPp1n9EN6jSdg_6K75rwN8Xe28vb4gvYTjoMWOg-wFBwS7KfGpL53114_Yhm1-BaKxiaO8PROpUE1au5UheS8dkZ0A6PIDzWYtD2BAcHFDNIaHq2OB_GHWoJXRU_Ie8Vpvg814KwaCBjBRcYRmNIvwuvM6LERMM_emyjx4xpWydMHB1uGZy77AtMLLaRW9AJVykq4-oWPmB46fqtspQucG17c-EGRZxSivlLA2evieOyOMNXI3PQF/w280-h280-p/12633720_1179451595415928_3654751649444577840_o.jpg', 'test', 'test@gmail.com', '$2y$12$jniETXMblositfIBwGGMv.GRicA4jUDoaaupu/vrhDRJ5Siw5fzWG', 'Test', 1586962829441, NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`product_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `product_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=869;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;


--
-- Table structure for table `products_images`
--

CREATE TABLE `products_images` (
  `productimages_id` int(11) NOT NULL,
  `product_fkid` int(11) DEFAULT NULL,
  `images` varchar(255) NOT NULL,
  `created_at` bigint(20) NOT NULL,
  `updated_at` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for table `products_images`
--
ALTER TABLE `products_images`
  ADD PRIMARY KEY (`productimages_id`),
  ADD KEY `product_id` (`product_fkid`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `products_images`
--
ALTER TABLE `products_images`
  MODIFY `productimages_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `products_images`
--
ALTER TABLE `products_images`
  ADD CONSTRAINT `products_images_ibfk_1` FOREIGN KEY (`product_fkid`) REFERENCES `products` (`product_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
