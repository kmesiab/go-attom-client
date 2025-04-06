# National Flood Data

Base Url

`https://api.nationalflooddata.com/v3/data`


Example:

```go

type Payload struct {
  Lat float64 `json:"lat"`
  Lng float64 `json:"lng"`
  Searchtype string `json:"searchtype"`
  Loma bool `json:"loma"`
  Elevation bool `json:"elevation"`
}

func main() {
  headers := map[string]string{
    "x-api-key": "YOUR KEY",
  }
  payload := Payload{
    Lat: 34.071783,
    Lng: -118.2596,
    Searchtype: "addresscoord",
    Loma: false,
    Elevation: false,
  }
  resp, err := http.Get("https://api.nationalflooddata.com/v3/data?" + url.Values(payload).Encode())
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(body))
}
```

Response: <a href="data.json">data.json</a>
