package server

import (
	"errors"
	"testing"

	"github.com/matryer/is"
)

func TestMeta(t *testing.T) {
	is := is.New(t)
	tables := []struct {
		slug        string
		institution string
		location    string
		country     string
		start_date  string
		end_date    string
		position    string
		occupation  string
		path        string
		err         error
	}{
		{
			"wobcom",
			"Wobcom GmbH.",
			"Wolsburg",
			"DE",
			"2023-05",
			"",
			"Software Developer Intern",
			"professional",
			"./static/wobcom/index.mdx",
			nil,
		},
	}
	for _, table := range tables {
		table := table
		t.Run(table.slug, func(t *testing.T) {
			var mdxFile MDXFile
			mdxFile.Read(table.path)
			if table.err == nil {
				is.NoErr(mdxFile.Meta.Err)
				is.Equal(mdxFile.Meta.Slug, table.slug)
				is.Equal(mdxFile.Meta.Slug, table.slug)
				is.Equal(mdxFile.Meta.Slug, table.slug)
				is.Equal(mdxFile.Meta.Slug, table.slug)
				is.Equal(mdxFile.Meta.Slug, table.slug)
				is.Equal(mdxFile.Meta.Slug, table.slug)
				is.Equal(mdxFile.Meta.Slug, table.slug)
				is.Equal(mdxFile.Meta.Slug, table.slug)
				is.Equal(mdxFile.Meta.Slug, table.slug)
			} else {
				is.True(mdxFile.Meta.Err != nil)
				is.True(errors.Is(mdxFile.Meta.Err, table.err))
			}
		})
	}
}
