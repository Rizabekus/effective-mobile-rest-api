# People Service API

## Overview

This service processes API requests for personal information, enriches the response with probable age, gender, and nationality using external APIs, and stores the data in a PostgreSQL database. It exposes various RESTful methods for managing people records.

## Project Summary

The People Service API is designed to provide a seamless interface for handling personal information. It utilizes external APIs to enhance the data with details such as age, gender, and nationality. The enriched data is stored in a PostgreSQL database, and the service offers RESTful methods for CRUD operations on people records. The project emphasizes clean code, efficient filtering, and pagination, along with comprehensive logging for monitoring and debugging.

## REST Methods

1. **Get People**
   - Endpoint: `/people`
   - Method: `GET`
   - Retrieves data with filters and pagination.

2. **Get Person by ID**
   - Endpoint: `/people/{id}`
   - Method: `GET`
   - Retrieves information for a specific person.

3. **Create Person**
   - Endpoint: `/people`
   - Method: `POST`
   - Adds a new person to the database.

4. **Update Person**
   - Endpoint: `/people/{id}`
   - Method: `PUT`
   - Modifies information for a specific person.

5. **Delete Person**
   - Endpoint: `/people/{id}`
   - Method: `DELETE`
   - Removes a person from the database.

## Enrichment and Database

- Age Enrichment: [Agify API](https://api.agify.io/?name=Dmitriy)
- Gender Enrichment: [Genderize API](https://api.genderize.io/?name=Dmitriy)
- Nationality Enrichment: [Nationalize API](https://api.nationalize.io/?name=Dmitriy)

Enriched data is stored in a PostgreSQL database, and the database structure is created through migrations.

## Logging

The code is extensively covered with debug and info logs to facilitate troubleshooting and monitoring.

## Configuration

Sensitive configuration data is stored in a `.env` file.
