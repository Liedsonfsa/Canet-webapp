package models

import "time"

// Publicacao representa uma publicação feita pelo usuário
type Publicacao struct {
	ID        		uint64 		`json:"id,omitempty"`
	Titulo    		string 		`json:"titulo,omitempty"`
	Conteudo  		string 		`json:"conteudo,omitempty"`
	AutorID   		uint64 		`json:"autorId,omitempty"`
	AutorNick 		string 		`json:"autorNick,omitempty"`
	Curtidas  		uint64 		`json:"curtidas"`
	CriadaEm  		time.Time 	`json:"criadaEm,omitempty"`
}

/*
insert into publicacoes (titulo, conteudo, autor_id) values
("Olá mundo", "Essa é a minha primeira publicação aqui", 15),
("Algém aqui", "Alguém paea fazer amizade?", 15),
("Reyes de Europa", "Espanha campeã da euro.", 15);
*/