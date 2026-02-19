from fastapi import APIRouter, HTTPException, Path
from pydantic import BaseModel
from typing import Dict, Any
from .templates import template_manager
from .triggers import get_trigger

router = APIRouter()

class TriggerPayload(BaseModel):
    variables: Dict[str, Any]

@router.post("/trigger/{trigger_id}")
def trigger_email(payload: TriggerPayload, trigger_id: str = Path(..., title="The ID of the trigger")):
    trigger = get_trigger(trigger_id)
    if not trigger:
        raise HTTPException(status_code=404, detail="Trigger not found")

    variables = payload.variables

    # Render Subject
    subject = template_manager.render(trigger["subject"], variables)

    # Render Body
    body = template_manager.render(trigger["template"], variables)

    # Mock sending email
    for recipient in trigger["recipients"]:
        print(f"Sending managed email to {recipient}")
        print(f"Subject: {subject}")
        print(f"Body: {body}")

    return {
        "status": "sent",
        "trigger_id": trigger_id,
        "recipients": trigger["recipients"],
        "rendered_subject": subject,
        "rendered_body": body
    }
