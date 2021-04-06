Based on documentation https://api-docs.transferwise.com/#payouts-guide-getting-started

Should:

- Use the wise exchange API to monitor once or twice a day the GBP-EUR conversion rate

- Should check if conversion rate is  more favourable than rate on 2-3-2021 (0.8660)

- Should transfer a variable amount of cash to the destination account


Technical requirements:

- Should run 27/7 and be hosted on AWS or similar

- Should be a chron job

- Should be container 

- Should not share any credentials publicly
