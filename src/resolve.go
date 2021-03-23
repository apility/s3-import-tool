package src

import (
	"fmt"
	"os"
	"path/filepath"
)

type Target struct {
	LocalFilename string
}

func (t Target) GetS3ObjectKey(c Configuration) (string, error) {

	f := filepath.Base(t.LocalFilename)
	return f, nil
}

func (t Target) ToString(c Configuration) string {
	path, err := t.GetS3ObjectKey(c)
	if err != nil {
		path = "error"
	}
	return fmt.Sprintf("%s -> s3://%s/%s", t.LocalFilename, c.BucketName, path)
}

func (c Configuration) CreateTargetsList() ([]Target, error) {
	targets := make([]Target, 0)

	for i, target := range c.Paths {
		abs, err := filepath.Abs(target)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse argument %d, %s", i, err)
		}
		files, err := expandGlob(abs, c.RecursiveSearch)
		if err != nil {
			return nil, fmt.Errorf("Unable to traverse paths for glob %d: %s", i, err)
		}
		for _, file := range files {
			info, err := os.Stat(file)
			if err != nil {
				fmt.Printf("Error: Could not check file: %s: %s\r\n", file, err.Error())
				continue
			}
			if info.IsDir() == false {
				targets = append(targets, Target{
					LocalFilename: file,
				})
			}
		}
	}
	return targets, nil
}

func expandGlob(path string, recursive bool) ([]string, error) {
	filenames, err := filepath.Glob(path)
	if err != nil {
		return nil, err
	}
	for _, filename := range filenames {
		fInfo, err := os.Stat(filename)
		if err != nil {
			return nil, err
		}
		if fInfo.IsDir() && (recursive) {
			var files []string
			err = filepath.Walk(filename, func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() {
					files = append(files, path)
				}
				return err
			})

			for _, file := range files {
				dirContents, err := expandGlob(file, recursive)
				if err != nil {
					return nil, fmt.Errorf("Unable to expand %s", file)
				}
				filenames = append(filenames, dirContents...)
			}

		} else if fInfo.IsDir() == false {
			//filenames = append(filenames, fInfo.Name())
		}
	}
	return filenames, nil
}
