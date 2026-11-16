# ADR 0005 — Uso seguro do banco e das dependências

- **Status:** proposta
- **Data:** 2026-11-16

## Contexto

Duas fontes comuns de falha em sistemas reais não estão no código que eu escrevo: o
acesso ao banco com permissões demais e as bibliotecas de terceiros, que podem ter falhas
ou até código malicioso (ataques à cadeia de suprimentos).

## Decisão

**Banco**
- A aplicação conecta com um usuário **sem permissão de alterar tabelas**; as migrações
  rodam com outro usuário; o superusuário não é usado pela aplicação.
- Credenciais só no `.env`, nunca no código.
- Consultas sempre parametrizadas (padrão do SQLAlchemy); nada de montar SQL com texto concatenado.
- Logs de eventos ficam no PostgreSQL no MVP. Quando os eventos crescerem (ex.: logins do
  Windows em volume), avaliar **OpenSearch** para logs, como fazem os SIEMs, mantendo o
  PostgreSQL para usuários, ativos e configuração.

**Dependências**
- Poucas bibliotecas, preferindo as mais conhecidas e mantidas.
- Arquivos de trava sempre no Git: `package-lock.json`, `uv.lock` (ou equivalente) e `go.sum`.
- Instalação reprodutível (`npm ci`, instalação a partir do lock no Python).
- Dependabot ligado no GitHub e auditoria de dependências no CI (`npm audit`, `pip-audit`, `govulncheck`).

## Consequências

- Mais configuração no início, mas é exatamente o que uma equipe de segurança cobraria.
- Vira uma seção do README sobre como o projeto trata a cadeia de suprimentos.
