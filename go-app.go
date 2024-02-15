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
	// Access secrets using a client configured with application default credentials
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx, option.WithCredentialsFile("path/to/credentials.json"))
	if err != nil {
		fmt.Printf("Error creating Secret Manager client: %v", err)
		return
	}

	// Access secrets directly without redundant environment variable calls
	projectIDSecret, err := client.AccessSecretVersion(ctx, &secretmanager.AccessSecretVersionRequest{
		Name: "projects/" + os.Getenv("GCP_PROJECT_ID") + "/secrets/GCP_PROJECT_ID/versions/latest",
	})
	if err != nil {
		fmt.Printf("Error accessing GCP_PROJECT_ID secret: %v", err)
		return
	}
	serviceAccountSecret, err := client.AccessSecretVersion(ctx, &secretmanager.AccessSecretVersionRequest{
		Name: "projects/" + os.Getenv("GCP_PROJECT_ID") + "/secrets/" + os.Getenv("GCP_SERVICE_ACCOUNT") + "/versions/latest",
	})
	if err != nil {
		fmt.Printf("Error accessing GCP_SERVICE_ACCOUNT secret: %v", err)
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
