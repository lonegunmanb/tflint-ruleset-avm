package interfaces

import (
	"github.com/matt-FFFFFF/tfvarcheck/varcheck"
	"github.com/zclconf/go-cty/cty"
)

var RoleAssignmentsTypeString = `map(object({
      role_definition_id_or_name             = string
      principal_id                           = string
      description                            = optional(string, null)
      skip_service_principal_aad_check       = optional(bool, false)
      condition                              = optional(string, null)
      condition_version                      = optional(string, null)
      delegated_managed_identity_resource_id = optional(string, null)
    }))`

var roleAssignmentsType = stringToTypeConstraintWithDefaults(RoleAssignmentsTypeString)

var RoleAssignments = AvmInterface{
	Name:          "role_assignments",
	VarType:       varcheck.NewVarCheck(roleAssignmentsType, cty.EmptyObjectVal, false),
	VarTypeString: RoleAssignmentsTypeString,
	Enabled:       true,
	Link:          "https://azure.github.io/Azure-Verified-Modules/specs/shared/interfaces/#role-assignments",
}
