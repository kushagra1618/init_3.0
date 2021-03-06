CREATE DATABASE init;
CREATE TABLE `apps` ( 
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `NAME` varchar(250) NOT NULL,
  `EMAIL` varchar(250) NOT NULL,
  `DOMAIN` varchar(250) NOT NULL,
  `CONTACT` varchar(15) NOT NULL,
  `RSAPUBKEY` varchar(3000) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `EMAIL` (`EMAIL`)
); 
/**
Table to manage username and public address pairing
*/
CREATE TABLE `users` (
  `USERNAME` varchar(100) DEFAULT NULL,
  `PUBADDR` varchar(100) DEFAULT NULL,
  UNIQUE KEY `USERNAME` (`USERNAME`)
);