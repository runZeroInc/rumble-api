package main

import (
	"context"
	"log"

	rc "github.com/RumbleDiscovery/rumble-api/go/rumble-client"
	"github.com/gofrs/uuid"
)

func main() {
	var err error

	ctx := context.Background()
	client := rc.NewClient(&rc.Config{})
	cnf := client.GetConfig()
	if len(cnf.DefaultHeader) == 0 {
		log.Printf("Set the RUMBLE_API_KEY environment variable")
		return
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

	// Build a map of SiteID to TaskID
	log.Printf("siteMap: %#v", siteAgentMap)
	log.Printf("taskMap: %#v", siteTaskMap)
}
