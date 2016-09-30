CREATE DATABASE  IF NOT EXISTS `gatorloop` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `gatorloop`;
-- MySQL dump 10.13  Distrib 5.7.9, for osx10.9 (x86_64)
--
-- Host: 127.0.0.1    Database: gatorloop
-- ------------------------------------------------------
-- Server version	5.7.10

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Acceleration`
--

DROP TABLE IF EXISTS `Acceleration`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Acceleration` (
  `idAcceleration` int(11) NOT NULL AUTO_INCREMENT,
  `acceleration` decimal(10,2) NOT NULL,
  PRIMARY KEY (`idAcceleration`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Acceleration`
--

LOCK TABLES `Acceleration` WRITE;
/*!40000 ALTER TABLE `Acceleration` DISABLE KEYS */;
/*!40000 ALTER TABLE `Acceleration` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `AuxiliaryBattery`
--

DROP TABLE IF EXISTS `AuxiliaryBattery`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `AuxiliaryBattery` (
  `idAuxiliaryBattery` int(11) NOT NULL AUTO_INCREMENT,
  `voltage` decimal(10,2) NOT NULL,
  `soc` decimal(10,2) NOT NULL,
  `temperature` decimal(10,2) NOT NULL,
  `amp_hour` decimal(10,2) NOT NULL,
  PRIMARY KEY (`idAuxiliaryBattery`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `AuxiliaryBattery`
--

LOCK TABLES `AuxiliaryBattery` WRITE;
/*!40000 ALTER TABLE `AuxiliaryBattery` DISABLE KEYS */;
/*!40000 ALTER TABLE `AuxiliaryBattery` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Position`
--

DROP TABLE IF EXISTS `Position`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Position` (
  `idPosition` int(11) NOT NULL AUTO_INCREMENT,
  `position` decimal(10,2) NOT NULL,
  PRIMARY KEY (`idPosition`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Position`
--

LOCK TABLES `Position` WRITE;
/*!40000 ALTER TABLE `Position` DISABLE KEYS */;
/*!40000 ALTER TABLE `Position` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `PrimaryBattery`
--

DROP TABLE IF EXISTS `PrimaryBattery`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `PrimaryBattery` (
  `idPrimaryBattery` int(11) NOT NULL AUTO_INCREMENT,
  `voltage` decimal(10,2) NOT NULL,
  `soc` decimal(10,2) NOT NULL,
  `temperature` decimal(10,2) NOT NULL,
  `amp_hour` decimal(10,2) NOT NULL,
  PRIMARY KEY (`idPrimaryBattery`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PrimaryBattery`
--

LOCK TABLES `PrimaryBattery` WRITE;
/*!40000 ALTER TABLE `PrimaryBattery` DISABLE KEYS */;
/*!40000 ALTER TABLE `PrimaryBattery` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Rotation`
--

DROP TABLE IF EXISTS `Rotation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Rotation` (
  `idRotation` int(11) NOT NULL AUTO_INCREMENT,
  `roll` decimal(10,2) NOT NULL,
  `pitch` decimal(10,2) NOT NULL,
  `yaw` decimal(10,2) NOT NULL,
  PRIMARY KEY (`idRotation`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Rotation`
--

LOCK TABLES `Rotation` WRITE;
/*!40000 ALTER TABLE `Rotation` DISABLE KEYS */;
/*!40000 ALTER TABLE `Rotation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Temperature`
--

DROP TABLE IF EXISTS `Temperature`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Temperature` (
  `idTemperature` int(11) NOT NULL AUTO_INCREMENT,
  `temperature` decimal(10,2) NOT NULL,
  PRIMARY KEY (`idTemperature`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Temperature`
--

LOCK TABLES `Temperature` WRITE;
/*!40000 ALTER TABLE `Temperature` DISABLE KEYS */;
/*!40000 ALTER TABLE `Temperature` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Velocity`
--

DROP TABLE IF EXISTS `Velocity`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Velocity` (
  `idVelocity` int(11) NOT NULL AUTO_INCREMENT,
  `velocity` decimal(10,5) NOT NULL,
  PRIMARY KEY (`idVelocity`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Velocity`
--

LOCK TABLES `Velocity` WRITE;
/*!40000 ALTER TABLE `Velocity` DISABLE KEYS */;
/*!40000 ALTER TABLE `Velocity` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-09-30 11:49:05
