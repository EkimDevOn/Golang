


GORM em Golang

O que é GORM?

GORM é um framework ORM (Object-Relational Mapping) para a linguagem Go. ORM é um padrão de projeto que permite que os desenvolvedores interajam com bancos de dados relacionais usando objetos e métodos de programação orientada a objetos, em vez de escrever SQL manualmente.
Principais características do GORM

Suporte a múltiplos bancos de dados: GORM suporta MySQL, PostgreSQL, SQLite, SQL Server e outros.

Modelos e Mapeamento: Facilita a definição de modelos (estruturas) que mapeiam tabelas de banco de dados.

CRUD completo: Oferece métodos para Create, Read, Update e Delete de registros.

Relacionamentos: Suporta relacionamentos como Has One, Has Many, Belongs To, Many To Many, entre outros.

Migrações automáticas: Pode migrar esquemas de banco de dados automaticamente.

Transações: Suporta transações, incluindo transações aninhadas e pontos de salvamento.

Preloading: Permite carregar relacionamentos de forma eficiente.

Logs e Debugging: Oferece suporte a logs e modo de depuração para ajudar no desenvolvimento.

Instalação

Para começar a usar GORM, você precisa instalar o pacote e o driver do banco de dados que você pretende usar:

go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u gorm.io/driver/mysql
go get -u gorm.io/driver/postgres
go get -u gorm.io/driver/sqlserver

Exemplo Básico
Aqui está um exemplo simples de como usar GORM para criar uma conexão com um banco de dados SQLite, criar uma tabela e inserir um registro:

package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Product struct {
    gorm.Model
    Code  string
    Price uint
}



func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrar schema
    db.AutoMigrate(&Product{})

    // Criar
    db.Create(&Product{Code: "D42", Price: 100})
}


Modelos e Mapeamento

GORM permite que você defina modelos usando estruturas Go. Você pode incluir campos especiais como gorm.Model para adicionar automaticamente campos como ID, CreatedAt, UpdatedAt e DeletedAt:

type User struct {
    gorm.Model
    Name  string
    Email string
}
Relacionamentos
GORM suporta uma variedade de relacionamentos entre modelos. Por exemplo, um relacionamento "Has One":
go
复制
type User struct {
    gorm.Model
    CreditCard CreditCard
}

type CreditCard struct {
    gorm.Model
    Number string
    UserID uint
}


Conclusão
GORM é um framework ORM poderoso e flexível para Go que facilita a interação com bancos de dados.
 Com suporte a múltiplos bancos de dados, recursos de CRUD completo, relacionamentos e muito mais,
  é uma ferramenta valiosa para desenvolvedores Go. Para mais informações e exemplos detalhados,
   você pode consultar a documentação oficial em https://gorm.io/