# CI/CD Pipeline with GitHub Actions workflow

This repository contains a Go application (`go-app.go`) and a GitHub Actions workflow (`ci.yml`) for setting up a secure CI/CD pipeline using Google Cloud services like Secret Manager and Google Cloud SDK.

## Go Application: `go-app.go`

The `go-app.go` file contains a simple Go web server that retrieves secrets from Google Secret Manager and uses them in the application logic. The main function starts an HTTP server listening on port 8080, with a handler function that retrieves secrets and responds with a "Hello, World!" message.

## GitHub Actions Workflow: `ci.yml`

The `ci.yml` GitHub Actions workflow sets up a CI/CD pipeline with the following jobs:

### 1. Build and Test
- **Trigger:** Triggered on push events to the `main` branch.
- **Steps:**
  1. Checkout the repository.
  2. Set up Go environment.
  3. Build the Go application.
  4. Run tests for the application.

### 2. Deploy App Engine
- **Trigger:** Depends on the success of the "Build and Test" job.
- **Steps:**
  1. Set the Google Cloud project using `gcloud config`.
  2. Access the service account key securely from Secret Manager and save it as `key.json`.
  3. Authenticate with Google Cloud SDK using the service account key.
  
## Usage
To use this CI/CD pipeline for your project, follow these steps:
1. Update the `go-app.go` file with your application logic.
2. Configure the `ci.yml` file as needed, especially the environment variables for Google Cloud credentials.
3. Commit and push your changes to the `main` branch.

## Security Considerations
- Ensure that sensitive information such as service account names and credentials are securely stored and managed, preferably using tools like Google Secret Manager.
- Use least privilege access for service accounts and APIs to minimize the risk of unauthorized access.
