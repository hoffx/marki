# Test

Lorem ipsum dolor sit amet, **consectetur** adipisici elit, sed eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquid ex ea commodi consequat. Quis aute iure reprehenderit in voluptate $\Gamma(n) = (n-1)!\quad\forall n\in\mathbb N$ velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint obcaecat cupiditat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

$$
\Gamma(z) = \int_0^\infty t^{z-1}e^{-t}dt\,.
$$

Lorem ipsum dolor sit amet, **consectetur** adipisici elit, sed eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquid ex ea commodi consequat. Quis aute iure reprehenderit in voluptate $\Gamma(n) = (n-1)!\quad\forall n\in\mathbb N$ velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint obcaecat cupiditat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

```go
package main

import (
	"os"

	"github.com/hoffx/marki/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Marki"
	app.Usage = "Markdown rendering server"
	app.Author = "Hoff Industires"
	app.Commands = []cli.Command{cmd.Web}
	app.Run(os.Args)
}
```
Lorem ipsum dolor sit amet, **consectetur** adipisici elit, sed eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquid ex ea commodi consequat. Quis aute iure reprehenderit in voluptate $\Gamma(n) = (n-1)!\quad\forall n\in\mathbb N$ velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint obcaecat cupiditat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
