# mia_template_service_name_placeholder

[![pipeline status][pipeline]][git-link]
[![coverage report][coverage]][git-link]

## Summary

%CUSTOM_PLUGIN_SERVICE_DESCRIPTION%

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

## Contributing

[git-link]: <replace with your git link>

[pipeline]: <replace with your git link>/badges/master/pipeline.svg
[coverage]: <replace with your git link>/badges/master/coverage.svg
