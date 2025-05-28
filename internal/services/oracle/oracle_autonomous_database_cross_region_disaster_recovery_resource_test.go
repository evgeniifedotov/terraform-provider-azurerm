// Copyright © 2025, Oracle and/or its affiliates. All rights reserved

package oracle_test

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/oracledatabase/2025-03-01/autonomousdatabases"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/oracle"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"testing"
)

type AdbsCrossRegionDisasterRecoveryResource struct{}

func (a AdbsCrossRegionDisasterRecoveryResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := autonomousdatabases.ParseAutonomousDatabaseID(state.ID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Oracle.OracleClient25.AutonomousDatabases.Get(ctx, *id)
	if err != nil {
		return nil, fmt.Errorf("retrieving adbs_crdr %s: %+v", id, err)
	}
	return pointer.To(resp.Model != nil), nil
}

func TestAdbsCrossRegionDisasterRecoveryResource_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, oracle.AutonomousDatabaseCrossRegionDisasterRecoveryResource{}.ResourceType(), "test")
	r := AdbsCrossRegionDisasterRecoveryResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func TestAdbsCrossRegionDisasterRecoveryResource_update(t *testing.T) {
	data := acceptance.BuildTestData(t, oracle.AutonomousDatabaseCrossRegionDisasterRecoveryResource{}.ResourceType(), "test")
	r := AdbsCrossRegionDisasterRecoveryResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.update(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func TestAdbsCrossRegionDisasterRecoveryResource_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, oracle.AutonomousDatabaseCrossRegionDisasterRecoveryResource{}.ResourceType(), "test")
	r := AdbsCrossRegionDisasterRecoveryResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func (a AdbsCrossRegionDisasterRecoveryResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`

%s

provider "azurerm" {
  features {}
}

resource "azurerm_oracle_autonomous_database_cross_region_disaster_recovery" "test" {
  
  name = "CRDRInstanceForEfOriginTest"

  display_name                     = "CRDRInstanceForEfOriginTest"
  resource_group_name              = "EF-Test-CRDR"
  location                         = "%[4]s"
  remote_disaster_recovery_type    = "Adg"
  database_type                    = "CrossRegionDisasterRecovery"
  source                           = "CrossRegionDisasterRecovery"
  source_id      				   = "/subscriptions/4aa7be2d-ffd6-4657-828b-31ca25e39985/resourceGroups/EF-Test-CRDR/providers/Oracle.Database/autonomousDatabases/EfTestOriginal"
  replicate_automatic_backups	   = true
  source_location				   = "%[3]s"
  license_model                    = "LicenseIncluded"
  backup_retention_period_in_days  = 12
  auto_scaling_enabled             = false
  auto_scaling_for_storage_enabled = false
  mtls_connection_required         = false
  data_storage_size_in_tbs         = 1
  compute_model                    = "ECPU"
  compute_count                    = 3
  db_workload                      = "DW"
  admin_password                   = "TestPass#2024#"
  db_version                       = "19c"
  character_set                    = "AL32UTF8"
  national_character_set           = "AL16UTF16"
  subnet_id                        = azurerm_subnet.test.id
  virtual_network_id               = azurerm_virtual_network.test.id
  customer_contacts                = ["test@test.com"]
}
`, a.template(data), "EF-Test-CRDR", "eastus", "germanywestcentral")
}

func (a AdbsCrossRegionDisasterRecoveryResource) update(data acceptance.TestData) string {
	return fmt.Sprintf(`

%s

provider "azurerm" {
  features {}
}

resource "azurerm_oracle_autonomous_database_cross_region_disaster_recovery" "test" {


  name = "CRDRInstanceForEfOriginTest"

  display_name                     = "CRDRInstanceForEfOriginTest"
  resource_group_name              = "EF-Test-CRDR"
  location                         = "%[4]s"
  remote_disaster_recovery_type    = "BackupBased"
  database_type                    = "CrossRegionDisasterRecovery"
  source                           = "BackupFromTimestamp"
  source_id      				   = "/subscriptions/4aa7be2d-ffd6-4657-828b-31ca25e39985/resourceGroups/EF-Test-CRDR/providers/Oracle.Database/autonomousDatabases/EfTestOriginal"
  replicate_automatic_backups	   = false
  source_location				   = "%[3]s"
  compute_model                    = "ECPU"
  compute_count                    = 3
  license_model                    = "LicenseIncluded"
  backup_retention_period_in_days  = 12
  auto_scaling_enabled             = false
  auto_scaling_for_storage_enabled = false
  mtls_connection_required         = false
  data_storage_size_in_tbs         = 1
  db_workload                      = "DW"
  admin_password                   = "TestPass#2024#"
  db_version                       = "19c"
  character_set                    = "AL32UTF8"
  national_character_set           = "AL16UTF16"
  subnet_id                        = azurerm_subnet.test.id
  virtual_network_id               = azurerm_virtual_network.test.id
}
`, a.template(data), "EF-Test-CRDR", "eastus", "germanywestcentral")
}

func (a AdbsCrossRegionDisasterRecoveryResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_oracle_autonomous_database_cross_region_disaster_recovery" "import" {
  name                             = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.name
  display_name                     = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.display_name
  resource_group_name              = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.resource_group_name
  location                         = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.location
  remote_disaster_recovery_type    = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.remote_disaster_recovery_type
  database_type                    = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.database_type
  source                           = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.source
  source_id      				   = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.source_id
  replicate_automatic_backups	   = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.replicate_automatic_backups
  source_ocid					   = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.source_ocid
  source_location				   = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.source_location
  compute_model                    = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.compute_model
  compute_count                    = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.compute_count
  license_model                    = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.license_model
  backup_retention_period_in_days  = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.backup_retention_period_in_days
  auto_scaling_enabled             = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.auto_scaling_enabled
  auto_scaling_for_storage_enabled = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.auto_scaling_for_storage_enabled
  mtls_connection_required         = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.mtls_connection_required
  data_storage_size_in_tbs         = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.data_storage_size_in_tbs
  db_workload                      = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.db_workload
  admin_password                   = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.admin_password
  db_version                       = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.db_version
  character_set                    = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.character_set
  national_character_set           = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.national_character_set
  subnet_id                        = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.subnet_id
  virtual_network_id               = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.virtual_network_id
}
`, a.complete(data))
}

func (a AdbsCrossRegionDisasterRecoveryResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`

data "azurerm_client_config" "current" {}

resource "azurerm_resource_group" "test" {
  name     = "%[1]s"
  location = "%[3]s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctest%[1]s_vnet"
  address_space       = ["10.0.0.0/16"]
  location            = "%[3]s"
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "test" {
  name                 = "eacctest%[1]s"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefixes     = ["10.0.1.0/24"]

  delegation {
    name = "delegation"

    service_delegation {
      actions = [
        "Microsoft.Network/networkinterfaces/*",
        "Microsoft.Network/virtualNetworks/subnets/join/action",
      ]
      name = "Oracle.Database/networkAttachments"
    }
  }
}

`, "EF-Test-CRDR", "eastus", "germanywestcentral")
}
