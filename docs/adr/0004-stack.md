# ADR 0004 — Stack: React + TypeScript, FastAPI, PostgreSQL e agente em Go

- **Status:** proposta
- **Data:** 2026-11-16

## Contexto

O projeto é focado em segurança (defesa) e tem três partes que conversam só por HTTP/JSON
e WebSocket: frontend, backend e agente. Cada parte pode usar a linguagem que melhor
serve ao seu papel. A segurança do sistema depende mais de como ele é construído do que
da linguagem, mas a escolha influencia a superfície de ataque e o valor do projeto para a
carreira.

## Decisão

| Parte | Escolha | Por quê |
|---|---|---|
| Frontend | React + TypeScript + Vite, TanStack Query, React Router; React Three Fiber depois da v1 | Já tenho experiência com React; tipagem ajuda a evitar erros |
| Backend | Python + FastAPI + Pydantic + SQLAlchemy + Alembic | Python é a linguagem do dia a dia em segurança; ferramentas de detecção (pySigma, bibliotecas do MITRE ATT&CK) estão em Python; FastAPI valida entradas pelos tipos e gera a documentação da API |
| Banco | PostgreSQL (Docker) | Maduro, open source; papéis e permissões finas, Row-Level Security, pgAudit, TLS |
| Agente | Go | Gera um único executável, sem runtime nem dependências instaladas na máquina monitorada: menor superfície de ataque no lugar que o sistema quer proteger |
| Testes | pytest (backend), Vitest (frontend), `go test` (agente) | Padrões de cada linguagem |

Tipos do frontend gerados a partir da documentação OpenAPI do FastAPI (`openapi-typescript`),
para front e back concordarem sobre o formato dos dados.

## Alternativas consideradas

- **Backend em Node (Fastify):** uma linguagem só e tipos compartilhados, mas perde a
  proximidade com o ecossistema de segurança.
- **Agente em Node ou Python:** mais rápido de começar, mas exige runtime e dependências
  em cada máquina monitorada.

## Consequências

- Três linguagens no projeto; Python e Go são novas para mim. Mitigação: cada uma fica
  isolada no seu papel, e o agente é a menor parte do sistema.
- Preciso entender cada trecho antes do commit, porque vou usar IA como apoio.
- Ver também a ADR 0005 (dependências e banco).
