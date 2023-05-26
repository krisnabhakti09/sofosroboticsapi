-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Waktu pembuatan: 05 Des 2022 pada 08.28
-- Versi server: 8.0.31-0ubuntu0.20.04.1
-- Versi PHP: 7.4.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sofosrobotics`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `robotmasters`
--

CREATE TABLE `robotmasters` (
  `id` int NOT NULL,
  `name` varchar(45) CHARACTER SET latin1 COLLATE latin1_swedish_ci DEFAULT NULL,
  `description` text NOT NULL,
  `image` varchar(255) DEFAULT NULL,
  `typemaster` varchar(255) NOT NULL,
  `price` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `robotmasters`
--

INSERT INTO `robotmasters` (`id`, `name`, `description`, `image`, `typemaster`, `price`) VALUES
(1, 'Robot Angkut', 'Teknologi terbaru Robot Handling untuk Bantu Angkut dan Pindahkan Barang', 'images/robotangkut.jpg', 'robotics', 3000000),
(10, 'Robot Rakit', 'Teknologi terbaru Robot Handling untuk Bantu Perakitan Barang', 'images/robotperakitan.jpg', 'automation', 6000000);

-- --------------------------------------------------------

--
-- Struktur dari tabel `robotorders`
--

CREATE TABLE `robotorders` (
  `id` int NOT NULL,
  `idrobotproductpartdevice` int NOT NULL,
  `userid` int NOT NULL,
  `kodeinvoice` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci DEFAULT NULL,
  `namecustomer` varchar(255) NOT NULL,
  `namarobotmaster` varchar(255) NOT NULL,
  `addresscustomer` text NOT NULL,
  `delivery` varchar(255) NOT NULL,
  `pricedelivery` int NOT NULL,
  `deliverydesc` varchar(255) NOT NULL,
  `pricemethod` varchar(255) NOT NULL,
  `pricemethodadmin` int NOT NULL,
  `pricemethodsn` varchar(255) NOT NULL,
  `subtotal` int NOT NULL,
  `ppn` int NOT NULL,
  `totals` int NOT NULL,
  `pesan` varchar(255) NOT NULL,
  `recommended` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `robotorders`
--

INSERT INTO `robotorders` (`id`, `idrobotproductpartdevice`, `userid`, `kodeinvoice`, `namecustomer`, `namarobotmaster`, `addresscustomer`, `delivery`, `pricedelivery`, `deliverydesc`, `pricemethod`, `pricemethodadmin`, `pricemethodsn`, `subtotal`, `ppn`, `totals`, `pesan`, `recommended`) VALUES
(502, 3, 6, 'INV615730000', 'mobile', 'Robot Angkut', 'Palembang ', 'JNE REG', 17000, 'Akan sampai tanggal 29 - 30 Okt', 'BNI', 3000, '8343543853884', 15700000, 10000, 15730000, 'Segera proses ya', ''),
(503, 1, 6, 'INV615730000', 'mobile', 'Robot Angkut', 'Palembang ', 'JNE REG', 17000, 'Akan sampai tanggal 29 - 30 Okt', 'BNI', 3000, '8343543853884', 15700000, 10000, 15730000, 'Segera proses ya', ''),
(504, 2, 6, 'INV615730000', 'mobile', 'Robot Angkut', 'Palembang ', 'JNE REG', 17000, 'Akan sampai tanggal 29 - 30 Okt', 'BNI', 3000, '8343543853884', 15700000, 10000, 15730000, 'Segera proses ya', ''),
(507, 1, 6, 'INV63030000', 'wandipratama edit', 'Robot Angkut', 'Ggg', 'JNE REG', 17000, 'Akan sampai tanggal 29 - 30 Okt', 'BNI', 3000, '8343543853884', 3000000, 10000, 3030000, '', 'ya'),
(508, 1, 6, 'INV63030000', 'wandipratama edit', 'Robot Angkut', 'Ggg', 'JNE REG', 17000, 'Akan sampai tanggal 29 - 30 Okt', 'BNI', 3000, '8343543853884', 3000000, 10000, 3030000, '', 'ya');

-- --------------------------------------------------------

--
-- Struktur dari tabel `robotpartdevices`
--

CREATE TABLE `robotpartdevices` (
  `id` int NOT NULL,
  `name` varchar(45) CHARACTER SET latin1 COLLATE latin1_swedish_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `robotpartdevices`
--

INSERT INTO `robotpartdevices` (`id`, `name`) VALUES
(1, 'Mekanik'),
(2, 'Software'),
(3, 'Fitur');

-- --------------------------------------------------------

--
-- Struktur dari tabel `robotproductpartdevices`
--

CREATE TABLE `robotproductpartdevices` (
  `id` int NOT NULL,
  `idrobotmaster` int NOT NULL,
  `idrobotpartdevice` int NOT NULL,
  `namerobotmaster` varchar(255) NOT NULL,
  `namerobotpartdevice` varchar(255) NOT NULL,
  `namecomponent` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci DEFAULT NULL,
  `descriptioncomponent` text NOT NULL,
  `image` varchar(255) DEFAULT NULL,
  `price` int DEFAULT NULL,
  `unit` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci DEFAULT NULL,
  `required` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `robotproductpartdevices`
--

INSERT INTO `robotproductpartdevices` (`id`, `idrobotmaster`, `idrobotpartdevice`, `namerobotmaster`, `namerobotpartdevice`, `namecomponent`, `descriptioncomponent`, `image`, `price`, `unit`, `required`) VALUES
(1, 1, 2, 'Robot Angkut', 'Software', 'ESP 32', 'ESP32 adalah mikrokontroller berharga rendah dan hemat energi dengan wifi dan dual-mode bluetooth terintegrasi.', 'images/s1.jpg', 500000, 'pcs', NULL),
(2, 1, 2, 'Robot Angkut', 'Software', 'HC 05', 'Module Bluetooth HC-05 adalah module komunikasi nirkabel via bluetooth yang dimana beroperasi pada frekuensi 2.4GHz dengan pilihan dua mode', 'images/s2.jpg', 250000, 'pcs', NULL),
(3, 1, 1, 'Robot Angkut', 'Mekanik', 'Torsi motor sesuai dengan pilihan', 'Torsi sangat berperan di putaran mesin bawah saat motor anda melaju dari kondisi diam', 'images/f4.jpg', 15000000, 'Kg', 'required'),
(4, 1, 2, 'Robot Angkut', 'Software', 'ESP 32 internet acces', 'ESP32 bisa di fungsikan sebagai web server. Misal web server pada ESP32 mempunyai IP 192.167.123.4, user meminta koneksi ke web server', 'images/s3.jpg', 5000000, 'pcs', NULL),
(5, 1, 3, 'Robot Angkut', 'Fitur', 'HC 05', 'Module Bluetooth HC-05 adalah module komunikasi nirkabel via bluetooth yang dimana beroperasi pada frekuensi 2.4GHz', 'images/f7.jpg', 200000, 'pcs', NULL),
(6, 1, 3, 'Robot Angkut', 'Fitur', 'Infrared', 'radiasi elektromagnetik dari panjang gelombang lebih panjang dari cahaya tampak, tetapi lebih pendek dari radiasi gelombang radio. ', 'images/f10.jpg', 200000, 'pcs', NULL),
(7, 1, 3, 'Robot Angkut', 'Fitur', 'Line Sensor', 'Sensor merupakan indera bagi robot sehingga dapat mengenali berbagai parameter disekitar lingkungan', 'images/f1.jpg', 1000000, 'pcs', NULL),
(8, 1, 3, 'Robot Angkut', 'Fitur', 'SRF', 'SRF adalah bahan bakar yang dihasilkan dari limbah tidak berbahaya sesuai dengan standar Eropa EN 15359.', 'images/f11.png', 1000000, 'pcs', NULL),
(9, 1, 2, 'Robot Angkut', 'Fitur', 'Camera Module', 'produk yang digunakan untuk mengambil foto dan video dari perangkat seluler, seperti telepon pintar, mobil, dan peralatan rumah tangga pintar.', 'images/f1.jpg', 3500000, 'pcs', NULL),
(10, 1, 3, 'Robot Angkut', 'Fitur', 'GPS Module', 'Modul GPS berfungsi sebagai penerima GPS (Global Positioning System Receiver) yang dapat mendeteksi lokasi dengan menangkap dan memroses sinyal dari satelit', 'images/f3.jpg', 5000000, 'pcs', NULL),
(11, 1, 1, 'Robot Angkut', 'Mekanik', 'Lidar', 'IDAR (Light Detection and Ranging) adalah sebuah teknologi peraba jarak jauh optik yang mengukur properti cahaya yang tersebar untuk menemukan jarak ', 'images/f5.jpg', 800000, 'pcs', NULL),
(12, 1, 2, 'Robot Angkut', 'Software', 'RFID reader', 'Pengertian RFID reader adalah alat untuk membaca RFID tag yang sering disebut juga interrogator dan merupakan perangkat yang penting pada sebuah sistem RFID', 'images/f9.jpg', 800000, 'pcs', NULL),
(13, 1, 2, 'Robot Angkut', 'Software', 'Kompas Module', 'Kompas adalah navigator pola probabilistik tiga koordinat. Setiap koordinat memiliki kecepatan perjalanan yang unik dan keluaran pemicu khusus.', 'images/f6.jpg', 500000, 'pcs', NULL),
(14, 1, 2, 'Robot Angkut', 'Software', 'Navigation Module', 'Modul navigasi memungkinkan Anda menentukan struktur navigasi untuk situs web Anda, dan menampilkan satu atau lebih menu di halaman Anda, berdasarkan struktur ini.', 'images/f2.jpg', 800000, 'pcs', NULL),
(15, 1, 3, 'Robot Angkut', 'Fitur', 'Rotary OMRON', 'Rotary Encoder mengukur jumlah rotasi, sudut rotasi, dan posisi rotasi. Tambahan. Encoder Inkremental mengeluarkan string pulsa', 'images/f7.jpg', 800000, 'pcs', NULL),
(16, 1, 3, 'Robot Angkut', 'Fitur', 'Voice Module', 'Modul Pengenalan Suara adalah papan pengenalan suara yang ringkas dan mudah dikendalikan. ', 'images/f5.jpg', 200000, 'pcs', NULL),
(17, 1, 1, 'Robot Angkut', 'Mekanik', 'Speaker', 'Speaker adalah perangkat keras output yang berfungsi mengeluarkan hasil pemrosesan oleh CPU berupa audio/suara. ', 'images/f8.jpg', 200000, 'pcs', NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `avatar` varchar(255) NOT NULL,
  `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `number`, `email`, `password`, `avatar`, `role`) VALUES
