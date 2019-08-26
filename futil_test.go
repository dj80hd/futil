package futil

import "testing"
import "io/ioutil"
import "os"
import "github.com/stretchr/testify/assert"

func TestGetLines(t *testing.T) {
	lines, err := GetLines("test/data/123.txt")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(lines), "should be 3 lines")
	assert.Equal(t, "one", lines[0], "")
	assert.Equal(t, "two", lines[1], "")
	assert.Equal(t, "three", lines[2], "")
}

func TestGetContent(t *testing.T) {
	content, err := GetContent("test/data/123.txt")
	assert.Nil(t, err)
	assert.Equal(t, "one\ntwo\nthree\n", content)
}

func TestSetContent(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "789.txt")
	assert.Nil(t, err)
	defer os.Remove(tmpfile.Name()) // clean up

	content := "seven\neight\nnine\n"
	err = SetContent(tmpfile.Name(), content)
	assert.Nil(t, err)
	lines, err := GetLines(tmpfile.Name())
	assert.Nil(t, err)
	assert.Equal(t, "nine", lines[2])
}

func TestSetLines(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "456.txt")
	assert.Nil(t, err)
	defer os.Remove(tmpfile.Name()) // clean up

	err = SetLines(tmpfile.Name(), []string{"four", "five", "six"})
	assert.Nil(t, err)
	content, err := GetContent(tmpfile.Name())
	assert.Nil(t, err)
	assert.Equal(t, "four\nfive\nsix\n", content)
}

func TestBadFileName(t *testing.T) {
	_, err := GetLines("not-a-real-filename")
	assert.NotNil(t, err)

	_, err = GetContent("not-a-real-filename")
	assert.NotNil(t, err)

	err = SetContent("/tmp", "you will never see this")
	assert.NotNil(t, err)

	err = SetLines("/tmp", []string{"line1", "line2"})
	assert.NotNil(t, err)
}
