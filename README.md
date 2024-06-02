# LinguagensDeProgramacao
Repositório destinado para o desenvolvimento do trabalho da disciplina Linguagens de Programção

# Data de Apresentação
05/06/2024

# Professor
Orientador: Marco Rodrigo Costa

# Integrantes do Grupo
'''

- Henrique Augusto Rodrigues - Paradigma e Características
O paradigma de programação predominante em Go é o da programação concorrente sendo enssencialmente imperativo. A linguagem foi projetada com suporte nativo para concorrência, o que significa que é fácil escrever programas que executam tarefas simultaneamente, aproveitando os recursos de hardware modernos, como múltiplos núcleos de processamento. [et. al]
Os nomes são tão importantes no Go quanto em qualquer outro idioma. Eles até têm efeito semântico: a visibilidade de um nome fora de um pacote é determinada pelo fato de seu primeiro caractere ser maiúsculo. [GO. Documentação de Effective Go. Disponível em: https://go.dev/doc/effective_go. Acesso em: 20 fev. 2024.].
A programação simultânea consiste em organizar um programa de forma que diversos cálculos possam ser realizados ao mesmo tempo, durante períodos de tempo sobrepostos. Essa abordagem enfatiza a simultaneidade especificada diretamente pelo programador, em oposição àquela implícita por otimizações do compilador ou pelo hardware do sistema. Com o término do crescimento exponencial da velocidade dos núcleos de processador na metade da década de 2000, devido a limitações energéticas e térmicas, os fabricantes passaram a adotar a estratégia de colocar múltiplos núcleos em um único processador. Essa mudança possibilitou a execução de mais operações ao mesmo tempo. A escalabilidade do paralelismo não é recente, sendo comum em máquinas de grande porte utilizadas há décadas, sobretudo na computação científica e análise de dados. Mesmo em computadores com um único núcleo, sistemas operacionais e interpretadores conseguem simular a simultaneidade através da alternância de contextos entre diferentes tarefas, permitindo a execução de vários programas ao mesmo tempo. Atualmente, com a tendência de aumento do número de núcleos de processador, é crucial que os aplicativos sejam projetados de forma a explorar o paralelismo para otimizar a velocidade de processamento dentro de um único programa. [KAMIL, Amir. Curso de Linguagens de Programação na Universidade de Michigan, outono de 2016.] 
Rob Pike, membro da equipe Go na Google, discute padrões de concorrência em Go, enfatizando sua importância na estruturação de programas para interagir efetivamente com o mundo real. Ele esclarece que a concorrência envolve a composição de computações executando independentemente, contrastando-a com o paralelismo. A concorrência facilita o código limpo e compreensível que pode simular interações do mundo real sem exigir expertise em conceitos de baixo nível como barreiras de memória.

	Pike traça a linhagem de ideias e linguagens concorrentes que levaram ao Go, destacando o papel fundamental do artigo de Tony Hoare, "Communicating Sequential Processes", como base. Ele menciona linguagens como Occam, Erlang e Concurrent ML, observando a posição do Go na sequência Newsqueak-Limbo-Alef. Ao contrário de Erlang, que se comunica diretamente com processos, o Go usa canais como valores de primeira classe para comunicação.

	Esses exemplos mostram aspectos fundamentais da concorrência em Go, incluindo computações executando independentemente e comunicação entre goroutines por meio de canais.
	
	Rob Pike destaca que, embora um modelo não seja melhor que o outro, eles são formalmente equivalentes. Ele enfatiza o uso de canais em Go e decide focar em exemplos práticos, mesmo que tediosos, para demonstrar a concorrência de forma mais clara. Pike mostra um programa que imprime lentamente "la la la" uma vez por segundo, explicando que ele é executado em segundo plano enquanto o programa principal continua. Ele introduz a ideia de goroutines, que são funções que podem ser executadas independentemente, sem esperar pelo retorno. Pike também menciona canais como um meio fundamental de comunicação entre goroutines em Go, e como eles são declarados e inicializados no código.

	Pike explica o uso de canais em Go para comunicação e sincronização entre goroutines. Ao enviar um valor em um canal, utiliza-se o operador de seta para a esquerda, enquanto para receber um valor, utiliza-se o operador de seta para a direita. Ele mostra um exemplo prático onde uma função envia mensagens para um canal e outra função recebe essas mensagens, garantindo comunicação e sincronização entre as goroutines. Pike destaca que, ao enviar ou receber valores em um canal, a operação é bloqueante, o que sincroniza as goroutines. Ele introduz o conceito de canais bufferizados, mas opta por não utilizá-los em exemplos simples. Em vez disso, ele enfatiza o princípio de "não compartilhar memória, compartilhar comunicação" em Go, sugerindo o uso de canais para passagem de dados entre goroutines. Pike apresenta o padrão de concorrência de gerador, onde uma função retorna um canal, e demonstra como utilizar esse padrão para construir serviços concorrentes. Ele finaliza mostrando como combinar múltiplos canais em um único canal usando uma função multiplexadora, permitindo a execução independente e sincronizada de várias goroutines.

	Pike apresenta uma maneira de garantir que as goroutines operem de forma sincronizada e em lockstep. Ele demonstra como enviar um canal dentro de outro canal para controlar a sincronização entre as goroutines. Isso é feito construindo uma estrutura de mensagem que inclui tanto a mensagem a ser impressa quanto um canal de espera, que serve como um sinalizador para avançar. Pike mostra como implementar esse padrão nas funções concorrentes e como isso garante a sincronização entre as goroutines, mesmo quando a temporização é aleatória. Em seguida, ele introduz o uso da declaração select em Go, que é uma estrutura de controle semelhante a um switch, mas que permite controlar o comportamento do programa com base em quais comunicações podem prosseguir a qualquer momento. A declaração select é fundamental para o modelo de concorrência em Go, pois permite que o programa bloqueie até que uma comunicação possa prosseguir. Pike explica que, quando várias comunicações estão disponíveis ao mesmo tempo, a declaração select escolhe uma delas pseudoaleatoriamente. Ele compara esse comportamento com os comandos guardados de Dijkstra.

	Pike demonstra o uso da declaração select em Go para controlar o comportamento do programa com base em quais comunicações podem prosseguir. Ele compara a declaração select a um switch, onde cada caso representa uma comunicação possível. Isso permite que o programa bloqueie até que uma comunicação possa prosseguir. Ele mostra como usar a declaração select para implementar um limite de tempo em uma comunicação, garantindo que o programa não espere indefinidamente por uma resposta. Pike também introduz o uso de canais de sinalização para coordenar a finalização de operações assíncronas, garantindo que todas as tarefas estejam concluídas antes de encerrar o programa. Além disso, ele ilustra como utilizar canais para comunicação bidirecional entre goroutines, permitindo sincronização e troca de mensagens entre elas. Esses conceitos são exemplificados por meio de uma série de cenários de programação concorrente, mostrando diferentes maneiras de coordenar a execução de tarefas e controlar o fluxo do programa.

	Rob Pike apresenta uma série de exemplos para demonstrar o uso da concorrência em Go na construção de programas paralelos e robustos. Ele começa com exemplos simples, como o uso da declaração select para controlar comunicações assíncronas e a implementação de um sistema de busca fictício. Em seguida, ele avança para exemplos mais sofisticados, como a replicação de serviços de busca e a aplicação de limites de tempo para evitar a espera indefinida por respostas lentas. Ao longo do processo, Pike destaca a facilidade de uso das primitivas de concorrência em Go, enfatizando a simplicidade na construção de programas complexos sem a necessidade de preocupações com segurança ou barreiras de memória. Ele encerra com uma ressalva sobre o uso responsável dessas técnicas, alertando que, embora divertidas, elas devem ser aplicadas de forma criteriosa, utilizando-as como ferramentas para a construção eficiente de software concorrente.

	Rob Pike conclui destacando que goroutines e canais são características de concorrência essenciais em Go, que facilitam a construção de software concorrente interessante para resolver problemas do mundo real. Ele ressalta a importância de usar as ferramentas certas para cada tarefa e encoraja os desenvolvedores a equilibrar a estrutura do programa, combinando grandes componentes com as primitivas de concorrência apresentadas. Pike também menciona que, embora seja divertido lidar com a concorrência em Go, é importante não exagerar no seu uso. Ele sugere que, embora a linguagem ofereça maneiras poderosas de lidar com a concorrência, é fundamental aplicá-las de forma criteriosa, escolhendo a abordagem adequada para cada situação. Por fim, Pike responde a algumas perguntas da plateia sobre testes de código concorrente e possíveis melhorias na análise estática de programas em Go. [YouTube. Google I/O 2012 - Go Concurrency Patterns. Disponível em: <https://www.youtube.com/watch?v=f6kdp27TYZs>. Acesso em: 29 maio 2024, 6:34.]
 
 Dentre as características da linguagem, há a detecção de vulnerabilidaes onde, a linguagem em questão fornece ferramentas confiáveis e de baixo ruído para o aprendizado de vulnerabilidades para os desenvolvedores utilizando o *govulncheck* tendo uma extensão para o VS Code, baixando o VS Code Go [Tutorial: Comece com o VS Code Go. Disponível em: https://go.dev/doc/tutorial/govulncheck-ide. Acesso em 25 de maio de 2024 às 3:37].
 
 O Go é uma linguagem de programação que desempenha um papel fundamental tanto no **Desenvolvimento de Operações (DevOps)** quanto na **Engenharia de Confiabilidade do Site (SRE)**. Vamos dar uma olhada em como o Go beneficia essas áreas:

1. **DevOps**:
   - As equipes de DevOps automatizam tarefas e melhoram o processo de integração contínua e implantação contínua (CI/CD) nas organizações de engenharia.
   - O Go oferece **vantagens significativas** para os profissionais de DevOps:
     - **Biblioteca padrão robusta**: O Go possui uma extensa biblioteca padrão que inclui pacotes para necessidades comuns, como manipulação de HTTP, E/S de arquivos, formatação de tempo, expressões regulares e formatos JSON/CSV.
     - **Tempo de construção rápido**: O Go é conhecido por seus tempos de compilação e inicialização rápidos, permitindo que os profissionais de DevOps se concentrem na lógica de negócios sem atrasos excessivos.
     - **Digitação estática e tratamento explícito de erros**: A tipagem estática e o tratamento explícito de erros tornam até mesmo pequenos scripts mais robustos e fáceis de manter.

2. **SRE**:
   - A Engenharia de Confiabilidade do Site (SRE) foi desenvolvida no Google para tornar os sites de grande escala mais confiáveis, eficientes e escaláveis.
   - O Go é uma escolha natural para os profissionais de SRE pelas seguintes razões:
     - **Concorrência e rede**: Os recursos de concorrência e gerenciamento de rede do Go são ideais para ferramentas que lidam com implantação em nuvem. Isso permite a automação eficiente enquanto a infraestrutura de desenvolvimento cresce.
     - **Segurança e confiabilidade**: O Go é conhecido por sua segurança e confiabilidade, tornando-o adequado para sistemas críticos.
     - **Baixa pegada de memória e documentação automática**: O coletor de lixo do Go elimina preocupações com gerenciamento de memória, e o gerador automático de documentação (godoc) facilita a manutenção e o estabelecimento de boas práticas desde o início.

Em resumo, o Go é uma escolha sólida para profissionais de DevOps e SRE, oferecendo ferramentas poderosas para automação, escalabilidade e confiabilidade. [Operações de Desenvolvimento e Engenharia de Confiabilidade do Site. Disponível em: https://go.dev/solutions/devops. Acesso em 25 de maio de 2024 às 3:58]

Source: Conversation with Copilot, 31/05/2024
(1) Development Operations & Site Reliability Engineering - The Go .... https://go.dev/solutions/devops.
(2) What Is SRE? How Does It Relate to DevOps? - How-To Geek. https://www.howtogeek.com/devops/what-is-sre-how-does-it-relate-to-devops/.
(3) Supercharge your DevOps practice with SRE principles - Google Cloud. https://cloud.google.com/blog/products/devops-sre/supercharge-your-devops-practice-with-sre-principles.
(4) SRE in the 2022 State of DevOps report | Google Cloud Blog. https://cloud.google.com/blog/products/devops-sre/sre-in-the-2022-state-of-devops-report.   

Claro! Aqui estão algumas características principais da linguagem Go (GoLang) relacionadas ao uso de memória cache e concorrência:

1. **Concorrência Eficiente**:
   - Go possui suporte integrado para concorrência com goroutines e canais (channels).
   - Goroutines são "threads" leves gerenciadas pelo tempo de execução do Go, permitindo a execução concorrente eficiente de tarefas.
   - Canais facilitam a comunicação e sincronização entre goroutines, tornando a programação concorrente mais fácil e segura.

2. **Gerenciamento de Memória Eficiente**:
   - Go possui um coletor de lixo (garbage collector) automático que gerencia a alocação e desalocação de memória.
   - O coletor de lixo periodicamente escaneia a memória para identificar e liberar memória não utilizada, prevenindo vazamentos de memória e otimizando o uso de recursos.

3. **Padrões de Design para Memória Cache**:
   - Go oferece suporte para implementar padrões de design eficientes para memória cache.
   - Bibliotecas como "groupcache" e "go-cache" são comumente utilizadas para implementar mecanismos de cache eficientes em memória.
   - Essas bibliotecas permitem armazenar e recuperar dados frequentemente acessados na memória, melhorando o desempenho de aplicativos e sistemas.

4. **Sintaxe Simples e Limpa**:
   - Go possui uma sintaxe simples e limpa, facilitando a leitura e escrita de código.
   - A simplicidade da linguagem contribui para a eficiência do desenvolvimento e facilita a manutenção do código.

5. **Concorrência Segura**:
   - A linguagem Go é projetada para facilitar a escrita de código concorrente seguro.
   - O modelo de concorrência de Go, com goroutines e canais, promove práticas de programação seguras, evitando problemas comuns associados à concorrência, como condições de corrida e deadlocks.

Essas características tornam a linguagem Go uma escolha popular para o desenvolvimento de sistemas concorrentes e de alta performance, especialmente em ambientes onde o uso eficiente da memória e a concorrência são essenciais, como em servidores web e sistemas distribuídos. [et.al]

- Luís Augusto Lima de Oliveira
### História de Go
Go surgiu no final de 2007 como um projeto desenvolvido na Google por Robert Griesemer, Rob Pike, e Ken Thompson. A sua principal premissa inicial foi resolver problemas encontrados nas linguagens até então consolidadas no mercado para a utilização em servidores. O que foi observado é que linguagens como java, c, c++ e python não lidam bem com processadores multicore, sistemas em rede, clusters de computação maciça e o modelo de programação web atuais, problemas esses que não foram focos do período quando as linguagens foram estruturadas. 
Isso, combinado à escala de trabalho em servidores que se estende em milhares de linhas de código que demoram até mesmo horas para serem compilados resultava em um processo de desenvolvimento ineficiente e com grande margem de erro. Um exemplo disso foi uma compilação de C++ feita na google em 2007, que demorou cerca de 45 minutos para gerar o binário devido à má gestão de includes do compilador de C++. A versão de 2012 desse mesmo programa teve seu tempo de compilação diminuida quase que pela metade mesmo com mais recursos, compilando em 27 minutos.
Go foi, portanto, projetado e desenvolvido visando a Engenharia de Software na buscsa de tornar o trabalho neste ambiente de projetos de grande porte mais produtivo, assim como as aplicações mais eficientes. Em 2008, Go passou de um projeto de tempo parcial para um projeto de tempo integral no Google, e com isso, muitos outros engenheiros participaram de seu desenvolvimento a partir desse ponto. Em 2009 Go se tornou open-source, sendo sua primeira versão estável lançada em 2012. Em 2018 a logo atual da linguagem foi criada. A ultima versão até o momento é a 1.22.


#### Curiosidade
Go é comumente chamado de GoLang devido à antiga url do projeto, que era golang.org. Atualmente a nova url é go.dev, onde se encontram diversas informações assim como a documentação da linguagem. 

https://go.dev/blog/go-brand

https://go.dev/talks/2012/splash.article

https://en.wikipedia.org/wiki/Go_(programming_language)#Naming_dispute

https://www.ime.usp.br/~gold/cursos/2015/MAC5742/reports/GoLang.pdf

https://www.taylorfrancis.com/books/mono/10.1201/9781003309055/golang-sufyan-bin-uzayr

- Laura Caetano Costa
 ### Tutorial prático: vizualizando a concorrência em Go
## Passo 1: Instalação do Go
- Baixe o Go:
Vá para o site oficial do Go em https:	//golang.org/dl/.
- Baixe a versão apropriada para o seu sistema operacional.
- Instale o Go:
Siga as instruções de instalação fornecidas para o seu sistema operacional.
Verifique se o Go foi instalado corretamente executando o comando go version no terminal ou prompt de comando.

## Passo 2: Criar e Executar um Programa 
- Crie um Novo Diretório para o Projeto

- Crie um novo diretório onde você deseja criar o seu projeto Go

- Crie um Arquivo Fonte Go

- Abra o arquivo hello.go em um editor de texto ou IDE de sua escolha

- Escreva o seguinte código Go:
	
	import (
		"fmt"
		"time"
	)
	
	// função que calcula o quadrado de um número e envia o resultado para um canal
	func square(number int, ch chan int) {
		fmt.Printf("Goroutine iniciada para número: %d\n", number)
		result := number * number
		time.Sleep(time.Second) // Simula algum trabalho
		fmt.Printf("Goroutine finalizada para número: %d\n", number)
		ch <- result
		
		
	}
	
	func main() {
		fmt.Println("Iniciando programa")
	
		// Canal que pode receber e enviar inteiros
		results := make(chan int)
	
		// Cria goroutines para calcular o quadrado dos números de 1 a 10
		for i := 1; i <= 5; i++ {
			go square(i, results)
		}
	
		fmt.Println("Todas as goroutines foram enviadas")
	
		// Recebe e imprime os resultados das goroutines
		for i := 1; i <= 5; i++ {
			fmt.Println("Esperando pelo resultado", i)
			result := <-results
			fmt.Println("Resultado recebido:", result)
		}
	
		fmt.Println("Todos os resultados foram recebidos")
	}

- Salve e Feche o Arquivo

- Compile e Execute o Programa:
	go run hello.go


- Victor Ferraz de Moraes	
### Influências no GoLang

#### C

Uma das influências mais significativas no GoLang é a linguagem de programação C. A linguagem C é conhecida por seu desempenho e capacidades de programação de baixo nível. O GoLang herda muitos recursos do C, mas busca resolver suas complexidades.

- **Sintaxe e Estrutura:** O Go adota uma sintaxe estilo C, tornando-o familiar para muitos desenvolvedores.
- **Eficiência:** Assim como o C, Go é compilado para código de máquina, oferecendo alto desempenho adequado para programação de sistemas.
- **Influência na Programação de Sistemas:** O Go retém a eficiência do C, mas introduz recursos modernos como a coleta de lixo, que simplifica a gestão de memória.

#### Python

A simplicidade e legibilidade do Python influenciaram significativamente a filosofia de design do GoLang. O Go busca ser tão fácil de ler e escrever quanto o Python, promovendo a clareza e a manutenibilidade do código.

- **Legibilidade:** A sintaxe do Go é limpa e minimalista, semelhante à do Python, reduzindo a carga cognitiva sobre os desenvolvedores.
- **Biblioteca Padrão:** A biblioteca padrão abrangente do Python inspirou o Go a incluir um conjunto robusto de bibliotecas desde o início, simplificando tarefas comuns sem depender fortemente de dependências externas.

#### Java

O modelo de concorrência e os recursos orientados a objetos do Java forneceram lições para o modelo de concorrência do GoLang e sua simplicidade estrutural.

- **Concorrência:** Enquanto o Java usa threads para concorrência, o Go introduz goroutines, threads leves que são mais eficientes e fáceis de usar, facilitando a programação concorrente.
- **Simplicidade:** O Go evita a complexidade da extensa hierarquia de programação orientada a objetos do Java, optando por interfaces e sistemas de tipos mais simples.

#### Limbo

Limbo, projetada para o sistema operacional Inferno, também influenciou bastante o Go. O foco do Limbo em concorrência, sistemas em rede e portabilidade se reflete no GoLang.

- **Concorrência:** As threads leves do Limbo influenciaram o desenvolvimento das goroutines no Go.
- **Canais:** O conceito de canais para comunicação entre threads no Limbo é diretamente refletido no Go.
- **Portabilidade:** Assim como o Limbo, o Go busca ser portátil em várias plataformas, tornando-se uma ferramenta versátil para ambientes de computação modernos.

#### Referências

[1] "The Go Programming Language." Google. Disponível em: [https://golang.org/doc/](https://golang.org/doc/)

[2] Kernighan, B.W., Ritchie, D.M. (1988). The C Programming Language. Prentice Hall.

[3] Van Rossum, G. (1995). The Python Programming Language. USENIX Annual Technical Conference.

[4] Gosling, J., Joy, B., Steele, G., Bracha, G. (2014). The Java Language Specification. Oracle Corporation.

[5] "History of Go." Wikipedia. Disponível em: [https://en.wikipedia.org/wiki/Go_(programming_language)](https://en.wikipedia.org/wiki/Go_(programming_language))

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
[1] GO. Documentação de Effective Go. Disponível em: https://go.dev/doc/effective_go. Acesso em: 20 fev. 2024.
[2] GO. Documentação sobre IDEs para GoLang. Disponível em: https://go.dev/doc/editors. Acesso em: 10 mar. 2025.
KAMIL, Amir. Curso de Linguagens de Programação na Universidade de Michigan, outono de 2016.
