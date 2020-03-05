# Testing

First off, you want to install the [OPA CLI tool](https://www.openpolicyagent.org/docs/latest/#running-opa).

Then:

```
$ opa eval \
      --input s3_input_000.json \
      --data s3_rules_000.rego \
      's3_read_write_global' 
```