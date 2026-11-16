# Modelo de dados (MVP)

Primeira versão, pensada para o MVP. Vai mudar com a implementação; quando mudar,
este arquivo muda junto.

```mermaid
erDiagram
    USUARIO {
        uuid id PK
        string email UK
        string senha_hash
        timestamp criado_em
    }
    GRUPO {
        uuid id PK
        string nome UK
        timestamp criado_em
    }
    ATIVO {
        uuid id PK
        uuid grupo_id FK
        string nome
        string tipo "computador, servidor, impressora..."
        string ip
        string estado "nao_monitorado, sem_dados, normal, atencao, critico"
        timestamp estado_desde
        int grade_coluna "posição na vila (opcional)"
        int grade_linha "posição na vila (opcional)"
        timestamp criado_em
    }
    FONTE_DADOS {
        uuid id PK
        uuid ativo_id FK
        string tipo "agente"
        string token_hash
        timestamp ultimo_contato
        timestamp revogada_em "nulo = ativa"
        timestamp criado_em
    }
    LEITURA {
        bigint id PK
        uuid ativo_id FK
        uuid fonte_id FK
        string metrica "temperatura_cpu"
        float valor
        timestamp medido_em
        timestamp recebido_em
    }
    EVENTO {
        uuid id PK
        uuid ativo_id FK
        string tipo "temperatura_alta, normalizado, sem_dados..."
        string severidade "info, atencao, critico"
        string descricao
        jsonb detalhes
        timestamp ocorrido_em
    }
    LOG_AUDITORIA {
        bigint id PK
        string acao "login_falhou, token_invalido..."
        string origem_ip
        jsonb detalhes
        timestamp ocorrido_em
    }

    GRUPO ||--o{ ATIVO : contem
    ATIVO ||--o{ FONTE_DADOS : "monitorado por"
    ATIVO ||--o{ LEITURA : recebe
    FONTE_DADOS ||--o{ LEITURA : envia
    ATIVO ||--o{ EVENTO : gera
```

## Notas de projeto

- **`ATIVO.estado` é derivado, mas guardado.** A fonte da verdade são os eventos; o campo
  guardado existe para a lista carregar rápido. Toda mudança de estado nasce de um evento (RN03).
- **`grade_coluna` e `grade_linha` são só posição lógica na grade.** Nunca coordenadas 3D.
  A API não sabe que existe uma vila; uma tabela ou um mapa 2D usam o mesmo dado.
- **`LEITURA.metrica` é texto**, para que novas métricas (uso de memória, eventos de login)
  entrem sem mudar a tabela.
- **`medido_em` e `recebido_em` separados**: o relógio do agente pode estar errado, e a
  diferença entre os dois também é um sinal.
- **`FONTE_DADOS` separada de `ATIVO`** materializa a distinção entre descobrir e monitorar (ADR 0002).
- Ainda não modelado (depois da v1): aparelhos descobertos na rede (cidadãos e encapuzados),
  usuários com papéis diferentes, limites por ativo.
