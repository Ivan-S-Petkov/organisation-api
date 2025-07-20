# Organisation API

A REST API for managing users, licenses, and organizational plans built with Go, Gin, and PostgreSQL.

## Features

- **User Management**: Create, list, and filter users with role-based access
- **Plan Management**: Switch between organizational plans (Basic, Professional, Enterprise)
- **License Management**: Assign and unassign licenses to users
- **Search & Filtering**: Filter users by license status, role, and search by name/email

## Tech Stack

- **Language**: Go
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL
- **ORM**: GORM

## API Endpoints

### Authentication
- `POST /login` - User login

### Users
- `GET /users` - List users with pagination and filtering
- `POST /users` - Create new user (Admin only)

### Plans
- `GET /plan` - Get current organizational plan
- `POST /plan/switch` - Switch organizational plan (Admin only)

### Licenses
- `POST /licenses/assign` - Assign license to user (Admin only)
- `POST /licenses/unassign` - Remove license from user (Admin only)

### Query Parameters for `/users`

```bash
GET /users?page=1&perPage=10&hasLicense=true&role=admin&search=john
```

- `page` - Page number (default: 1)
- `perPage` - Items per page (default: 10)
- `hasLicense` - Filter by license status (true/false)
- `role` - Filter by user role (admin/user)
- `search` - Search in name and email fields
