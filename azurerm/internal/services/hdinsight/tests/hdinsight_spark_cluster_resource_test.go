package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

func TestAccAzureRMHDInsightSparkCluster_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_spark_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightSparkCluster_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account"),
		},
	})
}

func TestAccAzureRMHDInsightSparkCluster_gen2basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_spark_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightSparkCluster_gen2basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account"),
		},
	})
}

func TestAccAzureRMHDInsightSparkCluster_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_spark_cluster", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightSparkCluster_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.RequiresImportErrorStep(testAccAzureRMHDInsightSparkCluster_requiresImport),
		},
	})
}

func TestAccAzureRMHDInsightSparkCluster_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_spark_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightSparkCluster_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account"),
			{
				Config: testAccAzureRMHDInsightSparkCluster_updated(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account"),
		},
	})
}

func TestAccAzureRMHDInsightSparkCluster_sshKeys(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_spark_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightSparkCluster_sshKeys(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("storage_account",
				"roles.0.head_node.0.ssh_keys",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.ssh_keys",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.ssh_keys",
				"roles.0.zookeeper_node.0.vm_size"),
		},
	})
}

func TestAccAzureRMHDInsightSparkCluster_virtualNetwork(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_spark_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightSparkCluster_virtualNetwork(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account"),
		},
	})
}

func TestAccAzureRMHDInsightSparkCluster_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_spark_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightSparkCluster_complete(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account"),
		},
	})
}

func TestAccAzureRMHDInsightSparkCluster_tls(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_spark_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightSparkCluster_tls(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account"),
		},
	})
}

