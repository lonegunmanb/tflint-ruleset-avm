package waf

import "github.com/Azure/tflint-ruleset-avm/attrvalue"

var AzurermStorageAccountAccountReplicationType = func() *attrvalue.AttrStringValueRule {
	return attrvalue.NewAttrStringValueRule(
		"azurerm_storage_account",
		"account_replication_type",
		[]string{"ZRS"},
	)
}
