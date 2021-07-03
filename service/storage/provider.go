package storage

import (
	"os"
  "path"
	"time"

	anacrolixResource "github.com/anacrolix/missinggo/v2/resource"
)

type Provider struct {
  basedir string
  Instances []*Instance
}

func NewProvider(basedir string, lifetime time.Duration) *Provider {
  var provider *Provider
  provider = &Provider{
    basedir: basedir,
    Instances: make([]*Instance, 0),
  }

  go cleanup(provider, lifetime)

  return provider
}

func (p *Provider) NewInstance(filename string) (anacrolixResource.Instance, error) {
  filename = path.Join(p.basedir, filename)

  os.MkdirAll(path.Dir(filename), os.ModePerm)

  var instance *Instance
  instance = &Instance{
    filename: filename,
    activeCount: 0,
    lastaccess: time.Now(),
  }

  p.Instances = append(p.Instances, instance)

  return instance, nil
}  

func cleanup(p *Provider, lifetime time.Duration){
  for {
    var key int
    var instance *Instance
    for key_candidate, candidate := range p.Instances {
      if instance == nil {
        key = key_candidate
        instance = candidate
        continue
      }

      if candidate.LastActive().Before(instance.LastActive()) {
        key = key_candidate
        instance = candidate
        continue
      }
    }

    if instance == nil {
      time.Sleep(lifetime)
      continue
    }

    var sleep time.Duration
    sleep = instance.LastActive().Add(lifetime).Sub(time.Now())
    if sleep > 0 {
      time.Sleep(sleep)
      continue
    }

    if err := instance.Delete(); err == nil {
      p.Instances = append(p.Instances[:key], p.Instances[key+1:]...)
    }
  }
}
