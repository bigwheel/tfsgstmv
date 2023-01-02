package jsonplan

import (
	"encoding/json"
)

func Unmarshal(txt []byte) (error, plan) {
	var p plan
	err := json.Unmarshal(txt, &p)

	if err != nil {
		return err, p
	} else {
		return nil, p
	}
}

// https://github.com/hashicorp/terraform/blob/dec48a85106d5e151327f96e4418ccb1a7b80a98/internal/command/jsonplan/plan.go#L29-L46
type plan struct {
	FormatVersion   string           `json:"format_version,omitempty"`
	ResourceChanges []resourceChange `json:"resource_changes,omitempty"`
}

// https://github.com/hashicorp/terraform/blob/dec48a85106d5e151327f96e4418ccb1a7b80a98/internal/command/jsonplan/resource.go#L10-L42
type resourceChange struct {
	// Address is the absolute resource address
	Address string `json:"address,omitempty"`

	// ModuleAddress is the module portion of the above address. Omitted if the
	// instance is in the root module.
	ModuleAddress string `json:"module_address,omitempty"`

	Type         string `json:"type,omitempty"`
	Name         string `json:"name,omitempty"`
	ProviderName string `json:"provider_name,omitempty"`
	Change       change `json:"change,omitempty"`
}

// https://github.com/hashicorp/terraform/blob/dec48a85106d5e151327f96e4418ccb1a7b80a98/internal/command/jsonplan/plan.go#L62-L108
type change struct {
	// Actions are the actions that will be taken on the object selected by the
	// properties below. Valid actions values are:
	//    ["no-op"]
	//    ["create"]
	//    ["read"]
	//    ["update"]
	//    ["delete", "create"]
	//    ["create", "delete"]
	//    ["delete"]
	// The two "replace" actions are represented in this way to allow callers to
	// e.g. just scan the list for "delete" to recognize all three situations
	// where the object will be deleted, allowing for any new deletion
	// combinations that might be added in future.
	Actions []string `json:"actions,omitempty"`
}
