package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"google.golang.org/api/option"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Retrieve secrets from environment variables
	projectID := os.Getenv("GCP_PROJECT_ID")
	secretName := os.Getenv("GCP_SERVICE_ACCOUNT_NAME")

	// Access secrets directly using the secret name
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx, option.WithCredentialsFile("path/to/credentials.json"))
	if err != nil {
		fmt.Printf("Error creating Secret Manager client: %v", err)
		return
	}

	projectIDSecret, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: "projects/" + projectID + "/secrets/" + projectID + "/versions/latest",
	})
	if err != nil {
		fmt.Printf("Error accessing project ID secret: %v", err)
		return
	}

	serviceAccountSecret, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: "projects/" + projectID + "/secrets/" + secretName + "/versions/latest",
	})
	if err != nil {
		fmt.Printf("Error accessing service account secret: %v", err)
		return
	}

	// Use retrieved secrets for your application logic
	fmt.Fprintf(w, "Hello, World from project %s!", string(projectIDSecret.Payload.Data))

	payloadData := serviceAccountSecret.Payload.Data
	serviceAccountEmail := string(payloadData)
	if serviceAccountEmail == "" {
		fmt.Println("Error: Empty service account email")
		return
	}

	// Use serviceAccountEmail further...

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
