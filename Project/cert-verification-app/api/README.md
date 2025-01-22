# Certificate Verification API

This directory contains the Go RESTful API for the Certificate Verification application. The API is responsible for handling requests related to certificate verification and interacting with the backend services.

## Overview

The API provides endpoints to verify certificates and manage certificate-related data. It is designed to be lightweight and efficient, serving as a bridge between the frontend application and the backend services.

## Getting Started

To run the API, ensure you have Go installed on your machine. Follow these steps:

1. Clone the repository:
   ```
   git clone <repository-url>
   cd cert-verification-app/api
   ```

2. Install the necessary dependencies:
   ```
   go mod tidy
   ```

3. Run the API:
   ```
   go run main.go
   ```

The API will start on the default port (8080). You can change the port in the `main.go` file if needed.

## API Endpoints

- `POST /verify`: Verifies a certificate and returns the verification status.
- `GET /certificates`: Retrieves a list of all certificates.
- `GET /certificates/{id}`: Retrieves details of a specific certificate by ID.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.