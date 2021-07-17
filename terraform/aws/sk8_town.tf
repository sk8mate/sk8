# S3
resource "aws_s3_bucket" "sk8_town" {
  bucket = "sk8.town"
  acl    = "public-read"

  website {
    index_document = "index.html"
    routing_rules = <<JSON
[
  {
    "RoutingRuleCondition": {
      "KeyPrefixEquals": "/"
    },
    "Redirect": {
      "HostName": "github.com",
      "ReplaceKeyWith": "sk8mate/sk8/"
    }
  }
]
JSON
  }
}

# Route53
resource "aws_route53_zone" "sk8_town" {
  name = "sk8.town"
}

resource "aws_route53_record" "sk8_town_s3" {
  zone_id = aws_route53_zone.sk8_town.zone_id
  name    = "sk8.town"
  type    = "A"

  alias {
    name                   = aws_cloudfront_distribution.sk8_town.domain_name
    zone_id                = aws_cloudfront_distribution.sk8_town.hosted_zone_id
    evaluate_target_health = true
  }
}

# Certficiate
resource "aws_acm_certificate" "sk8_town" {
  provider                  = aws.us_east_1
  domain_name               = "sk8.town"
  subject_alternative_names = ["*.sk8.town"]
  validation_method         = "DNS"
  lifecycle {
    create_before_destroy = true
  }
}

locals {
  domain_to_zone_id = {
    "web.sk8.town" = aws_route53_zone.web_sk8_town.zone_id
  }
}

resource "aws_route53_record" "sk8_town" {
  for_each = {
    for dvo in aws_acm_certificate.sk8_town.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
      zone_id = lookup(local.domain_to_zone_id, dvo.domain_name, aws_route53_zone.sk8_town.zone_id)
    }
  }

  allow_overwrite = true
  name            = each.value.name
  records         = [each.value.record]
  ttl             = 60
  type            = each.value.type
  zone_id         = each.value.zone_id
}

resource "aws_acm_certificate_validation" "sk8_town" {
  provider                = aws.us_east_1
  certificate_arn         = aws_acm_certificate.sk8_town.arn
  validation_record_fqdns = [for record in aws_route53_record.sk8_town : record.fqdn]
}

# Cloudfront
locals {
  sk8_town_origin_id = "sk8_town"
}

resource "aws_cloudfront_distribution" "sk8_town" {
  origin {
    domain_name = aws_s3_bucket.sk8_town.website_endpoint
    origin_id   = local.sk8_town_origin_id
    custom_origin_config {
      http_port              = "80"
      https_port             = "443"
      origin_protocol_policy = "http-only"
      origin_ssl_protocols   = ["TLSv1", "TLSv1.1", "TLSv1.2"]
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  aliases = ["sk8.town"]

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
    target_origin_id = local.sk8_town_origin_id
    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  price_class = "PriceClass_100"

  restrictions {
    geo_restriction {
      restriction_type = "whitelist"
      locations        = ["PL"]
    }
  }

  viewer_certificate {
    acm_certificate_arn = aws_acm_certificate.sk8_town.arn
    ssl_support_method = "sni-only"
  }
}