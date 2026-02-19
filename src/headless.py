from fastapi import APIRouter, HTTPException, Header
from pydantic import BaseModel
from typing import Optional

router = APIRouter()

class EmailPayload(BaseModel):
    to: str
    subject: str
    body: str

# Mock database of valid API keys
VALID_API_KEYS = {"valid-api-key"}

def verify_api_key(api_key: str):
    if api_key not in VALID_API_KEYS:
        raise HTTPException(status_code=401, detail="Invalid API Key")
    return True

@router.post("/send/headless")
def send_email_headless(payload: EmailPayload, x_api_key: Optional[str] = Header(None)):
    if not x_api_key:
        raise HTTPException(status_code=401, detail="API Key missing")

    verify_api_key(x_api_key)

    # Mock email sending logic
    print(f"Sending email to {payload.to}")
    print(f"Subject: {payload.subject}")
    print(f"Body: {payload.body}")

    return {"status": "sent", "to": payload.to}
