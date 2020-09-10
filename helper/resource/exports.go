package resource

import (
	"context"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/internal/configs/configschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/internal/helper/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/internal/plans/objchange"
)

func UpgradeFlatmapState(ctx context.Context, version int, m map[string]string, res *schema.Resource, meta interface{}) (map[string]interface{}, int, error) {
	return plugin.UpgradeFlatmapState(ctx, version, m, res, meta)
}

func UpgradeJSONState(ctx context.Context, version int, m map[string]interface{}, res *schema.Resource, meta interface{}) (map[string]interface{}, error) {
	return plugin.UpgradeJSONState(ctx, version, m, res, meta)
}

func RemoveAttributes(v interface{}, ty cty.Type) {
	plugin.RemoveAttributes(v, ty)
}

func NormalizeObjectFromLegacySDK(val cty.Value, schema *configschema.Block) cty.Value {
	return objchange.NormalizeObjectFromLegacySDK(val, schema)
}
