package storage

import (
	"io"
	"io/fs"
	"os"
	"time"
)

type Instance struct {
  activeCount int
  filename string
  lastaccess time.Time
}

func (i *Instance) Get() (io.ReadCloser, error) {
  i.active()

  var session io.ReadCloser
  var err error
  session, err = os.Open(i.filename)
  if err != nil {
    i.deactive()
    return session, err
  }
  
  return NewReadCloserWrapper(session, i.deactive), nil
}

func (i *Instance) Put(r io.Reader) error {
  i.active()
  defer i.deactive()

  var file *os.File 
  var err error
  file, err = os.OpenFile(i.filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0640)
  if err != nil {
    return err
  }
  defer file.Close()

  _, err = io.Copy(file, r)
  return err
}

func (i *Instance) ReadAt(b []byte, off int64) (int, error) {
  i.active()
  defer i.deactive()

  var file *os.File
  var err error
  file, err = os.Open(i.filename)
  if err != nil {
    return 0, err
  }
  defer file.Close()

  return file.ReadAt(b, off)
}

func (i *Instance) WriteAt(b []byte, off int64) (int, error) {
  i.active()
  defer i.deactive()

  var file *os.File 
  var err error
  file, err = os.OpenFile(i.filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0640)
  if err != nil {
    return 0, err
  }
  defer file.Close()

  return file.WriteAt(b, off)
}

func (i *Instance) Stat() (fs.FileInfo, error) {
  i.active()
  defer i.deactive()

  return os.Stat(i.filename)
}

func (i *Instance) Delete() error {
  i.active()
  defer i.deactive()

  return os.Remove(i.filename)
}

func (i *Instance) Readdirnames() ([]string, error) {
  i.active()
  defer i.deactive()

  var file *os.File
  var err error
  file, err = os.Open(i.filename)
  if err != nil {
    return []string{}, err
  }
  defer file.Close()

  return file.Readdirnames(0)
}

func (i *Instance) active(){
  i.activeCount = i.activeCount + 1 
}

func (i *Instance) deactive(){
  i.activeCount = i.activeCount - 1

  if i.activeCount == 0 {
    i.lastaccess = time.Now()
  }
}

func (i *Instance) LastActive() time.Time {
  if i.activeCount != 0 {
    i.lastaccess = time.Now()
  }

  return i.lastaccess
}
