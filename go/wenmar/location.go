package wenmar

// ForLocation returns a scoped client that injects the X-Wenmar-Location
// header on every request. The parent client is not mutated. This is defined
// on *Client and returns a *Client so the full generated surface is available
// on the scoped client.
