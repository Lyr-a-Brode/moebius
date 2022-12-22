package api

//go:generate oapi-codegen -config ./server.cfg.yml ./spec/swagger.yml
//go:generate oapi-codegen -generate types -o server_types.gen.go --package api ./spec/swagger.yml
