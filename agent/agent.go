package agent

import "encoding/json"

// Constant represents a generated constant from the LLM.
type Constant struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Client is a stub that would interact with an LLM.
type Client struct{}

// FetchConstant simulates fetching a constant definition from an LLM.
func (c *Client) FetchConstant(prompt string) (Constant, error) {
	// This is a deterministic stub returning the same constant every call.
	data := `{"name": "GeneratedConst", "value": "generated"}`
	var resp Constant
	if err := json.Unmarshal([]byte(data), &resp); err != nil {
		return Constant{}, err
	}
	return resp, nil
}
