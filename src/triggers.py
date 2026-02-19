from typing import Dict, TypedDict, List

class TriggerData(TypedDict):
    template: str
    subject: str
    recipients: List[str]

# Mock database of triggers
TRIGGERS: Dict[str, TriggerData] = {
    "welcome-trigger": {
        "template": "Hello {{ name }}, welcome to our service!",
        "subject": "Welcome {{ name }}!",
        "recipients": ["admin@example.com"]
    },
    "alert-trigger": {
        "template": "Alert: {{ message }}",
        "subject": "System Alert",
        "recipients": ["ops@example.com"]
    }
}

def get_trigger(trigger_id: str) -> TriggerData:
    return TRIGGERS.get(trigger_id)
