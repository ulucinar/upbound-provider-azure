// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package alertsmanagement

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure alertsmanagement group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_monitor_smart_detector_alert_rule", func(r *config.Resource) {
		r.References["action_group.ids"] = config.Reference{
			TerraformName: "azurerm_monitor_action_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["scope_resource_ids"] = config.Reference{
			TerraformName: "azurerm_application_insights",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_monitor_alert_processing_rule_action_group", func(r *config.Resource) {
		r.References["add_action_group_ids"] = config.Reference{
			TerraformName: "azurerm_monitor_action_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["scopes"] = config.Reference{
			TerraformName: "azurerm_resource_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_monitor_alert_processing_rule_suppression", func(r *config.Resource) {
		r.References["scopes"] = config.Reference{
			TerraformName: "azurerm_resource_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
