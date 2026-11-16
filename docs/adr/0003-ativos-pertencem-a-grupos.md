# ADR 0003 — Ativos pertencem a grupos desde o início

- **Status:** proposta
- **Data:** 2026-11-16

## Contexto

Com poucas máquinas, a vila mostra uma casa por ativo. Com centenas, isso vira poluição
visual. A solução prevista é o zoom semântico: de longe, cada grupo vira um "bairro" com
o resumo dos seus ativos. Adicionar grupos só quando a rede crescer exigiria mudar banco,
API e cena ao mesmo tempo.

## Decisão

- Todo ativo pertence a exatamente um grupo desde a v1.
- O backend calcula o resumo de cada grupo: pior estado e contagem por estado (RN07).
- A API terá um endpoint de resumo por grupo, além do detalhe por ativo.

## Consequências

- Custo quase zero no MVP (um grupo padrão já resolve).
- A visualização escala sem refatoração no modelo de dados.
- A regra de agregação é lógica de negócio e fica testada no backend (ADR 0001).
