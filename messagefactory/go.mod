module github.com/project-eria/xaal-go/messagefactory

go 1.12

require (
	github.com/project-eria/logger v0.0.1
	github.com/project-eria/xaal-go/device v0.5.0
	github.com/project-eria/xaal-go/message v0.5.0
	github.com/project-eria/xaal-go/utils v0.5.0
	github.com/stretchr/testify v1.4.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
)

replace github.com/project-eria/xaal-go/device => ../device

replace github.com/project-eria/xaal-go/message => ../message

replace github.com/project-eria/xaal-go/utils => ../utils
