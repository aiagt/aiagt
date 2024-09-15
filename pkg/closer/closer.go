package closer

type Closer interface {
	Close() error
}

func Close(c Closer) {
	_ = c.Close()
}
