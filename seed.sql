INSERT INTO `categories`(`name`) VALUES ('Action');
INSERT INTO `categories`(`name`) VALUES ('RPG');
INSERT INTO `categories`(`name`) VALUES ('Квесты');
INSERT INTO `categories`(`name`) VALUES ('Онлайн-игры');
INSERT INTO `categories`(`name`) VALUES ('Стратегии');


INSERT INTO `attachments`(`path`) VALUES ('/static/img/cover/game-1.jpg');
INSERT INTO `attachments`(`path`) VALUES ('/static/img/cover/game-2.jpg');
INSERT INTO `attachments`(`path`) VALUES ('/static/img/cover/game-3.jpg');
INSERT INTO `attachments`(`path`) VALUES ('/static/img/cover/game-4.jpg');
INSERT INTO `attachments`(`path`) VALUES ('/static/img/cover/game-5.jpg');
INSERT INTO `attachments`(`path`) VALUES ('/static/img/cover/game-6.jpg');
INSERT INTO `attachments`(`path`) VALUES ('/static/img/cover/game-7.jpg');
INSERT INTO `attachments`(`path`) VALUES ('/static/img/cover/game-8.jpg');
INSERT INTO `attachments`(`path`) VALUES ('/static/img/cover/game-9.jpg');
INSERT INTO `attachments`(`path`) VALUES ('/static/img/news/ps4-pro_01.jpg');
INSERT INTO `attachments`(`path`) VALUES ('/static/img/news/ps_vr.jpg');



INSERT INTO `products`(`name`, `description`, `price`, `image`) VALUES ('The Witcher 3: Wild Hunt', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit','200',1);
INSERT INTO `products`(`name`, `description`, `price`, `image`) VALUES ('Overwatch','Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit','250',2);
INSERT INTO `products`(`name`, `description`, `price`, `image`) VALUES ('Deus Ex: Mankind Divided', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit','300',3);
INSERT INTO `products`(`name`, `description`, `price`, `image`) VALUES ('World of WarCraft', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit','400',4);
INSERT INTO `products`(`name`, `description`, `price`, `image`) VALUES ('Call of Duty: Black ops II', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit','600',5);
INSERT INTO `products`(`name`, `description`, `price`, `image`) VALUES ('Batman', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit','500',6);
INSERT INTO `products`(`name`, `description`, `price`, `image`) VALUES ('SuperMario', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit','500',7);

INSERT INTO `news`(`title`,`description`, `image`) VALUES ('О новых играх в режиме VR', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. ullamco laboris nisi ut aliquip ex ea commodo consequat.', 11);
INSERT INTO `news`(`title`,`description`, `image`) VALUES ('О новой PS4 Pro', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. ullamco laboris nisi ut aliquip ex ea commodo consequat.', 10);

INSERT INTO `category_product`( `product_id`, `category_id`) VALUES (1,1);
INSERT INTO `category_product`( `product_id`, `category_id`) VALUES (2,2);
INSERT INTO `category_product`( `product_id`, `category_id`) VALUES (3,3);
INSERT INTO `category_product`( `product_id`, `category_id`) VALUES (4,4);
INSERT INTO `category_product`( `product_id`, `category_id`) VALUES (5,5);
INSERT INTO `category_product`( `product_id`, `category_id`) VALUES (6,5);
INSERT INTO `category_product`( `product_id`, `category_id`) VALUES (7,5);