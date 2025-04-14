package methods

type GetUpdates struct {
	Offset         int      `json:"offset,omitempty"`          // First update to return (use offset > last update_id)
	Limit          int      `json:"limit,omitempty"`           // Max updates to retrieve (1-100, default 100)
	Timeout        int      `json:"timeout,omitempty"`         // Long polling timeout (seconds)
	AllowedUpdates []string `json:"allowed_updates,omitempty"` // List of update types to receive
}
