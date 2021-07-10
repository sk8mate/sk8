resource "aws_s3_bucket" "web_sk8_town" {
  bucket = "web.sk8.town"
  acl    = "public-read"
  policy = <<POLICY
{
  "Id": "Policy1625590597815",
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "Stmt1625590591483",
      "Action": [
        "s3:GetObject"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::web.sk8.town/*",
      "Principal": "*"
    }
  ]
}
POLICY

  website {
    index_document = "index.html"
  }
}
