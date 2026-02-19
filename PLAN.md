# Project Plan: Managed Email Service

## Overview
This project implements a managed email service with two modes:
1.  **Headless Mode**: Sends emails based on a full payload (recipient, subject, body).
2.  **Fully Managed Mode**: Triggers pre-defined templates with variable interpolation.

## Architecture

### Tech Stack
-   **Language**: Python 3.x
-   **Framework**: FastAPI
-   **Template Engine**: Jinja2
-   **Testing**: Pytest

### Directory Structure
```
.
├── README.md
├── requirements.txt
├── src/
│   ├── main.py        # Application entry point
│   ├── headless.py    # Headless mode logic
│   ├── managed.py     # Managed mode logic
│   ├── templates.py   # Template management
│   └── triggers.py    # Trigger storage and retrieval
└── tests/
    └── test_app.py    # Integration tests
```

## Implementation Steps

### 1. Project Initialization
-   Set up `requirements.txt` with dependencies: `fastapi`, `uvicorn`, `jinja2`, `python-dotenv`.
-   Create the `src/` directory and a basic `main.py` application.

### 2. Headless Mode
-   **Endpoint**: `POST /send/headless`
-   **Payload**:
    ```json
    {
      "api_key": "string",
      "to": "email@example.com",
      "subject": "string",
      "body": "string"
    }
    ```
-   **Logic**: Validate API key (mock validation), simulate sending email (print to console).

### 3. Fully Managed Mode
-   **Endpoint**: `POST /trigger/{trigger_id}`
-   **Payload**:
    ```json
    {
      "variables": {
        "key": "value"
      }
    }
    ```
-   **Logic**:
    -   Lookup trigger by ID to find the template and default recipient.
    -   Render the template using Jinja2 with provided variables.
    -   Simulate sending email.

### 4. Testing
-   Implement tests for both endpoints to ensure correct status codes and responses.

## Future Improvements
-   Add persistent storage (database).
-   Integrate with a real email provider (e.g., SendGrid, SES).
-   Add proper authentication and authorization.
