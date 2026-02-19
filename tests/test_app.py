from fastapi.testclient import TestClient
from src.main import app
from src.triggers import TRIGGERS

client = TestClient(app)

def test_read_root():
    response = client.get("/")
    assert response.status_code == 200
    assert response.json() == {"message": "Welcome to the Managed Email Service"}

def test_headless_mode_success():
    payload = {
        "to": "test@example.com",
        "subject": "Test Subject",
        "body": "Test Body"
    }
    headers = {"x-api-key": "valid-api-key"}
    response = client.post("/send/headless", json=payload, headers=headers)
    assert response.status_code == 200
    assert response.json() == {"status": "sent", "to": "test@example.com"}

def test_headless_mode_invalid_api_key():
    payload = {
        "to": "test@example.com",
        "subject": "Test Subject",
        "body": "Test Body"
    }
    headers = {"x-api-key": "invalid-key"}
    response = client.post("/send/headless", json=payload, headers=headers)
    assert response.status_code == 401
    assert response.json() == {"detail": "Invalid API Key"}

def test_headless_mode_missing_api_key():
    payload = {
        "to": "test@example.com",
        "subject": "Test Subject",
        "body": "Test Body"
    }
    response = client.post("/send/headless", json=payload)
    assert response.status_code == 401
    assert response.json() == {"detail": "API Key missing"}

def test_managed_mode_success():
    trigger_id = "welcome-trigger"
    payload = {"variables": {"name": "TestUser"}}
    response = client.post(f"/trigger/{trigger_id}", json=payload)
    assert response.status_code == 200
    data = response.json()
    assert data["status"] == "sent"
    assert data["trigger_id"] == trigger_id
    assert data["rendered_subject"] == "Welcome TestUser!"
    assert data["rendered_body"] == "Hello TestUser, welcome to our service!"

def test_managed_mode_trigger_not_found():
    trigger_id = "non-existent-trigger"
    payload = {"variables": {"name": "TestUser"}}
    response = client.post(f"/trigger/{trigger_id}", json=payload)
    assert response.status_code == 404
    assert response.json() == {"detail": "Trigger not found"}
