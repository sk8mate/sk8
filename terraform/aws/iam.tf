resource "aws_iam_user" "ci" {
  name = "ci"
  path = "/sk8/"
}

resource "aws_iam_access_key" "ci" {
  user = aws_iam_user.ci.name
}
