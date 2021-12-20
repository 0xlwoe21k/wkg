package safe


func Go(f func() error, c chan error) {
	go func() {
		err := Call(f)
		if c != nil {
			c <- err
		}
	}()
}