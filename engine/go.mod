module github.com/project-eria/xaal-go/engine

go 1.12

require (
	github.com/project-eria/logger v0.0.1
	github.com/project-eria/xaal-go/device v0.5.0
	github.com/project-eria/xaal-go/message v0.5.0
	github.com/project-eria/xaal-go/messagefactory v0.5.0
	github.com/project-eria/xaal-go/network v0.5.0
	github.com/project-eria/xaal-go/utils v0.5.0
	github.com/stretchr/testify v1.4.0
)

replace github.com/project-eria/xaal-go/device => ../device

replace github.com/project-eria/xaal-go/messagefactory => ../messagefactory

replace github.com/project-eria/xaal-go/message => ../message

replace github.com/project-eria/xaal-go/network => ../network

replace github.com/project-eria/xaal-go/utils => ../utils
