# Certificate Verification Application

This project is a Certificate Verification Application that consists of a frontend, backend, RESTful API, and a MySQL database. The application allows users to verify certificates through a web interface.

## Project Structure

```
cert-verification-app
├── frontend              # Frontend application files
│   ├── index.html       # Main HTML document
│   ├── styles.css       # CSS styles for the frontend
│   └── app.js           # JavaScript code for user interactions
├── backend               # Backend application files
│   ├── src              # Source code for the Spring Boot application
│   │   ├── main         # Main application code
│   │   │   ├── java     # Java source files
│   │   │   │   └── com
│   │   │   │       └── example
│   │   │   │           └── CertVerificationApp
│   │   │   │               ├── CertVerificationAppApplication.java
│   │   │   │               └── controller
│   │   │   │                   └── CertController.java
│   │   │   └── resources  # Resources for the Spring Boot application
│   │   │       ├── application.properties
│   │   │       └── static
│   │   └── test          # Test files for the Spring Boot application
│   │       └── java
│   │           └── com
│   │               └── example
│   │                   └── CertVerificationApp
│   │                       └── CertVerificationAppApplicationTests.java
│   ├── pom.xml          # Maven configuration file
│   └── README.md        # Documentation for the backend
├── api                  # Go RESTful API files
│   ├── main.go          # Entry point for the Go API
│   └── README.md        # Documentation for the Go API
├── db                   # Database setup files
│   ├── schema.sql       # SQL schema for MySQL database
│   └── README.md        # Documentation for the database
└── README.md            # Overall documentation for the project
```

## Setup Instructions

1. **Frontend**: Navigate to the `frontend` directory and open `index.html` in a web browser to access the application.

2. **Backend**: 
   - Navigate to the `backend` directory.
   - Build the application using Maven with the command: `mvn clean install`.
   - Run the Spring Boot application with: `mvn spring-boot:run`.

3. **API**: 
   - Navigate to the `api` directory.
   - Run the Go API with the command: `go run main.go`.

4. **Database**: 
   - Set up the MySQL database using the schema defined in `db/schema.sql`.
   - Update the database connection settings in `backend/src/main/resources/application.properties`.

## Usage

- Access the frontend application through your web browser.
- Use the provided interface to verify certificates.
- The backend and API handle the logic for certificate verification and database interactions.

## Contributing

Feel free to contribute to this project by submitting issues or pull requests.