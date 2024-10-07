package models

// Device represents a device in the system
type Device struct {
	ID          int64  `bun:"id,pk,autoincrement"`
	Brand       string `bun:"brand,notnull"`
	Model       string `bun:"model,notnull"`
	ReleaseYear int    `bun:"release_year,notnull"`
}

// DeviceIssue represents an issue that can occur with a device
type DeviceIssue struct {
	ID               int64  `bun:"id,pk,autoincrement" json:"id"`
	IssueDescription string `bun:"issue_description,notnull,unique" json:"description"`
}

type DeviceModel struct {
	ID    int64  `json:"id"`
	Model string `json:"model"`
}

type BrandModel struct {
	Brand  string        `json:"brand"`
	Models []DeviceModel `json:"models"`
}
