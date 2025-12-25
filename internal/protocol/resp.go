package protocol

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func readLine(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	line = strings.TrimSuffix(line, "\r\n")
	return line, nil
}

func readBulkString(r *bufio.Reader) (string, error) {
	line, err := readLine(r)
	if err != nil {
		return "", err
	}

	if len(line) == 0 || line[0] != '$' {
		return "", fmt.Errorf("expected a bulk string")
	}

	length, err := strconv.Atoi(line[1:])
	if err != nil {
		return "", err
	}

	buf := make([]byte, length+2)

	_, err = r.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:length]), nil
}

func ReadCommand(r *bufio.Reader) (*Command, error) {
	line, err := readLine(r)
	if err != nil {
		return nil, err
	}

	if len(line) == 0 || line[0] != '*' {
		return nil, errors.New("expected an array")
	}

	count, err := strconv.Atoi(line[1:])
	if err != nil {
		return nil, err
	}

	parts := make([]string, 0, count)

	for i := 0; i < count; i++ {
		s, err := readBulkString(r)
		if err != nil {
			return nil, err
		}

		parts = append(parts, s)
	}

	if len(parts) == 0 {
		return nil, errors.New("empty command provided")
	}

	return &Command{
		Name: strings.ToUpper(parts[0]),
		Args: parts[1:],
	}, nil
}
