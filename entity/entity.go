package entity

import "time"

// InsertDataModels ...
type InsertDataModels struct {
	Path      string `json:"path" yaml:"path"`
	Value     string `json:"value" yaml:"value"`
	IsEncrypt bool   `json:"isEncrypt" yaml:"isEncrypt"`
}

// TemplatesModels ...
type TemplatesModels struct {
	Stage      string             `json:"stage" yaml:"stage"`
	Name       string             `json:"name" yaml:"name"`
	Parameters []InsertDataModels `json:"parameters" yaml:"parameters"`
}

// GenerateOutputs ...
type GenerateOutputs struct {
	Name         string
	Value        string
	Type         string
	Arn          string
	LastModified *time.Time
	Version      int64
}
