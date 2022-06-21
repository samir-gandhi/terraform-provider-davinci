terraform {
  required_providers {
    davinci = {
      version = "0.0.1"
      source = "samir-gandhi/pingidentity/davinci"
    }
  }
}

# variable "customer_name" {
#   type    = string
#   default = "Vagrante espresso"
# }

# data "davinci_customers" "customer" {}

# output "all_customers" {
#   value = data.davinci_customers.all.customers
# }

# output "customers" {
#   value = data.davinci_customers.all.customers
# }

# output "customer" {
#   value = {
#     for customer in data.davinci_customers.all.customers :
#     customer.id => customer
#     if customer.name == var.customer_name
#   }
#   # value = data.davinci_customers_all.customers
# }