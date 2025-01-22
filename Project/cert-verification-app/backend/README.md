# Cert Verification App - Backend

This document provides an overview of the backend component of the Cert Verification App, which is built using Spring Boot.

## Project Structure

- `src/main/java/com/example/CertVerificationApp`: Contains the main application code.
  - `CertVerificationAppApplication.java`: The entry point of the Spring Boot application.
  - `controller/CertController.java`: Handles HTTP requests related to certificate verification.
  
- `src/main/resources`: Contains configuration and static resources.
  - `application.properties`: Configuration properties for the application.
  - `static`: Directory for serving static files.

- `src/test/java/com/example/CertVerificationApp`: Contains test cases for the application.
  - `CertVerificationAppApplicationTests.java`: Unit tests for the application.

- `pom.xml`: Maven configuration file for managing dependencies and build settings.

## Getting Started

1. **Prerequisites**: Ensure you have Java and Maven installed on your machine.

2. **Clone the Repository**: 
   ```
   git clone <repository-url>
   cd cert-verification-app/backend
   ```

3. **Build the Application**: 
   ```
   mvn clean install
   ```

4. **Run the Application**: 
   ```
   mvn spring-boot:run
   ```

5. **Access the API**: The backend API will be available at `http://localhost:8080`.

## API Endpoints

- **POST /verify**: Endpoint to verify certificates. (Details to be defined in `CertController.java`)

## Testing

Run the tests using Maven:
```
mvn test
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.