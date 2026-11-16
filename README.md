# Monitoramento 3D (nome provisório)

Sistema web de monitoramento de segurança de redes. Cada rede vira uma vila neo-medieval
em 3D: computadores são casas, o servidor é o castelo, o firewall é a muralha e os aparelhos
desconhecidos andam pelas ruas como encapuzados.

> Status: **fase de documentação do MVP**. Ainda não há código de produção.

## Por que este projeto existe

Portfólio para entrar na área de segurança (defesa) e aprender cibersegurança construindo
um sistema de detecção de ponta a ponta. A visualização é a assinatura do projeto; a
engenharia por trás dela é o que ele quer provar.

## Estrutura

| Pasta | O que tem |
|---|---|
| `agent/` | Agente que roda na máquina monitorada e envia leituras ao backend (Go) |
| `backend/` | API, regras de negócio, autenticação, persistência (Python + FastAPI + PostgreSQL) |
| `frontend/` | Interface web (React + TypeScript; 3D com React Three Fiber mais adiante) |
| `docs/` | Visão, requisitos, modelo de dados e decisões de arquitetura |
| `docs/adr/` | Registro de decisões de arquitetura (ADRs) |
| `docs/referencia-visual/` | Protótipo visual da vila (Three.js puro, dados falsos). Referência, não código do produto |
| `spikes/` | Testes rápidos e descartáveis para tirar dúvidas técnicas |

## Por onde ler

1. [Visão](docs/00-visao.md)
2. [Definição de pronto da v1](docs/01-definicao-de-pronto-v1.md)
3. [Requisitos do MVP](docs/02-requisitos-mvp.md)
4. [Modelo de dados](docs/03-modelo-de-dados.md)
5. [Decisões de arquitetura](docs/adr/) — todas como **proposta** até eu revisar e aceitar
6. [Ideias para a visualização](docs/ideias-visualizacao.md) (backlog, fora da v1)
