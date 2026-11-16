# Definição de pronto da v1

A v1 está **pronta** quando todos os itens abaixo forem verdade. Nada fora desta lista é
necessário para chamar a v1 de pronta, por mais legal que seja.

> Regra de ouro: se surgir uma ideia nova durante a v1, ela vai para
> [ideias-visualizacao.md](ideias-visualizacao.md) ou para a seção "Depois da v1" no fim
> deste arquivo, e não para o código.

## Funciona de ponta a ponta

- [ ] Um agente rodando no meu PC envia a temperatura do processador ao backend a cada intervalo configurado.
- [ ] O backend só aceita leituras de agentes com token válido.
- [ ] As leituras ficam salvas no PostgreSQL.
- [ ] Uma regra transforma leitura em evento (ex.: temperatura acima do limite) e o evento muda o estado do ativo.
- [ ] Um usuário faz login no frontend e vê a lista de ativos com o estado de cada um.
- [ ] A lista atualiza sozinha (WebSocket) quando um estado muda, sem recarregar a página.
- [ ] Ao clicar num ativo, vejo detalhes: estado, última leitura e eventos recentes.

## Qualidade

- [ ] Testes automatizados das regras de estado e das rotas principais da API.
- [ ] Testes de autorização: sem token ou com token errado, a API recusa.
- [ ] O projeto inteiro sobe com um comando (`docker compose up`).
- [ ] Nenhuma senha ou token no código ou no histórico do Git.

## Documentação

- [ ] README explica o que é, como rodar e como testar.
- [ ] Requisitos e modelo de dados atualizados com o que foi realmente construído.
- [ ] Pelo menos 3 ADRs com as decisões principais.
- [ ] Um GIF ou vídeo curto no README mostrando o fluxo funcionando.

## Depois da v1 (não fazem parte da v1)

- Vila 3D ligada aos dados reais.
- Simulador de dados falsos no formato da API.
- Eventos de login do Windows (força bruta).
- Descoberta de aparelhos na rede (cidadãos e encapuzados).
- Agrupamento por bairros (zoom semântico).
- Modos de treino.
