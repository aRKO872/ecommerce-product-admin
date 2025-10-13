INSERT INTO `users` (`id`, `username`, `phone`, `email`, `address`)
VALUES
(UUID(), 'frank', '5432109876', 'frank@example.com', '12 Elm St, CityF'),
(UUID(), 'gina', NULL, 'gina@example.com', '34 Cedar Rd, CityG'),
(UUID(), 'harry', '4321098765', NULL, '56 Spruce Ln, CityH'),
(UUID(), 'irene', '3210987654', 'irene@example.com', '78 Willow Ave, CityI'),
(UUID(), 'jack', '2109876543', 'jack@example.com', '90 Poplar Blvd, CityJ');

INSERT INTO `sellers` (`id`, `name`, `address`, `phone`, `email`) VALUES
(UUID(), 'TechWorld', '1 Tech Park, CityA', '1112223333', 'contact@techworld.com'),
(UUID(), 'ToyHouse', '22 Toy Lane, CityB', NULL, 'sales@toyhouse.com'),
(UUID(), 'FashionHub', '33 Fashion St, CityC', '4445556666', NULL),
(UUID(), 'DailyNeeds', '44 Essential Ave, CityD', '7778889999', 'support@dailyneeds.com');

INSERT INTO `products` (`id`, `name`, `category`, `img_url`)
VALUES
(UUID(), 'Laptop Pro', 'ELECTRONICS', 'https://example.com/img/laptop.png'),
(UUID(), 'Tablet Z', 'ELECTRONICS', 'https://example.com/img/tablet.png'),
(UUID(), 'Headphones Max', 'ELECTRONICS', 'https://example.com/img/headphones.png'),
(UUID(), 'Robot Toy', 'TOYS', 'https://example.com/img/robot.png'),
(UUID(), 'Puzzle Game', 'TOYS', 'https://example.com/img/puzzle.png'),
(UUID(), 'T-Shirt V', 'CLOTHES', 'https://example.com/img/tshirt.png'),
(UUID(), 'Jacket Winter', 'CLOTHES', 'https://example.com/img/jacket.png'),
(UUID(), 'Dress Summer', 'CLOTHES', 'https://example.com/img/dress.png'),
(UUID(), 'Shampoo 1L', 'DAILYCARE', 'https://example.com/img/shampoo1l.png'),
(UUID(), 'Soap Pack', 'DAILYCARE', 'https://example.com/img/soap.png'),
(UUID(), 'Toothpaste Gel', 'DAILYCARE', 'https://example.com/img/toothpaste.png'),
(UUID(), 'Notebook XL', 'ESSENTIALS', 'https://example.com/img/notebookxl.png'),
(UUID(), 'Pen Set', 'ESSENTIALS', 'https://example.com/img/penset.png'),
(UUID(), 'Marker Pack', 'ESSENTIALS', 'https://example.com/img/markers.png'),
(UUID(), 'Sneakers Run', 'CLOTHES', 'https://example.com/img/sneakers.png'),
(UUID(), 'Action Figure', 'TOYS', 'https://example.com/img/actionfigure.png'),
(UUID(), 'Smart Watch', 'ELECTRONICS', 'https://example.com/img/smartwatch.png'),
(UUID(), 'Face Cream', 'DAILYCARE', 'https://example.com/img/facecream.png'),
(UUID(), 'Eraser Pack', 'ESSENTIALS', 'https://example.com/img/eraser.png'),
(UUID(), 'Hat Summer', 'CLOTHES', 'https://example.com/img/hat.png');

INSERT INTO `inventory` (`id`, `product_id`, `seller_id`, `sold`)
SELECT UUID(), p.id, s.id, 0
FROM `products` p
JOIN `sellers` s
WHERE
    (p.category = 'ELECTRONICS' AND s.name = 'TechWorld') OR
    (p.category = 'TOYS' AND s.name = 'ToyHouse') OR
    (p.category = 'CLOTHES' AND s.name = 'FashionHub') OR
    (p.category IN ('DAILYCARE','ESSENTIALS') AND s.name = 'DailyNeeds');

