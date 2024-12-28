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

SMTP_SERVER=smtp.gmail.com</br>
SMTP_PORT=587</br>
SMTP_USERNAME=your_email@gmail.com (eq Sender)</br>
SMTP_PASSWORD=your_app_password</br>
HTTP_PORT=8999</br>

*Server side envs (Dokploy, Coolify, etc..) are possible.*

----
HTML template files are located in the `templates` directory. You can modify these files to customize the email content.

---- 

### API Endpoints and required parameters
All endpoints are **POST** requests. </br>

`Single or multiple recipients can be added to the recipients array.`

**/mailservice/registration** 

```json
{
  "recipients": [
    {
      "email": "john@doe.com",
      "fname": "John",
      "lname": "Doe"
    },
    {
      "email": "jane@doe.com",
      "fname": "Jane",
      "lname": "Doe"
    }
  ]
}
```
**/mailservice/orderconf**

```json
{
  "recipients": [
    {
      "email": "john@doe.com",
      "fname": "John",
      "lname": "Doe"
    },
    {
      "email": "jane@doe.com",
      "fname": "Jane",
      "lname": "Doe"
    }
  ]
}
```
#### Note: Following endpoints are not implemented yet.
**/mailservice/notifyshipment**

```json
{
  "email": "john@doe.com",
  "fname": "John",
  "lname": "Doe"
}
```

**/mailservice/sendinvoice**

```json
{
  "email": "john@doe.com",
  "fname": "John",
  "lname": "Doe"
}
```
----
### Response Body

```json
{
    "message": "Email sent successfully!",
    "status": "success"
}
```
