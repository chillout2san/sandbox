CREATE TABLE `users` (
    `id` varchar(26) NOT NULL COMMENT 'ユーザーのid',
    `name` varchar(255) NOT NULL COMMENT 'ユーザーの名前',
    `mail` varchar(255) NOT NULL COMMENT 'ユーザーのメールアドレス',
    `imagePath` varchar(255) NOT NULL COMMENT 'ユーザーのプロフィール写真のパス',
    `pass` varchar(255) NOT NULL COMMENT 'ユーザーのパスワード',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;