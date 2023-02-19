package funfacts

import (
  
  "github.com/Nicoletapa/funtemps/conv"
)

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra

  Sett inn alle Funfucts i en struktur
  type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
  }
*/
type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
}

func GetFunFacts(about string) []string {
  switch about {
  case "sun":
    return sun()
  case "luna" :
    return luna()
  case "terra" :
    return terra()
  default:
    return []string{}
  }
}


func sun() [] string{
  return []string{
  "Temperatur i Solens kjerne er 15000000C",
  "Temperatur på ytre lag av Solen",
  }
}



func luna() []string {
  return []string {
    "Temperatur på Månens overflate om natten er -183C",
    "Temperatur på Månens overflate om dagen er 106C",
  }
}
func terra() []string  {
  return []string {
    "Laveste temperatur målt på Jordens overflate er -89.4C",
    "Temperatur i Jordens indre kjerne er 9392K",
  }
}

