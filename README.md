# tamil-wordle

## Run locally

```bash
go run tamil-wordle.go
```

### Generate admin key

```bash
openssl genrsa -out admin.rsa 4096
openssl rsa -in admin.rsa -pubout > admin.rsa.pub
```

### Generate Firebase service account key

To generate a private key file for your service account:

In the Firebase console, open Settings > Service Accounts.

Click Generate New Private Key, then confirm by clicking Generate Key.

Securely store the JSON file containing the key.

GOOGLE_APPLICATION_CREDENTIALS=<path-to-service-account-key-file> go run tamil-wordle.go
