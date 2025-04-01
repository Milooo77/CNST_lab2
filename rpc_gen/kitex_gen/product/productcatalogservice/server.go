// Code generated by Kitex v0.9.1. DO NOT EDIT.
package productcatalogservice

import (
	product "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler product.ProductCatalogService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler product.ProductCatalogService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
