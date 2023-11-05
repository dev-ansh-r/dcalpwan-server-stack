# The dcalpwan server for LoRaWAN Networks.

## The server is built using the The Things Stack as the base of the development. The server has been configured completely for the dcalpwan network requirements.

LoRaWAN is a protocol for low-power wide area networks. It allows for large scale Internet of Things deployments where low-powered devices efficiently communicate with Internet-connected applications over long range wireless connections.

## Features

- LoRaWAN Network Server
  - [x] Supports LoRaWAN 1.0
  - [x] Supports LoRaWAN 1.0.1
  - [x] Supports LoRaWAN 1.0.2
  - [x] Supports LoRaWAN 1.0.3
  - [x] Supports LoRaWAN 1.0.4
  - [x] Supports LoRaWAN 1.1
  - [x] Supports LoRaWAN Regional Parameters 1.0
  - [x] Supports LoRaWAN Regional Parameters 1.0.2 rev B
  - [x] Supports LoRaWAN Regional Parameters 1.0.3 rev A
  - [x] Supports LoRaWAN Regional Parameters 1.1 rev A
  - [x] Supports LoRaWAN Regional Parameters 1.1 rev B
  - [x] Supports Class A devices
  - [x] Supports Class B devices
  - [x] Supports Class C devices
  - [x] Supports OTAA devices
  - [x] Supports ABP devices
  - [x] Supports MAC Commands
  - [x] Supports Adaptive Data Rate
  - [ ] Implements LoRaWAN Back-end Interfaces 1.0
- LoRaWAN Application Server
  - [x] Payload conversion of well-known payload formats
  - [x] Payload conversion using custom JavaScript functions
  - [x] MQTT pub/sub API
  - [x] HTTP Webhooks API
- LoRaWAN Join Server
  - [x] Supports OTAA session key derivation
  - [x] Supports external crypto services
  - [x] Implements LoRaWAN Back-end Interfaces 1.0
  - [x] Implements LoRaWAN Back-end Interfaces 1.1 draft 3
- OAuth 2.0 Identity Server
  - [x] User management
  - [x] Entity management
  - [x] ACLs
- GRPC APIs
- HTTP APIs
- Command-Line Interface
  - [x] Create account and login
  - [x] Application management and traffic
  - [x] End device management, status and traffic
  - [x] Gateway management and status
- Web Interface (Console)
  - [x] Create account and login
  - [x] Application management and traffic
  - [x] End device management, status and traffic
  - [x] Gateway management, status and traffic