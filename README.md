# AtualizaLucedata

Crie um arquivo chamado rsrc.syso que conterá os recursos, incluindo o ícone.

Você pode usar uma ferramenta online para converter seu ícone para um arquivo .ico e, em seguida, convertê-lo para o formato binário usando o rsrc:
`rsrc -ico icon.ico -o rsrc.syso`

Para gerar o exe, rode dentro da pasta do projeto: `go build`
