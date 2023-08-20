package utils

func (i *Input) ReadChunk() (string, error) {
	lines := ""
	l := 0
	i.rdr.Buffer(make([]byte, i.BatchSize), i.BatchSize)
	for l < i.BatchSize && i.rdr.Scan() {
		lines += i.rdr.Text()
		l++
	}
	if err := i.rdr.Err(); err != nil {
		return "", err
	}
	return lines, nil
}
