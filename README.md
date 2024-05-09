# Terminal Golang Channel

Um chat em tempo real por linha de comando implementado na linguagem Go e Sockets TCP.

---

## Features

- Baseado em mensagens de texto
- Chat anônimo
- Multi boards
- Implementado em Sockets TCP puros

---

## Como utilizar

Clone este repositório na sua máquina e no diretório do projeto, execute o comando:

```bash
go run . 8080
```

Com isso, o servidor do chat já estará executando e escutando na porta informada ou na 8080 por padrão.
Agora você pode conectar quantos clientes quiser atráves do comando `telnet`:

```bash
telnet localhost 8080
```

E pronto! Você já pode se comunicar via terminal entre dispositivos. Use `.help` para listar os comandos disponíveis. ex.: renomear nickname, listar e navegar entre as boards etc...
> Obs.:
> - Se certifique que o telnet está instalado na sua máquina. Caso não esteja você pode instalá-lo via o gerenciador de pacotes da sua distro Linux, ou ativar nos recursos do Windows.
> - Para se conectar em um servidor que está rodando em outro computador, substitua "localhost" pelo IP local da máquina no qual você deseja se conectar.
