resource "aws_route53_zone" "sk8_town" {
  name = "sk8.town"
}

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
    name                   = "s3-website-eu-west-1.amazonaws.com"
    zone_id                = aws_s3_bucket.web_sk8_town.hosted_zone_id
    evaluate_target_health = true
  }
}