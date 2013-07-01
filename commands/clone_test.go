package commands

import (
	"github.com/bmizerany/assert"
	"github.com/jingweno/gh/github"
	"os"
	"path/filepath"
	"testing"
)

func TestTransformCloneArgs(t *testing.T) {
	github.DefaultConfigFile = "./test_support/clone_gh"
	config := github.Config{User: "jingweno", Token: "123"}
	github.SaveConfig(&config)
	defer os.RemoveAll(filepath.Dir(github.DefaultConfigFile))

	args := NewArgs([]string{"foo/gh"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.Size())
	assert.Equal(t, "git://github.com/foo/gh.git", args.First())

	args = NewArgs([]string{"-p", "foo/gh"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.Size())
	assert.Equal(t, "git@github.com:foo/gh.git", args.First())

	args = NewArgs([]string{"jingweno/gh"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.Size())
	assert.Equal(t, "git@github.com:jingweno/gh.git", args.First())

	args = NewArgs([]string{"-p", "acl-services/devise-acl"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.Size())
	assert.Equal(t, "git@github.com:acl-services/devise-acl.git", args.First())

	args = NewArgs([]string{"-p", "jekyll_and_hyde"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.Size())
	assert.Equal(t, "git@github.com:jingweno/jekyll_and_hyde.git", args.First())
}