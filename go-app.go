package main

import (
	"context"
	"fmt"
	"net/http"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
)

// Replace with your actual project ID
var projectId = "go-app-414412"

func handler(w http.ResponseWriter, r *http.Request) {
	// Access secrets and use them securely
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		fmt.Printf("Error creating Secret Manager client: %v", err)
		return
	}

	// Access GCP_PROJECT_ID
	projectIDSecretName := "projects/" + projectId + "/secrets/GCP_PROJECT_ID/versions/latest"
	accessRequest := &secretmanagerpb.AccessSecretRequest{Name: projectIDSecretName}
	projectIDResult, err := client.AccessSecret(ctx, accessRequest)
	if err != nil {
		fmt.Printf("Error accessing GCP_PROJECT_ID secret: %v", err)
		return
	}

	// Access GCP_SERVICE_ACCOUNT (consider using environment variables for this)
	serviceAccountSecretName := "projects/" + projectId + "/secrets/GCP_SERVICE_ACCOUNT/versions/latest"
	accessRequest = &secretmanagerpb.AccessSecretRequest{Name: serviceAccountSecretName}
	serviceAccountResult, err := client.AccessSecret(ctx, accessRequest)
	if err != nil {
		fmt.Printf("Error accessing GCP_SERVICE_ACCOUNT secret: %v", err)
		return
	}

	// Use retrieved secrets and projectID for your application logic
	fmt.Fprintf(w, "Hello, World from project %s!", string(projectIDResult.Payload.Data))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
