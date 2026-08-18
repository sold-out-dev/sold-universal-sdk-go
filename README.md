# Universal SDK — Go

Go SDK for **Akamai Bot Manager**: sensor data generation, pixel and SBSD challenges, sec-cpt solving and `_abck` cookie validation.

## 🔑 Getting API Access

Before using this SDK you need an API key:

1. Go to [sold-out.dev](https://sold-out.dev/), create an account and link your API key.
2. You can even ask for a **free trial** on our [Discord](https://discord.gg/NMRfsunxZ4).

## 📦 Installation

```bash
go get github.com/sold-out-dev/sold-universal-sdk-go
```

## 🔧 Basic Usage

```go
package main

import (
	"context"
	"fmt"

	universalsdk "github.com/sold-out-dev/sold-universal-sdk-go"
)

func main() {
	session := universalsdk.NewSession("your-api-key")

	sensorData, sensorContext, err := session.GenerateSensorData(context.Background(), &universalsdk.SensorInput{
		// Configure sensor parameters
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(sensorData, sensorContext)
}
```

### Session options

```go
session := universalsdk.NewSession("your-api-key").
	WithClient(customHTTPClient).
	WithCompression(universalsdk.CompressionGzip)
```

For drop-in compatibility with the upstream SDK, `WithJwtKey` and `WithOrganization` still
exist but are no-ops: this API authenticates with the API key alone.

### Custom base URL

The API base url defaults to `universalsdk.DefaultBaseUrl` (`https://sold-out.dev`). Override it with `WithBaseUrl`:

```go
session := universalsdk.NewSession("your-api-key").
	WithBaseUrl("https://akamai.example.com") // trailing slashes are stripped
```

Setting it to an empty string falls back to the default.

## 🛡️ Akamai Bot Manager

### Generating Sensor Data

```go
sensorData, sensorContext, err := session.GenerateSensorData(ctx, &universalsdk.SensorInput{
	// Configure sensor parameters
})
```

### Parsing Script Path

```go
scriptPath, err := akamai.ParseScriptPath(reader)
```


### Handling Sec-Cpt Challenges

- `ParseSecCptChallenge`: parse sec-cpt challenges from HTML
- `ParseSecCptChallengeFromJson`: parse from JSON responses
- `GenerateSecCptPayload`: generate challenge response payloads
- `Sleep` / `SleepWithContext`: handle challenge timing requirements

### Cookie Validation

- `IsCookieValid`: check `_abck` cookie validity for a request count
- `IsCookieInvalidated`: determine if more sensors are needed


### Pixel Challenge Solving

```go
pixelData, err := session.GeneratePixelData(ctx, &universalsdk.PixelInput{
	// Pixel challenge parameters
})
```

Pixel parsing helpers: `ParsePixelHtmlVar`, `ParsePixelScriptURL`, `ParsePixelScriptVar`.

### SBSD Challenge Solving

```go
sbsdData, err := session.GenerateSbsdData(ctx, &universalsdk.SbsdInput{
	// SBSD parameters
})
```

## 📄 License

MIT — see [LICENSE](LICENSE).

---

Fork of the Hyper Solutions SDK, trimmed down to the Akamai part and pointed at our own API.
