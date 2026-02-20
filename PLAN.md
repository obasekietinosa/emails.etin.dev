# Email Service Project Plan

## Overview
This project aims to build a managed email service with two modes of operation:
1.  **Fully Managed Mode**: Users define templates with variables. Emails are triggered via a unique URL.
2.  **Headless Mode**: Users send raw email content via an API, authenticated with an API Key.

## Technology Stack
-   **Backend**: Go (Golang)
-   **Frontend**: React TypeScript
-   **Database**: SQLite (via GORM)
-   **API Framework**: Chi
-   **Authentication**: JWT (for Web Client), API Keys (for Headless Mode)

---

## Detailed Roadmap

### Phase 1: Backend API (Go)

#### 1. Initialization
-   [ ] Initialize Go module `github.com/tinrab/emails` (or similar).
-   [ ] Set up directory structure:
    -   `cmd/api`: Main application entry point.
    -   `internal/database`: Database connection and GORM setup.
    -   `internal/handlers`: HTTP request handlers.
    -   `internal/models`: Database models.
    -   `internal/service`: Business logic (Auth, Email).
    -   `internal/middleware`: Custom middleware (Auth).

#### 2. Database & Models
-   [ ] Define Models:
    -   `User`: ID, Email, PasswordHash.
    -   `Template`: ID, UserID, Name, Subject, Body, TriggerToken.
    -   `ApiKey`: ID, UserID, Key, Name.
    -   `EmailLog`: ID, UserID, Recipient, Subject, Status, Mode.
-   [ ] Setup SQLite connection and Auto-migration.

#### 3. Core Services
-   [ ] **Auth Service**: JWT generation and validation.
-   [ ] **Email Service**: Interface `EmailSender` with a `MockSender` implementation that logs to stdout.

#### 4. API Endpoints
-   [ ] **Authentication**:
    -   `POST /register`: User registration.
    -   `POST /login`: User login (returns JWT).
-   [ ] **Managed Mode (Templates)**:
    -   `GET /templates`: List user templates.
    -   `POST /templates`: Create a new template.
    -   `GET /templates/{id}`: Get template details.
    -   `PUT /templates/{id}`: Update template.
    -   `DELETE /templates/{id}`: Delete template.
    -   `POST /trigger/{token}`: Trigger an email from a template (public/unauthenticated or token-based).
-   [ ] **Headless Mode**:
    -   `POST /send`: Send raw email (Protected by API Key).
    -   `POST /api-keys`: Generate API Key (Protected by JWT).

#### 5. Middleware
-   [ ] **JWT Middleware**: Protects web client endpoints.
-   [ ] **API Key Middleware**: Protects headless mode endpoints.

---

### Phase 2: Frontend Web Client (React TS)

#### 1. Setup
-   [ ] Initialize project with Vite (React + TypeScript).
-   [ ] Setup React Router.
-   [ ] Setup UI Library (e.g., Tailwind CSS).

#### 2. Features
-   [ ] **Authentication**: Login and Register pages.
-   [ ] **Dashboard**: Overview of recent activity.
-   [ ] **Template Management**:
    -   List view of templates.
    -   Create/Edit form with preview.
    -   View Trigger URL for each template.
-   [ ] **API Key Management**:
    -   Generate and view API keys.
-   [ ] **Email Logs**: View history of sent emails.

---

### Phase 3: Integration & Polish
-   [ ] Connect Frontend to Backend API.
-   [ ] End-to-end testing of both modes.
-   [ ] Finalize documentation.
