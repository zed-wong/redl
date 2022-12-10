package main

import (
	"errors"
        "github.com/urfave/cli/v2"
)

func Single(c *cli.Context) error {
	if (len(c.String("id")) == 0) {
		return errors.New("Course ID (-i) is required");
	}
	if (len(c.String("token")) == 0) {
		return errors.New("JWT token (-t) is required");
	}
	dir := c.String("dir")
	if (dir[len(dir)-1:] != "/") {
		dir+="/"
	}
	DownloadSingleCourse(dir, c.String("base"), c.Int("id"), c.String("token"))
	return nil
}
func Range(c *cli.Context) error {
	return nil
}
func All(c *cli.Context) error {
	return nil
}
