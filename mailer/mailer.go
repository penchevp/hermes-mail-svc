package mailer

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/rs/zerolog/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	// Subject The subject line for the email.
	Subject = "Amazon SES Test (AWS SDK for Go)"

	// CharSet The character encoding for the email.
	CharSet = "UTF-8"
)

type Config struct {
	AccessKeyID     string
	SecretAccessKey string
	Sender          string
	Region          string
}

func Send(config Config, recipient string, body string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(config.AccessKeyID, config.SecretAccessKey, ""),
	})

	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(config.Sender),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aErr, ok := err.(awserr.Error); ok {
			switch aErr.Code() {
			case ses.ErrCodeMessageRejected:
				log.Err(aErr).Msg(ses.ErrCodeMessageRejected)
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				log.Err(aErr).Msg(ses.ErrCodeMailFromDomainNotVerifiedException)
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				log.Err(aErr).Msg(ses.ErrCodeConfigurationSetDoesNotExistException)
			default:
				log.Err(aErr)
			}
		} else {
			log.Err(err)
		}

		return err
	}

	// can save the resulting message id as a receipt, maybe send to a sink connector in elastic?
	fmt.Println(result)

	return nil
}
