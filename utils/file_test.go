package utils

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestReadFileValidFile(t *testing.T) {
	//ReadFile(fileName string) ([]byte, error)
	_, err := ReadFile("../config/config.yml")
	require.Nil(t, err)
}

func TestReadFileInValidFile(t *testing.T) {
	_, err := ReadFile("./xxxxx.yml")
	require.NotNil(t, err)
}

