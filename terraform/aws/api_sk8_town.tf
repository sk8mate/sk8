resource "aws_elastic_beanstalk_application" "api_sk8_town" {
  name        = "api-sk8-town"
}

resource "aws_elastic_beanstalk_environment" "api_sk8_town" {
  name                = "api-sk8-town"
  application         = aws_elastic_beanstalk_application.api_sk8_town.name
  solution_stack_name = "64bit Amazon Linux 2 v3.4.4 running Docker"

  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name = "IamInstanceProfile"
    value = "aws-elasticbeanstalk-ec2-role"
  }

  setting {
      namespace = "aws:ec2:instances"
      name = "InstanceTypes"
      value = "t2.micro"
  }

  setting {
      namespace = "aws:elasticbeanstalk:environment"
      name = "EnvironmentType"
      value = "SingleInstance"
  }

  setting {
    namespace = "aws:elasticbeanstalk:application:environment"
    name = "DOMAIN_LINK"
    value = "api.sk8.town"
  }

  setting {
    namespace = "aws:elasticbeanstalk:application:environment"
    name = "PORT"
    value = "8080"
  }

  /* SET MANUALLY THROUGH THE CONSOLE */
  # setting {
  #   namespace = "aws:elasticbeanstalk:application:environment"
  #   name = "EMAIL_LINK"
  #   value = ""
  # }
}


# IAM Policy
resource "aws_iam_user_policy_attachment" "ci_beanstalk_web_tier" {
  user       = aws_iam_user.ci.name
  policy_arn = "arn:aws:iam::aws:policy/AWSElasticBeanstalkWebTier"
}
resource "aws_iam_user_policy_attachment" "ci_beanstalk_managed_updates_customer_role_policy" {
  user       = aws_iam_user.ci.name
  policy_arn = "arn:aws:iam::aws:policy/AWSElasticBeanstalkManagedUpdatesCustomerRolePolicy"
}

# Route 53
resource "aws_route53_zone" "api_sk8_town" {
  name = "api.sk8.town"
}

data "aws_elastic_beanstalk_hosted_zone" "current" {}

resource "aws_route53_record" "api_sk8_town" {
  zone_id = aws_route53_zone.api_sk8_town.zone_id
  name    = aws_route53_zone.api_sk8_town.name
  type    = "A"

  alias {
    name                   = aws_elastic_beanstalk_environment.api_sk8_town.cname
    zone_id                = data.aws_elastic_beanstalk_hosted_zone.current.id
    evaluate_target_health = true
  }
}

resource "aws_route53_record" "api_sk8_town_ns" {
  zone_id = aws_route53_zone.sk8_town.zone_id
  name    = "api.sk8.town"
  type    = "NS"
  ttl     = "30"
  records = aws_route53_zone.api_sk8_town.name_servers
}
