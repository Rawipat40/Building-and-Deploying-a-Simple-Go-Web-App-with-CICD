# Go App with Google Cloud Secret Manager Integration

This Go application demonstrates integration with Google Cloud Secret Manager to securely access and use secrets in a web application deployed to Google App Engine. The application retrieves sensitive information such as the `GCP_PROJECT_ID` and `GCP_SERVICE_ACCOUNT` from Secret Manager and uses them securely in the application logic.

## Package Used

- `cloud.google.com/go/secretmanager/apiv1`: This package provides functionalities to interact with Google Cloud Secret Manager from a Go application. It allows the application to securely access and retrieve secrets stored in Secret Manager.

## How to Use

1. **Set Up Google Cloud Secret Manager:**
   - Create a new secret in Google Cloud Secret Manager for each sensitive information your application needs to access (e.g., `GCP_PROJECT_ID`, `GCP_SERVICE_ACCOUNT`).

2. **Update Project ID and Service Account in the Code:**
   - Replace the placeholder `projectId` variable in `go-app.go` with your actual Google Cloud project ID.
   - Update the `handler` function in `go-app.go` to use the retrieved secrets (`GCP_PROJECT_ID` and `GCP_SERVICE_ACCOUNT`) in your application logic.

3. **GitHub Actions Workflow:**
   - The provided GitHub Actions workflow (`ci.yml`) automates the build, test, and deployment process of the Go application to Google App Engine.
   - Make sure to add your Google Cloud service account key as a secret named `GCP_SERVICE_ACCOUNT` in your GitHub repository.

4. **Deploy to Google App Engine:**
   - The deployment step in the workflow deploys the application to Google App Engine using the secrets retrieved from Secret Manager.

## Deployment

- The application is deployed to Google App Engine using Google Cloud's infrastructure.
- The deployment is triggered automatically by pushing changes to the main branch of the GitHub repository.

## Additional Notes

- Ensure proper IAM permissions are assigned to the service account used for authentication with Google Cloud Secret Manager.
- Consider using environment variables or other secure methods to manage sensitive information in production environments.

Feel free to customize and extend the code and workflow according to your requirements.

