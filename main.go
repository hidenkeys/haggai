package main

import (
	"fmt"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"log"
	"net/mail"
)

func main() {
	// Initialize PocketBase app
	app := pocketbase.New()

	// Fires when a "user" record is created
	app.OnRecordCreate("TasterBoxInquiry", "bespokecakeinquiry", "weddingcakeinquiry", "workshopbooking\n").BindFunc(func(e *core.RecordEvent) error {
		// Ensure next handler is processed
		if err := e.Next(); err != nil {
			return err
		}

		fmt.Println(e.App.Settings().Meta.SenderAddress)
		fmt.Println(e.App.Settings().Meta.SenderName)
		// Construct the email message
		message := &mailer.Message{
			From: mail.Address{
				Address: e.App.Settings().Meta.SenderAddress,
				Name:    e.App.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: "teniolasobande04@gmail.com"}},         // Sending email to the created user's email
			Subject: "Welcome to Our Platform",                                       // Set the subject for the email
			HTML:    fmt.Sprintf("<h1>New User Record</h1><pre>%s</pre>", *e.Record), // JSON formatted body

		}

		// Send the email using PocketBase's built-in mailer
		if err := e.App.NewMailClient().Send(message); err != nil {
			return fmt.Errorf("failed to send email: %w", err)
		}

		// Continue processing the event
		return nil
	})

	// Start the app
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
