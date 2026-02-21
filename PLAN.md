# Email Service Project Plan

## Overview
This project aims to build a managed email service with two modes of operation:
1.  **Fully Managed Mode**: Users define templates with variables. Emails are triggered via a unique URL.
2.  **Headless Mode**: Users send raw email content via an API, authenticated with an API Key.

## Technology Stack
-   **Backend**: Go (Golang)
-   **Frontend**: React TypeScript
-   **Database**: PostgreSQL (via GORM). All environments (Dev/Prod) use Postgres.
-   **API Framework**: Chi
-   **Authentication**: JWT (for Web Client), API Keys (for Headless Mode)

---

## Detailed Roadmap

### Phase 1: Backend API (Go)

#### 1. Initialization
-   [x] Initialize Go module `github.com/tinrab/emails`.
-   [x] Set up directory structure.
-   [x] Install dependencies.

#### 2. Database & Models
-   [x] Define Models: `User`, `Template`, `ApiKey`, `EmailLog`.
-   [x] Setup PostgreSQL connection.

#### 3. Core Services
-   [x] **Auth Service**: JWT generation and validation.
-   [x] **Email Service**: Interface `EmailSender` with `MockSender` implementation.

#### 4. API Endpoints
-   [x] **Authentication**: `/register`, `/login`.
-   [x] **Managed Mode**: Template CRUD, Trigger endpoint.
-   [x] **Headless Mode**: `/send` endpoint, API Key generation.

#### 5. Middleware
-   [x] **JWT Middleware**: Protects web client endpoints.
-   [x] **API Key Middleware**: Protects headless mode endpoints.

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
