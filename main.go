package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"log"
	"net/mail"
	"os"
	"strconv"
)

// Function to create email body for TasterBoxInquiry
func createTasterBoxInquiryEmail(e *core.RecordEvent) string {
	return fmt.Sprintf(`
		<h1>TasterBox Inquiry Details</h1>
		<p><strong>Event Type:</strong> %s</p>
		<p><strong>Event Date:</strong> %s</p>
		<p><strong>Event Venue:</strong> %s</p>
		<p><strong>Delivery Time:</strong> %s</p>
		<p><strong>No of Taster Boxes:</strong> %d</p>
		<p><strong>Flavours List:</strong> %s</p>
		<p><strong>Dietary Needs:</strong> %s</p>
		<p><strong>Company Name:</strong> %s</p>
		<p><strong>Contact Name:</strong> %s</p>
		<p><strong>Email:</strong> %s</p>
		<p><strong>Phone Number:</strong> %s</p>
		<p><strong>Additional Details:</strong> %s</p>
	`,
		e.Record.GetString("event_type"),
		e.Record.GetString("event_date"),
		e.Record.GetString("event_venue"),
		e.Record.GetString("delivery_time"),
		e.Record.GetInt("no_of_taster_boxes"),
		e.Record.GetString("flavlors_list"),
		e.Record.GetString("dietary_needs"),
		e.Record.GetString("company_name"),
		e.Record.GetString("contact_name"),
		e.Record.GetString("email"),
		e.Record.GetString("phone_number"),
		e.Record.GetString("additional_details"),
	)
}

// Function to create email body for BespokeCakeInquiry
func createBespokeCakeInquiryEmail(e *core.RecordEvent) string {
	return fmt.Sprintf(`
		<h1>Bespoke Cake Inquiry Details</h1>
		<p><strong>Event Date:</strong> %s</p>
		<p><strong>Event Time:</strong> %s</p>
		<p><strong>Theme:</strong> %s</p>
		<p><strong>Number of Tiers:</strong> %d</p>
		<p><strong>Tier Shape:</strong> %s</p>
		<p><strong>Size:</strong> %s</p>
		<p><strong>Flavours:</strong> %s</p>
		<p><strong>Dietary Needs:</strong> %s</p>
		<p><strong>Design Inspiration:</strong> %s</p>
		<p><strong>Couple's Names:</strong> %s</p>
		<p><strong>Email:</strong> %s</p>
		<p><strong>Phone Number:</strong> %s</p>
		<p><strong>Additional Details:</strong> %s</p>
	`,
		e.Record.GetString("event_date"),
		e.Record.GetString("event_time"),
		e.Record.GetString("theme"),
		e.Record.GetInt("number_of_tiers"),
		e.Record.GetString("tier_shape"),
		e.Record.GetString("size"),
		e.Record.GetString("flavours"),
		e.Record.GetString("dietary_needs"),
		e.Record.GetString("design_inspiration"),
		e.Record.GetString("couples_names"),
		e.Record.GetString("email"),
		e.Record.GetString("phone_number"),
		e.Record.GetString("additional_details"),
	)
}

// Function to create email body for WeddingCakeInquiry
func createWeddingCakeInquiryEmail(e *core.RecordEvent) string {
	return fmt.Sprintf(`
		<h1>Wedding Cake Inquiry Details</h1>
		<p><strong>Wedding Date:</strong> %s</p>
		<p><strong>Wedding Venue:</strong> %s</p>
		<p><strong>Number of Tiers:</strong> %d</p>
		<p><strong>Number of Guests:</strong> %s</p>
		<p><strong>Cake Flavours:</strong> %s</p>
		<p><strong>Dietary Needs:</strong> %s</p>
		<p><strong>Design Inspiration:</strong> %s</p>
		<p><strong>Email:</strong> %s</p>
		<p><strong>Phone Number:</strong> %s</p>
		<p><strong>Additional Details:</strong> %s</p>
	`,
		e.Record.GetString("wedding_date"),
		e.Record.GetString("wedding_venue"),
		e.Record.GetInt("number_of_tiers"),
		e.Record.GetString("number_of_guests"),
		e.Record.GetString("cake_flavours"),
		e.Record.GetString("dietary_needs"),
		e.Record.GetString("design_inspiration"),
		e.Record.GetString("email"),
		e.Record.GetString("phone_number"),
		e.Record.GetString("additional_details"),
	)
}

