// Package telegram specifies the types used to interact with a telegram server
package telegram

// updatesResponce is used to parse the updates JSON data
type updatesResponce struct {
	Ok          bool     `json:"ok"`
	Result      []Update `json:"result"`
	Description string   `json:"description"`
}

// Update stores the update data
type Update struct {
	ID      int     `json:"update_id"`
	Message Message `json:"message"`
}

// Message stores the message details
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

// Chat stores the chat information (tied to a message)
type Chat struct {
	ID int `json:"id"`
}
