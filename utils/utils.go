package utils

import (
	"bufio"
	"os"
	"time"
)

func ReadFileLineByLine(filename string, handle func(string) error) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if err := handle(scanner.Text()); err != nil {
			return err
		}
	}

	return nil
}

func Retry(f func() error) (err error) {
	if err = f(); err == nil {
		return
	}

	m := 100 * time.Millisecond
	t := m

	for i := 1; i < 10; i++ {
		time.Sleep(t)
		t += m

		if err = f(); err == nil {
			return
		}
	}

	return
}
