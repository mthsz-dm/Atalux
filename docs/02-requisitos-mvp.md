# Requisitos do MVP

Escopo: **um ativo monitorado atravessando a arquitetura inteira.**
Fonte de dados → backend → banco → API → frontend → representação.

Itens marcados com **[decidir]** ainda precisam de uma decisão minha antes do código.

## Atores

| Ator | Quem é |
|---|---|
| Administrador | Pessoa que usa o sistema: cadastra ativos, gera tokens, acompanha estados |
| Agente | Programa instalado numa máquina monitorada; só envia dados, não consulta nada |

## Requisitos funcionais

| ID | Requisito |
|---|---|
| RF01 | O administrador faz login com e-mail e senha. |
| RF02 | O administrador cadastra um grupo (ex.: "Laboratório", "Administrativo"). |
| RF03 | O administrador cadastra um ativo com nome, tipo, grupo e endereço IP. |
| RF04 | O administrador registra uma fonte de dados (agente) para um ativo e recebe um token, exibido uma única vez. |
| RF05 | O administrador pode revogar o token de um agente. |
| RF06 | O agente envia leituras de temperatura do processador ao backend em intervalos regulares. |
| RF07 | O sistema armazena cada leitura com data, hora e ativo de origem. |
| RF08 | O sistema aplica regras às leituras e gera eventos quando um limite é ultrapassado ou quando a situação volta ao normal. |
| RF09 | O estado do ativo é recalculado a partir dos eventos (ver regras de negócio). |
| RF10 | O administrador vê a lista de ativos com nome, grupo e estado. |
| RF11 | O administrador vê o detalhe de um ativo: estado, última leitura e últimos eventos. |
| RF12 | A interface recebe mudanças de estado em tempo real, sem recarregar a página. |
| RF13 | O sistema marca o ativo como "sem dados" quando o agente fica sem enviar leituras por um tempo. |

## Regras de negócio

| ID | Regra |
|---|---|
| RN01 | Um ativo só é **monitorado** se tiver pelo menos uma fonte de dados ativa (token não revogado). Sem isso, o estado é **não monitorado**. |
| RN02 | Estados possíveis: `nao_monitorado`, `sem_dados`, `normal`, `atencao`, `critico`. |
| RN03 | Todo estado `atencao` ou `critico` precisa estar ligado a pelo menos um evento que o justifique. |
| RN04 | Limites padrão de temperatura: atenção a partir de **[decidir]** °C, crítico a partir de **[decidir]** °C. Sugestão inicial: 80 °C e 90 °C, configuráveis por ativo depois. |
| RN05 | Para evitar alarme por pico isolado, o estado só piora após **[decidir]** leituras seguidas acima do limite. Sugestão: 3. |
| RN06 | Após **[decidir]** minutos sem leitura, o ativo vai para `sem_dados`. Sugestão: 3× o intervalo de envio do agente. |
| RN07 | O estado de um grupo é o pior estado entre seus ativos, acompanhado da contagem por estado (prepara os "bairros" da vila). |

## Requisitos não funcionais

| ID | Requisito |
|---|---|
| RNF01 | Senhas guardadas com hash forte (argon2). Nunca em texto puro. |
| RNF02 | Tokens de agente guardados apenas como hash; o valor original só aparece no momento da criação. |
| RNF03 | Toda entrada da API é validada por esquema (Pydantic) antes de chegar às regras. |
| RNF04 | Rotas do administrador exigem sessão válida; rotas do agente exigem token válido; uma não serve para a outra (menor privilégio). |
| RNF05 | Erros não expõem detalhes internos (stack trace, SQL) para quem chama a API. |
| RNF06 | Tentativas de login e de envio com token inválido são registradas para auditoria. |
| RNF07 | Regras de estado ficam no backend; o frontend apenas apresenta. |
| RNF08 | O projeto sobe com `docker compose up`, sem passos manuais além de um arquivo `.env`. |
| RNF09 | Regras de negócio e rotas principais têm testes automatizados. |
| RNF10 | Comunicação entre agente e backend preparada para HTTPS (em desenvolvimento local, HTTP é aceito). |
| RNF11 | Tentativas de login limitadas por tempo e por origem (proteção contra força bruta). |
| RNF12 | Se o login usar cookie (ver dúvidas abertas), as rotas que alteram dados têm proteção contra CSRF; o cookie é `httpOnly`, `Secure` e `SameSite`. |
| RNF13 | Respostas com cabeçalhos de segurança (ex.: Content-Security-Policy, X-Content-Type-Options). |
| RNF14 | Acesso ao banco com menor privilégio e dependências controladas (ver ADR 0005). |
| RNF15 | Os requisitos de segurança seguem como referência o **OWASP ASVS** (nível 1 no MVP). |

## Fora do MVP

Vila 3D, simulador, descoberta de aparelhos na rede, eventos de login do Windows,
múltiplos usuários com papéis diferentes, notificações, modos de treino.

## Dúvidas abertas

- **[decidir]** Login com sessão por cookie ou JWT? Sugestão: cookie de sessão `httpOnly`, mais simples de proteger num app web próprio.
- **[decidir]** Intervalo de envio do agente. Sugestão: 30 segundos.
- **[decidir]** Por quanto tempo guardar as leituras brutas? Sugestão: 30 dias no MVP.
