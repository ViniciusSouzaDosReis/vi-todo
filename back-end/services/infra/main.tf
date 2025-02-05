terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "sa-east-1"
}

resource "aws_dynamodb_table" "vi_todo_dev_table" {
  name         = "vi_todo_dev"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "user_id"
  range_key    = "group_id"

  attribute {
    name = "user_id"
    type = "S"
  }

  attribute {
    name = "group_id"
    type = "S"
  }

  global_secondary_index {
    name               = "group_id_index"
    hash_key           = "group_id"
    projection_type    = "INCLUDE"
    non_key_attributes = ["group_description", "tasks", "users", "group_title", "invetes"]
  }
}
