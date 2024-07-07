package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario cria um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao ler o corpo da resquisição!"))
		return
	}

	var usuario usuario

	if err := json.Unmarshal(request, &usuario); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao coverter o usuário para struct!"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao conectar com banco de dados."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	insertion, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, err := insertion.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao obter o ID inserido!"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}

// BuscarUsuarios retorna uma lista de todos os usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar com banco de dados!"))
		return
	}
	defer db.Close()

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Erro ao buscar usuários!"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Erro ao escanear banco de dados!"))
			return
		}
		usuarios = append(usuarios, usuario)

	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao retornar json com a lista de usuários!"))
		return
	}

}

// BuscarUsuario retorna apenas um usuário
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	id, err := strconv.ParseInt(parameters["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
	}

	db, err := banco.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar com banco!"))
		return
	}
	defer db.Close()

	linha, err := db.Query("select * from usuarios where id = ?", id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}
	defer linha.Close()

	var usuario usuario
	if linha.Next() {
		if err := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Erro ao escanear o usuário!"))
		}
	}

	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Erro ao retornar json!"))
		return
	}
}

// AtualizaUsario edita as informações de usuários
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parâmetro para inteiro."))
		return
	}

	corpoDaRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao ler corpo da requisição!"))
		return
	}

	var usuario usuario
	if err := json.Unmarshal(corpoDaRequisicao, &usuario); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter usuário para struct!"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome=?, email=? where id = ?")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao criar statement!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Email, ID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario deleta os usuarios do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao deletar usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
