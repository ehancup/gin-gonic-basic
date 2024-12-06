CREATE TABLE `users` (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL,
    address varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    born_date timestamp,
    updated_date timestamp,
    UNIQUE KEY(email)
);

-- $ migrate -path database/migrations -database "mysql://root:hancup20@tcp(localhost:3306)/go_gin_gonic" up
-- $ migrate -path database/migrations -database "mysql://root:hancup20@tcp(localhost:3306)/go_gin_gonic" down
-- $ migrate create -ext sql -dir database/migrations -seq remove_unique
