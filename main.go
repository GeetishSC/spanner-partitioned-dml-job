package main

import (
	"cloud.google.com/go/spanner"
	"context"
	"fmt"
	"google.golang.org/api/option"
	sppb "google.golang.org/genproto/googleapis/spanner/v1"
	"log"
	"os"
	"strconv"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Replace your_project_id with your actual Google Cloud Project ID
	projectID := "sharechat-production"
	instanceID := "production-spanner-6"
	databaseID := "production-db"

	// Set up the context with your Google Cloud credentials
	ctx := context.Background()
	timeout, _ := strconv.Atoi(os.Args[1])
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(int64(timeout))*time.Second)
	defer cancel()

	// Provide the path to your service account key JSON file for authentication.
	// Replace "path/to/your/service_account_key.json" with the actual path to your key file.
	// If running in a GCP environment, this step might not be necessary as the default credentials are used.
	credsOption := option.WithCredentialsFile("/home/tag-chat-sa-prod.json")

	// Create a Spanner client.
	client, err := spanner.NewClient(ctx, fmt.Sprintf("projects/%s/instances/%s/databases/%s",
		projectID, instanceID, databaseID), credsOption)
	if err != nil {
		log.Fatalf("Failed to create Spanner client: %v", err)
	}

	defer client.Close()

	stmt := spanner.Statement{SQL: "DELETE from chatMessage@{FORCE_INDEX=chatMessage_threadId_createdOn_deleted} where createdOn < 1658392721000"}
	rowCount, err := client.PartitionedUpdateWithOptions(ctx, stmt, spanner.QueryOptions{Priority: sppb.RequestOptions_PRIORITY_LOW})
	if err != nil {
		log.Println("Error executing partitioned DML", err)
	}
	fmt.Println("%d record(s) updated.\n", rowCount)
	return
}
