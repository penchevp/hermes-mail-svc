# hermes-mail-svc

hermes-mail-svc is a microservice that sends email notifications to opted-in customers using Amazon SES. hermes-mail-svc holds a copy of customer data in its own database (eventual consistency).

Fault tolerant.

# Environment variables

| Name | Description |
| :--- | :--- |
| DB_HOST | Database server to connect to |
| DB_PORT | Database server port |
| DB_NAME | Database to connect to |
| DB_USERNAME | User for database connection |
| DB_PASSWORD | User password for database connection |
| MAIL_ACCESS_KEY | AWS access key id |
| MAIL_SECRET_ACCESS_KEY | AWS secret access key |
| MAIL_SENDER | Sender on whose behalf emails are sent |
| MAIL_REGION | AWS region where SES is set up |

# Dependencies
* `github.com/segmentio/kafka-go`
* `github.com/aws/aws-sdk-go`
* `github.com/rs/zerolog`
* `gorm.io/gorm`
* `gorm.io/driver/postgres`
* `github.com/google/uuid`
