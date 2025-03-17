package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx           context.Context
	agentAPIClient *AgentAPIClient
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		agentAPIClient: NewAgentAPIClient("http://localhost:8000"),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetTools returns all available agent tools
func (a *App) GetTools() ([]Tool, error) {
	return a.agentAPIClient.GetTools()
}

// ListAgents returns all agents
func (a *App) ListAgents() ([]Agent, error) {
	return a.agentAPIClient.ListAgents()
}

// GetAgent returns a specific agent by ID
func (a *App) GetAgent(agentID string) (*Agent, error) {
	return a.agentAPIClient.GetAgent(agentID)
}

// CreateAgent creates a new agent
func (a *App) CreateAgent(name string, instructions string, handoffDescription string, model string, temperature float64, topP float64, maxTokens int, tools []string) (*Agent, error) {
	createReq := CreateAgentRequest{
		Name:               name,
		Instructions:       instructions,
		HandoffDescription: handoffDescription,
		Model:              model,
		ModelSettings: ModelSettings{
			Temperature: temperature,
			TopP:        topP,
			MaxTokens:   maxTokens,
		},
		Tools: tools,
	}
	
	return a.agentAPIClient.CreateAgent(createReq)
}

// DeleteAgent deletes an agent by ID
func (a *App) DeleteAgent(agentID string) error {
	return a.agentAPIClient.DeleteAgent(agentID)
}

// RunAgent runs an agent with input and context
func (a *App) RunAgent(agentID string, input string, context map[string]interface{}) (*RunAgentResponse, error) {
	runReq := RunAgentRequest{
		Input:   input,
		Context: context,
	}
	
	return a.agentAPIClient.RunAgent(agentID, runReq)
}

// CloneAgent clones an agent with updates
func (a *App) CloneAgent(agentID string, updates map[string]interface{}) (*Agent, error) {
	return a.agentAPIClient.CloneAgent(agentID, updates)
}

// SetAPIURL changes the API base URL
func (a *App) SetAPIURL(baseURL string) string {
	a.agentAPIClient = NewAgentAPIClient(baseURL)
	return fmt.Sprintf("API URL changed to: %s", baseURL)
}

// GetAPIURL returns the current API base URL
func (a *App) GetAPIURL() string {
	return a.agentAPIClient.BaseURL
}
