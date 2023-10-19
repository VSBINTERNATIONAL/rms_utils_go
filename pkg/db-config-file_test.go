package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDbConfigFile(t *testing.T) {
	dbConfigFile := "db_config.json"
	config, err := NewConfig(dbConfigFile)

	require.Nil(t, err)
	require.Equal(t, "web", config.Database.DBUser)
}
