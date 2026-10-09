# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Environment Setup
```bash
# Initial setup - create environment files
cd api && cp .env.sample .env
cd ../mysql && cp .env.sample .env

# Start development environment
docker compose up -d

# Check container status
docker compose ps
```

### Testing
```bash
# Run all tests
cd api && go test ./...

# Run tests for specific package
go test ./domain/user
go test ./usecase/account/login
go test ./usecase/account/register

# Run tests with verbose output
go test -v ./...
```

### Local Development
```bash
# Access application locally
# POST to http://localhost:8080/account/check_auth with:
# {"id": "id", "token": "token"}
# Should return auth error indicating proper setup
```

### Deployment
```bash
# Deploy to Google Cloud (requires main branch)
./deploy.sh

# Manual Cloud Build
gcloud builds submit --region=us-central1 --config cloudbuild.yaml
```

## Architecture Overview

This is a Clean Architecture Go application implementing user authentication with three core functions: registration, login, and authentication checking.

### Layer Structure

**Domain Layer** (`/api/domain/`):
- User entity accessed through interface pattern (not exported struct)
- Password value object with automatic SHA256 hashing
- Repository interfaces defining data contracts
- Domain services for business logic (e.g., duplicate user checking)

**Usecase Layer** (`/api/usecase/`):
- Registration and login workflows with comprehensive validation
- DTOs for clean input/output boundaries
- Dependency injection using repository interfaces

**Infrastructure Layer** (`/api/infrastructure/`):
- Dual database support: Docker MySQL (local) + Cloud SQL (production)
- Repository implementations with transaction management
- Environment-driven connection switching via `DB_FLAG`

**Presentation Layer** (`/api/presentation/`):
- Chi router with comprehensive middleware stack
- Generic input mapper using Go generics
- Account controller handling three endpoints

### Key Patterns

- **Interface Segregation**: User entity hidden behind interface
- **Factory Pattern**: Domain entities with validation
- **Repository Pattern**: Clean data access abstraction
- **Generic Programming**: Type-safe request parsing middleware

### Security Implementation

- ULID for user IDs (not sequential)
- SHA256 password hashing
- JWT tokens with 11-minute expiration
- Secure cookie configuration (HttpOnly, Secure, SameSite)

### Database Schema

Single `users` table with fields: id (ULID), name, mail (unique), imagePath, pass (hashed). Schema auto-applied via Docker entrypoint from `/mysql/schema/`.

### Environment Configuration

Critical environment variables:
- `DB_FLAG`: "GCP" for Cloud SQL, anything else for local MySQL
- `JWT_SECRET_KEY`: Token signing key
- `ALLOW_ORIGIN`: CORS configuration
- Database credentials for Cloud SQL deployment

### Testing Strategy

Uses testify/mock and testify/assert with table-driven tests. Mock repositories provided for unit testing. All tests include comprehensive error case coverage and use Japanese test case labels.

### Deployment Architecture

Cloud Run service with auto-scaling (0-5 instances), Cloud Build pipeline, and Artifact Registry. Deployment script includes safety checks for main branch requirement and GCP authentication validation.