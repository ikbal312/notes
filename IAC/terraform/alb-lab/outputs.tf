output "alb_dns_name" {
  description = "the dns name of the load balancer"
  value       = aws_lb.my_alb.dns_name
}
