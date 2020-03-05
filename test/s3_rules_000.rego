package dev.eunomia.testing

default s3_read_write_global = false

s3_read_write_global {
  action := input.iam[_].Statement[_].Action
  glob.match("s3:*", [], action)
}