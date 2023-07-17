resource "volcengine_vpn_gateway" "foo" {
  vpc_id = "vpc-12b31m7z2kc8w17q7y2fih9ts"
  subnet_id = "subnet-12bh8g2d7fshs17q7y2nx82uk"
  bandwidth = 20
  vpn_gateway_name = "tf-test"
  description = "tf-test"
  period = 2
  project_name = "default"
}