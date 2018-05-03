package errorLogging

type srvImpl struct {
	d dal
}

func newSrv(d dal) srv {
	return &srvImpl{d}
}

func (s *srvImpl) Feature1() error {
	err := s.d.Op1()
	if err != nil {
		return err
	}

	err = s.d.Op2()
	if err != nil {
		return err
	}

	return nil
}

func (s *srvImpl) Feature1Panic() {
	s.d.Op1Panic()
	s.d.Op2Panic()
}
