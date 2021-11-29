curl -s https://learn.01founders.co/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Cassidy.Hall94\"}}){id}}"}'|jq .id

