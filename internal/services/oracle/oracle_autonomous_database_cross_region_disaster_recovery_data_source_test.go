// Copyright © 2025, Oracle and/or its affiliates. All rights reserved

package oracle_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/oracle"
)

type AutonomousDatabaseCrossRegionDisasterRecoveryDataSource struct{}

func TestAdbsCrossRegionDisasterRecoveryDataSource_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, oracle.AutonomousDatabaseCrossRegionDisasterRecoveryDataSource{}.ResourceType(), "test")
	r := AutonomousDatabaseCrossRegionDisasterRecoveryDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(

				check.That(data.ResourceName).Key("remote_disaster_recovery_type").Exists(),
				check.That(data.ResourceName).Key("database_type").Exists(),
				check.That(data.ResourceName).Key("source").Exists(),
				check.That(data.ResourceName).Key("source_id").Exists(),
				check.That(data.ResourceName).Key("location").Exists(),
				check.That(data.ResourceName).Key("name").Exists(),
				check.That(data.ResourceName).Key("resource_group_name").Exists(),
				check.That(data.ResourceName).Key("display_name").Exists(),
			),
		},
	})
}

func (d AutonomousDatabaseCrossRegionDisasterRecoveryDataSource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

data "azurerm_oracle_autonomous_database_cross_region_disaster_recovery" "test" {
  name                = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.name
  resource_group_name = azurerm_oracle_autonomous_database_cross_region_disaster_recovery.test.resource_group_name
}
`, AdbsCrossRegionDisasterRecoveryResource{}.complete(data))
}
