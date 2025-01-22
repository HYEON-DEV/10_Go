# Database Documentation for Cert Verification App

This README file provides information about the database setup and usage for the Cert Verification App.

## Database Overview

The Cert Verification App uses a MySQL database to store and manage certificate data. The database schema is defined in the `schema.sql` file located in this directory.

## Setup Instructions

1. **Install MySQL**: Ensure that you have MySQL installed on your machine. You can download it from the official MySQL website.

2. **Create Database**: Create a new database for the Cert Verification App. You can do this using the MySQL command line or a GUI tool like MySQL Workbench.

   ```sql
   CREATE DATABASE cert_verification_db;
   ```

3. **Run Schema**: Execute the `schema.sql` file to set up the necessary tables and relationships.

   ```sql
   SOURCE path/to/cert-verification-app/db/schema.sql;
   ```

4. **Configure Connection**: Update the database connection settings in the `application.properties` file located in the backend resources directory to point to your MySQL database.

## Usage

Once the database is set up, the Cert Verification App will interact with it to store and retrieve certificate information as needed. Ensure that the database server is running when you start the application.

## Additional Information

For more details on the database schema and the specific tables created, refer to the `schema.sql` file.