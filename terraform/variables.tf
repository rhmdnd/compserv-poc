variable "username" {
  description = "Database user name"
  type        = string
  default     = "postgres"
  sensitive   = true
}

variable "table_name" {
  description = "The name of the database table for compliance data"
  type        = string
  default     = "compliance"
}

variable "storage" {
  description = "Allocated storage in gigabytes"
  type        = number
  default     = 20
}
