
provider "aws" {
  region = "us-east-1"
}

resource "aws_db_instance" "test_db" {
  identifier           = "test-db-instance"
  engine               = "mysql"
  engine_version       = "5.7"
  instance_class       = "db.t2.micro"
  allocated_storage    = 10
  name                 = "test_db"
  username             = "rdstest"
  password             = "password"
  parameter_group_name = "default.mysql5.7"
  skip_final_snapshot  = true
}
