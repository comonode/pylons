package types

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEntriesListSerialize(t *testing.T) {
	entries := GenEntries("chair", "Raichu")

	cases := map[string]struct {
	}{
		"basic flow test": {},
	}
	for testName := range cases {
		t.Run(testName, func(t *testing.T) {
			data, err := json.Marshal(entries)
			require.NoError(t, err)
			var serializedEntries EntriesList
			err = json.Unmarshal(data, &serializedEntries)
			require.NoError(t, err)
			require.True(t, reflect.DeepEqual(entries, serializedEntries))
		})
	}
}
