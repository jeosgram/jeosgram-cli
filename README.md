# Jeosgram-CLI

El directorio del CLI es `{HOME}/.jeosgram/`, aca se almacenan las configuraciones y demas cosas...

El CLI esta construido sobre:
- [cobra](https://github.com/spf13/cobra) - cosas de cli
- [pterm](https://github.com/pterm/pterm) - terminal guapa
- [survey](https://github.com/AlecAivazis/survey) - input key


Todavia no hay instalador, sin embargo se puede compilar como cualquier programa de go
```bash
go build -ldflags "-s -w" -o jeosgram main.go
```

En caso de windows
```bash
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o jeosgram.exe main.go
```


### TODO

- [ ] script para administrar los lanzamientos en distintas plataformas o aprender a usar [goreleaser](https://goreleaser.com)

