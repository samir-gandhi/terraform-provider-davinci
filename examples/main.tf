terraform {
  required_providers {
    davinci = {
      version = "0.0.1"
      source = "samir-gandhi/pingidentity/davinci"
    }
  }
}

provider "davinci" {
  # username = ""
  # password = ""
}

data "davinci_customers" "customers" {}

output "customers" {
  value = data.davinci_customers.customers
  
}

# module "tdf" {
#   source = "./customers"
#   customer_name = "tempdvflows"
# }

# output "psl" {
#   value = module.tdf.customer
# }