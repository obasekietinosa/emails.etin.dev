# emails.etin.dev
Managed email service with two modes of operation: headless mode and fully managed.

In fully managed mode, the user provides a template they want sent to a specific email address or group of email addresses when a trigger is hit and can specify variables to be interpolated into this template. We provide them a unique, generated URL and they can hit that will kick off the flow.

In headless mode, the user provides the full email content, along with the receiver email address and all the trappings of an email. They hit the generic provided endpoint (same for all users) with the information making up the email along with their authentication method, like an API key, since we expect to be authenicating a server.
