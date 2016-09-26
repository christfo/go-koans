package go_koans

import (
	"bytes"
	"io"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		written, _ := io.Copy(out, in)
		assert(written == 11)
		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		written, _ := io.CopyN(out, in, 5)
		assert(written == 5)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
