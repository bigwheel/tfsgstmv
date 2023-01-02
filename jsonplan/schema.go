package jsonplan



type Plan struct {
	FormatVersion    string      `json:"format_version,omitempty"`
	TerraformVersion string      `json:"terraform_version,omitempty"`
}
