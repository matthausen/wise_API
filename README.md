## WISE API

This app is using the Wise (ex TransferWise) API to monitor exchange rate between currency and notify the user about a favourable exchange rate.
The exchange rate is compared every 12h.
Based on documentation https://api-docs.transferwise.com/#payouts-guide-getting-started


## Requirements

You should create a .env file in the root directory with your credentials:

- A Wise Token for authentication
- Your email address and password to enable notifications
- The endpoint to fetch your profile information and quotes. These can be from the sanbox environment or from the prod environment

### How to run:

Build with docker:

- `docker build -t wise .`

And run:

- `docker run -p 8080:8080 -i -t wise`

### TODO
1. Should check different scenarios for favourable exchange rate: e.g. 0% better to 5% or better
2. Should have a basic UI
3. Should take user input to confirm a money transfer
