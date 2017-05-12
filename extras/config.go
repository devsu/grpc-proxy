package extras

import (
  "encoding/json"
  "fmt"
  "os"
  "io/ioutil"
  "strings"
)

type Config struct {
  Listen string       `json:"listen"`
  Verbose bool        `json:"verbose"`
  CertFile string     `json:"certFile"`
  KeyFile string      `json:"keyFile"`
  Backends []Backend  `json:"backends"`
}

type Backend struct {
  Filter string       `json:"filter"`
  Backend string      `json:"backend"`
  BackendEnv string   `json:"backendEnv"`
  CertFile string     `json:"certFile"`
  ServerName string   `json:"serverName"`
}

func GetConfiguration(file string) Config {
  raw, err := ioutil.ReadFile(file)

  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  var config Config
  if err := json.Unmarshal(raw, &config); err != nil {
    panic(err)
  }

  fmt.Printf("Proxy configuration read from file %q \n%s", file, ToNiceJson(config));

  config.Backends = ReplaceEnvironmentVariables(config.Backends);

  return config
}

func (backend Backend) ToString() string {
  return ToJson(backend)
}

func ToNiceJson(conf interface{}) string {
  str := ToJson(conf)
  str = strings.Replace(str, "},", "},\n\t", -1)
  str = strings.Replace(str, "[", "[\n\t", -1)
  str = strings.Replace(str, "]", "\n]", -1)
  str = strings.Replace(str, "],", "],\n", -1)
  return str + "\n";
}

func ToJson(conf interface{}) string {
  bytes, err := json.Marshal(conf)
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  return string(bytes)
}

func ReplaceEnvironmentVariables(backends []Backend) []Backend {
  var modified []Backend
  modified = make([]Backend, len(backends))
  for i, backend := range backends {
    modified[i] = backend.ReplaceEnvironmentVariables()
  }
  return modified;
}

func (backend Backend) ReplaceEnvironmentVariables() Backend {
  if backend.BackendEnv != "" && os.Getenv(backend.BackendEnv) != "" {
    backend.Backend = os.Getenv(backend.BackendEnv)
  }
  return backend
}