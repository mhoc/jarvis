
// Interface into datum storage
// Datums are user-generated strings which we want to store and retrieve
// Aliases are provided for each datum so we can match against multiple things
// and still store it under the same key

package data

import (
  "github.com/mhoc/jarvis/log"
)

type Datum struct {
  Key string
  UserSpec bool
  Aliases []string
}

var RegisteredDatums = []Datum {
  Datum{
    Key: "user-birthday-",
    UserSpec: true,
    Aliases: []string{"my birthday", "my day of birth", "my birthdate"},
  },
  Datum{
    Key: "user-zipcode-",
    UserSpec: true,
    Aliases: []string{"my zipcode", "my zip code", "my zip"},
  },
}

func GetDatumFromAlias(target string) (bool, Datum) {
  for _, dat := range RegisteredDatums {
    for _, alias := range dat.Aliases {
      if alias == target {
        return true, dat
      }
    }
  }
  return false, Datum{}
}

func StoreDatum(trigger string, value string, user string) bool {
  log.Trace("Storing datum %v", trigger)
  in, dat := GetDatumFromAlias(trigger)
  if in && dat.UserSpec {
    set(dat.Key + user, value)
  } else if in {
    set(dat.Key, value)
  } else {
    return false
  }
  return true
}

func GetDatum(trigger string, user string) (bool, string) {
  log.Trace("Getting datum %v", trigger)
  in, dat := GetDatumFromAlias(trigger)
  if in && dat.UserSpec {
    return get(dat.Key + user)
  } else if in {
    return get(dat.Key)
  } else {
    return false, ""
  }
}