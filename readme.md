# Go Mailservice

A lightweight Go-based mail service that provides a REST API for sending emails. Configuration is managed via environment variables, which can be supplied through a `.env` file or server-side setup.

---

## Requirements

- Go (1.23.4)
- Docker (optional for deployment)

---

## Setup

### Environment Variables

The service requires the following environment variables:

- `SMTP_SERVER` - SMTP server address (e.g., `smtp.gmail.com`)
- `SMTP_PORT` - SMTP server port (e.g., `587`)
- `SMTP_USERNAME` - SMTP username (your email)
- `SMTP_PASSWORD` - SMTP password or app password
- `SMTP_SENDER` - Email address used as sender

#### Using a `.env` file

Create a `.env` file in the project directory:

SMTP_SERVER=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your_email@gmail.com (eq Sender)
SMTP_PASSWORD=your_app_password
HTTP_PORT=8999

Server side enviroments (Dokploy) possible.

---- 

### Request Body

{
  "recipient": "recipient@example.com",
  "subject": "Test Email",
  "message": "This is a test email."
}

### Response Body

{
    "message": "Email sent successfully!",
    "status": "success"
}