(1, 'wandi', '0831456767', 'wandi2@gmail.com', '$2b$10$XxvDthkMzQAwpaQ5huSrG.dMAkY/e0pqyBNFdp5Zl2GaqSdtnCzo.', 'images/default/defaultavatar.png', 'user'),
(2, 'pratama', '082622626', 'pratama@gmail.com', '$2b$10$BfIHcu9Cn3J1quRzdKmOmenzc2aoscRwhHN2yZdlMx/pgvfOobqsu', 'images/default/defaultavatar.png', 'user'),
(3, 'admin', '082622626', 'admin@gmail.com', '$2b$10$BfIHcu9Cn3J1quRzdKmOmenzc2aoscRwhHN2yZdlMx/pgvfOobqsu', 'images/default/defaultavatar.png', 'admin'),
(4, 'diwa', '0934349934', 'diwa@gmail.com', '$2a$04$b27dDuOuvo.DfnNH9kbyyum2MP74C29UZ5KUEjdgRUo6Yl62lOV3G', 'images/default/defaultavatar.png', 'user'),
(6, 'wandipratama edit', '0934349934', 'mobile@gmail.com', '$2a$04$ry8yQvoE9Bh0qX5nqmYZFuaYyRwY0bCBBrRltnpipoo1mBk1KYbCy', 'images/6-db6d91a8-7dcd-4373-9eea-87831c173d94.jpeg', 'user');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `robotmasters`
--
ALTER TABLE `robotmasters`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `robotorders`
--
ALTER TABLE `robotorders`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `robotpartdevices`
--
ALTER TABLE `robotpartdevices`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `robotproductpartdevices`
--
ALTER TABLE `robotproductpartdevices`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `robotmasters`
--
ALTER TABLE `robotmasters`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT untuk tabel `robotorders`
--
ALTER TABLE `robotorders`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=509;

--
-- AUTO_INCREMENT untuk tabel `robotpartdevices`
--
ALTER TABLE `robotpartdevices`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `robotproductpartdevices`
--
ALTER TABLE `robotproductpartdevices`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
