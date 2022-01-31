CREATE DATABASE `inventorydb`;

CREATE TABLE `inventorydb`.`products` (
    `productId` INT NOT NULL AUTO_INCREMENT,
    `manufacturer` VARCHAR(255) NOT NULL, 
    `sku` VARCHAR(60) NOT NULL,  
    `upc` VARCHAR(60) NOT NULL,  
    `pricePerUnit` DECIMAL(13, 2) NOT NULL, 
    `quantityOnHand` INT NOT NULL,   
    `productName` VARCHAR(255) NOT NULL,  
    PRIMARY KEY (`productId`));

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`)
VALUES ("VietDuc", "221212hj1h2j12", "1287873", "400.26", 5005, "alphabet board");

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`)
VALUES ("Uni", "221212hj1sasah2j12", "2121212", "325.12", 2000, "guitar");

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`)
VALUES ("Nissan", "1212saa", "12232323", "436.26", 5905, "lamp shade");