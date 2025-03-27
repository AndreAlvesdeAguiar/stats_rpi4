Este repositório contém um projeto em Go que exibe informações do sistema do Raspberry Pi 4 em um display OLED SSD1306. O objetivo é criar um dashboard em tempo real para monitorar estatísticas como uso de CPU, memória, e temperatura, exibindo essas informações em um display conectado ao Raspberry Pi 4.

# Pré-requisitos

Raspberry Pi 4 com Ubuntu instalado.

Display OLED SSD1306.

Go instalado (versão mínima recomendada: 1.20).

# Instalação

## Clone o repositório:

Clonando o repositório
### git clone https://github.com/AndreAlvesdeAguiar/stats_rpi4.git

Entre no diretório do projeto:

## Acessando o diretório do projeto
 cd stats_rpi4

Instale as dependências do Go:

## Instalando dependências
 go mod tidy

# Uso

Compile o programa:

# Compilando o programa
 go build -o stats_rpi4

Execute o programa:

# Executando o programa
 ./stats_rpi4

O programa irá coletar dados do sistema e exibi-los no display OLED conectado ao Raspberry Pi 4.

### Estrutura do Projeto

#├── main.go               # Arquivo principal do projeto
#├── go.mod                # Arquivo de dependências do Go
#├── README.md             # Documentação do projeto

### Funcionalidades Implementadas

Exibição da temperatura do sistema.

Exibição do uso da CPU.

Exibição do uso da memória.

### Configuração do Display OLED SSD1306

Certifique-se de que o display OLED SSD1306 esteja conectado corretamente ao seu Raspberry Pi 4.
As conexões devem ser feitas usando os pinos I2C (SCL e SDA).

# Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo LICENSE para mais detalhes.

### Contato

Criado por Andre Alves de Aguiar - GitHub

Se você tiver alguma dúvida ou sugestão, sinta-se à vontade para entrar em contato.
