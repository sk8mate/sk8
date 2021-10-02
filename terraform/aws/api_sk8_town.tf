variable "email_link" {
  type = string
  sensitive = true
}

resource "aws_ssm_parameter" "email_link" {
  name        = "/production/email_link"
  type        = "SecureString"
  value       = var.email_link

  lifecycle {
    ignore_changes = [value]
  }
}

variable "places_tomtom_api_key" {
  type = string
  sensitive = true
}

resource "aws_ssm_parameter" "places_tomtom_api_key" {
  name        = "/production/places_tomtom_api_key"
  type        = "SecureString"
  value       = var.places_tomtom_api_key

  lifecycle {
    ignore_changes = [value]
  }
}

variable "db_host" {
  type = string
  sensitive = true
}

resource "aws_ssm_parameter" "db_host" {
  name        = "/production/db_host"
  type        = "SecureString"
  value       = var.db_host

  lifecycle {
    ignore_changes = [value]
  }
}

variable "db_port" {
  type = string
  sensitive = true
}

resource "aws_ssm_parameter" "db_port" {
  name        = "/production/db_port"
  type        = "SecureString"
  value       = var.db_port

  lifecycle {
    ignore_changes = [value]
  }
}

variable "db_user" {
  type = string
  sensitive = true
}

resource "aws_ssm_parameter" "db_user" {
  name        = "/production/db_user"
  type        = "SecureString"
  value       = var.db_user

  lifecycle {
    ignore_changes = [value]
  }
}

variable "db_password" {
  type = string
  sensitive = true
}

resource "aws_ssm_parameter" "db_password" {
  name        = "/production/db_password"
  type        = "SecureString"
  value       = var.db_password

  lifecycle {
    ignore_changes = [value]
  }
}

variable "db_name" {
  type = string
  sensitive = true
}

resource "aws_ssm_parameter" "db_name" {
  name        = "/production/db_name"
  type        = "SecureString"
  value       = var.db_name

  lifecycle {
    ignore_changes = [value]
  }
}

resource "aws_elastic_beanstalk_application" "api_sk8_town" {
  name        = "api-sk8-town"
}

resource "aws_elastic_beanstalk_environment" "api_sk8_town" {
  name                = "api-sk8-town"
  application         = aws_elastic_beanstalk_application.api_sk8_town.name
  solution_stack_name = "64bit Amazon Linux 2 v3.4.5 running Docker"

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
    value = aws_route53_zone.api_sk8_town.name
  }

  setting {
    namespace = "aws:elasticbeanstalk:application:environment"
    name = "EMAIL_LINK"
    value = aws_ssm_parameter.email_link.value
  }

  setting {
    namespace = "aws:elasticbeanstalk:application:environment"
    name = "SK8_PLACES_TOMTOM_API_KEY"
    value = aws_ssm_parameter.places_tomtom_api_key.value
  }

  setting {
    namespace = "aws:elasticbeanstalk:application:environment"
    name = "SK8_DB_HOST"
    value = aws_ssm_parameter.db_host.value
  }

  setting {
    namespace = "aws:elasticbeanstalk:application:environment"
    name = "SK8_DB_PORT"
    value = aws_ssm_parameter.db_port.value
  }

  setting {
    namespace = "aws:elasticbeanstalk:application:environment"
    name = "SK8_DB_USER"
    value = aws_ssm_parameter.db_user.value
  }

  setting {
    namespace = "aws:elasticbeanstalk:application:environment"
    name = "SK8_DB_PASSWORD"
    value = aws_ssm_parameter.db_password.value
  }

  setting {
    namespace = "aws:elasticbeanstalk:application:environment"
    name = "SK8_DB_NAME"
    value = aws_ssm_parameter.db_name.value
  }
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
