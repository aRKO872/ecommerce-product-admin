CREATE TABLE `users` (
  `id` CHAR(36) NOT NULL DEFAULT (UUID()),
  `username` VARCHAR(255) NOT NULL,
  `phone` VARCHAR(50),
  `email` VARCHAR(255),
  `address` VARCHAR(1024) NOT NULL,
  `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_users_username` (`username`),
  KEY `idx_users_phone` (`phone`),
  KEY `idx_users_email` (`email`),
  KEY `idx_users_created_at` (`created_at`),
  CONSTRAINT `chk_users_phone_or_email` CHECK ((`phone` IS NOT NULL) OR (`email` IS NOT NULL))
);

CREATE TABLE `sellers` (
  `id` CHAR(36) NOT NULL DEFAULT (UUID()),
  `name` VARCHAR(255) NOT NULL,
  `address` VARCHAR(1024) NOT NULL,
  `phone` VARCHAR(50),
  `email` VARCHAR(255),
  PRIMARY KEY (`id`),
  KEY `idx_sellers_phone` (`phone`),
  KEY `idx_sellers_email` (`email`),
  CONSTRAINT `chk_sellers_phone_or_email` CHECK ((`phone` IS NOT NULL) OR (`email` IS NOT NULL))
);

CREATE TABLE `products` (
  `id` CHAR(36) NOT NULL DEFAULT (UUID()),
  `name` VARCHAR(512) NOT NULL,
  `category` ENUM('CLOTHES','TOYS','ELECTRONICS','ESSENTIALS','DAILYCARE') NOT NULL DEFAULT 'CLOTHES',
  `img_url` VARCHAR(2048),
  `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`),
  KEY `idx_products_category` (`category`),
  KEY `idx_products_created_at` (`created_at`)
);

CREATE TABLE `inventory` (
  `id` CHAR(36) NOT NULL DEFAULT (UUID()),
  `product_id` CHAR(36) NOT NULL,
  `seller_id` CHAR(36) NOT NULL,
  `sold` BOOLEAN NOT NULL DEFAULT FALSE,
  PRIMARY KEY (`id`),
  KEY `idx_inventory_product_id` (`product_id`),
  KEY `idx_inventory_seller_id` (`seller_id`),
  KEY `idx_inventory_sold` (`sold`),
  CONSTRAINT `fk_inventory_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT `fk_inventory_seller` FOREIGN KEY (`seller_id`) REFERENCES `sellers` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE `orders` (
  `id` CHAR(36) NOT NULL DEFAULT (UUID()),
  `product_id` CHAR(36) NOT NULL,
  `placed_by` CHAR(36) NOT NULL,
  `dateOfOrder` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `tracking_phase` ENUM('ORDER_PLACED','SHIPPED','OUT_FOR_DELIVERY','DELIVERED','ORDER_NOT_PLACED','ORDER_CANCELLED') NOT NULL DEFAULT 'ORDER_PLACED',
  PRIMARY KEY (`id`),
  KEY `idx_orders_product_id` (`product_id`),
  KEY `idx_orders_placed_by` (`placed_by`),
  KEY `idx_orders_tracking_phase` (`tracking_phase`),
  KEY `idx_orders_dateOfOrder` (`dateOfOrder`),
  CONSTRAINT `fk_orders_inventory` FOREIGN KEY (`product_id`) REFERENCES `inventory` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT `fk_orders_users` FOREIGN KEY (`placed_by`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
);
