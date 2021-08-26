# S3
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

# Route53
resource "aws_route53_zone" "web_sk8_town" {
  name = "web.sk8.town"
}

resource "aws_route53_record" "web_sk8_town_ns" {
  zone_id = aws_route53_zone.sk8_town.zone_id
  name    = "web.sk8.town"
  type    = "NS"
  ttl     = "30"
  records = aws_route53_zone.web_sk8_town.name_servers
}

resource "aws_route53_record" "web_sk8_town_s3" {
  zone_id = aws_route53_zone.web_sk8_town.zone_id
  name    = "web.sk8.town"
  type    = "A"

  alias {
    name                   = aws_cloudfront_distribution.web_sk8_town.domain_name
    zone_id                = aws_cloudfront_distribution.web_sk8_town.hosted_zone_id
    evaluate_target_health = true
  }
}

# Cloudfront
locals {
  web_sk8_town_origin_id = "web_sk8_town"
}

resource "aws_cloudfront_distribution" "web_sk8_town" {
  origin {
    domain_name = aws_s3_bucket.web_sk8_town.bucket_regional_domain_name
    origin_id   = local.web_sk8_town_origin_id
  }

  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = "index.html"

  aliases = ["web.sk8.town"]

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD"]
    cached_methods   = ["GET", "HEAD"]
    compress         = true
    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }
    target_origin_id = local.web_sk8_town_origin_id
    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  price_class = "PriceClass_100"

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    acm_certificate_arn = aws_acm_certificate.sk8_town.arn
    ssl_support_method = "sni-only"
  }

  custom_error_response {
    error_code = 404
    response_code = 200
    response_page_path = "/index.html"
    error_caching_min_ttl = 300
  }
}

# IAM Policy
resource "aws_iam_user_policy" "web_sk8_town" {
  name = "web_sk8_town"
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
    },
    {
      "Sid": "Stmt1625607934827",
      "Action": [
        "cloudfront:CreateInvalidation"
      ],
      "Effect": "Allow",
      "Resource": "${aws_cloudfront_distribution.web_sk8_town.arn}"
    }
  ]
}
POLICY
}
