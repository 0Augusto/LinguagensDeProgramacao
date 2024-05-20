# LinguagensDeProgramacao
Repositório destinado para o desenvolvimento do trabalho da disciplina Linguagens de Programção

# Data de Apresentação
05/06/2024

# Professor
Marco Rodrigo Costa

# Integrantes do Grupo
'''

- Henrique Augusto Rodrigues - Paradigma e Características
O paradigma de programação predominante em Go é o da programação concorrente. A linguagem foi projetada com suporte nativo para concorrência, o que significa que é fácil escrever programas que executam tarefas simultaneamente, aproveitando os recursos de hardware modernos, como múltiplos núcleos de processamento. [et. al]

- Luís Augusto Lima de Oliveira
- Laura Caetano Costa
- Victor Ferraz de Moraes

'''
# Tema do Trabalho
Linguagem de Programação Go

# Sumário
'''
- Introdução                                                                          **(página inicial ~ página final)**
- Roteiro                                                                             **(página inicial ~ página final)**
- Agenda                                                                              **(página inicial ~ página final)**
- Paradigma(s) a que pertence (características principais, etc)                       **(página inicial ~ página final)**
- Características mais marcantes da linguagem                                         **(página inicial ~ página final)**
- Linguagens relacionadas (influenciadores, influenciadas, similares, “opostas”, etc) **(página inicial ~ página final)**
- Exemplo(s) de programa(s)                                                           **(página inicial ~ página final)**
- PRÁTICA: tutoriais de instalação, uso e programação + exemplos]                     **(página inicial ~ página final)**
- Considerações finais                                                                **(página inicial ~ página final)**
- Bibliografia (livros, www, artigos, outras publicações)                             **(página inicial ~ página final)**
- Apêndice (em caso de trabalho com mais de 1 componente) – manifestação              **(página inicial ~ página final)**
individual de cada aluno(a) sobre o que mais lhe chamou a atenção em sua              **(página inicial ~ página final)**
linguagem tema, na pesquisa realizada. A redação de cada aluno(a) deve preencher      **(página inicial ~ página final)**
em torno de 1 página A4, deve ser objetiva e pode ter ilustrações, códigos ou outras  **(página inicial ~ página final)**
formas de expressar o destaque escolhido.                                             **(página inicial ~ página final)**

'''

## Imports mais Comuns da Linguagem ##

import (
    "fmt" // Pacote fmt é usado para formatação de saída, como impressão na tela.
    "os" // Pacote os fornece funcionalidades para interagir com o sistema operacional, como manipulação de arquivos.
    "net/http" // Pacote net/http é usado para criar servidores HTTP e clientes HTTP.
    "time" // Pacote time é usado para manipular datas, horários e temporizadores.
    "math/rand" // Pacote math/rand é usado para gerar números aleatórios.
    "encoding/json" // Pacote encoding/json é usado para codificar e decodificar dados no formato JSON.
    "database/sql" // Pacote database/sql é usado para interagir com bancos de dados SQL.
    _ "github.com/go-sql-driver/mysql" // Importa o driver MySQL para uso com o pacote database/sql.
    "github.com/gorilla/mux" // Pacote gorilla/mux é usado para criar roteadores HTTP mais avançados em Go.
    "github.com/dgrijalva/jwt-go" // Pacote jwt-go é usado para trabalhar com tokens JWT em Go.
    "github.com/jinzhu/gorm" // Pacote gorm é uma biblioteca ORM para interagir com bancos de dados relacionais em Go.
    "github.com/stretchr/testify/assert" // Pacote testify/assert é usado para escrever testes em Go.
    "github.com/spf13/viper" // Pacote viper é usado para lidar com configurações de aplicativos em Go.
    "github.com/go-redis/redis/v8" // Pacote redis é usado para interagir com um banco de dados NoSQL do Redis em Go.
    "github.com/joho/godotenv" // Pacote godotenv é usado para carregar variáveis de ambiente de arquivos .env em Go.
    "github.com/sirupsen/logrus" // Pacote logrus é usado para logging estruturado em Go.
    "github.com/pkg/errors" // Pacote errors fornece funções para manipulação de erros em Go.
)

## Criação da Linguagem ##

Frequentemente utilizada para servidores WEB, microserviços, aplicação de alto desempenho. Foi designada a Google em 2007 para aprimorar a produtividade de programação, na era multicore, networked, máquinas e grandes bases de códigos.
- Digitação estática e eficiência em tempo de execução (como C)
- Legibilidade e usabilidade (como Python)[24]
- Rede de alto desempenho e multiprocessamento

## Guia de Instalação ##

# LINUX: #

1. Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:
   $ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz

2. Add /usr/local/go/bin to the PATH environment variable.
You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):
export PATH=$PATH:/usr/local/go/bin

3. Verify that you've installed Go by opening a command prompt and typing the following command:
   go version

4. Confirm that the command prints the installed version of Go.

# Mac #

1. Open the package file you downloaded and follow the prompts to install Go.
The package installs the Go distribution to /usr/local/go. The package should put the /usr/local/go/bin directory in your PATH environment variable. You may need to restart any open Terminal sessions for the change to take effect.

2. Verify that you've installed Go by opening a command prompt and typing the following command:
   go version

3. Confirm that the command prints the installed version of Go.

# Windows #

1. Open the MSI file you downloaded and follow the prompts to install Go.
By default, the installer will install Go to Program Files or Program Files (x86). You can change the location as needed. After installing, you will need to close and reopen any open command prompts so that changes to the environment made by the installer are reflected at the command prompt.

2. Verify that you've installed Go.
- In Windows, click the Start menu.
- In the menu's search box, type cmd, then press the Enter key.
- In the Command Prompt window that appears, type the following command:
      $ go version

## Referência Bibliográfica ##

Documentação Go: https://go.dev/doc/effective_go
IDE GoLang:      https://go.dev/doc/editors
