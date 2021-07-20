# slow-http
Simulate slow http responses

## Usage

In docker compose add this

```yaml
version: "3.8"

services:
  slow-http:
    image: kishanb/slow-http:1.0.0
    ports:
      - "8080:8080"
    environment:
      DELAY_RESPONSE: 20s
```

You can configure DELAY_RESPONSE to your desired delay value