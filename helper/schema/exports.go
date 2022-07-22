package schema

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/internal/configs/configschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/internal/configs/hcl2shim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/internal/plans/objchange"
)

func NormalizeObjectFromLegacySDK(val cty.Value, schema *configschema.Block) cty.Value {
	return objchange.NormalizeObjectFromLegacySDK(val, schema)
}

func HCL2ValueFromConfigValue(v interface{}) cty.Value {
	return hcl2shim.HCL2ValueFromConfigValue(v)
}
