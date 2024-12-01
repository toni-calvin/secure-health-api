# TopDoctors Challenge

Welcome to the TopDoctors Challenge! This is a backend application built with Go (golang) and PostgreSQL, designed to run seamlessly in a Dockerized environment.

---

## Requirements
Before you start, ensure you have the following installed:

 - 🐳 Docker - to containerize and run the application.
 - 🛠️ Docker Compose - to orchestrate multi-container applications.

---

## Getting Started

### 1. Clone the Repository
Begin by cloning the repository to your local machine:

```bash
git clone https://github.com/toni-calvin/topdoctors-challenge.git
cd topdoctors-challenge
```

### 2. Configure Admin User
An admin user is required to bootstrap the system.

1. Copy the .env.example file to .env:
```bash 
cp .env.example .env 
```
2. Open the .env file and set your admin username and admin password:
```bash 
ADMIN_USER=your_admin_username
ADMIN_PASSWORD=your_admin_password
```
This admin user will be used to create additional users within the application to later be able to login. 

### 3. Build and Run the Application
Use the provided Makefile to build and run the application:
1. Build the Docker image 
```bash
make build 
```

2. Run the application 
```bash
make run 
```

The application will now be running on http://localhost:8080.

