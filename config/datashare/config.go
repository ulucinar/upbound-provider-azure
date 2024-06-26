// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package datashare

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures datashare group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_data_share_account", func(r *config.Resource) {
		r.Kind = "Account"
	})

	p.AddResourceConfigurator("azurerm_data_share", func(r *config.Resource) {
		r.Kind = "DataShare"
		r.References["account_id"] = config.Reference{
			TerraformName: "azurerm_data_share_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_data_share_dataset_blob_storage", func(r *config.Resource) {
		r.References["data_share_id"] = config.Reference{
			TerraformName: "azurerm_data_share",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["container_name"] = config.Reference{
			TerraformName: "azurerm_storage_container",
		}
		r.References["storage_account.name"] = config.Reference{
			TerraformName: "azurerm_storage_account",
		}
		r.References["storage_account.resource_group_name"] = config.Reference{
			TerraformName: "azurerm_resource_group",
		}
	})

	p.AddResourceConfigurator("azurerm_data_share_dataset_data_lake_gen2", func(r *config.Resource) {
		r.References["share_id"] = config.Reference{
			TerraformName: "azurerm_data_share",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["storage_account_id"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["file_system_name"] = config.Reference{
			TerraformName: "azurerm_storage_data_lake_gen2_filesystem",
		}
	})

	p.AddResourceConfigurator("azurerm_data_share_dataset_kusto_cluster", func(r *config.Resource) {
		r.References["share_id"] = config.Reference{
			TerraformName: "azurerm_data_share",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["kusto_cluster_id"] = config.Reference{
			TerraformName: "azurerm_kusto_cluster",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_data_share_dataset_kusto_database", func(r *config.Resource) {
		r.References["share_id"] = config.Reference{
			TerraformName: "azurerm_data_share",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["kusto_database_id"] = config.Reference{
			TerraformName: "azurerm_kusto_database",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
