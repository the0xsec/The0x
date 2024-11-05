resource "aws_cloudformation_stack" "example" {
  name          = "example-stack"
  template_body = <<EOF
  {
    "Resources": {
      "MyBucket": {
        "Type": "AWS::S3::Bucket"
      }
    }
  }
  EOF
  # Termination protection is not set
}
