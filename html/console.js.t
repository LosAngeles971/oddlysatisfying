function scrollToBottom () {
   var elem = document.getElementById('terminal');
   elem.scrollTop = elem.scrollHeight
}

var screen = document.getElementById('terminal');

var typeToScreen = new Typewriter(screen, {
    loop: false,
    delay: 1,
    autoStart: true,
    cursor: '█',
    strings: []
});

typeToScreen
    .deleteAll()
    {{- range .Log }}
    .typeString('{{ .Line }}<br>')
    .callFunction(scrollToBottom)
    {{- end }}
    .callFunction(scrollToBottom);