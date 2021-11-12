package filter

import (
	severity2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/severity"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTagBasedFilter(t *testing.T) {
	{
		filter := New(&Config{
			Tags: []string{"cves", "2021", "jira"},
		})

		t.Run("true", func(t *testing.T) {
			matched, _ := filter.Match([]string{"jira"}, []string{"pdteam"}, severity2.Low, nil)
			require.True(t, matched, "could not get correct match")
		})
		t.Run("false", func(t *testing.T) {
			matched, _ := filter.Match([]string{"consul"}, []string{"pdteam"}, severity2.Low, nil)
			require.False(t, matched, "could not get correct match")
		})
		t.Run("match-extra-tags-positive", func(t *testing.T) {
			matched, _ := filter.Match([]string{"cves", "vuln"}, []string{"pdteam"}, severity2.Low, []string{"vuln"})
			require.True(t, matched, "could not get correct match")
		})
		t.Run("match-extra-tags-negative", func(t *testing.T) {
			matched, _ := filter.Match([]string{"cves"}, []string{"pdteam"}, severity2.Low, []string{"vuln"})
			require.False(t, matched, "could not get correct match")
		})
	}

	t.Run("not-match-excludes", func(t *testing.T) {
		filter := New(&Config{
			ExcludeTags: []string{"dos"},
		})
		matched, err := filter.Match([]string{"dos"}, []string{"pdteam"}, severity2.Low, nil)
		require.False(t, matched, "could not get correct match")
		require.Equal(t, ErrExcluded, err, "could not get correct error")
	})
	t.Run("match-includes", func(t *testing.T) {
		filter := New(&Config{
			Tags:        []string{"cves", "fuzz"},
			ExcludeTags: []string{"dos", "fuzz"},
			IncludeTags: []string{"fuzz"},
		})
		matched, err := filter.Match([]string{"fuzz"}, []string{"pdteam"}, severity2.Low, nil)
		require.Nil(t, err, "could not get match")
		require.True(t, matched, "could not get correct match")
	})
	t.Run("match-includes", func(t *testing.T) {
		filter := New(&Config{
			Tags:        []string{"fuzz"},
			ExcludeTags: []string{"fuzz"},
		})
		matched, err := filter.Match([]string{"fuzz"}, []string{"pdteam"}, severity2.Low, nil)
		require.Nil(t, err, "could not get match")
		require.True(t, matched, "could not get correct match")
	})
	t.Run("match-author", func(t *testing.T) {
		filter := New(&Config{
			Authors: []string{"pdteam"},
		})
		matched, _ := filter.Match([]string{"fuzz"}, []string{"pdteam"}, severity2.Low, nil)
		require.True(t, matched, "could not get correct match")
	})
	t.Run("match-severity", func(t *testing.T) {
		filter := New(&Config{
			Severities: severity2.Severities{severity2.High},
		})
		matched, _ := filter.Match([]string{"fuzz"}, []string{"pdteam"}, severity2.High, nil)
		require.True(t, matched, "could not get correct match")
	})
	t.Run("match-exclude-with-tags", func(t *testing.T) {
		filter := New(&Config{
			Tags:        []string{"tag"},
			ExcludeTags: []string{"another"},
		})
		matched, _ := filter.Match([]string{"another"}, []string{"pdteam"}, severity2.High, nil)
		require.False(t, matched, "could not get correct match")
	})
	t.Run("match-conditions", func(t *testing.T) {
		filter := New(&Config{
			Authors:    []string{"pdteam"},
			Tags:       []string{"jira"},
			Severities: severity2.Severities{severity2.High},
		})
		matched, _ := filter.Match([]string{"jira", "cve"}, []string{"pdteam", "someOtherUser"}, severity2.High, nil)
		require.True(t, matched, "could not get correct match")

		matched, _ = filter.Match([]string{"jira"}, []string{"pdteam"}, severity2.Low, nil)
		require.False(t, matched, "could not get correct match")

		matched, _ = filter.Match([]string{"jira"}, []string{"random"}, severity2.Low, nil)
		require.False(t, matched, "could not get correct match")

		matched, _ = filter.Match([]string{"consul"}, []string{"random"}, severity2.Low, nil)
		require.False(t, matched, "could not get correct match")
	})
}
