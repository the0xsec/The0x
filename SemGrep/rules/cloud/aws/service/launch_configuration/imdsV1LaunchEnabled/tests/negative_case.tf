resource "aws_launch_configuration" "secured" {
  name          = "secured-launch-configuration"
  image_id      = "ami-87654321"
  instance_type = "t2.micro"

  metadata_options {
    http_tokens = "required"  # IMDSv2 is being used
  }
}
