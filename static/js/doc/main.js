document.addEventListener("DOMContentLoaded", function() {
  console.log("hallo");
  renderMathInElement(document.body, {
    displayMode: true,
    delimiters:[
      {left: "$$", right: "$$", display: true},
      {left: "$", right: "$", display: false},
      {left: "\\[", right: "\\]", display: true},
      {left: "\\(", right: "\\)", display: false}
      ]
  });
});