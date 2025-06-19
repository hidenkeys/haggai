Hegai Cakes -

Prerequisites
Before running the application, ensure you have the following:

Go installed

PocketBase installed and running

A Stripe account for handling payments

Installation and Setup
1. Clone the Repository
Clone this repository to your local machine:
git clone https://github.com/your-repository/haggai.git
cd haggai

3. Install Dependencies
Run the following command to install the required packages and tidy up the Go modules:
go mod tidy

4. Build the Application
To build the application, use the following command:
go build -o pocketbase1 .

5. Run the Application Locally
To run the application locally, use the command:
./pocketbase1 serve
This will start the server locally at http://localhost:8090.

5. Run the Application on a Server
For deployment on a server, bind the application to 0.0.0.0 to make it publicly accessible:
./pocketbase1 serve --http=0.0.0.0:8090

7. Access the Admin Dashboard
The admin dashboard is accessible via the following URLs:

Locally: http://localhost:8090/_/

Hosted: "http://<ip address>:8090/_/ or http://<domain>/_/"

7. Edit the Domain in main.go
In the root folder, open main.go and update the following line with your current domain:

pocketBaseDomain := "https://orca-app-h75k3.ondigitalocean.app/" // Replace this with your domain

9. Set Email Service
To configure your email service, navigate to the Settings section of the PocketBase admin dashboard and configure the Mail Settings under the Email section.
This will allow you to send emails for notifications such as order confirmations or payment receipts.
