# Fridge replenishment server

  
Server module of the system for forecasting replenishment of goods stored in refrigeration equipment at distributed points of retail networks. 

The main functional unit of the module is the web server that implements the REST API for interfacing with the front part of the system, providing readings of data from the message queue and storage in the database

The system is designed to be deployed in the IBM Сloud(Bluemix) using services:

- Storing data in the PostgreSQL database
- MQTT Message Queuing - IBM IoT HuB
