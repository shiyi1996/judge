package g

var (
	GitVer    = "v0.0.1"
	BuildTime = "2017-10-01"
)

func Version() string {
	return GitVer
}

func BuildInfo() string {
	return BuildTime
}
