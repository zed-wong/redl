package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

func Single(c *cli.Context) error {
	if len(c.String("id")) == 0 {
		return errors.New("Course ID (-i) is required")
	}
	if len(c.String("token")) == 0 {
		return errors.New("JWT token (-t) is required")
	}
	dir := fmtDir(c.String("dir"))
	DownloadSingleCourse(dir, c.String("base"), c.String("token"), c.Int("id"))
	return nil
}
func Range(c *cli.Context) error {
	if len(c.String("range")) == 0 {
		return errors.New("Range (-r) is required")
	}
	if len(c.String("token")) == 0 {
		return errors.New("JWT token (-t) is required")
	}
	if !strings.Contains(c.String("range"), "-") {
		return errors.New("Invalid range, must contain '-'. (e.g. 1-1000)")
	}
	dir := fmtDir(c.String("dir"))
	ranges := c.String("range")

	rangel := strings.Split(ranges, "-")
	if len(rangel) != 2 {
		return errors.New("Invalid range, must contain from and to. (e.g. 100-200)")
	}
	r0, err := strconv.Atoi(rangel[0])
	if err != nil {
		return errors.New("Invalid from, must be int. (e.g. 100-200)")
	}
	r1, err := strconv.Atoi(rangel[1])
	if err != nil {
		return errors.New("Invalid to, must be int. (e.g. 100-200)")
	}
	if r1 < r0 {
		return errors.New("Invalid range, from must be smaller than to. (e.g. 1-2)")
	}

	DownloadRange(dir, c.String("base"), c.String("token"), r0, r1)
	return nil
}
func List(c *cli.Context) error {
	list := c.String("range")
	if len(c.String("range")) == 0 {
		return errors.New("Range (-r) is required")
	}
	if len(c.String("token")) == 0 {
		return errors.New("JWT token (-t) is required")
	}
	dir := fmtDir(c.String("dir"))
	arr := strings.Split(list, ",")
	if len(arr) == 0 {
		return errors.New("Please use comma between each course id, e.g. '1,234,567'")
	}
	if len(arr) == 1 {
		return errors.New("To download a single course, pls use -single")
	}
	for i := 0; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			return errors.New(fmt.Sprintf("Please use number only (%s)", arr[i]))
		}
		DownloadSingleCourse(dir, c.String("base"), c.String("token"), num)
	}
	return nil
}

func All(c *cli.Context) error {
	if len(c.String("token")) == 0 {
		return errors.New("JWT token (-t) is required")
	}
	dir := fmtDir(c.String("dir"))
	DownloadAll(dir, c.String("base"), c.String("token"))
	return nil
}

func fmtDir(dir string) string {
	if len(dir) != 0 {
		if dir[len(dir)-1:] != "/" {
			return dir + "/"
		}
	}
	return dir
}
