package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Vai retornar a aplicação de linha de comando para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App de linha de comando"
	app.Usage = "Busca ips e nomes"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca ips de endereços",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "rancher.szanetwork.com",
				},
			},
			Action: func(c *cli.Context) {

			},
		},
	}

	return app
}

func buscarIP(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
