output "s3_bucket_arn" {
  description = "The ARN of the created S3 bucket"
  value       = aws_s3_bucket.this.arn
}

output "s3_bucket_domain_name" {
  description = "The domain name of the created S3 bucket"
  value       = aws_s3_bucket.this.bucket_domain_name
}