// Function to create email body for WorkshopBooking
func createWorkshopBookingEmail(e *core.RecordEvent) string {
	return fmt.Sprintf(`
		<h1>Workshop Booking Details</h1>
		<p><strong>Full Name:</strong> %s</p>
		<p><strong>Email:</strong> %s</p>
		<p><strong>Phone Number:</strong> %s</p>
		<p><strong>Number of Participants:</strong> %d</p>
		<p><strong>Preferred Date:</strong> %s</p>
		<p><strong>Type and Flavour:</strong> %s</p>
		<p><strong>Desired Outcome:</strong> %s</p>
		<p><strong>Additional Details:</strong> %s</p>
	`,
		e.Record.GetString("full_name"),
		e.Record.GetString("email"),
		e.Record.GetString("phone_number"),
		e.Record.GetInt("Number_of_participants"),
		e.Record.GetString("preferred_date"),
		e.Record.GetString("type_and_flavour"),
		e.Record.GetString("desired_outcome"),
		e.Record.GetString("additional_details"),
	)
}

func main() {
	// Initialize PocketBase app
	_ = godotenv.Load()
	app := pocketbase.New()

	//fires when TasterBoxInquiry is created
	app.OnRecordCreate("TasterBoxInquiry").BindFunc(func(e *core.RecordEvent) error {
		// Ensure next handler is processed
		if err := e.Next(); err != nil {
			return err
		}

		// Get the collection name (which triggered the event)
		collectionName := e.Record.Collection().Name

		// Prepare the email body based on the collection name
		var emailBody string
		emailBody = createTasterBoxInquiryEmail(e)

		// Construct the email message
		message := &mailer.Message{
			From: mail.Address{
				Address: e.App.Settings().Meta.SenderAddress,
				Name:    e.App.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: "teniolasobande04@gmail.com"}}, // You can replace this with the email from the record itself
			Subject: "New Record Created - " + collectionName,                // Subject is dynamic based on collection name
			HTML:    emailBody,                                               // Constructed email body based on the schema
		}

		// Send the email using PocketBase's built-in mailer
		if err := e.App.NewMailClient().Send(message); err != nil {
			return fmt.Errorf("failed to send email: %w", err)
		}

		// Continue processing the event
		return nil
	})

	//fires when bespoke cake inquiry is created
	app.OnRecordCreate("bespokecakeinquiry").BindFunc(func(e *core.RecordEvent) error {
		// Ensure next handler is processed
		if err := e.Next(); err != nil {
			return err
		}

		// Get the collection name (which triggered the event)
		collectionName := e.Record.Collection().Name

		// Prepare the email body based on the collection name
		var emailBody string
		emailBody = createBespokeCakeInquiryEmail(e)

		// Construct the email message
		message := &mailer.Message{
			From: mail.Address{
				Address: e.App.Settings().Meta.SenderAddress,
				Name:    e.App.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: "teniolasobande04@gmail.com"}}, // You can replace this with the email from the record itself
			Subject: "New Record Created - " + collectionName,                // Subject is dynamic based on collection name
			HTML:    emailBody,                                               // Constructed email body based on the schema
		}

		// Send the email using PocketBase's built-in mailer
		if err := e.App.NewMailClient().Send(message); err != nil {
			return fmt.Errorf("failed to send email: %w", err)
		}

		// Continue processing the event
		return nil
	})

	//fires when wedding cake inquiry is created
	app.OnRecordCreate("weddingcakeinquiry").BindFunc(func(e *core.RecordEvent) error {
		// Ensure next handler is processed
		if err := e.Next(); err != nil {
			return err
		}

		// Get the collection name (which triggered the event)
		collectionName := e.Record.Collection().Name

		// Prepare the email body based on the collection name
		var emailBody string
		emailBody = createWeddingCakeInquiryEmail(e)

		// Construct the email message
		message := &mailer.Message{
			From: mail.Address{
				Address: e.App.Settings().Meta.SenderAddress,
				Name:    e.App.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: "teniolasobande04@gmail.com"}}, // You can replace this with the email from the record itself
			Subject: "New Record Created - " + collectionName,                // Subject is dynamic based on collection name
			HTML:    emailBody,                                               // Constructed email body based on the schema
		}

		// Send the email using PocketBase's built-in mailer
		if err := e.App.NewMailClient().Send(message); err != nil {
			return fmt.Errorf("failed to send email: %w", err)
		}

		// Continue processing the event
		return nil
	})

	//fires when workshop booking is created
	app.OnRecordCreate("workshopbooking").BindFunc(func(e *core.RecordEvent) error {
		// Ensure next handler is processed
		if err := e.Next(); err != nil {
			return err
		}

		// Get the collection name (which triggered the event)
		collectionName := e.Record.Collection().Name

		// Prepare the email body based on the collection name
		var emailBody string
		emailBody = createWorkshopBookingEmail(e)

		// Construct the email message
		message := &mailer.Message{
			From: mail.Address{
				Address: e.App.Settings().Meta.SenderAddress,
				Name:    e.App.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: "teniolasobande04@gmail.com"}}, // You can replace this with the email from the record itself
			Subject: "New Record Created - " + collectionName,                // Subject is dynamic based on collection name
			HTML:    emailBody,                                               // Constructed email body based on the schema
		}

		// Send the email using PocketBase's built-in mailer
		if err := e.App.NewMailClient().Send(message); err != nil {
			return fmt.Errorf("failed to send email: %w", err)
		}

		// Continue processing the event
		return nil
	})

	// Payment link creation
	app.OnRecordCreate("cart").BindFunc(func(e *core.RecordEvent) error {
		// Payment method could be passed in the record, e.g., `payment_method`
		// Get the payment method from the record or the request
		paymentMethod := e.Record.GetString("payment_method")
		totalAmountStr := e.Record.GetString("total") // Assuming total_amount is a string

		// Convert total_amount to float64 (you might need to handle any possible conversion errors here)
		totalAmount, err := strconv.ParseFloat(totalAmountStr, 64)
		if err != nil {
			return fmt.Errorf("failed to parse total amount: %w", err)
		}

		// Convert the amount to cents for Stripe (1 unit = 100 cents)
		amountInCents := int64(totalAmount * 100)

		// Check if the payment method is Stripe
		if paymentMethod == "stripe" {
			// Set your Stripe API key
			stripe.Key = os.Getenv("STRIPESK") // Use your secret key here

			// Create a Stripe Checkout Session
			params := &stripe.CheckoutSessionParams{
				PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
				LineItems: []*stripe.CheckoutSessionLineItemParams{
					{
						PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
							Currency:   stripe.String(string(stripe.CurrencyUSD)),
							UnitAmount: stripe.Int64(amountInCents),
							ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
								Name: stripe.String("Cart Total"),
							},
						},
						Quantity: stripe.Int64(1),
					},
				},
				Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
				SuccessURL: stripe.String("http://localhost:8090/success"), // Replace with your success URL
				CancelURL:  stripe.String("http://localhost:8090/cancel"),  // Replace with your cancel URL
			}

			session, err := session.New(params)
			if err != nil {
				return fmt.Errorf("failed to create checkout session: %w", err)
			}

			// Send the Checkout Session URL to the frontend
			fmt.Println("Checkout Session URL:", session.URL)

			// You can then send this `session.URL` to the frontend to redirect the user to the payment page
		} else if paymentMethod == "paypal" {
			// For PayPal, you'll need to use PayPal's SDK or make API calls to their endpoints
			// You can create a PayPal order like this (example in HTTP request format):
			// PayPal's SDK is required here
			// You'll need to send the order ID or payment link to the frontend for PayPal to process
			fmt.Println("Create PayPal payment link here (not implemented in this example)")
		}

		// Continue processing the event
		return e.Next()
	})

	pocketBaseDomain := "https://orca-app-h75k3.ondigitalocean.app/"
	// Fires when a record list is fetched
	app.OnRecordsListRequest("Shop").BindFunc(func(e *core.RecordsListRequestEvent) error {
		log.Println("Records List Request triggered")
		// Iterate through each record in the list
		for _, record := range e.Records {
			// Get the record ID
			recordID := record.GetString("id")

			// Construct the URLs for the images (image1, image2, image3, image4)
			for i := 1; i <= 4; i++ {
				// Check if the image field exists and is not empty
				imageField := fmt.Sprintf("image%d", i)
				imageFile := record.GetString(imageField)
				if imageFile != "" {
					// Construct the file URL
					imageURL := fmt.Sprintf("%s/api/files/%s/%s/%s", pocketBaseDomain, record.Collection().Name, recordID, imageFile)

					// Replace the image field value with the constructed URL
					record.Set(imageField, imageURL)
					log.Println("Image field:", imageURL)
				}
			}
		}

		log.Println("Modified records:", e.Records)

		// Continue processing the event and return the modified records
		return e.Next()
	})
	// Start the app
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
