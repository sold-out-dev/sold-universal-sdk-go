package universalsdk

type SensorInput struct {
	// Abck is the _abck cookie retrieved from your cookiejar.
	// Make sure you always retrieve it fresh from the cookiejar as it gets updated while making requests.
	Abck string `json:"abck"`

	// Bmsz is the bm_sz cookie retrieved from your cookiejar, make sure you always retrieve it fresh from the cookiejar.
	Bmsz string `json:"bmsz"`

	// Version is the akamai version, this will usually be "2"
	Version string `json:"version"`

	// PageUrl is the page url that loaded the akamai script, it is the same URL as the referer header on the sensor posts
	PageUrl string `json:"pageUrl"`

	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent string `json:"userAgent"`

	// Script is mutually exclusive with [SensorInput.Context], the first sensor request should include the script field.
	// 	Subsequent request should only include the Context.
	Script string `json:"script"`

	// ScriptUrl is the full URL where you are posting sensor data to
	ScriptUrl string `json:"scriptUrl"`

	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip"`
	Context        string `json:"context"`
}

type PixelInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent      string `json:"userAgent"`
	HTMLVar        string `json:"htmlVar"`
	ScriptVar      string `json:"scriptVar"`
	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip"`
}

type SbsdInput struct {
	Index int `json:"index"`
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent      string `json:"userAgent"`
	Uuid           string `json:"uuid"`
	PageUrl        string `json:"pageUrl"`
	OCookie        string `json:"o"`
	Script         string `json:"script"`
	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip"`
}

type DynamicInput struct {
	// Script is the akamai script's contents.
	Script string `json:"script"`
}

type apiResponse struct {
	Payload  string   `json:"payload"`
	Swhanedl string   `json:"swhanedl,omitempty"`
	Context  string   `json:"context,omitempty"`
	TimeZone string   `json:"timeZone,omitempty"`
	ClientId string   `json:"clientId,omitempty"`
	Headers  *Headers `json:"headers"`
	Error    string   `json:"error"`
}

type Headers struct {
	DeviceMemory    string `json:"sec-ch-device-memory"`
	Mobile          string `json:"sec-ch-ua-mobile"`
	Arch            string `json:"sec-ch-ua-arch"`
	Platform        string `json:"sec-ch-ua-platform"`
	Model           string `json:"sec-ch-ua-model"`
	FullVersionList string `json:"sec-ch-ua-full-version-list"`
}
