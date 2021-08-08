package service

import (
  "regexp"
)

func isMagnet(magnet string) bool {
  var re *regexp.Regexp
  re = regexp.MustCompile("(?i)^magnet:\\?xt=urn:[a-z0-9]+:[a-z0-9]{32,40}(?:&dn=[^&]+)?(?:&(?:tr|xs)=[^&]+)*$")
  
  var match bool
  match = re.Match([]byte(magnet)) 

  return match
}
