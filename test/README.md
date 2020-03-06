# Testing

First off, you want to install the [OPA CLI tool](https://www.openpolicyagent.org/docs/latest/#running-opa).

Then:

```
$ opa eval \
      --input s3_input_000.json \
      --data s3_rules_000.rego \
      --package dev.eunomia.testing \
      's3_read_write_global[msg]' 
{
  "result": [
    {
      "expressions": [
        {
          "value": "Policy 'one fun policy' allows to read/write all S3 buckets",
          "text": "s3_read_write_global[msg]",
          "location": {
            "row": 1,
            "col": 1
          }
        }
      ],
      "bindings": {
        "msg": "Policy 'one fun policy' allows to read/write all S3 buckets"
      }
    }
  ]
}
```