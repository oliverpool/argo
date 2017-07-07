package argo

// GID represents a JSON-RPC response to a request
type GID struct {
	GID string // GID of the download
	ID  string `json:"id"`
}

func (g GID) ResultAndIDPointers() (interface{}, *string) {
	return &g.GID, &g.ID
}

// Ok represents a JSON-RPC response to a request
type Ok string
