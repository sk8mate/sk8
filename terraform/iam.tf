resource "aws_iam_user" "ci" {
  name = "ci"
  path = "/sk8/"
}

resource "aws_iam_access_key" "ci" {
  user = aws_iam_user.ci.name
}

resource "aws_iam_user_policy" "ci" {
  name = "ci"
  user = aws_iam_user.ci.name

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "Stmt1546506260896",
      "Action": "s3:ListBucket",
      "Effect": "Allow",
      "Resource": "${aws_s3_bucket.web_sk8_town.arn}"
    },
    {
      "Sid": "Stmt1625607934826",
      "Action": [
        "s3:DeleteObject",
        "s3:GetObject",
        "s3:PutObject"
      ],
      "Effect": "Allow",
      "Resource": "${aws_s3_bucket.web_sk8_town.arn}/*"
    }
  ]
}
POLICY
}
