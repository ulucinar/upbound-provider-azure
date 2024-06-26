// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package healthcareapis

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure healthcareapis resource group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_healthcare_medtech_service", func(r *config.Resource) {
		r.References["workspace_id"] = config.Reference{
			TerraformName: "azurerm_healthcare_workspace",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["eventhub_namespace_name"] = config.Reference{
			TerraformName: "azurerm_eventhub_namespace",
		}
		r.References["eventhub_name"] = config.Reference{
			TerraformName: "azurerm_eventhub",
		}
		r.References["eventhub_consumer_group_name"] = config.Reference{
			TerraformName: "azurerm_eventhub_consumer_group",
		}
	})

	p.AddResourceConfigurator("azurerm_healthcare_medtech_service_fhir_destination", func(r *config.Resource) {
		r.References["medtech_service_id"] = config.Reference{
			TerraformName: "azurerm_healthcare_medtech_service",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["destination_fhir_service_id"] = config.Reference{
			TerraformName: "azurerm_healthcare_fhir_service",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
