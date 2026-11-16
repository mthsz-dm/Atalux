# Visão

## O problema

Ferramentas de monitoramento de segurança mostram a rede como tabelas e gráficos. Elas são
precisas, mas difíceis de entender para quem está começando e pouco intuitivas para ter
uma noção rápida do estado da rede como um todo (consciência situacional).

## A proposta

Um sistema de monitoramento de verdade, com uma visualização espacial por cima:

- cada rede é uma **vila** numa ilha flutuante;
- **casas** são computadores monitorados, o **castelo** é o servidor central, o **portão** é o
  roteador, as **torres de vigia** são câmeras ou sensores;
- a **muralha com barreira de luz** representa o firewall;
- **cidadãos** são aparelhos conhecidos (celulares, notebooks); **encapuzados** são aparelhos
  desconhecidos;
- o **estado** de cada ativo aparece em luz: azul (normal), amarelo (atenção), vermelho (crítico).

Estilo: **neo-medieval**. Silhuetas medievais de pedra escura, com a informação carregada
por uma camada de luz cyberpunk.

## Arquitetura em uma frase

**Um núcleo, várias caras.** Backend, regras e dados são únicos. Lista, painel, vila 3D e
futuros modos de treino são apenas formas diferentes de apresentar o mesmo núcleo.

## Objetivo

1. Portfólio para vagas de segurança (defesa).
2. Aprender, na prática, como detecção funciona.
3. Vender, só se um dia fizer sentido. Não guia nenhuma decisão agora.

## Princípios

- **Engenharia antes de estética.** A vila só entra depois que o núcleo funciona.
- **Vermelho exige evidência.** Nenhum estado muda sem um evento ou regra que justifique.
- **Descobrir não é monitorar.** Um aparelho só é monitorado com fonte de dados autorizada.
- **Segurança faz parte da construção**, não é só o tema do projeto.
- **Fatias finas.** Primeiro um ativo atravessando a arquitetura inteira; depois o resto.

## Fora do escopo por enquanto

Modos de treino (fácil, difícil, extremo), PvP red team × blue team, aparelho de coleta
dedicado, monitoramento de clientes reais, modelo de negócio. Estão registrados como ideias
futuras e não influenciam o MVP, exceto onde uma decisão de agora evita refatoração depois
(ver ADRs).
