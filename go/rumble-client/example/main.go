package main

import (
	"context"
	"log"
	"net/http"

	rc "github.com/RumbleDiscovery/rumble-api/go/rumble-client"
)

func main() {

	var res *http.Response
	var err error

	ctx := context.Background()
	client := rc.NewClient(&rc.Config{})

	log.Printf("Rumble Components")
	log.Printf("=================")
	ver, res, err := client.PublicAPI.GetLatestAgentVersion(ctx)
	if err != nil {
		log.Fatalf("failed to get agent version: %s", err)
	}
	log.Printf("     Agent: %s", ver.Version)

	ver, res, err = client.PublicAPI.GetLatestScannerVersion(ctx)
	if err != nil {
		log.Fatalf("failed to get scanner version: %s", err)
	}
	log.Printf("   Scanner: %s", ver.Version)

	ver, res, err = client.PublicAPI.GetLatestPlatformVersion(ctx)
	if err != nil {
		log.Fatalf("failed to get platform version: %s", err)
	}
	log.Printf("  Platform: %s", ver.Version)

	cnf := client.GetConfig()
	if len(cnf.DefaultHeader) == 0 {
		log.Printf("Set the RUMBLE_API_KEY environment variable to test authenticated APIs")
		return
	}

	org, res, err := client.OrganizationAPI.GetOrganization(ctx)
	if err != nil {
		log.Fatalf("failed to get organization %s", err)
	}
	log.Printf("")
	log.Printf("Organization %s has %d assets", org.GetName(), int64(org.GetAssetCount()))

	_ = res

}
