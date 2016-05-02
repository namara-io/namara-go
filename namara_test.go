package namara

import "testing"

var (
  logTemplate = "%s() failed...\nactual:   %s\nexpected: %v\n"
  dataSet = "5885fce0-92c4-4acb-960f-82ce5a0a4650"
  version = "en-1"
  nam = New("myapikey", false)
)

func TestBasePath(t *testing.T) {
  expectedBasePath := "https://api.namara.io/v0/data_sets/5885fce0-92c4-4acb-960f-82ce5a0a4650/data/en-1"

  if basePath := nam.BasePath(dataSet, version); basePath != expectedBasePath {
    t.Logf(logTemplate, "BasePath", basePath, expectedBasePath)
    t.Fail()
  }
}

func TestPath(t *testing.T) {
  expectedPath := "https://api.namara.io/v0/data_sets/5885fce0-92c4-4acb-960f-82ce5a0a4650/data/en-1/?api_key=myapikey&limit=1&offset=0&select=market_nam%2C+town&where=town+%3D+%27Toronto%27"
  allOptions := Options{"market_nam, town", "town = 'Toronto'", "1", "0", ""}

  if path := nam.Path(dataSet, version, &allOptions); path != expectedPath {
    t.Logf(logTemplate, "Path", path, expectedPath)
    t.Fail()
  }
}

func TestUrlEncode(t *testing.T) {
  expectedEncoded := "limit=1&offset=0&select=market_nam%2C+town&where=town+%3D+%27Toronto%27"

  allOptions := Options{"market_nam, town", "town = 'Toronto'", "1", "0", ""}

  if encoded := allOptions.urlEncode(); encoded != expectedEncoded {
    t.Logf(logTemplate, "urlEncode", encoded, expectedEncoded)
    t.Fail()
  }
}

func TestIsAggregation(t *testing.T) {
  without := Options{Operation: "count(*)"}
  with := Options{Select: "town"}

  if check := without.isAggregation(); check != true {
    t.Logf(logTemplate, "isAggregation", check, true)
    t.Fail()
  }

  if check := with.isAggregation(); check != false {
    t.Logf(logTemplate, "isAggregation", check, false)
    t.Fail()
  }
}