-- Assign 15 orders per user
INSERT INTO orders (product_id, placed_by, dateOfOrder, tracking_phase)
SELECT i.id, u.id, NOW() - INTERVAL FLOOR(RAND()*30) DAY,
       CASE FLOOR(RAND()*6)
            WHEN 0 THEN 'ORDER_PLACED'
            WHEN 1 THEN 'SHIPPED'
            WHEN 2 THEN 'OUT_FOR_DELIVERY'
            WHEN 3 THEN 'DELIVERED'
            WHEN 4 THEN 'ORDER_NOT_PLACED'
            WHEN 5 THEN 'ORDER_CANCELLED'
       END
FROM inventory i
JOIN users u ON u.username = 'frank'
WHERE i.sold = FALSE
LIMIT 15;

INSERT INTO orders (product_id, placed_by, dateOfOrder, tracking_phase)
SELECT i.id, u.id, NOW() - INTERVAL FLOOR(RAND()*30) DAY,
       CASE FLOOR(RAND()*6)
            WHEN 0 THEN 'ORDER_PLACED'
            WHEN 1 THEN 'SHIPPED'
            WHEN 2 THEN 'OUT_FOR_DELIVERY'
            WHEN 3 THEN 'DELIVERED'
            WHEN 4 THEN 'ORDER_NOT_PLACED'
            WHEN 5 THEN 'ORDER_CANCELLED'
       END
FROM inventory i
JOIN users u ON u.username = 'jack'
WHERE i.sold = FALSE
LIMIT 15;

INSERT INTO orders (product_id, placed_by, dateOfOrder, tracking_phase)
SELECT i.id, u.id, NOW() - INTERVAL FLOOR(RAND()*30) DAY,
       CASE FLOOR(RAND()*6)
            WHEN 0 THEN 'ORDER_PLACED'
            WHEN 1 THEN 'SHIPPED'
            WHEN 2 THEN 'OUT_FOR_DELIVERY'
            WHEN 3 THEN 'DELIVERED'
            WHEN 4 THEN 'ORDER_NOT_PLACED'
            WHEN 5 THEN 'ORDER_CANCELLED'
       END
FROM inventory i
JOIN users u ON u.username = 'irene'
WHERE i.sold = FALSE
LIMIT 15;

INSERT INTO orders (product_id, placed_by, dateOfOrder, tracking_phase)
SELECT i.id, u.id, NOW() - INTERVAL FLOOR(RAND()*30) DAY,
       CASE FLOOR(RAND()*6)
            WHEN 0 THEN 'ORDER_PLACED'
            WHEN 1 THEN 'SHIPPED'
            WHEN 2 THEN 'OUT_FOR_DELIVERY'
            WHEN 3 THEN 'DELIVERED'
            WHEN 4 THEN 'ORDER_NOT_PLACED'
            WHEN 5 THEN 'ORDER_CANCELLED'
       END
FROM inventory i
JOIN users u ON u.username = 'harry'
WHERE i.sold = FALSE
LIMIT 15;

INSERT INTO orders (product_id, placed_by, dateOfOrder, tracking_phase)
SELECT i.id, u.id, NOW() - INTERVAL FLOOR(RAND()*30) DAY,
       CASE FLOOR(RAND()*6)
            WHEN 0 THEN 'ORDER_PLACED'
            WHEN 1 THEN 'SHIPPED'
            WHEN 2 THEN 'OUT_FOR_DELIVERY'
            WHEN 3 THEN 'DELIVERED'
            WHEN 4 THEN 'ORDER_NOT_PLACED'
            WHEN 5 THEN 'ORDER_CANCELLED'
       END
FROM inventory i
JOIN users u ON u.username = 'gina'
WHERE i.sold = FALSE
LIMIT 15;

UPDATE `inventory` i
JOIN `orders` o ON i.id = o.product_id
SET i.sold = TRUE;

