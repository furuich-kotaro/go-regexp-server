## what this
API to perform regexp.FindAllStringSubmatch

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -H "Authentication: xxxx" \
  -d  \'{"text": "<https://github.com/>", "regex": "github"}' https://xxxxxx/regex

{
  "Result": [
    [
      "github"
    ]
  ]
}
```
