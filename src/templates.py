from jinja2 import Environment, BaseLoader

class TemplateManager:
    def __init__(self):
        # In a real app, this might load from files or a DB
        self.env = Environment(loader=BaseLoader())

    def render(self, template_str: str, variables: dict) -> str:
        template = self.env.from_string(template_str)
        return template.render(**variables)

template_manager = TemplateManager()
