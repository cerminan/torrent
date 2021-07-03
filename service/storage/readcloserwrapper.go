package storage

import "io"


type ReadCloserWrapper struct {
  session io.ReadCloser
  closecallaback func()
}

func NewReadCloserWrapper(session io.ReadCloser, closecallaback func()) ReadCloserWrapper {
  return ReadCloserWrapper{
    session: session,
    closecallaback: closecallaback,
  }
}

func (rcw ReadCloserWrapper) Read(b []byte) (int, error) {
  return rcw.session.Read(b)
}

func (rcw ReadCloserWrapper) Close() error {
  var err error
  err = rcw.session.Close()
  if err != nil {
    return err
  }
  
  rcw.closecallaback()
  return nil
}
