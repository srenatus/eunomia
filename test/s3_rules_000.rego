package dev.eunomia.testing

s3_read_write_global[msg] {
  stmt := input.iam[_].Statement[_]
  stmt.Effect == "Allow"
  stmt.Action == "s3:*"
  msg := sprintf("Policy '%v' allows to read/write all S3 buckets", [stmt.Sid])
}

