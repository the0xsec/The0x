resource "aws_db_instance" "secured" {
  identifier              = "secured-db-instance"
  engine                 = "mysql"
  instance_class         = "db.t2.micro"
  allocated_storage       = 20
  storage_type           = "gp2"
  
  publicly_accessible = false  # RDS instance is not publicly accessible
}
