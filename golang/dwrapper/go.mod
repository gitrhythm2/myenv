module dwrapper

go 1.16

require (
	github.com/fatih/color v1.11.0
	local.packages/myenv v0.0.0-00010101000000-000000000000
)

replace local.packages/myenv => ../lib
