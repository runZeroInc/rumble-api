package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	openapi "github.com/RumbleDiscovery/rumble-api/go"
	rc "github.com/RumbleDiscovery/rumble-api/go/rumble-client"
	"github.com/gofrs/uuid"
)

func main() {
	var err error

	var flagTest = flag.Bool("D", false, "dry run mode, do not actually schedule scans")
	var flagScanTime = flag.String("t", "00:00:00", "scan start time as hh:mm:ss in UTC")
	var flagScanFreq = flag.String("f", "daily", "scan frequency as once, hourly, daily, weekly, monthly, or continuous")
	var flagDebug = flag.Bool("d", false, "enable debug mode")
	flag.Parse()

	ctx := context.Background()
	client := rc.NewClient(&rc.Config{})
	cnf := client.GetConfig()

	if *flagDebug {
		cnf.Debug = true
	}

	if len(cnf.DefaultHeader) == 0 {
		log.Printf("Set the RUMBLE_API_KEY environment variable")
		return
	}

	fspec := "2006-01-02 15:04:05"
	cs := time.Now().UTC().Format(fspec)
	bits := strings.SplitN(cs, " ", 2)
	ns := fmt.Sprintf("%s %s", bits[0], *flagScanTime)
	ts, err := time.Parse(fspec, ns)
	if err != nil {
		log.Fatalf("could not parse scan time: %s", err)
	}

	// Build a list of connected agents
	getAgentsReq := client.OrganizationApi.GetAgents(ctx)
	agents, _, err := getAgentsReq.Execute()
	if err != nil {
		log.Fatalf("failed to get agent list %s", err)
	}

	// Build a map of SiteID to AgentID
	siteAgentMap := make(map[string]string)
	for _, agent := range agents {
		if agent.GetInactive() {
			continue
		}
		if !agent.GetConnected() {
			log.Printf("warning: agent %s [%s] is offline and will be ignored", agent.GetName(), agent.GetId())
			continue
		}
		if agent.GetSiteId() == uuid.Nil.String() {
			log.Printf("warning: agent %s [%s] is not assigned to a specific site and will be ignored", agent.GetName(), agent.GetId())
			continue
		}
		if v, ok := siteAgentMap[agent.GetSiteId()]; ok {
			log.Printf("warning: agent %s [%s] is ignored as %s is already assigned to this site", agent.GetName(), agent.GetId(), v)
			continue
		}
		siteAgentMap[agent.GetSiteId()] = agent.GetId()
	}

	// Build a list of recurring tasks
	getTasksReq := client.OrganizationApi.GetTasks(ctx)
	tasks, _, err := getTasksReq.Execute()
	if err != nil {
		log.Fatalf("failed to get task list %s", err)
	}

	// Build a map of SiteID to TaskID
	siteTaskMap := make(map[string]string)
	for _, task := range tasks {
		if !task.GetRecur() {
			continue
		}
		if task.GetType() != "scan" {
			continue
		}
		siteTaskMap[task.GetSiteId()] = task.GetId()
	}

	// Build a list of sites
	getSitesReq := client.OrganizationApi.GetSites(ctx)
	sites, _, err := getSitesReq.Execute()
	if err != nil {
		log.Fatalf("failed to get sites list %s", err)
	}

	scansCurrent := 0
	todoSites := []openapi.Site{}
	for _, site := range sites {
		site := site
		if _, ok := siteTaskMap[site.GetId()]; ok {
			scansCurrent++
			continue
		}

		if _, ok := siteAgentMap[site.GetId()]; !ok {
			log.Printf("warning: site %s [%s] has no available agent and will be ignored", site.GetName(), site.GetId())
			continue
		}

		todoSites = append(todoSites, site)
	}

	log.Printf("there are %d recurring scans across %d sites", scansCurrent, len(sites))
	if len(todoSites) == 0 {
		log.Printf("there is no work to do, exiting...")
		return
	}

	log.Printf("adding %d new %s scans starting at %s...", len(todoSites), *flagScanFreq, ts)

	for _, site := range todoSites {

		if *flagTest {
			log.Printf("dry run: we would have added scan for site %s [%s]..", site.GetName(), site.GetId())
			continue
		}

		log.Printf("adding scan for site %s [%s]..", site.GetName(), site.GetId())

		newScanReq := client.OrganizationApi.CreateScan(ctx, site.GetId())
		newScanReq = newScanReq.ScanOptions(openapi.ScanOptions{
			Targets:       "defaults",
			ScanFrequency: newString(*flagScanFreq),
			ScanStart:     newString(fmt.Sprintf("%d", ts.Unix())),
		})
		info, _, err := newScanReq.Execute()
		if err != nil {
			log.Fatalf("failed to get create scan for %s [%s]: %s", site.GetName(), site.GetId(), err)
		}
		log.Printf("scheduled a scan on site %s with ID %s", site.GetName(), info.GetId())
	}
}

func newString(s string) *string {
	cs := s
	return &cs
}
