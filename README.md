# Columba

Columba is a shipping rate calculator created for Shopify stores that provides shipping rates for
indonesian Couriers. Designed to be run on Google Cloud Function. It is used as a
[CarrierService](https://help.shopify.com/en/api/reference/shipping-and-fulfillment/carrierservice)
on Shopify.

## Features

1. Calculated shipping rates for TIKI, JNE, Pos Indonesia (powered by [RajaOngkir](https://rajaongkir.com/).
Rates are calculated based on weight, origin, and shipping location.
2. Monitoring using [Sentry](https://sentry.io)

## Usage
The simplest way to deploy your own columba instance is as a Google Cloud Function.

### Pre-requisites
1. RajaOngkir account and API key
2. Sentry account and API key (optional, only if you want monitoring)

### Deployment
This is my preferred way to deploy:

1. Fork this repo
2. [Mirror](https://cloud.google.com/source-repositories/docs/mirroring-a-github-repository) your fork on Google's Cloud Source
for use with Cloud Function.
3. [Deploy](https://cloud.google.com/functions/docs/deploying/console) to Google Cloud Function using Cloud Source Repository.
    * While creating the function, also add RajaOngkir API as environment variable named `RAJA_ONGKIR_KEY`.
    * To turn sentry on, just add your sentry DSN as envronment variable named `SENTRY_DSN`
