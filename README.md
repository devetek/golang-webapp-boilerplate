## Prerequisite

- go1.20.6

## Description

Golang Web Application Boilerplate, run command below:

1. Without hot reload

```sh
make run
```

2. With hot reload, depend on [air](https://github.com/cosmtrek/air#installation)

```sh
make run-hot
```

or use go binary command `export ENV=development && go run cmd/webapp/*.go`. Finally, open http://localhost:3000/

## Todo

Add more use case examples

- [ ] Connect with real data (REST / Graphql)
- [ ] More Rich UI, e.g: (infinite scroll, slider, transition, etc)
