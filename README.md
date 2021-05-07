# Yuriko

A simple search server for Ghost, using Typesense engine.

## Installation.

```bash
go get -v github.com/Damillora/yuriko
```

## Configuration

Yuriko is configured using environment variables.

* `TYPESENSE_API_URL`: URL of the Typesense engine
* `TYPESENSE_API_KEY`: API key of Typesense engine
* `WEBHOOK_API_KEY`: API key that will be use to authenticate webhooks

## Usage

Yuriko accepts publish-time webhooks from Ghost with the endpoint `{YURIKO_URL}/api/webhook?key={WEBHOOK_API_KEY}`

## Contributing

Yuriko is still in an early stage, however pull requests are welcome! 

## License
[MIT](https://choosealicense.com/licenses/mit/)