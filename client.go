package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// API Client for the Agent REST API
type AgentAPIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// Tool represents an agent tool from the API
type Tool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ModelSettings represents model configuration options
type ModelSettings struct {
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
}

// Agent represents an agent from the API
type Agent struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	Instructions       string   `json:"instructions"`
	HandoffDescription string   `json:"handoff_description,omitempty"`
	Model              string   `json:"model,omitempty"`
	Tools              []string `json:"tools"`
}

// CreateAgentRequest represents the request to create a new agent
type CreateAgentRequest struct {
	Name               string        `json:"name"`
	Instructions       string        `json:"instructions"`
	HandoffDescription string        `json:"handoff_description,omitempty"`
	Model              string        `json:"model,omitempty"`
	ModelSettings      ModelSettings `json:"model_settings,omitempty"`
	Tools              []string      `json:"tools"`
}

// RunAgentRequest represents the request to run an agent
type RunAgentRequest struct {
	Input   string                 `json:"input"`
	Context map[string]interface{} `json:"context,omitempty"`
}

// RunAgentResponse represents the response from running an agent
type RunAgentResponse struct {
	AgentID   string                 `json:"agent_id"`
	Result    string                 `json:"result"`
	RawOutput map[string]interface{} `json:"raw_output,omitempty"`
}

// NewAgentAPIClient creates a new client for the Agent API
func NewAgentAPIClient(baseURL string) *AgentAPIClient {
	return &AgentAPIClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

// GetTools fetches all available tools from the API
func (c *AgentAPIClient) GetTools() ([]Tool, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/tools", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var tools []Tool
	err = json.Unmarshal(body, &tools)
	if err != nil {
		return nil, err
	}

	return tools, nil
}

// ListAgents fetches all agents from the API
func (c *AgentAPIClient) ListAgents() ([]Agent, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/agents", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var agents []Agent
	err = json.Unmarshal(body, &agents)
	if err != nil {
		return nil, err
	}

	return agents, nil
}

// GetAgent fetches a specific agent by ID
func (c *AgentAPIClient) GetAgent(agentID string) (*Agent, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/agents/%s", c.BaseURL, agentID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var agent Agent
	err = json.Unmarshal(body, &agent)
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

// CreateAgent creates a new agent
func (c *AgentAPIClient) CreateAgent(createReq CreateAgentRequest) (*Agent, error) {
	reqBody, err := json.Marshal(createReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/agents", c.BaseURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var agent Agent
	err = json.Unmarshal(body, &agent)
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

// DeleteAgent deletes an agent by ID
func (c *AgentAPIClient) DeleteAgent(agentID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/agents/%s", c.BaseURL, agentID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

// RunAgent runs an agent with the provided input and context
func (c *AgentAPIClient) RunAgent(agentID string, runReq RunAgentRequest) (*RunAgentResponse, error) {
	reqBody, err := json.Marshal(runReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/agents/%s/run", c.BaseURL, agentID), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response RunAgentResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CloneAgent clones an agent with updates
func (c *AgentAPIClient) CloneAgent(agentID string, updates map[string]interface{}) (*Agent, error) {
	reqBody, err := json.Marshal(updates)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/agents/%s/clone", c.BaseURL, agentID), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var agent Agent
	err = json.Unmarshal(body, &agent)
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

// doRequest performs the HTTP request and returns the response body
func (c *AgentAPIClient) doRequest(req *http.Request) ([]byte, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s (status code: %d)", body, resp.StatusCode)
	}

	return body, nil
}
