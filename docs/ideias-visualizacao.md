# Ideias para a visualização (backlog)

Nada aqui faz parte da v1. A referência visual aprovada é
[vila-neo-medieval.html](referencia-visual/vila-neo-medieval.html) (abra no navegador).

## Ajustes observados no protótipo

- **Janelas como dado.** Hoje ficam sempre azuis. Ideia: acesas quando o computador está
  ligado e em uso, apagadas quando desligado. Um PC aceso às 3h numa escola já merece atenção.
- **Trilhas de dados confusas.** Todas aparecem o tempo todo. Ideia: mostrar a trilha só
  do ativo selecionado ou dos ativos em alerta.
- **PC e impressora parecidos.** A forma deve dizer o tipo e a cor deve dizer o estado.
  Antes de desenhar, fechar a lista de tipos de aparelho que o sistema reconhece.

## Vocabulário visual

| Na rede | Na vila |
|---|---|
| Computador monitorado | Casa |
| Servidor central | Castelo (2×2) |
| Roteador | Portão |
| Câmera / sensor | Torre de vigia |
| Impressora | Oficina |
| Firewall | Muralha com barreira de luz |
| Aparelho conhecido (celular, notebook) | Cidadão |
| Aparelho desconhecido | Encapuzado |
| Máquina em quarentena | Masmorra (ideia) |
| Grupo / sub-rede | Bairro (zoom semântico) |

## Ideias maiores

- Zoom semântico: de longe, bairros com pior estado e contagem; de perto, as casas.
- Busca que leva a câmera até um ativo.
- Filtro "só o que tem alerta".
- Origem do ataque por geolocalização de IP (GeoLite2) e intruso vindo daquele lado.
- Visão simples (2D) para máquina fraca, com as mesmas cores.
- Desempenho: `frameloop="demand"`, `InstancedMesh`, carregar o 3D sob demanda.
- Modos de treino com o mesmo jeito de usar do modo profissional.
- PvP no formato red team × blue team (CTF de ataque e defesa).
