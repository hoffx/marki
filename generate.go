package main

import (
	_ "github.com/tdewolff/minify"
)

//go:generate echo ------------------------------------------------------------------------
//go:generate echo
//go:generate echo This file (generate.go) includes all necessary go generate commands of
//go:generate echo this project. While some of them only have to be run when initializing
//go:generate echo the project on your machine and others more frequently. Just disable
//go:generate echo the ones you don't need by removing the "go:generate" part.
//go:generate echo
//go:generate echo ------------------------------------------------------------------------

// ------ project initialization ------------------------------------------
//
//go:generate go get .
//go:generate wget -O static/extensions/katex/katex.min.js https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.9.0-alpha2/katex.min.js
//go:generate wget -O static/extensions/katex/katex.min.css https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.9.0-alpha2/katex.min.css
//go:generate wget -O static/extensions/katex-autorender/katex-autorender.min.js https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.9.0-alpha2/contrib/auto-render.min.js

// ------ others ----------------------------------------------------------
//
//go:generate minify -r --match .+\.js static/js/doc static/extensions/katex static/js static/extensions/katex-autorender -o static/js/doc.min.js
