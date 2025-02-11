resource "aws_eks_cluster" "example" {
  name     = "example-cluster"
  role_arn = "arn:aws:iam::123456789012:role/eks-cluster-role"
  version  = "1.21"
  vpc_config {
    subnet_ids = ["subnet-12345678", "subnet-87654321"]
  }
}

resource "aws_cloudwatch_log_group" "example" {
  name              = "/aws/eks/example-cluster"
  retention_in_days = 7
}
