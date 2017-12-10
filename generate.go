package main

import (
	_ "github.com/tdewolff/minify"
)

// go get -u github.com/tdewolff/minify/cmd/minify

// wget -O static/extensions/katex/katex.min.js https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.9.0-alpha2/katex.min.js
// wget -O static/extensions/katex/katex.min.css https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.9.0-alpha2/katex.min.css
// wget -O static/extensions/katex-autorender/katex-autorender.min.js https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.9.0-alpha2/contrib/auto-render.min.js

//go:generate minify -r --match .+\.js static/js/doc static/extensions/katex static/js static/extensions/katex-autorender -o static/js/doc.min.js
