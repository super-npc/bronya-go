package bronya_go

import "embed"

//go:embed resources/static/*
var SysStatic embed.FS

//go:embed resources/public/*
var SysPublic embed.FS
