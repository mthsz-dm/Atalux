# ADR 0002 — Descobrir um aparelho não é monitorá-lo

- **Status:** proposta
- **Data:** 2026-11-16

## Contexto

Um aparelho pode aparecer na rede (foi visto) sem que exista autorização ou uma fonte
confiável de dados sobre ele. Tratar "visto" como "monitorado" levaria a duas falhas:
mostrar como normal algo sobre o qual não sabemos nada, e coletar dados de máquinas sem
autorização.

## Decisão

- Ativo e fonte de dados são entidades separadas (`ATIVO` e `FONTE_DADOS`).
- Um ativo só é monitorado com uma fonte ativa, registrada pelo administrador, com token próprio.
- Sem fonte ativa, o estado é `nao_monitorado`, nunca `normal`.

## Consequências

- O sistema nunca afirma que algo está bem sem evidência.
- Na vila, isso vira três níveis visíveis: casas monitoradas, cidadãos conhecidos e
  encapuzados desconhecidos.
- Revogar um token tira o ativo do monitoramento sem apagar seu histórico.
