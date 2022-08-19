# Go Amadeus

Amadeus SOAP wrapper written in go.

Currently supports Availability and Description actions.

Can be used for dev purposes.

## Running locally
1. Create `config.yaml` file
2. Add `CONFIG_FILE=path/to/config.yaml` to your env
3. `go run cmd/main.go`

It will also create `logs/` dir in the your cwd to store XML request/response.
There is no way you can disable it for now.

### Using the released version
You can also grab a released version from Releases and run executable.

## Config file reference

```yaml
username: XXXXXX
password: P@$$w0rd
pseudoCityCode: NYC1Z123
agentDutyCode: SU
requestorType: U
posType: 1
url: https://nodeD1.test.webservices.amadeus.com/
wsap: 1XXXXWYZZYZY
```

Usually provided by Amadeus team.

## API ref
Running the server spins us a local sever on `0.0.0.0:3000`.
Use `PORT` env-var to set a different port

### Availability
- `POST /api/availability`

```json
{
  "checkIn": "2022-09-10",
  "checkOut": "2022-09-11",
  "adults": 1,
  "currency": "USD",
  "hotels": [
    "HYSEAHSR"
  ]
}
```

- `POST /api/description`

```json
{
  "hotels": [
    "HYSEAHSR"
  ]
}
```

## Notes
Initial conversion is done using [xmltogo tool](https://www.onlinetool.io/xmltogo/)
