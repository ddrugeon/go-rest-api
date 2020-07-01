package model

type Version struct {
	GitCommit string `json:"git"`
	BuildDate string `json:"buildDate"`
	Version   string `json:"version"`
	DBVersion string `json:"db"`
}

func NewVersion(gitCommit, buildDate, version string, db string) *Version {
	return &Version{
		GitCommit: gitCommit,
		BuildDate: buildDate,
		Version:   version,
		DBVersion: db,
	}
}
