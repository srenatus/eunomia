package dev.eunomia.testing

s3_read_write_global[msg] {
  some i, j  
  sid := input.iam[i].Statement[j].Sid  
  effect := input.iam[i].Statement[j].Effect  
  action := input.iam[i].Statement[j].Action
  effect == "Allow"
  action == "s3:*"
  msg := sprintf("Policy %v allows to read/write", sid)
}