func testAccAzureRMHDInsightSparkCluster_basic(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightSparkCluster_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_hdinsight_spark_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"

  component_version {
    spark = "2.3"
  }

  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }

  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }

  roles {
    head_node {
      vm_size  = "Standard_A4_V2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }

    worker_node {
      vm_size               = "Standard_A4_V2"
      username              = "acctestusrvm"
      password              = "AccTestvdSC4daf986!"
      target_instance_count = 3
    }

    zookeeper_node {
      vm_size  = "Medium"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMHDInsightSparkCluster_gen2basic(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightSparkCluster_gen2template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_hdinsight_spark_cluster" "test" {
  depends_on = [azurerm_role_assignment.test]

  name                = "acctesthdi-%d"
  resource_group_name = "${azurerm_resource_group.test.name}"
  location            = "${azurerm_resource_group.test.location}"
  cluster_version     = "3.6"
  tier                = "Standard"

  component_version {
    spark = "2.3"
  }

  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }

  storage_account_gen2 {
    storage_resource_id          = azurerm_storage_account.gen2test.id
    filesystem_id                = azurerm_storage_data_lake_gen2_filesystem.gen2test.id
    managed_identity_resource_id = azurerm_user_assigned_identity.test.id
    is_default                   = true
  }

  roles {
    head_node {
      vm_size  = "Standard_A4_V2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }

    worker_node {
      vm_size               = "Standard_A4_V2"
      username              = "acctestusrvm"
      password              = "AccTestvdSC4daf986!"
      target_instance_count = 3
    }

    zookeeper_node {
      vm_size  = "Medium"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMHDInsightSparkCluster_requiresImport(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightSparkCluster_basic(data)
	return fmt.Sprintf(`
%s

resource "azurerm_hdinsight_spark_cluster" "import" {
  name                = azurerm_hdinsight_spark_cluster.test.name
  resource_group_name = azurerm_hdinsight_spark_cluster.test.resource_group_name
  location            = azurerm_hdinsight_spark_cluster.test.location
  cluster_version     = azurerm_hdinsight_spark_cluster.test.cluster_version
  tier                = azurerm_hdinsight_spark_cluster.test.tier
  dynamic "component_version" {
    for_each = azurerm_hdinsight_spark_cluster.test.component_version
    content {
      spark = component_version.value.spark
    }
  }
  dynamic "gateway" {
    for_each = azurerm_hdinsight_spark_cluster.test.gateway
    content {
      enabled  = gateway.value.enabled
      password = gateway.value.password
      username = gateway.value.username
    }
  }
  dynamic "storage_account" {
    for_each = azurerm_hdinsight_spark_cluster.test.storage_account
    content {
      is_default           = storage_account.value.is_default
      storage_account_key  = storage_account.value.storage_account_key
      storage_container_id = storage_account.value.storage_container_id
    }
  }
  dynamic "roles" {
    for_each = azurerm_hdinsight_spark_cluster.test.roles
    content {
      dynamic "head_node" {
        for_each = lookup(roles.value, "head_node", [])
        content {
          password           = lookup(head_node.value, "password", null)
          ssh_keys           = lookup(head_node.value, "ssh_keys", null)
          subnet_id          = lookup(head_node.value, "subnet_id", null)
          username           = head_node.value.username
          virtual_network_id = lookup(head_node.value, "virtual_network_id", null)
          vm_size            = head_node.value.vm_size
        }
      }

      dynamic "worker_node" {
        for_each = lookup(roles.value, "worker_node", [])
        content {
          min_instance_count    = lookup(worker_node.value, "min_instance_count", null)
          password              = lookup(worker_node.value, "password", null)
          ssh_keys              = lookup(worker_node.value, "ssh_keys", null)
          subnet_id             = lookup(worker_node.value, "subnet_id", null)
          target_instance_count = worker_node.value.target_instance_count
          username              = worker_node.value.username
          virtual_network_id    = lookup(worker_node.value, "virtual_network_id", null)
          vm_size               = worker_node.value.vm_size
        }
      }

      dynamic "zookeeper_node" {
        for_each = lookup(roles.value, "zookeeper_node", [])
        content {
          password           = lookup(zookeeper_node.value, "password", null)
          ssh_keys           = lookup(zookeeper_node.value, "ssh_keys", null)
          subnet_id          = lookup(zookeeper_node.value, "subnet_id", null)
          username           = zookeeper_node.value.username
          virtual_network_id = lookup(zookeeper_node.value, "virtual_network_id", null)
          vm_size            = zookeeper_node.value.vm_size
        }
      }
    }
  }
}
`, template)
}

func testAccAzureRMHDInsightSparkCluster_sshKeys(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightSparkCluster_template(data)
	return fmt.Sprintf(`
%s

variable "ssh_key" {
  default = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCqaZoyiz1qbdOQ8xEf6uEu1cCwYowo5FHtsBhqLoDnnp7KUTEBN+L2NxRIfQ781rxV6Iq5jSav6b2Q8z5KiseOlvKA/RF2wqU0UPYqQviQhLmW6THTpmrv/YkUCuzxDpsH7DUDhZcwySLKVVe0Qm3+5N2Ta6UYH3lsDf9R9wTP2K/+vAnflKebuypNlmocIvakFWoZda18FOmsOoIVXQ8HWFNCuw9ZCunMSN62QGamCe3dL5cXlkgHYv7ekJE15IA9aOJcM7e90oeTqo+7HTcWfdu0qQqPWY5ujyMw/llas8tsXY85LFqRnr3gJ02bAscjc477+X+j/gkpFoN1QEmt terraform@demo.tld"
}

resource "azurerm_hdinsight_spark_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"

  component_version {
    spark = "2.3"
  }

  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }

  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }

  roles {
    head_node {
      vm_size  = "Standard_A4_V2"
      username = "acctestusrvm"
      ssh_keys = [var.ssh_key]
    }

    worker_node {
      vm_size               = "Standard_A4_V2"
      username              = "acctestusrvm"
      ssh_keys              = [var.ssh_key]
      target_instance_count = 3
    }

    zookeeper_node {
      vm_size  = "Medium"
      username = "acctestusrvm"
      ssh_keys = [var.ssh_key]
    }
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMHDInsightSparkCluster_updated(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightSparkCluster_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_hdinsight_spark_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"

  component_version {
    spark = "2.3"
  }

  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }

  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }

  roles {
    head_node {
      vm_size  = "Standard_A4_V2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }

    worker_node {
      vm_size               = "Standard_A4_V2"
      username              = "acctestusrvm"
      password              = "AccTestvdSC4daf986!"
      target_instance_count = 5
    }

    zookeeper_node {
      vm_size  = "Medium"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }
  }

  tags = {
    Hello = "World"
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMHDInsightSparkCluster_virtualNetwork(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightSparkCluster_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_virtual_network" "test" {
  name                = "acctestvirtnet%d"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.0.2.0/24"
}

resource "azurerm_hdinsight_spark_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"

  component_version {
    spark = "2.3"
  }

  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }

  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }

  roles {
    head_node {
      vm_size            = "Standard_A4_V2"
      username           = "acctestusrvm"
      password           = "AccTestvdSC4daf986!"
      subnet_id          = azurerm_subnet.test.id
      virtual_network_id = azurerm_virtual_network.test.id
    }

    worker_node {
      vm_size               = "Standard_A4_V2"
      username              = "acctestusrvm"
      password              = "AccTestvdSC4daf986!"
      target_instance_count = 3
      subnet_id             = azurerm_subnet.test.id
      virtual_network_id    = azurerm_virtual_network.test.id
    }

    zookeeper_node {
      vm_size            = "Medium"
      username           = "acctestusrvm"
      password           = "AccTestvdSC4daf986!"
      subnet_id          = azurerm_subnet.test.id
      virtual_network_id = azurerm_virtual_network.test.id
    }
  }
}
`, template, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func testAccAzureRMHDInsightSparkCluster_complete(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightSparkCluster_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_virtual_network" "test" {
  name                = "acctestvirtnet%d"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.0.2.0/24"
}

resource "azurerm_hdinsight_spark_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"

  component_version {
    spark = "2.3"
  }

  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }

  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }

  roles {
    head_node {
      vm_size            = "Standard_A4_V2"
      username           = "acctestusrvm"
      password           = "AccTestvdSC4daf986!"
      subnet_id          = azurerm_subnet.test.id
      virtual_network_id = azurerm_virtual_network.test.id
    }

    worker_node {
      vm_size               = "Standard_A4_V2"
      username              = "acctestusrvm"
      password              = "AccTestvdSC4daf986!"
      target_instance_count = 3
      subnet_id             = azurerm_subnet.test.id
      virtual_network_id    = azurerm_virtual_network.test.id
    }

    zookeeper_node {
      vm_size            = "Medium"
      username           = "acctestusrvm"
      password           = "AccTestvdSC4daf986!"
      subnet_id          = azurerm_subnet.test.id
      virtual_network_id = azurerm_virtual_network.test.id
    }
  }

  tags = {
    Hello = "World"
  }
}
`, template, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func testAccAzureRMHDInsightSparkCluster_template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_storage_account" "test" {
  name                     = "acctestsa%s"
  resource_group_name      = azurerm_resource_group.test.name
  location                 = azurerm_resource_group.test.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_storage_container" "test" {
  name                  = "acctest"
  storage_account_name  = azurerm_storage_account.test.name
  container_access_type = "private"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomString)
}

func testAccAzureRMHDInsightSparkCluster_gen2template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_storage_account" "gen2test" {
  name                     = "accgen2test%s"
  resource_group_name      = azurerm_resource_group.test.name
  location                 = azurerm_resource_group.test.location
  account_kind             = "StorageV2"
  account_tier             = "Standard"
  account_replication_type = "LRS"
  is_hns_enabled           = true
}

resource "azurerm_storage_data_lake_gen2_filesystem" "gen2test" {
  name               = "acctest"
  storage_account_id = azurerm_storage_account.gen2test.id
}

resource "azurerm_user_assigned_identity" "test" {
  resource_group_name = "${azurerm_resource_group.test.name}"
  location            = "${azurerm_resource_group.test.location}"

  name = "test-identity"
}

data "azurerm_subscription" "primary" {}

resource "azurerm_role_assignment" "test" {
  scope                = "${data.azurerm_subscription.primary.id}"
  role_definition_name = "Storage Blob Data Owner"
  principal_id         = "${azurerm_user_assigned_identity.test.principal_id}"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomString)
}

func testAccAzureRMHDInsightSparkCluster_tls(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightSparkCluster_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_hdinsight_spark_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"
  tls_min_version     = "1.2"

  component_version {
    spark = "2.3"
  }

  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }

  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }

  roles {
    head_node {
      vm_size  = "Standard_A4_V2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }

    worker_node {
      vm_size               = "Standard_A4_V2"
      username              = "acctestusrvm"
      password              = "AccTestvdSC4daf986!"
      target_instance_count = 3
    }

    zookeeper_node {
      vm_size  = "Medium"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }
  }
}
`, template, data.RandomInteger)
}
