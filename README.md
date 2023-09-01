# Order Delay Manager

Order Delay Manager is a rest api service for manage delayed orders and assign each of them to agent. It is built with Go, using the Cobra CLI library for command-line functionality, MySQL for database storage, Echo as the web framework, and Swagger for API documentation. Docker and Docker Compose are used for containerization and orchestration.

## Table of Contents

- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
- [Usage](#usage)
    - [CLI Commands](#cli-commands)
    - [API Endpoints](#api-endpoints)

## Getting Started

### Prerequisites

Ensure you have the following software installed on your system:

- Go: [Installation Guide](https://golang.org/doc/install)
- Docker: [Installation Guide](https://docs.docker.com/get-docker/)
- Docker Compose: [Installation Guide](https://docs.docker.com/compose/install/)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/yourproject.git
   cd yourproject
2. Build the Docker containers:
   ```bash
   docker-compose build
3. Start the Docker containers:
   ```bash
   docker-compose up -d
4. Access the application at [http://localhost:PORT]() where PORT is the port specified in your env.

### Usage
CLI Commands

To use the CLI, navigate to the project's root directory and run the following commands:
1. Run Rest API Command
   ```bash
   ./task-app rest
2. Run Seed database Command
   ```bash
   ./task-app seed
3. Run Migration Command
   ```bash
   ./task-app migration up/down

### API Endpoints
The API documentation can be found at [http://localhost:PORT/swagger/index.html]() after starting the Docker containers.

### Notice
If you want to use redis version of Queue management instead of in memory, you need to add it to docker compose and call initialize it in startup