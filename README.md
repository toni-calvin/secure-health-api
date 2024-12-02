# TopDoctors Challenge

Welcome to the TopDoctors Challenge! This is a backend application built with Go (golang) and PostgreSQL, designed to run seamlessly in a Dockerized environment.

---

## Requirements
Before you start, ensure you have the following installed:

 - 🐳 Docker - to containerize and run the application.
 - 🛠️ Docker Compose - to orchestrate multi-container applications.
 - 🐍 Python 3.x - for running the provided scenario testing script.
 - 📦 pip - for installing Python dependencies.

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
The admin user will be used to create additional users within the application for subsequent operations.

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
Before running any tests or scenarios, ensure that the Docker console logs the following message:

```bash
app-1  | <timestamp> Server running on port 8080
```

The application will now be running on http://localhost:8080 and is ready to accept requests.

## Running the Scenario Script (in another shell)
### 1. Create and activate the Virtual Environment 
```bash
python3 -m venv venv
source venv/bin/activate
```
### 2. Install required dependencies
```bash
pip install -r requirements.txt
```

### 3. Run the Scenario Script
```bash
pip install -r requirements.txt
```

### 3. Expected Behavior
On the first run, the script will:
- Log in with the admin credentials.
- Create internal and external users.
- Create patients and diagnoses.
- List patients and diagnoses with filters.

On subsequent runs, you will see errors for duplicate creations (e.g., duplicate users, patients, or diagnoses already existing in the database).

This behavior is intentional to demonstrate the handling of unique constraints and application logic.

## Running Test (in another shell)

## Additional Notes 
- To rebuild the applicatin you can run ```bash make rebuild ```

