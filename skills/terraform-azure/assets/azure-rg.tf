resource "azurerm_resource_group" "main" {
  name     = "main-rg"
  location = "East US"
  tags = { environment = "production" }
}
