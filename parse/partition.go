package parse

import (
	"bufio"
	"fmt"
	"os"
)

func ParititionFile(path string, num int)  error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	ofs := make([]*os.File, num, num)
	writers := make([]*bufio.Writer, num, num)

	for i := 0; i < num; i++ {
		ofs[i], err = os.Create(fmt.Sprintf("%v_part%vof%v", path, i, num))
		if err != nil {
			return err
		}
		defer ofs[i].Close()
		writers[i] = bufio.NewWriter(ofs[i])
	}

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		_,err = writers[i].WriteString(line + "\n")
		if err != nil {
			return err
		}
		i = (i + 1) % num
	}

	return nil
}
