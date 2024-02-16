package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"google.golang.org/api/option"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Retrieve secrets using environment variables injected by Cloud Code
	projectID := os.Getenv("PROJECT_ID")
	secretName := os.Getenv("GCP_SERVICE_ACCOUNT_NAME")

	// Access secrets directly using the secret name
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx, option.WithCredentialsFile("path/to/credentials.json"))
	if err != nil {
		fmt.Printf("Error creating Secret Manager client: %v", err)
		return
	}

	projectIDSecret, err := client.AccessSecretVersion(ctx, &secretmanager.AccessSecretVersionRequest{
		Name: "projects/" + projectID + "/secrets/" + projectID + "/versions/latest",
	})
	if err != nil {
		fmt.Printf("Error accessing project ID secret: %v", err)
		return
	}

	serviceAccountSecret, err := client.AccessSecretVersion(ctx, &secretmanager.AccessSecretVersionRequest{
		Name: "projects/" + projectID + "/secrets/" + secretName + "/versions/latest",
	})
	if err != nil {
		fmt.Printf("Error accessing service account secret: %v", err)
		return
	}

	// Use retrieved secrets for your application logic
	fmt.Fprintf(w, "Hello, World from project %s!", string(projectIDSecret.Payload.Data))

	// Access service account information if needed
	serviceAccountEmail := string(serviceAccountSecret.Payload.Data)
	// ...
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
