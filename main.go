package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

const (
	outputContrib = `{{range $idx, $v := .}}{{if $idx}}, {{end}}@{{$v.Login}} <{{$v.Name}}>{{end}}
`
)

type Contributor struct {
	Login string
	Name  string
}

func initFlag() (string, string) {
	since := flag.String("since", "", "commit/tag to start from (if not specified takes latest release tag)")
	until := flag.String("until", "master", "commit/tag to end on")
	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage of %[1]s: %[1]s [commit/tag range options] [owner]/[repo]\n\n"+
				"Retrieves contributors from a github repository between 2 commits or tags.\n\n"+
				"Examples:\n"+
				"\t%[1]s 42wim/matterbridge\n"+
				"\t%[1]s --since v1.11.0 --until v1.12.0 42wim/matterbridge\n"+
				"\t%[1]s --since 6f131250f1f48fb3898ee4c6717d9299a215ff67 --until v1.12.0 42wim/matterbridge\n"+
				"\t%[1]s --since 6f131250 --until v1.12.0 42wim/matterbridge\n\n"+
				"Options for commit/tag range:\n\n",
			os.Args[0],
		)
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() > 1 || flag.Arg(0) == "" {
		flag.Usage()
		os.Exit(0)
	}
	return *since, *until
}

func initStats() *Stats {
	sinceTag, untilTag := initFlag()
	repo := flag.Arg(0)

	repoinfo := strings.Split(repo, "/")
	if len(repoinfo) != 2 {
		fmt.Println("incorrect repo")
		flag.Usage()
		os.Exit(1)
	}
	s := &Stats{
		User:     repoinfo[0],
		Repo:     repoinfo[1],
		SinceTag: sinceTag,
		UntilTag: untilTag,
		Token:    os.Getenv("GITHUB_TOKEN"),
	}
	s.Init()
	return s
}

func main() {
	s := initStats()
	m, err := s.GetCommits()
	if err != nil {
		fmt.Println(err)
		return
	}

	cbs := []Contributor{}
	for login, commit := range m {
		cbs = append(cbs, Contributor{
			Name:  commit.GetCommit().GetAuthor().GetName(),
			Login: login,
		})
	}
	t := template.Must(template.New("output").Parse(outputContrib))
	err = t.Execute(os.Stdout, cbs)
	if err != nil {
		fmt.Println("executing template:", err)
		os.Exit(1)
	}
}
