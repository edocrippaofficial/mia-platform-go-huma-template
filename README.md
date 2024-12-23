# MiaPlatform Go Template with Huma

## Summary

This is a template to be used in Mia Platform console to create a new microservice in Go with the [Huma](https://huma.rocks/) framework, leveraging its automated openapi spec generator.

## Local Development

To develop the service locally you need:

- go 1.23+

Once you have all the dependency in place, you can launch:

```shell
go mod tidy
make test
```

This two commands, will install the dependencies and run the tests, ensuring that the service is working as expected.

To start the development server, execute this command:

```shell
make
```

## OpenAPI

Once running, head to `http://localhost:3000/documentation/` to navigate the APIs.

You can also download the openapi spec in both json and yaml format from `http://localhost:3000/documentation/openapi.[json/yaml]`
