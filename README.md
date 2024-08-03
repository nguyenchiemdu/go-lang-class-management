# Class Management API

## Overview

This project is a Go-based API for managing classes, teachers, and students using ArangoDB as the database. The API includes endpoints for CRUD operations on students, teachers, and classes, as well as functionality to manage class assignments and teacher-student associations.

## Project Structure

The project is structured as follows:
- **/config:** Contains configuration files and utilities for loading and managing application settings.
- **/controller:** Contains HTTP handlers and business logic for various endpoints, including user authentication and CRUD operations.
- **/dbservice:** Manages database interactions, including initialization, data retrieval, and updates.
- **/models:** Defines data models and types used throughout the application.
- **/utils:** Provides utility functions for tasks such as password hashing and token creation.

## Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/class-management.git
   cd class-management

1. **Install Dependencies**

   ```bash
   go mod tidy
   go get .

3. **Setup Environment Variables**

   ```sh
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_PORT=8529
   DB_NAME=class_management
   JWT_SECRET=your_jwt_secret
   JWT_EXPIRE=24h

4. **Run the Application**
   ```bash
   go run main.go

## Naming Conventions

This project follows the Go Naming Conventions:

	•	Package Names: Use short, lowercase names without underscores or camelCase. Example: models, controller.
	•	Variable Names: Use camelCase for variable names. Example: classID, teacherID.
	•	Function Names: Use camelCase for function names. Functions that are exported (public) should start with an uppercase letter. Example: GetClassByID, CreateUser.
	•	Constants: Use UPPER_SNAKE_CASE for constants. Example: JWT_SECRET, DB_PORT.
	•	Struct Names: Use CamelCase for struct names. Example: Class, Student.
	•	Error Variables: Use err for error variables. Example: err, dbError.
References: https://www.mohitkhare.com/blog/go-naming-conventions/

## API Endpoints

Refer to the Postman documentation for detailed API endpoint information and usage.

[JSON Postman file](postman.json)

## License
This project is licensed under the MIT License
