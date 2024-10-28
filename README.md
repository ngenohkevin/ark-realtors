
---
# Ark Realtors

`This is a hobby project and it's still ongoing. I'm working on it in my free time. I'm open to any suggestions and contributions. Feel free to contact me.`

## Project Overview

Ark Realtors is a real estate management application that allows users to list, search, and manage properties available for rent or sale. Built with a robust backend in Golang and a PostgreSQL database, Ark Realtors supports secure user management, session handling, and property image storage.

## Features

- **User Authentication**: Secure user management with hashed passwords and session control.
- **Property Listings**: Manage property listings with details like type, price, bedroom and bathroom count, and location.
- **Property Images**: Attach images to properties with descriptions to enhance property listings.
- **Session Management**: Track user sessions with refresh tokens, user-agent, IP address, and expiration time.

## Database Schema

This project uses PostgreSQL as the primary database, with the following tables:

### 1. `users`
Stores registered users with secure hashed passwords and roles.
- **Columns**:
    - `id` (UUID, Primary Key)
    - `username` (String, Unique, Not Null)
    - `full_name` (String, Not Null)
    - `email` (String, Unique, Not Null)
    - `hashed_password` (String, Not Null)
    - `role` (String, Default: `user`)
    - `password_changed_at` (Timestamp with Timezone, Default: `0001-01-01 00:00:00Z`)
    - `created_at` (Timestamp with Timezone, Default: `now()`)

### 2. `property`
Stores details about properties listed for rent or sale.
- **Columns**:
    - `id` (UUID, Primary Key, Foreign Key referencing `users.id`)
    - `type` (String, Not Null, e.g., `rent` or `sale`)
    - `price` (Numeric(7,2), Not Null)
    - `status` (String, Default: `available`)
    - `bedroom` (Integer, Not Null)
    - `bathroom` (Integer, Not Null)
    - `location` (String, Not Null)
    - `size` (String, Not Null)
    - `contact` (String, Not Null)
    - `created_at` (Timestamp with Timezone, Default: `now()`)

### 3. `pictures`
Stores image URLs for property listings.
- **Columns**:
    - `id` (UUID, Primary Key)
    - `property_id` (UUID, Foreign Key referencing `property.id`)
    - `img_url` (String, Not Null)
    - `description` (String, Not Null)

### 4. `sessions`
Manages user session information for secure access control.
- **Columns**:
    - `id` (UUID, Primary Key)
    - `username` (String, Foreign Key referencing `users.username`, Not Null)
    - `refresh_token` (String, Not Null)
    - `user_agent` (String, Not Null)
    - `client_ip` (String, Not Null)
    - `is_blocked` (Boolean, Default: `false`)
    - `expires_at` (Timestamp with Timezone, Not Null)
    - `created_at` (Timestamp with Timezone, Default: `now()`)

## Tech Stack

- **Golang**: Backend development
- **Gin Gonic**: HTTP web framework
- **sqlc**: SQL query generation
- **Docker & Docker Compose**: Containerization for application and database services
- **DBML**: Database modeling
- **PostgreSQL**: Database
- **Database Migration**: Smooth database schema changes
- **GitHub Actions**: CI/CD for automated testing and deployment

## Getting Started

### Prerequisites

- **Docker & Docker Compose**
- **Go**
- **PostgreSQL**

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/arkrealtors.git
   cd arkrealtors
   ```

2. Set up environment variables:
   Create a `.env` file with necessary configurations, such as database URL, ports, etc.

3. Run database migrations:
   ```bash
   make migrate-up
   ```

4. Build and run the application:
   ```bash
   docker-compose up --build
   ```

5. Access the API:
   The API will be available at `http://localhost:<PORT>`.

## Database Management

Database migrations are handled through a migration tool to ensure schema changes are implemented smoothly. To apply migrations:
```bash
make migrate-up
```

## CI/CD

GitHub Actions are used for continuous integration and deployment. Each push or pull request triggers tests and build processes to ensure stability.

---

**Enjoy managing your real estate listings with Ark Realtors!**