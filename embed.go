package bronya_go

import "embed"

//go:embed resources/public/*
var DistFS embed.FS
