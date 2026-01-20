package parsecsv

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type CSV struct {
	Header   []string
	Data     map[string][]string
	RowCount int
}

func (c *CSV) Unmarshal(f *os.File) error {
	data := make(map[string][]string, 0)
	header := []string{}
	line_count := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if line_count == 0 {
			header = strings.Split(line, ",")
		} else {
			sl := strings.Split(line, ",")
			for i, v := range sl {
				data[header[i]] = append(data[header[i]], v)
			}
		}
		line_count++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	*c = CSV{
		Header:   header,
		Data:     data,
		RowCount: line_count,
	}

	return nil
}

func (c *CSV) Marshal() ([]byte, error) {
	var sb strings.Builder

	{
		_, err := sb.WriteString(strings.Join(c.Header, ",") + "\n")
		if err != nil {
			return nil, err
		}
	}

	for row := range c.RowCount - 1 {
		line := make([]string, 0, len(c.Header))
		for _, v := range c.Header {
			line = append(line, c.Data[v][row])
		}

		_, err := sb.WriteString(strings.Join(line, ",") + "\n")
		if err != nil {
			return nil, err
		}
	}

	_, err := sb.WriteString("\n")
	if err != nil {
		return nil, err
	}

	return []byte(sb.String()), nil
}

func readFile(p string) (*os.File, error) {
	file, err := os.Open(p)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return file, nil
}

func Run() {
	file, err := readFile("./people.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var csv CSV

	if err := csv.Unmarshal(file); err != nil {
		panic(err)
	}

	os.Remove("./output.csv")

	contents, err := csv.Marshal()
	if err != nil {
		panic(err)
	}

	o, err := os.OpenFile("./output.csv", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	_, err = o.Write(contents)
	if err != nil {
		panic(err)
	}
}
