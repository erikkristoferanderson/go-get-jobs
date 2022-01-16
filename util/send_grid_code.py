"""Sends an email using the sendgrid api
Relies on environment variable SENDGRID_API_KEY"""

import os
from sendgrid import SendGridAPIClient
from sendgrid.helpers.mail import Mail


def send_email(sender, recipient, subject, html_content):
    message = Mail(
        from_email=sender,
        to_emails=recipient,
        subject=subject,
        html_content=html_content
    )
    try:
        sg = SendGridAPIClient(os.environ.get('SENDGRID_API_KEY'))
        response = sg.send(message)
        print(response.status_code)
        print(response.body)
        print(response.headers)
    except Exception as e:
        print(e.message)
