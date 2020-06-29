package model

type Version struct {
	GitCommit string `json:"git"`
	BuildDate string `json:"buildDate"`
	Version   string `json:"version"`
}

func NewVersion(gitCommit, buildDate, version string) *Version {
	return &Version{
		GitCommit: gitCommit,
		BuildDate: buildDate,
		Version:   version,
	}
}
