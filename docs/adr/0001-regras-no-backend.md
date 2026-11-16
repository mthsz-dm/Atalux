# ADR 0001 — Regras de negócio ficam no backend

- **Status:** proposta
- **Data:** 2026-11-16

## Contexto

O projeto terá várias formas de apresentar os mesmos dados: lista, painel, vila 3D e, no
futuro, modos de treino. Se cada uma calcular estado, cor ou alerta por conta própria, elas
vão discordar entre si, e a regra de segurança fica espalhada pelo frontend, onde qualquer
pessoa pode alterá-la.

## Decisão

Todo cálculo de estado, aplicação de limites e geração de eventos acontece no backend.
O frontend recebe dados prontos (ex.: `estado: "atencao"`) e decide apenas **como** mostrar.

## Consequências

- Lista e vila mostram sempre a mesma coisa.
- As regras são testáveis sem interface.
- O simulador dos modos de treino só precisa produzir dados no formato da API.
- O frontend fica mais simples, mas depende da API para qualquer mudança de regra.
